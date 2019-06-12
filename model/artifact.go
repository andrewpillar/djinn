package model

import (
	"database/sql"

	"github.com/andrewpillar/thrall/errors"

	"github.com/jmoiron/sqlx"
)

type Artifact struct {
	model

	BuildID int64          `db:"build_id"`
	JobID   int64          `db:"job_id"`
	Hash    string         `db:"hash"`
	Source  string         `db:"source"`
	Name    string         `db:"name"`
	Size    sql.NullInt64  `db:"size"`
	Type    sql.NullString `db:"type"`
	MD5     []byte         `db:"md5"`
	SHA256  []byte         `db:"sha256"`

	Build *Build
	Job   *Job
}

type ArtifactStore struct {
	*sqlx.DB

	Build *Build
	Job   *Job
}

func (a *Artifact) Create() error {
	stmt, err := a.Prepare(`
		INSERT INTO artifacts (build_id, job_id, hash, source, name, size, type, md5, sha256)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`)

	if err != nil {
		return errors.Err(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(a.BuildID, a.JobID, a.Hash, a.Source, a.Name, a.Size, a.Type, a.MD5, a.SHA256)

	return errors.Err(row.Scan(&a.ID, &a.CreatedAt, &a.UpdatedAt))
}

func (a *Artifact) Update() error {
	stmt, err := a.Prepare(`
		UPDATE artifacts
		SET size = $1, type = $2, md5 = $3, sha256 = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`)

	if err != nil {
		return errors.Err(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(a.Size, a.Type, a.MD5, a.SHA256, a.ID)

	return errors.Err(row.Scan(&a.UpdatedAt))
}

func (as ArtifactStore) All() ([]*Artifact, error) {
	aa := make([]*Artifact, 0)

	query := "SELECT * FROM artifacts"
	args := []interface{}{}

	if as.Build != nil {
		query += " WHERE build_id = $1"
		args = append(args, as.Build.ID)
	}

	if as.Job != nil {
		if as.Build != nil {
			query += " AND job_id = $2"
		} else {
			query += " WHERE job_id = $1"
		}

		args = append(args, as.Job.ID)
	}

	err := as.Select(&aa, query, args...)

	if err == sql.ErrNoRows {
		err = nil
	}

	for _, a := range aa {
		a.DB = as.DB
		a.Build = as.Build
		a.Job = as.Job
	}

	return aa, errors.Err(err)
}

func (as ArtifactStore) Find(id int64) (*Artifact, error) {
	a := &Artifact{
		model: model{
			DB: as.DB,
		},
		Build: as.Build,
		Job:   as.Job,
	}

	query := "SELECT * FROM artifacts WHERE id = $1"
	args := []interface{}{id}

	if as.Build != nil {
		query += " AND build_id = $2"
		args = append(args, as.Build.ID)
	}

	if as.Job != nil {
		if as.Build != nil {
			query += " AND job_id = $3"
		} else {
			query += " AND job_id = $2"
		}

		args = append(args, as.Job.ID)
	}

	err := as.Get(a, query, args...)

	if err == sql.ErrNoRows {
		err = nil
	}

	return a, errors.Err(err)
}

func (as ArtifactStore) FindByHash(hash string) (*Artifact, error) {
	a := &Artifact{
		model: model{
			DB: as.DB,
		},
		Build: as.Build,
		Job:   as.Job,
	}

	query := "SELECT * FROM artifacts WHERE hash = $1"
	args := []interface{}{hash}

	if as.Build != nil {
		query += " AND build_id = $2"
		args = append(args, as.Build.ID)
	}

	if as.Job != nil {
		if as.Build != nil {
			query += " AND job_id = $3"
		} else {
			query += " AND job_id = $2"
		}

		args = append(args, as.Job.ID)
	}

	err := as.Get(a, query, args...)

	if err == sql.ErrNoRows {
		err = nil
	}

	return a, errors.Err(err)
}

func (as ArtifactStore) New() *Artifact {
	a := &Artifact{
		model: model{
			DB: as.DB,
		},
		Build: as.Build,
		Job:   as.Job,
	}

	if as.Build != nil {
		a.BuildID = as.Build.ID
	}

	if as.Job != nil {
		a.JobID = as.Job.ID
	}

	return a
}
