package model

import (
	"database/sql"
	"strings"
	"time"

	"github.com/andrewpillar/thrall/errors"

	"github.com/jmoiron/sqlx"

	"github.com/lib/pq"
)

type User struct {
	model

	Email     string       `db:"email"`
	Username  string       `db:"username"`
	Password  []byte       `db:"password"`
	DeletedAt *pq.NullTime `db:"deleted_at"`
}

type UserStore struct {
	*Store
}

func NewUserStore(s *Store) *UserStore {
	return &UserStore{
		Store: s,
	}
}

func (us UserStore) New() *User {
	u := &User{
		model: model{
			DB: us.DB,
		},
	}

	return u
}

func (us UserStore) In(ids ...int64) ([]*User, error) {
	uu := make([]*User, 0)

	if len(ids) == 0 {
		return uu, nil
	}

	query, args, err := sqlx.In("SELECT * FROM users WHERE id IN (?)", ids)

	if err != nil {
		return uu, errors.Err(err)
	}

	err = us.Select(&uu, us.Rebind(query), args...)

	if err == sql.ErrNoRows {
		err = nil
	}

	for _, u := range uu {
		u.DB = us.DB
		u.Email = strings.TrimSpace(u.Email)
		u.Username = strings.TrimSpace(u.Username)
	}

	return uu, errors.Err(err)
}

func (us UserStore) Find(id int64) (*User, error) {
	u := &User{
		model: model{
			DB: us.DB,
		},
	}

	err := us.Get(u, "SELECT * FROM users WHERE id = $1", id)

	if err == sql.ErrNoRows {
		err = nil
	}

	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.TrimSpace(u.Username)

	return u, errors.Err(err)
}

func (us UserStore) FindByEmail(email string) (*User, error) {
	u := &User{
		model: model{
			DB: us.DB,
		},
	}

	err := us.Get(u, "SELECT * FROM users WHERE email = $1", email)

	if err == sql.ErrNoRows {
		err = nil
	}

	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.TrimSpace(u.Username)

	return u, errors.Err(err)
}

func (us UserStore) FindByUsername(username string) (*User, error) {
	u := &User{
		model: model{
			DB: us.DB,
		},
	}

	err := us.Get(u, "SELECT * FROM users WHERE username = $1", username)

	if err == sql.ErrNoRows {
		err = nil
	}

	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.TrimSpace(u.Username)

	return u, errors.Err(err)
}

func (us UserStore) FindByHandle(handle string) (*User, error) {
	u := &User{
		model: model{
			DB: us.DB,
		},
	}

	err := us.Get(u, "SELECT * FROM users WHERE username = $1 OR email = $2", handle, handle)

	if err == sql.ErrNoRows {
		err = nil
	}

	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.TrimSpace(u.Username)

	return u, errors.Err(err)
}

func (u *User) BuildStore() BuildStore {
	return BuildStore{
		Store: &Store{
			DB: u.DB,
		},
		user: u,
	}
}

func (u *User) NamespaceStore() NamespaceStore {
	return NamespaceStore{
		Store: &Store{
			DB: u.DB,
		},
		user: u,
	}
}

func (u *User) ObjectStore() ObjectStore {
	return ObjectStore{
		Store: &Store{
			DB: u.DB,
		},
		user: u,
	}
}

func (u *User) Create() error {
	stmt, err := u.Prepare(`
		INSERT INTO users (email, username, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`)

	if err != nil {
		return errors.Err(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(u.Email, u.Username, u.Password)

	return errors.Err(row.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt))
}

func (u *User) Update() error {
	stmt, err := u.Prepare(`
		UPDATE users
		SET email = $1, username = $2, password = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
	`)

	if err != nil {
		return errors.Err(err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(u.Email, u.Username, u.Password, u.ID)

	return errors.Err(row.Scan(&u.UpdatedAt))
}

func (u *User) Destroy() error {
	u.DeletedAt = &pq.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	stmt, err := u.Prepare("UPDATE users SET deleted_at = $1 WHERE id = $2")

	if err != nil {
		return errors.Err(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.DeletedAt)

	return errors.Err(err)
}

func (u *User) Deleted() bool {
	return u.DeletedAt != nil && u.DeletedAt.Valid
}

func (u *User) IsZero() bool {
	return u.ID == 0 &&
           u.Email == "" &&
           u.Username == "" &&
           len(u.Password) == 0 &&
           u.CreatedAt == nil &&
           u.UpdatedAt == nil &&
           u.DeletedAt == nil
}
