package web

import (
	"context"
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/log"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/namespace"
	"github.com/andrewpillar/thrall/oauth2"
	"github.com/andrewpillar/thrall/user"

	"github.com/andrewpillar/query"

	"github.com/gorilla/mux"

	"github.com/jmoiron/sqlx"
)

type Middleware struct {
	Handler
}

type modelFunc func(int64) (model.Model, error)

type errHandler func(http.ResponseWriter, string, int)

// Gate serves as a stripped down middleware handler function that will be
// passed the current user in the request, if any, along with the request
// itself. This will determine whether the given user can access whatever is on
// the other end of the current endpoint, hence the bool return value.
type Gate func(u *user.User, r *http.Request) (*http.Request, bool, error)

// Resource returns whether the current user has access to the given resource.
// The resource's ID will be taken from the request based on the name, this is
// passed back to the modelFunc which will return the underlying model for that
// resource.
func Resource(db *sqlx.DB, name string, r *http.Request, get modelFunc) (bool, error) {
	val := r.Context().Value("user")
	u, ok := val.(*user.User)

	if !ok {
		return false, errors.New("user not in request context")
	}

	base := filepath.Base(r.URL.Path)

	if base == name + "s" || base == "create" {
		return !u.IsZero(), nil
	}

	vars := mux.Vars(r)

	id, _ := strconv.ParseInt(vars[name], 10, 64)

	m, err := get(id)

	if err != nil {
		return false, errors.Err(err)
	}

	if m.IsZero() {
		return false, nil
	}

	namespaceId, _ := m.Values()["namespace_id"].(sql.NullInt64)

	if !namespaceId.Valid {
		userId, ok := m.Values()["user_id"].(int64)

		if !ok {
			return false, nil
		}
		return u.ID == userId, nil
	}

	root, err := namespace.NewStore(db).Get(
		query.WhereQuery("root_id", "=", namespace.SelectRootID(namespaceId.Int64)),
		query.WhereQuery("id", "=", namespace.SelectRootID(namespaceId.Int64)),
	)

	if err != nil {
		return false, errors.Err(err)
	}

	cc, err := namespace.NewCollaboratorStore(db, root).All()

	if err != nil {
		return false, errors.Err(err)
	}

	root.LoadCollaborators(cc)
	return root.AccessibleBy(u), nil
}

// Get the currently authenticated user from the request. Check for token
// auth first, then fallback to cookie.
func (h Middleware) auth(w http.ResponseWriter, r *http.Request) (*user.User, bool) {
	if _, ok := r.Header["Authorization"]; ok {
		u, t, err := h.UserToken(r)

		if err != nil {
			log.Error.Println(r.Method, r.URL.Path, errors.Err(err))
			return u, false
		}

		if !u.IsZero() {
			u.Permissions = t.Permissions()
		}
		return u, !u.IsZero()
	}

	u, err := h.UserCookie(r)

	if err != nil {
		cause := errors.Cause(err)

		if strings.Contains(cause.Error(), "expired timestamp") {
			c := &http.Cookie{
				Name:     "user",
				HttpOnly: true,
				Path:     "/",
				Expires:  time.Unix(0, 0),
			}
			http.SetCookie(w, c)
		} else {
			log.Error.Println(errors.Err(err))
		}
		return u, false
	}

	if !u.IsZero() {
		for _, res := range oauth2.Resources {
			u.Permissions[res.String()+":read"] = struct{}{}
			u.Permissions[res.String()+":write"] = struct{}{}
			u.Permissions[res.String()+":delete"] = struct{}{}
		}
	}
	return u, !u.IsZero()
}

func (h Middleware) Guest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := h.auth(w, r); ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, ok := h.auth(w, r)

		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "user", u))

		next.ServeHTTP(w, r)
	})
}

func (h Middleware) Gate(gates ...Gate) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, _ := h.auth(w, r)

			var (
				ok   bool
				err  error
				errh errHandler = HTMLError
			)

			if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
				errh = JSONError
			}

			r = r.WithContext(context.WithValue(r.Context(), "user", u))

			for _, g := range gates {
				r, ok, err = g(u, r)

				if err != nil {
					errh(w, "Something went wrong", http.StatusInternalServerError)
					return
				}

				if !ok {
					errh(w, "Not found", http.StatusNotFound)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
