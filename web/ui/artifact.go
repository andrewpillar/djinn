package ui

import (
	"net/http"
	"os"
	"strconv"

	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/filestore"
	"github.com/andrewpillar/thrall/log"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
	"github.com/andrewpillar/thrall/template/build"
	"github.com/andrewpillar/thrall/web"

	"github.com/gorilla/mux"
)

type Artifact struct {
	web.Handler

	Artifacts model.ArtifactStore
	FileStore filestore.FileStore
}

func (h Artifact) Index(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	buildId, _ := strconv.ParseInt(vars["build"], 10, 64)

	b, err := u.BuildStore().Find(buildId)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	search := r.URL.Query().Get("search")

	aa, err := b.ArtifactStore().Index(model.Search("name", search))

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	p := &build.ArtifactIndexPage{
		ShowPage: build.ShowPage{
			Page: template.Page{
				URI: r.URL.Path,
			},
			Build:  b,
		},
		Search:    search,
		Artifacts: aa,
	}

	d := template.NewDashboard(p, r.URL.Path, h.Alert(w, r))

	web.HTML(w, template.Render(d), http.StatusOK)
}

func (h Artifact) Show(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	buildId, _ := strconv.ParseInt(vars["build"], 10, 64)

	artifactId, err := strconv.ParseInt(vars["artifact"], 10, 64)

	if err != nil {
		web.HTMLError(w, "Not found", http.StatusNotFound)
		return
	}

	b, err := u.BuildStore().Find(buildId)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	a, err := b.ArtifactStore().Find(artifactId)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if a.IsZero() || a.Name != vars["name"] {
		web.HTMLError(w, "Not found", http.StatusNotFound)
		return
	}

	f, err := h.FileStore.Open(a.Hash)

	if err != nil {
		if os.IsNotExist(errors.Cause(err)) {
			web.HTMLError(w, "Not found", http.StatusNotFound)
			return
		}

		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	defer f.Close()

	http.ServeContent(w, r, a.Name, *a.UpdatedAt, f)
}
