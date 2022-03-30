package http

import (
	"net/http"

	"djinn-ci.com/alert"
	"djinn-ci.com/errors"
	"djinn-ci.com/namespace"
	"djinn-ci.com/server"
	"djinn-ci.com/template"
	"djinn-ci.com/user"
	userhttp "djinn-ci.com/user/http"
	"djinn-ci.com/variable"
	variabletemplate "djinn-ci.com/variable/template"

	"github.com/andrewpillar/webutil"

	"github.com/gorilla/csrf"
)

type UI struct {
	*Handler
}

func (h UI) Index(u *user.User, w http.ResponseWriter, r *http.Request) {
	sess, save := h.Session(r)

	vv, paginator, err := h.IndexWithRelations(u, r)

	if err != nil {
		h.InternalServerError(w, r, errors.Err(err))
		return
	}

	unmasked := variable.GetUnmasked(sess.Values)

	for _, v := range vv {
		if _, ok := unmasked[v.ID]; ok && v.Masked {
			if err := variable.Unmask(h.AESGCM, v); err != nil {
				alert.Flash(sess, alert.Danger, "Could not unmask variable")
				h.Log.Error.Println(r.Method, r.URL, "could not unmask variable", errors.Err(err))
			}
			continue
		}

		if v.Masked {
			v.Value = variable.MaskString
		}
	}

	if err := variable.LoadNamespaces(h.DB, vv...); err != nil {
		h.InternalServerError(w, r, errors.Err(err))
		return
	}

	csrf := csrf.TemplateField(r)

	p := &variabletemplate.Index{
		BasePage: template.BasePage{
			URL:  r.URL,
			User: u,
		},
		CSRF:      csrf,
		Search:    r.URL.Query().Get("search"),
		Paginator: paginator,
		Unmasked:  unmasked,
		Variables: vv,
	}
	d := template.NewDashboard(p, r.URL, u, alert.First(sess), csrf)
	save(r, w)
	webutil.HTML(w, template.Render(d), http.StatusOK)
}

func (h UI) Create(u *user.User, w http.ResponseWriter, r *http.Request) {
	sess, save := h.Session(r)

	csrf := csrf.TemplateField(r)

	p := &variabletemplate.Create{
		Form: template.Form{
			CSRF:   csrf,
			Errors: webutil.FormErrors(sess),
			Fields: webutil.FormFields(sess),
		},
	}
	d := template.NewDashboard(p, r.URL, u, alert.First(sess), csrf)
	save(r, w)
	webutil.HTML(w, template.Render(d), http.StatusOK)
}

func (h UI) Store(u *user.User, w http.ResponseWriter, r *http.Request) {
	sess, _ := h.Session(r)

	v, f, err := h.StoreModel(u, r)

	if err != nil {
		cause := errors.Cause(err)

		errs := webutil.NewValidationErrors()

		switch err := cause.(type) {
		case webutil.ValidationErrors:
			if errs, ok := err["fatal"]; ok {
				h.Log.Error.Println(r.Method, r.URL, errors.Slice(errs))
				alert.Flash(sess, alert.Danger, "Failed to create variable")
				h.RedirectBack(w, r)
				return
			}
			errs = err
		case *namespace.PathError:
			errs.Add("namespace", err)
		default:
			h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
			alert.Flash(sess, alert.Danger, "Failed to create variable")
			h.RedirectBack(w, r)
			return
		}

		webutil.FlashFormWithErrors(sess, f, errs)
		h.RedirectBack(w, r)
		return
	}

	alert.Flash(sess, alert.Success, "Variable has been added: "+v.Key)
	h.Redirect(w, r, "/variables")
}

func (h UI) Mask(u *user.User, v *variable.Variable, w http.ResponseWriter, r *http.Request) {
	sess, _ := h.Session(r)

	if u.ID != v.AuthorID {
		h.RedirectBack(w, r)
		return
	}

	unmasked := variable.GetUnmasked(sess.Values)

	if _, ok := unmasked[v.ID]; ok {
		delete(unmasked, v.ID)
		variable.PutUnmasked(sess.Values, unmasked)
	}

	h.RedirectBack(w, r)
}

func (h UI) Unmask(u *user.User, v *variable.Variable, w http.ResponseWriter, r *http.Request) {
	h.Log.Debug.Println(r.Method, r.URL, "unmasking variable", v.ID)
	sess, _ := h.Session(r)

	if u.ID != v.AuthorID {
		h.RedirectBack(w, r)
		return
	}

	unmasked := variable.GetUnmasked(sess.Values)
	unmasked[v.ID] = struct{}{}

	h.RedirectBack(w, r)
}

func (h UI) Destroy(u *user.User, v *variable.Variable, w http.ResponseWriter, r *http.Request) {
	sess, _ := h.Session(r)

	if err := h.DeleteModel(r.Context(), v); err != nil {
		alert.Flash(sess, alert.Success, "Failed to delete variable")
		h.RedirectBack(w, r)
		return
	}

	alert.Flash(sess, alert.Success, "Variable has been deleted")
	h.RedirectBack(w, r)
}

func RegisterUI(srv *server.Server) {
	user := userhttp.NewHandler(srv)

	ui := UI{
		Handler: NewHandler(srv),
	}

	sr := srv.Router.PathPrefix("/variables").Subrouter()
	sr.HandleFunc("", user.WithUser(ui.Index)).Methods("GET")
	sr.HandleFunc("/create", user.WithUser(ui.Create)).Methods("GET")
	sr.HandleFunc("", user.WithUser(ui.Store)).Methods("POST")
	sr.HandleFunc("/{variable:[0-9]+}/mask", user.WithUser(ui.WithVariable(ui.Mask))).Methods("PATCH")
	sr.HandleFunc("/{variable:[0-9]+}/unmask", user.WithSudo(ui.WithVariable(ui.Unmask))).Methods("GET")
	sr.HandleFunc("/{variable:[0-9]+}", user.WithUser(ui.WithVariable(ui.Destroy))).Methods("DELETE")
	sr.Use(srv.CSRF)
}