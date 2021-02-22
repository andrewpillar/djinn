package cron

import (
	"context"
	"fmt"

	"github.com/andrewpillar/djinn/build"
	"github.com/andrewpillar/djinn/database"
	"github.com/andrewpillar/djinn/errors"
	"github.com/andrewpillar/djinn/namespace"
	"github.com/andrewpillar/djinn/user"

	"github.com/andrewpillar/query"

	"github.com/jmoiron/sqlx"

	"github.com/mcmathja/curlyq"
)

// Batcher provides a way of retrieving batches of cron jobs that are ready
// to be executed.
type Batcher struct {
	err       error
	paginator database.Paginator
	store     *Store
	builds    *build.Store
	batch     []*Cron
}

// NewBatcher returns a new Batcher using the given Store to retrieve cron jobs
// from, and setting the size of each batch to the given limit.
func NewBatcher(db *sqlx.DB, limit int64) *Batcher {
	return &Batcher{
		store:  NewStore(db),
		builds: build.NewStore(db),
		paginator: database.Paginator{
			Page:  1,
			Limit: limit,
		},
	}
}

// Next will load in the next batch of cron jobs to be executed
// (WHERE NOW() >= next_run). This will return false if it reaches the end of
// the batches in the table. If the end of the table is reached, or if an error
// happens then false is returned.
func (b *Batcher) Next() bool {
	paginator, err := b.store.Paginate(b.paginator.Page, query.Where("NOW()", ">=", query.Lit("next_run")))

	if err != nil {
		b.err = errors.Err(err)
		return false
	}

	cc, err := b.store.All(
		query.Where("NOW()", ">=", query.Lit("next_run")),
		query.Limit(paginator.Limit),
		query.Offset(paginator.Offset),
	)

	if err != nil {
		b.err = errors.Err(err)
		return false
	}

	if len(cc) == 0 {
		return false
	}

	mm := database.ModelSlice(len(cc), Model(cc))

	uu, err := user.NewStore(b.store.DB).All(query.Where("id", "IN", query.List(database.MapKey("user_id", mm)...)))

	if err != nil {
		b.err = errors.Err(err)
		return false
	}

	users := make(map[int64]*user.User)

	for _, u := range uu {
		users[u.ID] = u
	}

	for _, c := range cc {
		c.User = users[c.UserID]
	}

	b.paginator = paginator
	b.batch = cc

	return paginator.Page == paginator.Pages[len(paginator.Pages)-1]
}

func (b *Batcher) Batch() []*Cron { return b.batch }

// Err returns the current error, if any, that occurred when loading a batch.
func (b *Batcher) Err() error { return b.err }

// Invoke will submit a build for each job in the current batch.
func (b *Batcher) Invoke(ctx context.Context, producers map[string]*curlyq.Producer) (int, error) {
	namespaces := namespace.NewStore(b.store.DB)
	users := user.NewStore(b.store.DB)

	errs := make([]error, 0, len(b.batch))
	n := 0

	for _, c := range b.batch {
		if c.NamespaceID.Valid {
			n, err := namespaces.Get(query.Where("id", "=", query.Arg(c.NamespaceID.Int64)))

			if err != nil {
				errs = append(errs, fmt.Errorf("failed to get namespace: %v", errors.Err(err)))
				continue
			}

			u, err := users.Get(query.Where("user_id", "=", query.Arg(n.UserID)))

			if err != nil {
				errs = append(errs, fmt.Errorf("failed to get namespace owner: %v", errors.Err(err)))
				continue
			}
			c.Manifest.Namespace = n.Path + "@" + u.Username
		}

		bld, err := b.store.Invoke(c)

		if err != nil {
			errs = append(errs, fmt.Errorf("failed to invoke cron: %v", errors.Err(err)))
			continue
		}

		queue, ok := producers[bld.Manifest.Driver["type"]]

		if !ok {
			errs = append(errs, fmt.Errorf("invalid build driver: %v", bld.Manifest.Driver["type"]))
			continue
		}

		if err := b.builds.Submit(ctx, queue, "djinn-scheduler", bld); err != nil {
			errs = append(errs, fmt.Errorf("failed to submit build: %v", errors.Err(err)))
			continue
		}
		n++
	}
	return n, errors.Slice(errs)
}
