package build

import (
	"github.com/andrewpillar/djinn/block"
	"github.com/andrewpillar/djinn/database"
	"github.com/andrewpillar/djinn/errors"
	"github.com/andrewpillar/djinn/log"
	"github.com/andrewpillar/djinn/user"

	"github.com/andrewpillar/query"

	"github.com/jmoiron/sqlx"
)

type curationRecord struct {
	artifact int64
	hash     string
}

// Curator is used for removing old build artifacts whose total size exceed
// the configured limit.
type Curator struct {
	limit     int64
	artifacts block.Store
	store     *ArtifactStore
	users     *user.Store
}

var ErrCuration = errors.New("artifact curation failed")

// NewCurator creates a new curator for cleaning up old artifacts from the
// given block store.
func NewCurator(db *sqlx.DB, artifacts block.Store, limit int64) Curator {
	return Curator{
		limit:     limit,
		artifacts: artifacts,
		store:     NewArtifactStore(db),
		users:     user.NewStore(db),
	}
}

// Invoke will remove any artifacts whose total size exceeds the configured
// limit. This will only do it for users who have "cleanup" enabled on their
// account.
func (c *Curator) Invoke(log *log.Logger) error {
	uu, err := c.users.All(query.Where("cleanup", "=", query.Arg(true)))

	if err != nil {
		return errors.Err(err)
	}

	mm := database.ModelSlice(len(uu), user.Model(uu))

	aa, err := c.store.All(
		query.Where("size", ">", query.Arg(0)),
		query.Where("user_id", "IN", database.List(database.MapKey("id", mm)...)),
		query.Where("deleted_at", "IS NOT", query.Lit("NULL")),
		query.OrderAsc("created_at"),
	)

	if err != nil {
		return errors.Err(err)
	}

	sums := make(map[int64]int64)
	curated := make(map[int64][]curationRecord)
	deleted := make([]int64, 0, len(aa))

	for _, a := range aa {
		sum := sums[a.UserID]
		sum += a.Size.Int64

		if sum >= c.limit {
			curated[a.UserID] = append(curated[a.UserID], curationRecord{
				artifact: a.ID,
				hash:     a.Hash,
			})
		}
	}

	for _, records := range curated {
		for _, r := range records {
			log.Debug.Println("removing artifact", r.hash)

			if err := c.artifacts.Remove(r.hash); err != nil {
				err = ErrCuration
				log.Error.Println("failed to remove artifact", r.hash, err)
				continue
			}
			deleted = append(deleted, r.artifact)
		}
	}

	if err := c.store.Deleted(deleted...); err != nil {
		return errors.Err(err)
	}
	return err
}
