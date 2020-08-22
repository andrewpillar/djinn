// Package provider providers the database.Model implementation for the Provider
// entity, and implementations for the oauth2.Provider interface for the
// different Git providers that can be used to authenticate against.
package provider

import (
	"encoding/json"
	"database/sql"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/andrewpillar/thrall/database"
	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/user"

	"github.com/andrewpillar/query"

	"github.com/jmoiron/sqlx"
)

// Provider is the type that represents an external Git hosting provider that
// has been connected to a user's account.
type Provider struct {
	ID             int64         `db:"id"`
	UserID         int64         `db:"user_id"`
	ProviderUserID sql.NullInt64 `db:"provider_user_id"`
	Name           string        `db:"name"`
	AccessToken    []byte        `db:"access_token"`
	RefreshToken   []byte        `db:"refresh_token"`
	Connected      bool          `db:"connected"`
	ExpiresAt      time.Time     `db:"expires_at"`
	AuthURL        string        `db:"-"`

	User *user.User `db:"-"`
}

// Store is the type for creating and modifying Provider models in the database.
type Store struct {
	database.Store

	// User is the bound user.User model. If not nil this will bind the
	// user.User model to any Provider models that are created. If not nil this
	// will append a WHERE clause on the user_id column for all SELECT queries
	// performed.
	User *user.User
}

var (
	_ database.Model  = (*Provider)(nil)
	_ database.Binder = (*Provider)(nil)

	table = "providers"
)

// NewStore returns a new Store for querying the providers table. Each database
// passed to this function will be bound to the returned Store.
func NewStore(db *sqlx.DB, mm ...database.Model) *Store {
	s := &Store{
		Store: database.Store{DB: db},
	}
	s.Bind(mm...)
	return s
}

// Model is called along with database.ModelSlice to convert the given slice of
// Provider models to a slice of database.Model interfaces.
func Model(pp []*Provider) func(int) database.Model {
	return func(i int) database.Model {
		return pp[i]
	}
}

// Select returns a query that selects the given column from the providers
// table, with each given query.Option applied to the returned query.
func Select(col string, opts ...query.Option) query.Query {
	return query.Select(append([]query.Option{
		query.Columns(col),
		query.From(table),
	}, opts...)...)
}

// Bind implements the database.Binder interface. This will only bind the model
// if it is a pointer to a user.User model.
func (p *Provider) Bind(mm ...database.Model) {
	for _, m := range mm {
		switch m.(type) {
		case *user.User:
			p.User = m.(*user.User)
		}
	}
}

// SetPrimary implements the database.Model interface.
func (p *Provider) SetPrimary(id int64) { p.ID = id }

// Primary implements the database.Model interface.
func (p *Provider) Primary() (string, int64) { return "id", p.ID }

// Endpoint implements the database.Model interface. This is a stub method and
// returns an empty string.
func (*Provider) Endpoint(_ ...string) string { return "" }

// IsZero implements the database.Model interface.
func (p *Provider) IsZero() bool {
	return p == nil || p.ID == 0 &&
		p.UserID == 0 &&
		!p.ProviderUserID.Valid &&
		p.Name == "" &&
		len(p.AccessToken) == 0 &&
		len(p.RefreshToken) == 0 &&
		!p.Connected &&
		p.ExpiresAt == time.Time{}
}

// JSON implements the database.Model interface. This is a stub method and
// returns an empty map.
func (*Provider) JSON(_ string) map[string]interface{} { return map[string]interface{}{} }

// Values implements the database.Model interface. This will return a map with
// the following values, user_id, provider_user_id, name, access_token,
// refresh_token, connected, and expires_at.
func (p *Provider) Values() map[string]interface{} {
	return map[string]interface{}{
		"user_id":          p.UserID,
		"provider_user_id": p.ProviderUserID,
		"name":             p.Name,
		"access_token":     p.AccessToken,
		"refresh_token":    p.RefreshToken,
		"connected":        p.Connected,
		"expires_at":       p.ExpiresAt,
	}
}

// ToggleRepo will with add or remove a hook for the given repository hosted on
// the current provider. This will either set/unset the HookID field on the
// given Repo struct, and will toggle the Enabled field depending on whether a
// hook was added or removed.
func (p *Provider) ToggleRepo(clients *Registry, r *Repo) error {
	_, cli, err := clients.Get(p.Name)

	if err != nil {
		return errors.Err(err)
	}

	tok, err := cli.block.Decrypt(p.AccessToken)

	if err != nil {
		return errors.Err(err)
	}

	switch p.Name {
	case "github":
		err = toggleGitHubRepo(cli, string(tok), r)
	case "gitlab":
		err = toggleGitLabRepo(cli, string(tok), r)
	}
	return errors.Err(err)
}

// Repos get's the repositories from the current provider's API endpoint. The
// given crypto.Block is used to decrypt the access token that is used to
// authenticate against the API. The given page is used to get the repositories
// on that given page.
func (p *Provider) Repos(clients *Registry, page int64) ([]*Repo, database.Paginator, error) {
	paginator := database.Paginator{}

	_, cli, err := clients.Get(p.Name)

	if err != nil {
		return nil, paginator, errors.Err(err)
	}

	tok, err := cli.block.Decrypt(p.AccessToken)

	if err != nil {
		return nil, paginator, errors.Err(err)
	}

	spage := strconv.FormatInt(page, 10)

	var (
		endpoint  string
		unmarshal func(io.Reader, int64, int64) []*Repo
	)

	switch p.Name {
	case "github":
		endpoint = "/user/repos?sort=updated&part=" + spage
		unmarshal = unmarshalGitHubRepos
	case "gitlab":
		resp0, err := cli.Get(string(tok), "/user")

		if err != nil {
			return nil, paginator, errors.Err(err)
		}

		defer resp0.Body.Close()

		if resp0.StatusCode != http.StatusOK {
			return nil, paginator, errors.New("unexpected http status: " + resp0.Status)
		}

		u := struct {
			ID int64
		}{}

		json.NewDecoder(resp0.Body).Decode(&u)

		endpoint = "/users/" + strconv.FormatInt(u.ID, 10) + "/projects?simple=true&order_by=updated_at&page=" + spage
		unmarshal = unmarshalGitLabRepos
	}

	resp, err := cli.Get(string(tok), endpoint)

	if err != nil {
		return nil, paginator, errors.Err(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, paginator, errors.New("unexpected http status: " + resp.Status)
	}

	rr := unmarshal(resp.Body, p.UserID, p.ID)

	for i := range rr {
		rr[i].Provider = p
	}

	next, prev := getNextAndPrev(resp.Header.Get("Link"))

	paginator.Next = next
	paginator.Prev = prev
	paginator.Pages = []int64{next, prev}

	return rr, paginator, errors.Err(err)
}

// New returns a new Provider binding any non-nil models to it from the current
// Store.
func (s *Store) New() *Provider {
	p := &Provider{
		User: s.User,
	}

	if s.User != nil {
		p.UserID = s.User.ID
	}
	return p
}

// Bind implements the database.Binder interface. This will only bind the model
// if it is a pointer to a user.User model.
func (s *Store) Bind(mm ...database.Model) {
	for _, m := range mm {
		switch m.(type) {
		case *user.User:
			s.User = m.(*user.User)
		}
	}
}

// Create creates a new provider of the given name, and with the given access
// and refresh tokens. The given userId parameter should be the ID of the user's
// account from the provider we're connecting to.
func (s *Store) Create(userId int64, name string, access, refresh []byte, connected bool) (*Provider, error) {
	p := s.New()
	p.ProviderUserID = sql.NullInt64{
		Int64: userId,
		Valid: userId > 0,
	}
	p.Name = name
	p.AccessToken = access
	p.RefreshToken = refresh
	p.Connected = connected

	err := s.Store.Create(table, p)
	return p, errors.Err(err)
}

// Update updates the provider in the database for the given id. This will set
// the userId, name, tokens, and connected status to the given values.
func (s *Store) Update(id, userId int64, name string, access, refresh []byte, connected bool) error {
	q := query.Update(
		query.Table(table),
		query.Set("provider_user_id", sql.NullInt64{
			Int64: userId,
			Valid: userId > 0,
		}),
		query.Set("name", name),
		query.Set("access_token", access),
		query.Set("refresh_token", refresh),
		query.Set("connected", connected),
		query.Where("id", "=", id),
	)

	_, err := s.DB.Exec(q.Build(), q.Args()...)
	return errors.Err(err)
}

// Delete deletes the given Provider models from the providers table.
func (s *Store) Delete(pp ...*Provider) error {
	mm := database.ModelSlice(len(pp), Model(pp))
	return errors.Err(s.Store.Delete(table, mm...))
}

// Get returns a single Provider database, applying each query.Option that is
// given. The database.Where option is applied to the *user.User bound database.
func (s *Store) Get(opts ...query.Option) (*Provider, error) {
	p := &Provider{
		User: s.User,
	}

	opts = append([]query.Option{
		database.Where(s.User, "user_id"),
	}, opts...)

	err := s.Store.Get(p, table, opts...)

	if err == sql.ErrNoRows {
		err = nil
	}
	return p, errors.Err(err)
}

// All returns a slice of Provider models, applying each query.Option that is
// given. The database.Where option is applied to the *user.User bound database.
func (s *Store) All(opts ...query.Option) ([]*Provider, error) {
	pp := make([]*Provider, 0)

	opts = append([]query.Option{
		database.Where(s.User, "user_id"),
	}, opts...)

	err := s.Store.All(&pp, table, opts...)

	if err == sql.ErrNoRows {
		err = nil
	}
	return pp, nil
}
