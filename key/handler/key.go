package handler

import (
	"net/http"

	"github.com/andrewpillar/djinn/crypto"
	"github.com/andrewpillar/djinn/database"
	"github.com/andrewpillar/djinn/errors"
	"github.com/andrewpillar/djinn/key"
	"github.com/andrewpillar/djinn/namespace"
	"github.com/andrewpillar/djinn/user"
	"github.com/andrewpillar/djinn/web"

	"github.com/andrewpillar/webutil"
)

// Key is the base handler that provides shared logic for the UI and API
// handlers for key creation, and management.
type Key struct {
	web.Handler

	loaders *database.Loaders
	block   *crypto.Block
}

func New(h web.Handler, block *crypto.Block) Key {
	loaders := database.NewLoaders()
	loaders.Put("author", h.Users)
	loaders.Put("namespace", namespace.NewStore(h.DB))

	return Key{
		Handler: h,
		loaders: loaders,
		block:   block,
	}
}

// IndexWithRelations retrieves a slice of *key.Key models for the user in
// the given request context. All of the relations for each key will be loaded
// into each model we have. If any of the keys have a bound namespace, then
// the namespace's user will be loaded too. A database.Paginator will also be
// returned if there are multiple pages of keys.
func (h Key) IndexWithRelations(r *http.Request) ([]*key.Key, database.Paginator, error) {
	u, ok := user.FromContext(r.Context())

	if !ok {
		return nil, database.Paginator{}, errors.New("user not in request context")
	}

	kk, paginator, err := key.NewStore(h.DB, u).Index(r.URL.Query())

	if err != nil {
		return kk, paginator, errors.Err(err)
	}

	if err := key.LoadRelations(h.loaders, kk...); err != nil {
		return kk, paginator, errors.Err(err)
	}

	nn := make([]database.Model, 0, len(kk))

	for _, k := range kk {
		if k.Namespace != nil {
			nn = append(nn, k.Namespace)
		}
	}

	err = h.Users.Load("id", database.MapKey("user_id", nn), database.Bind("user_id", "id", nn...))
	return kk, paginator, errors.Err(err)
}

// StoreModel unmarshals the request's data into a key, validates it and
// stores it in the database. Upon success this will return the newly created
// key. This also returns the form for creating a key.
func (h Key) StoreModel(r *http.Request) (*key.Key, key.Form, error) {
	f := key.Form{}

	u, ok := user.FromContext(r.Context())

	if !ok {
		return nil, f, errors.New("no user in request context")
	}

	keys := key.NewStoreWithBlock(h.DB, h.block, u)

	f.Resource = namespace.Resource{
		Author:     u,
		Namespaces: namespace.NewStore(h.DB, u),
	}
	f.Keys = keys

	if err := webutil.UnmarshalAndValidate(&f, r); err != nil {
		return nil, f, errors.Err(err)
	}

	k, err := keys.Create(f.Author.ID, f.Name, f.PrivateKey, f.Config)
	return k, f, errors.Err(err)
}

// ShowWithRelations retrieves the *key.Keys model from the context of the
// given request. All of the relations for the key will be loaded into the
// model we have. If the key has a namespace bound to it, then the
// namespace's user will be loaded to the namespace.
func (h Key) ShowWithRelations(r *http.Request) (*key.Key, error) {
	var err error

	k, ok := key.FromContext(r.Context())

	if !ok {
		return nil, errors.New("failed to get key from context")
	}

	if err := key.LoadRelations(h.loaders, k); err != nil {
		return k, errors.Err(err)
	}

	if k.Namespace != nil {
		err = h.Users.Load(
			"id", []interface{}{k.Namespace.Values()["user_id"]}, database.Bind("user_id", "id", k.Namespace),
		)
	}
	return k, errors.Err(err)
}

// UpdateModel unmarshals the request's data into a key, validates it and
// updates the existing key in the database with the content in the form. Upon
// success the updated key is returned. This also returns the form for
// modifying a key.
func (h Key) UpdateModel(r *http.Request) (*key.Key, key.Form, error) {
	f := key.Form{}

	u, ok := user.FromContext(r.Context())

	if !ok {
		return nil, f, errors.New("no user in request context")
	}

	k, ok := key.FromContext(r.Context())

	if !ok {
		return k, f, errors.New("no key in request context")
	}

	keys := key.NewStore(h.DB, u)

	f.Resource = namespace.Resource{
		Author:     u,
		Namespaces: namespace.NewStore(h.DB, u),
	}
	f.Key = k
	f.Keys = keys

	if err := webutil.UnmarshalAndValidate(&f, r); err != nil {
		return k, f, errors.Err(err)
	}

	namespaceId := k.NamespaceID.Int64

	if f.Namespace != "" {
		n, err := namespace.NewStore(h.DB).GetByPath(f.Namespace)

		if err != nil {
			return k, f, errors.Err(err)
		}
		namespaceId = n.ID
	}

	err := keys.Update(k.ID, namespaceId, f.Config)
	return k, f, errors.Err(err)
}

// DeleteModel removes the key in the given request context from the database.
func (h Key) DeleteModel(r *http.Request) error {
	k, ok := key.FromContext(r.Context())

	if !ok {
		return errors.New("no key in request context")
	}
	return errors.Err(key.NewStore(h.DB).Delete(k.ID))
}
