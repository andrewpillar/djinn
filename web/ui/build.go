package ui

import (
	"database/sql"
	"fmt"
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

	srv *machinery.Server
}

func NewBuild(h web.Handler) Build {
	return Build{
		Handler: h,
	}
}

func (h Build) Index(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	builds := u.BuildStore()

	var bb []*model.Build

	status := r.URL.Query().Get("status")

	if status != "" {
		bb, err = builds.ByStatus(status)
	} else {
		bb, err = builds.All()
	}

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err := builds.LoadRelations(bb); err != nil {
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

	b := u.BuildStore().New()
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

	manifest, _ := config.DecodeManifest(strings.NewReader(f.Manifest))

	// Create initial setup stage. Will contain the output of driver creation
	// and cloning of source repositories.
	setupStage := b.StageStore().New()
	setupStage.Name = fmt.Sprintf("setup - #%v", b.ID)

	if err := setupStage.Create(); err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	createJob := setupStage.JobStore().New()
	createJob.Name = "create driver"

	if err := createJob.Create(); err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	for _, obj := range manifest.Objects {
		o, err := u.ObjectStore().FindByName(obj.Source)

		if err != nil {
			log.Error.Println(errors.Err(err))
			web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		if o.IsZero() {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		bo := o.BuildObjectStore().New()
		bo.BuildID = b.ID
		bo.Source = obj.Source

		if err := bo.Create(); err != nil {
			log.Error.Println(errors.Err(err))
			web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	for i, src := range manifest.Sources {
		name := fmt.Sprintf("clone.%d", i + 1)

		commands := []string{
			"git clone " + src.URL + " " + src.Dir,
			"cd " + src.Dir,
			"git checkout -q " + src.Ref,
		}

		if src.Dir != "" {
			commands = append([]string{"mkdir -p " + src.Dir}, commands...)
		}

		j := setupStage.JobStore().New()
		j.Name = name
		j.Commands = strings.Join(commands, "\n")

		if err := j.Create(); err != nil {
			log.Error.Println(errors.Err(err))
			web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	for _, name := range manifest.Stages {
		canFail := false

		for _, allowed := range manifest.AllowFailures {
			if name == allowed {
				canFail = true
				break
			}
		}

		s := b.StageStore().New()
		s.Name = name
		s.CanFail = canFail

		if err := s.Create(); err != nil {
			log.Error.Println(errors.Err(err))
			web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		jobId := 1

		for _, mj := range manifest.Jobs {
			if mj.Stage != s.Name {
				continue
			}

			if mj.Name == "" {
				mj.Name = fmt.Sprintf("%s.%d", mj.Stage, jobId)
			}

			j := s.JobStore().New()
			j.Name = mj.Name
			j.Commands = strings.Join(mj.Commands, "\n")

			if err := j.Create(); err != nil {
				log.Error.Println(errors.Err(err))
				web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
				return
			}

			for _, ma := range mj.Artifacts {
				a := j.ArtifactStore().New()
				a.Source = ma.Source
				a.Name = ma.Destination

				if err := a.Create(); err != nil {
					log.Error.Println(errors.Err(err))
					web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
					return
				}
			}
		}
	}

	tt := make([]*model.Tag, len(f.Tags), len(f.Tags))

	for i, name := range f.Tags {
		if name == "" {
			continue
		}

		tt[i] = b.TagStore().New()
		tt[i].Name = strings.TrimSpace(name)

		if err := tt[i].Create(); err != nil {
			log.Error.Println(errors.Err(err))
			web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
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

	b, err := u.BuildStore().Find(id)

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

	if err := b.LoadRelations(); err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if err := b.StageStore().LoadJobs(b.Stages); err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	d := template.NewDashboard(p, r.URL.Path)

	web.HTML(w, template.Render(d), http.StatusOK)
}