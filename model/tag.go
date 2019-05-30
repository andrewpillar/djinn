package model

import (
	"database/sql"

	"github.com/andrewpillar/thrall/errors"

	"github.com/jmoiron/sqlx"
)

type Tag struct {
	model

	UserID  int64  `db:"user_id"`
	BuildID int64  `db:"build_id"`
	Name    string `db:"name"`

	User  *User
	Build *Build
}

type TagStore struct {
	*sqlx.DB

	User  *User
	Build *Build
}

func (t *Tag) Create() error {
	stmt, err := t.Prepare(`
		INSERT INTO tags (user_id, build_id, name)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`)

	if err != nil {
		return errors.Err(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(t.UserID, t.BuildID, t.Name)

	return errors.Err(row.Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt))
}

func (ts TagStore) All() ([]*Tag, error) {
	tt := make([]*Tag, 0)

	query := "SELECT * FROM tags"
	args := []interface{}{}

	if ts.Build != nil {
		query += " WHERE build_id = $1"
		args = append(args, ts.Build.ID)
	}

	query += " ORDER BY name ASC"

	err := ts.Select(&tt, query, args...)

	if err == sql.ErrNoRows {
		err = nil
	}

	for _, t := range tt {
		t.DB = ts.DB

		if ts.Build != nil {
			t.Build = ts.Build
		}
	}

	return tt, errors.Err(err)
}

func (ts TagStore) InBuildID(ids ...int64) ([]*Tag, error) {
	tt := make([]*Tag, 0)

	if len(ids) == 0 {
		return tt, nil
	}

	query, args, err := sqlx.In("SELECT * FROM tags WHERE build_id IN (?)", ids)

	if err != nil {
		return tt, errors.Err(err)
	}

	err = ts.Select(&tt, ts.Rebind(query), args...)

	if err == sql.ErrNoRows {
		err = nil
	}

	for _, t := range tt {
		t.DB = ts.DB
	}

	return tt, errors.Err(err)
}

func (ts TagStore) New() *Tag {
	t := &Tag{
		model: model{
			DB: ts.DB,
		},
		User:  ts.User,
		Build: ts.Build,
	}

	if ts.Build != nil {
		t.BuildID = ts.Build.ID
	}

	if ts.User != nil {
		t.UserID = ts.User.ID
	}

	return t
}
