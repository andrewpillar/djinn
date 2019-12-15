package core

import (
	"net/http"
	"strconv"

	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/web"

	"github.com/andrewpillar/query"

	"github.com/gorilla/mux"
)

type Job struct {
	web.Handler
}

func (h Job) Build(r *http.Request) *model.Build {
	val := r.Context().Value("build")

	b, _ := val.(*model.Build)

	return b
}

func (h Job) Show(r *http.Request) (*model.Job, error) {
	b := h.Build(r)

	if err := b.LoadNamespace(); err != nil {
		return &model.Job{}, errors.Err(err)
	}

	if err := b.Namespace.LoadUser(); err != nil {
		return &model.Job{}, errors.Err(err)
	}

	if err := b.LoadTags(); err != nil {
		return &model.Job{}, errors.Err(err)
	}

	if err := b.LoadTrigger(); err != nil {
		return &model.Job{}, errors.Err(err)
	}

	vars := mux.Vars(r)

	id, _ := strconv.ParseInt(vars["job"], 10, 64)

	j, err := b.JobStore().Get(query.Where("id", "=", id))

	if err != nil {
		return j, errors.Err(err)
	}

	if err := j.LoadStage(); err != nil {
		return j, errors.Err(err)
	}

	err = j.LoadArtifacts()

	return j, errors.Err(err)
}
