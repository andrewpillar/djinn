package handler

import (
	"net/http"
	"net/url"

	"djinn-ci.com/build"
	"djinn-ci.com/cron"
	"djinn-ci.com/database"
	"djinn-ci.com/errors"
	"djinn-ci.com/namespace"
	"djinn-ci.com/user"
	"djinn-ci.com/web"

	"github.com/andrewpillar/query"
	"github.com/andrewpillar/webutil"
)

// Cron is the base handler that provides shared logic for the UI and API
// handlers for cron job creation, and management.
type Cron struct {
	web.Handler

	loaders *database.Loaders
}

func New(h web.Handler) Cron {
	loaders := database.NewLoaders()
	loaders.Put("user", h.Users)
	loaders.Put("author", h.Users)
	loaders.Put("namespace", namespace.NewStore(h.DB))
	loaders.Put("build_tag", build.NewTagStore(h.DB))
	loaders.Put("build_trigger", build.NewTriggerStore(h.DB))

	return Cron{
		Handler: h,
		loaders: loaders,
	}
}

// IndexWithRelations retrieves a slice of *cron.Cron models for the user in
// the given request context. All of the relations for each cron will be loaded
// into each model we have. If any of the crons have a bound namespace, then
// the namespace's user will be loaded too. A database.Paginator will also be
// returned if there are multiple pages of cron jobs.
func (h Cron) IndexWithRelations(s *cron.Store, vals url.Values) ([]*cron.Cron, database.Paginator, error) {
	cc, paginator, err := s.Index(vals)

	if err != nil {
		return nil, paginator, errors.Err(err)
	}

	if err := cron.LoadRelations(h.loaders, cc...); err != nil {
		return nil, paginator, errors.Err(err)
	}

	nn := make([]database.Model, 0, len(cc))

	for _, c := range cc {
		if c.Namespace != nil {
			nn = append(nn, c.Namespace)
		}
	}

	err = h.Users.Load("id", database.MapKey("user_id", nn), database.Bind("user_id", "id", nn...))
	return cc, paginator, errors.Err(err)
}

// StoreModel unmarshals the request's data into a cron job, validates it and
// stores it in the database. Upon success this will return the newly created
// cron. This also returns the form for creating a cron job.
func (h Cron) StoreModel(r *http.Request) (*cron.Cron, cron.Form, error) {
	var f cron.Form

	ctx := r.Context()

	u, ok := user.FromContext(ctx)

	if !ok {
		return nil, f, errors.New("no user in request context")
	}

	crons := cron.NewStore(h.DB, u)

	f.Crons = crons

	if err := webutil.UnmarshalAndValidate(&f, r); err != nil {
		return nil, f, errors.Err(err)
	}

	c, err := crons.Create(u.ID, f.Name, f.Schedule, f.Manifest)

	if err != nil {
		return nil, f, errors.Err(err)
	}

	c.Author, err = user.NewStore(h.DB).Get(query.Where("id", "=", query.Arg(c.AuthorID)))

	if err != nil {
		return nil, f, errors.Err(err)
	}

	h.Queues.Produce(ctx, "events", &cron.Event{
		Cron:   c,
		Action: "created",
	})
	return c, f, nil
}

// ShowWithRelations retrieves the *cron.Cron model from the context of the
// given request. All of the relations for the cron will be loaded into the
// model we have. If the cron has a namespace bound to it, then the
// namespace's user will be loaded to the namespace.
func (h Cron) ShowWithRelations(r *http.Request) (*cron.Cron, error) {
	var err error

	c, ok := cron.FromContext(r.Context())

	if !ok {
		return nil, errors.New("no cron in request context")
	}

	if err := cron.LoadRelations(h.loaders, c); err != nil {
		return nil, errors.Err(err)
	}

	if c.Namespace != nil {
		err = h.Users.Load(
			"id", []interface{}{c.Namespace.Values()["user_id"]}, database.Bind("user_id", "id", c.Namespace),
		)
	}
	return c, errors.Err(err)
}

// UpdateModel unmarshals the request's data into a cron job, validates it and
// updates the existing cron job in the database with the content in the form.
// Upon success the updated cron job is returned. This also returns the form
// for modifying a cron job.
func (h Cron) UpdateModel(r *http.Request) (*cron.Cron, *cron.Form, error) {
	ctx := r.Context()
	f := &cron.Form{}

	u, ok := user.FromContext(ctx)

	if !ok {
		return nil, f, errors.New("no user in request context")
	}

	c, ok := cron.FromContext(ctx)

	if !ok {
		return nil, f, errors.New("no cron in request context")
	}

	crons := cron.NewStore(h.DB, u)

	f.Cron = c
	f.Crons = crons

	if err := webutil.UnmarshalAndValidate(f, r); err != nil {
		return nil, f, errors.Err(err)
	}

	if err := crons.Update(c.ID, f.Name, f.Schedule, f.Manifest); err != nil {
		return nil, f, errors.Err(err)
	}

	c.Name = f.Name
	c.Schedule = f.Schedule
	c.Manifest = f.Manifest

	var err error

	c.Author, err = user.NewStore(h.DB).Get(query.Where("id", "=", query.Arg(c.AuthorID)))

	if err != nil {
		return nil, f, errors.Err(err)
	}

	h.Queues.Produce(ctx, "events", &cron.Event{
		Cron:   c,
		Action: "updated",
	})
	return c, f, nil
}

// DeleteModel removes the cron in the given request context from the database.
func (h Cron) DeleteModel(r *http.Request) error {
	ctx := r.Context()

	c, ok := cron.FromContext(ctx)

	if !ok {
		return errors.New("no cron in request context")
	}

	if err := cron.NewStore(h.DB).Delete(c.ID); err != nil {
		return errors.Err(err)
	}

	var err error

	c.Author, err = user.NewStore(h.DB).Get(query.Where("id", "=", query.Arg(c.AuthorID)))

	if err != nil {
		return errors.Err(err)
	}

	h.Queues.Produce(ctx, "events", &cron.Event{
		Cron:   c,
		Action: "deleted",
	})
	return nil
}
