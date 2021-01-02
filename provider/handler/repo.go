package handler

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andrewpillar/djinn/crypto"
	"github.com/andrewpillar/djinn/database"
	"github.com/andrewpillar/djinn/errors"
	"github.com/andrewpillar/djinn/provider"
	providertemplate "github.com/andrewpillar/djinn/provider/template"
	"github.com/andrewpillar/djinn/template"
	"github.com/andrewpillar/djinn/user"
	"github.com/andrewpillar/djinn/web"

	"github.com/andrewpillar/query"
	"github.com/andrewpillar/webutil"

	"github.com/gorilla/csrf"

	"github.com/go-redis/redis"
)

// Repo is the handler that handles the adding and removing of webhooks to a
// repo on the provider.
type Repo struct {
	web.Handler

	redis     *redis.Client
	block     *crypto.Block
	providers *provider.Registry
}

var cacheKey = "repos-%s-%v-%v"

func NewRepo(h web.Handler, redis *redis.Client, block *crypto.Block, providers *provider.Registry) Repo {
	return Repo{
		Handler:   h,
		redis:     redis,
		block:     block,
		providers: providers,
	}
}

func (h Repo) cachePut(name string, id int64, rr []*provider.Repo, paginator database.Paginator) error {
	var (
		buf bytes.Buffer
		err error
	)

	enc := gob.NewEncoder(&buf)

	rr1 := make([]*provider.Repo, 0, len(rr))

	for _, r := range rr {
		rr1 = append(rr1, &provider.Repo{
			ID:             r.ID,
			UserID:         r.UserID,
			ProviderID:     r.ProviderID,
			ProviderUserID: r.ProviderUserID,
			HookID:         r.HookID,
			RepoID:         r.RepoID,
			Enabled:        r.Enabled,
			Name:           r.Name,
			Href:           r.Href,
		})
	}

	if err := enc.Encode(rr1); err != nil {
		return errors.Err(err)
	}

	if err := enc.Encode(paginator); err != nil {
		return errors.Err(err)
	}

	_, err = h.redis.Set(fmt.Sprintf(cacheKey, name, id, paginator.Page), buf.String(), time.Hour).Result()
	return errors.Err(err)
}

func (h Repo) cacheGet(name string, id, page int64) ([]*provider.Repo, database.Paginator, error) {
	var err error

	rr := make([]*provider.Repo, 0)
	paginator := database.Paginator{}

	s, err := h.redis.Get(fmt.Sprintf(cacheKey, name, id, page)).Result()

	if err != nil {
		if err == redis.Nil {
			return rr, paginator, nil
		}
		return rr, paginator, errors.Err(err)
	}

	r := strings.NewReader(s)
	dec := gob.NewDecoder(r)

	if err := dec.Decode(&rr); err != nil {
		return rr, paginator, errors.Err(err)
	}

	err = dec.Decode(&paginator)
	return rr, paginator, errors.Err(err)
}

func (h Repo) loadRepos(u *user.User, name string, page int64, reload bool) ([]*provider.Repo, *provider.Provider, []*provider.Provider, database.Paginator, error) {
	var paginator database.Paginator

	pp, err := provider.NewStore(h.DB, u).All(query.OrderAsc("name"))

	if err != nil {
		return nil, nil, nil, paginator, errors.Err(err)
	}

	accounts := make([]*provider.Provider, 0, len(pp))
	providerLookup := make(map[string]*provider.Provider)

	for _, p := range pp {
		if p.MainAccount {
			accounts = append(accounts, p)
		}
		providerLookup[p.Name + strconv.FormatInt(p.ProviderUserID.Int64, 10)] = p
	}

	p := accounts[0]

	if name != "" {
		for _, acc := range accounts {
			if acc.Name == name {
				p = acc
				break
			}
		}
	}

	name = p.Name

	rr, paginator, err := h.cacheGet(p.Name, u.ID, page)

	if len(rr) == 0 || reload {
		rr, paginator, err = p.Repos(h.block, h.providers, page)

		if err != nil {
			return nil, nil, nil, paginator, errors.Err(err)
		}

		if err := h.cachePut(name, u.ID, rr, paginator); err != nil {
			return nil, nil, nil, paginator, errors.Err(err)
		}
	}

	for _, r := range rr {
		r.Provider = providerLookup[p.Name + strconv.FormatInt(r.ProviderUserID, 10)]
		r.ProviderID = r.Provider.ID
	}

	enabled, err := provider.NewRepoStore(h.DB, u).All(query.Where("enabled", "=", query.Arg(true)))

	if err != nil {
		return nil, nil, nil, paginator, errors.Err(err)
	}

	enabledLookup := make(map[int64]map[int64]int64)

	for _, r := range enabled {
		if _, ok := enabledLookup[r.ProviderID]; !ok {
			enabledLookup[r.ProviderID] = make(map[int64]int64)
		}
		enabledLookup[r.ProviderID][r.RepoID] = r.ID
	}

	for _, r := range rr {
		if id, ok := enabledLookup[r.ProviderID][r.RepoID]; ok {
			r.ID = id
			r.Enabled = true
		}
	}
	return rr, p, accounts, paginator, nil
}

// Index serves the HTML response detailing the repositories retrieved from the
// provider in the request. It will first attempt to get the repositores from
// the cache, if the cache is empty then the API is hit, and the cache stores
// the response for 1 hour.
func (h Repo) Index(w http.ResponseWriter, r *http.Request) {
	sess, save := h.Session(r)

	u, ok := user.FromContext(r.Context())

	if !ok {
		h.Log.Error.Println(r.Method, r.URL, "failed to get user from request context")
	}

	q := r.URL.Query()

	page, err := strconv.ParseInt(q.Get("page"), 10, 64)

	if err != nil {
		page = 1
	}

	rr, p, pp, paginator, err := h.loadRepos(u, q.Get("name"), page, false)

	if err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	csrfField := string(csrf.TemplateField(r))

	pg := &providertemplate.RepoIndex{
		BasePage: template.BasePage{
			URL:  r.URL,
			User: u,
		},
		CSRF:      csrfField,
		Paginator: paginator,
		Repos:     rr,
		Provider:  p,
		Providers: pp,
	}
	d := template.NewDashboard(pg, r.URL, u, web.Alert(sess), csrfField)
	save(r, w)
	webutil.HTML(w, template.Render(d), http.StatusOK)
}

// Update will reload all of the provider repos in the cache.
func (h Repo) Update(w http.ResponseWriter, r *http.Request) {
	sess, _ := h.Session(r)

	u, ok := user.FromContext(r.Context())

	if !ok {
		h.Log.Error.Println(r.Method, r.URL, "failed to get user from request context")
	}

	q := r.URL.Query()

	page, err := strconv.ParseInt(q.Get("page"), 10, 64)

	if err != nil {
		page = 1
	}

	if _, _, _, _, err := h.loadRepos(u, q.Get("provider"), page, true); err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to refresh repository cache",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	sess.AddFlash(template.Alert{
		Level:   template.Success,
		Close:   true,
		Message: "Successfully reloaded repository cache",
	}, "alert")
	h.RedirectBack(w, r)
}

// Store will add a webhook to the provider's repo in the given request.
func (h Repo) Store(w http.ResponseWriter, r *http.Request) {
	sess, _ := h.Session(r)

	u, ok := user.FromContext(r.Context())

	if !ok {
		h.Log.Error.Println(r.Method, r.URL, "failed to get user from request context")
	}

	f := &provider.RepoForm{}

	if err := webutil.Unmarshal(f, r); err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to enable repository hooks",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	p, err := provider.NewStore(h.DB, u).Get(query.Where("id", "=", query.Arg(f.ProviderID)))

	if err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to enable repository hooks",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	repos := provider.NewRepoStore(h.DB, u, p)

	repo, err := repos.Get(query.Where("repo_id", "=", query.Arg(f.RepoID)))

	if err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to enable repository hooks",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	if repo.IsZero() {
		repo.ProviderID = f.ProviderID
		repo.RepoID = f.RepoID
		repo.Name = f.Name
	}

	if err := p.ToggleRepo(h.block, h.providers, repo); err != nil {
		cause := errors.Cause(err)

		if cause == provider.ErrLocalhost {
			sess.AddFlash(template.Alert{
				Level:   template.Danger,
				Close:   true,
				Message: "Failed to enable repository hooks: "+cause.Error(),
			}, "alert")
			h.RedirectBack(w, r)
			return
		}
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to enable repository hooks",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	if err := repos.Touch(repo); err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to enable repository hooks",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	sess.AddFlash(template.Alert{
		Level:   template.Success,
		Close:   true,
		Message: "Repository hooks enabled",
	}, "alert")
	h.RedirectBack(w, r)
}

// Destroy will remove the webhook from the repo in the request.
func (h Repo) Destroy(w http.ResponseWriter, r *http.Request) {
	sess, _ := h.Session(r)

	repo, ok := provider.RepoFromContext(r.Context())

	if !ok {
		h.Log.Error.Println(r.Method, r.URL, "failed to get repo from request context")
	}

	if !repo.Enabled {
		sess.AddFlash(template.Alert{
			Level:   template.Success,
			Close:   true,
			Message: "Repository hooks disabled",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	p, err := provider.NewStore(h.DB).Get(query.Where("id", "=", query.Arg(repo.ProviderID)))

	if err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to disable repository hooks",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	if err := p.ToggleRepo(h.block, h.providers, repo); err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to disable repository hooks",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	if err := provider.NewRepoStore(h.DB).Update(repo); err != nil {
		h.Log.Error.Println(r.Method, r.URL, errors.Err(err))
		sess.AddFlash(template.Alert{
			Level:   template.Danger,
			Close:   true,
			Message: "Failed to disable repository hooks",
		}, "alert")
		h.RedirectBack(w, r)
		return
	}

	sess.AddFlash(template.Alert{
		Level:   template.Success,
		Close:   true,
		Message: "Repository hooks disabled",
	}, "alert")
	h.RedirectBack(w, r)
}
