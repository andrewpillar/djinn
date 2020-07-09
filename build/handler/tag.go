package handler

import (
	"net/http"
	"strconv"

	"github.com/andrewpillar/thrall/build"
	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/form"
	"github.com/andrewpillar/thrall/user"
	"github.com/andrewpillar/thrall/web"

	"github.com/gorilla/mux"
)

type Tag struct {
	web.Handler
}

func (h Tag) StoreModel(r *http.Request) ([]*build.Tag, error) {
	ctx := r.Context()

	u, ok := user.FromContext(ctx)

	if !ok {
		return nil, errors.New("failed to get user from context")
	}

	b, ok := build.FromContext(ctx)

	if !ok {
		return nil, errors.New("failed to get build from context")
	}

	f := &build.TagForm{}

	if err := form.Unmarshal(f, r); err != nil {
		return nil, errors.Err(err)
	}

	tt, err := build.NewTagStore(h.DB, b).Create(u.ID, f.Tags...)
	return tt, errors.Err(err)
}

func (h Tag) DeleteModel(r *http.Request) error {
	id, _ := strconv.ParseInt(mux.Vars(r)["tag"], 10, 64)
	return errors.Err(build.NewTagStore(h.DB).Delete(id))
}
