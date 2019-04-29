package model

import (
	"database/sql"

	"github.com/andrewpillar/thrall/errors"

	"github.com/lib/pq"
)

type Object struct {
	model

	UserID    int64        `db:"user_id"`
	Name      string       `db:"name"`
	Filename  string       `db:"filename"`
	Type      string       `db:"type"`
	Size      int64        `db:"size"`
	MD5       []byte       `db:"md5"`
	SHA256    []byte       `db:"sha256"`
	DeletedAt *pq.NullTime `db:"deleted_at"`

	User *User
}

type ObjectStore struct {
	*Store

	user *User
}

func (os ObjectStore) New() *Object {
	o := &Object{
		model: model{
			DB: os.DB,
		},
	}

	if os.user != nil {
		o.UserID = os.user.ID
		o.User = os.user
	}

	return o
}

func (os ObjectStore) Find(id int64) (*Object, error) {
	o := &Object{
		model: model{
			DB: os.DB,
		},
	}

	query := "SELECT * FROM objects WHERE id = $1"
	args := []interface{}{id}

	if os.user != nil {
		query += " AND user_id = $2"
		args = append(args, os.user.ID)

		o.User = os.user
	}

	err := os.Get(o, query, args...)

	if err == sql.ErrNoRows {
		err = nil
	}

	return o, errors.Err(err)
}

func (os ObjectStore) FindByName(name string) (*Object, error) {
	o := &Object{
		model: model{
			DB: os.DB,
		},
	}

	query := "SELECT * FROM objects WHERE name = $1"
	args := []interface{}{name}

	if os.user != nil {
		query += " AND user_id = $2"
		args = append(args, os.user.ID)

		o.User = os.user
	}

	err := os.Get(o, query, args...)

	if err == sql.ErrNoRows {
		err = nil
	}

	return o, errors.Err(err)
}

func (o *Object) BuildObjectStore() BuildObjectStore {
	return BuildObjectStore{
		Store: &Store{
			DB: o.DB,
		},
		object: o,
	}
}

func (o *Object) Create() error {
	stmt, err := o.Prepare(`
		INSERT INTO objects (user_id, name, filename, type, size, md5, sha256)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`)

	if err != nil {
		return errors.Err(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(o.UserID, o.Name, o.Filename, o.Type, o.Size, o.MD5, o.SHA256)

	return errors.Err(row.Scan(&o.ID, &o.CreatedAt, &o.UpdatedAt))
}

func (o *Object) IsZero() bool {
	return o.ID == 0 &&
           o.UserID == 0 &&
           o.Name == "" &&
           o.Filename == "" &&
           o.Type == "" &&
           o.Size == 0 &&
           len(o.MD5) == 0 &&
           len(o.SHA256) == 0 &&
           o.CreatedAt == nil &&
           o.UpdatedAt == nil &&
           o.DeletedAt == nil
}
