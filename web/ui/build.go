package ui

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/andrewpillar/thrall/config"
	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/form"
	"github.com/andrewpillar/thrall/log"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
	"github.com/andrewpillar/thrall/template/build"
	"github.com/andrewpillar/thrall/web"

	"github.com/gorilla/mux"

	"github.com/RichardKnop/machinery/v1"
)

type Build struct {
	web.Handler

	queues     map[string]*machinery.Server
	namespaces *model.NamespaceStore
}

func NewBuild(h web.Handler, queues map[string]*machinery.Server, namespaces *model.NamespaceStore) Build {
	return Build{
		Handler:    h,
		queues:     queues,
		namespaces: namespaces,
	}
}

func (h Build) Index(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	status := r.URL.Query().Get("status")

	bb, err := u.BuildList(status)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	p := &build.IndexPage{
		Page: &template.Page{
			URI: r.URL.Path,
		},
		Builds: bb,
		Status: status,
	}

	d := template.NewDashboard(p, r.URL.Path)

	web.HTML(w, template.Render(d), http.StatusOK)
}

func (h Build) Create(w http.ResponseWriter, r *http.Request) {
	p := &build.CreatePage{
		Errors: h.Errors(w, r),
		Form:   h.Form(w, r),
	}

	d := template.NewDashboard(p, r.URL.RequestURI())

	web.HTML(w, template.Render(d), http.StatusOK)
}

func (h Build) Store(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	f := &form.Build{}

	if err := h.ValidateForm(f, w, r); err != nil {
		if _, ok := err.(form.Errors); ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	m, _ := config.DecodeManifest(strings.NewReader(f.Manifest))

	srv, ok := h.queues[m.Driver.Type]

	if !ok {
		errs := form.NewErrors()
		errs.Put("manifest", errors.New("Driver " + m.Driver.Type + " is not supported"))

		h.FlashForm(w, r, f)
		h.FlashErrors(w, r, errs)

		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	b := u.BuildStore().New()
	b.User = u
	b.Manifest = f.Manifest

	if f.Namespace != "" {
		n, err := u.NamespaceStore().FindOrCreate(f.Namespace)

		if err != nil {
			log.Error.Println(errors.Err(err))
			web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		b.NamespaceID = sql.NullInt64{
			Int64: n.ID,
			Valid: true,
		}
	}

	if err := b.Create(); err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err := b.Submit(srv); err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h Build) Show(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["build"], 10, 64)

	if err != nil {
		web.HTMLError(w, "Not found", http.StatusNotFound)
		return
	}

	b, err := u.BuildShow(id)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if b.IsZero() || u.ID != b.UserID {
		web.HTMLError(w, "Not found", http.StatusNotFound)
		return
	}

	p := &build.ShowPage{
		Page: &template.Page{
			URI: r.URL.Path,
		},
		Build: b,
	}

	if filepath.Base(r.URL.Path) == "manifest" {
		mp := &build.ShowManifestPage{
			ShowPage: p,
		}

		d := template.NewDashboard(mp, r.URL.Path)

		web.HTML(w, template.Render(d), http.StatusOK)
		return
	}

	if filepath.Base(r.URL.Path) == "raw" {
		web.Text(w, b.Manifest, http.StatusOK)
		return
	}

	d := template.NewDashboard(p, r.URL.Path)

	web.HTML(w, template.Render(d), http.StatusOK)
}
