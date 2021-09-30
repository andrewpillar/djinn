// Package object implements the database.Model interface for the Object entity.
package object

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"djinn-ci.com/database"
	"djinn-ci.com/env"
	"djinn-ci.com/errors"
	"djinn-ci.com/event"
	"djinn-ci.com/fs"
	"djinn-ci.com/namespace"
	"djinn-ci.com/queue"
	"djinn-ci.com/user"

	"github.com/andrewpillar/query"

	"github.com/jmoiron/sqlx"
)

// Object is the type that represents an object that has been uploaded by a user.
type Object struct {
	ID          int64         `db:"id"`
	UserID      int64         `db:"user_id"`
	AuthorID    int64         `db:"author_id"`
	NamespaceID sql.NullInt64 `db:"namespace_id"`
	Hash        string        `db:"hash"`
	Name        string        `db:"name"`
	Type        string        `db:"type"`
	Size        int64         `db:"size"`
	MD5         []byte        `db:"md5"`
	SHA256      []byte        `db:"sha256"`
	CreatedAt   time.Time     `db:"created_at"`
	DeletedAt   sql.NullTime  `db:"deleted_at"`

	Author    *user.User           `db:"-"`
	User      *user.User           `db:"-"`
	Namespace *namespace.Namespace `db:"-"`
}

type Event struct {
	dis event.Dispatcher

	Object *Object
	Action string
}

// Store is the type for creating and modifying Object models in the database.
// The Store type can have an underlying fs.Store implementation that is used
// for storing the contents of an object.
type Store struct {
	database.Store

	blockStore fs.Store

	// User is the bound user.User model. If not nil this will bind the
	// user.User model to any Object models that are created. If not nil this
	// will append a WHERE clause on the user_id column for all SELECT queries
	// performed.
	User *user.User

	// Namespace is the bound namespace.Namespace model. If not nil this will
	// bind the namespace.Namespace model to any Object models that are created.
	// If not nil this will append a WHERE clause on the namespace_id column for
	// all SELECT queries performed.
	Namespace *namespace.Namespace
}

var (
	_ database.Model  = (*Object)(nil)
	_ database.Binder = (*Store)(nil)
	_ database.Loader = (*Store)(nil)

	_ queue.Job = (*Event)(nil)

	table     = "objects"
	relations = map[string]database.RelationFunc{
		"author":    database.Relation("author_id", "id"),
		"user":      database.Relation("user_id", "id"),
		"namespace": database.Relation("namespace_id", "id"),
	}
)

// NewStore returns a new Store for querying the objects table. Each database
// passed to this function will be bound to the returned Store.
func NewStore(db *sqlx.DB, mm ...database.Model) *Store {
	s := &Store{
		Store: database.Store{DB: db},
	}
	s.Bind(mm...)
	return s
}

// NewStoreWithBlockStore is functionally the same as NewStore, however it sets
// the fs.Store to use on the returned Store. This will allow f or an object
// file to be stored.
func NewStoreWithBlockStore(db *sqlx.DB, blockStore fs.Store, mm ...database.Model) *Store {
	s := NewStore(db, mm...)
	s.blockStore = blockStore
	return s
}

func (s *Store) SetBlockStore(store fs.Store) { s.blockStore = store }

// FromContext returns the Object model from the given context, if any.
func FromContext(ctx context.Context) (*Object, bool) {
	o, ok := ctx.Value("object").(*Object)
	return o, ok
}

// LoadRelations loads all of the available relations for the given Object
// models using the given loaders available.
func LoadRelations(loaders *database.Loaders, oo ...*Object) error {
	mm := database.ModelSlice(len(oo), Model(oo))
	return errors.Err(database.LoadRelations(relations, loaders, mm...))
}

// Model is called along with database.ModelSlice to convert the given slice of Object
// models to a slice of database.Model interfaces.
func Model(oo []*Object) func(int) database.Model {
	return func(i int) database.Model {
		return oo[i]
	}
}

func InitEvent(dis event.Dispatcher) queue.InitFunc {
	return func(j queue.Job) {
		if ev, ok := j.(*Event); ok {
			ev.dis = dis
		}
	}
}

func (ev *Event) Name() string { return "event:" + event.Objects.String() }

func (ev *Event) Perform() error {
	if ev.dis == nil {
		return event.ErrNilDispatcher
	}

	payload := map[string]interface{}{
		"object": ev.Object.JSON(env.DJINN_API_SERVER),
		"action": ev.Action,
	}
	return errors.Err(ev.dis.Dispatch(event.New(ev.Object.NamespaceID, event.Objects, payload)))
}

// Bind implements the database.Binder interface. This will only bind the model
// if they are pointers to either user.User or namespace.Namespace.
func (o *Object) Bind(mm ...database.Model) {
	for _, m := range mm {
		switch v := m.(type) {
		case *user.User:
			o.Author = v

			if o.UserID == v.ID {
				o.User = v
			}
		case *namespace.Namespace:
			o.Namespace = v
		}
	}
}

// SetPrimary implements the database.Model interface.
func (o *Object) SetPrimary(id int64) { o.ID = id }

// Primary implements the database.Model interface.
func (o *Object) Primary() (string, int64) { return "id", o.ID }

// Endpoint returns the endpoint for the current Object. Each URI part in the
// given variadic list will be appended to the final returned string.
func (o *Object) Endpoint(uri ...string) string {
	if len(uri) > 0 {
		return "/objects/" + strconv.FormatInt(o.ID, 10) + "/" + strings.Join(uri, "/")
	}
	return "/objects/" + strconv.FormatInt(o.ID, 10)
}

// IsZero implements the database.Model interface.
func (o *Object) IsZero() bool {
	return o == nil || o.ID == 0 &&
		o.UserID == 0 &&
		!o.NamespaceID.Valid &&
		o.Hash == "" &&
		o.Name == "" &&
		o.Type == "" &&
		o.Size == 0 &&
		len(o.MD5) == 0 &&
		len(o.SHA256) == 0 &&
		o.CreatedAt == time.Time{} &&
		!o.DeletedAt.Valid
}

// JSON implements the database.Model interface. This will return a map with
// the current Object values under each key. If any of the User, or Namespace
// bound models exist on the Object, then the JSON representation of these
// models will be returned in the map, under the user, and namespace keys
// respectively.
func (o *Object) JSON(addr string) map[string]interface{} {
	json := map[string]interface{}{
		"id":           o.ID,
		"author_id":    o.AuthorID,
		"user_id":      o.UserID,
		"namespace_id": nil,
		"name":         o.Name,
		"type":         o.Type,
		"size":         o.Size,
		"md5":          hex.EncodeToString(o.MD5),
		"sha256":       hex.EncodeToString(o.SHA256),
		"created_at":   o.CreatedAt.Format(time.RFC3339),
		"url":          addr + o.Endpoint(),
	}

	if o.NamespaceID.Valid {
		json["namespace_id"] = o.NamespaceID.Int64
	}

	for name, m := range map[string]database.Model{
		"author":    o.Author,
		"user":      o.User,
		"namespace": o.Namespace,
	} {
		if !m.IsZero() {
			json[name] = m.JSON(addr)
		}
	}
	return json
}

// Values implements the database.Model interface. This will return a map with
// the following values, user_id, namespace_id, hash, name, type, size, md5,
// sha256, and deleted_at.
func (o *Object) Values() map[string]interface{} {
	return map[string]interface{}{
		"user_id":      o.UserID,
		"author_id":    o.AuthorID,
		"namespace_id": o.NamespaceID,
		"hash":         o.Hash,
		"name":         o.Name,
		"type":         o.Type,
		"size":         o.Size,
		"md5":          o.MD5,
		"sha256":       o.SHA256,
		"deleted_at":   o.DeletedAt,
	}
}

// Bind implements the database.Binder interface. This will only bind the model
// if they are pointers to either user.User or namespace.Namespace.
func (s *Store) Bind(mm ...database.Model) {
	for _, m := range mm {
		switch v := m.(type) {
		case *user.User:
			s.User = v
		case *namespace.Namespace:
			s.Namespace = v
		}
	}
}

// Create stores a new object with the given name, and hash. The given
// io.ReadSeeker is used to determine the content type of the object being
// stored, and used to copy the contents of the object to the underlying
// fs.Store. It is expected for the Store to have a fs.Store set on it,
// otherwise it will error.
func (s *Store) Create(authorId int64, name, hash string, rs io.ReadSeeker) (*Object, error) {
	if s.blockStore == nil {
		return nil, errors.New("nil block store")
	}

	header := make([]byte, 512)

	if _, err := rs.Read(header); err != nil {
		return nil, errors.Err(err)
	}

	if _, err := rs.Seek(0, io.SeekStart); err != nil {
		return nil, errors.Err(err)
	}

	md5 := md5.New()
	sha256 := sha256.New()

	tee := io.TeeReader(rs, io.MultiWriter(md5, sha256))

	dst, err := s.blockStore.Create(hash)

	if err != nil {
		return nil, errors.Err(err)
	}

	defer dst.Close()

	size, err := io.Copy(dst, tee)

	if err != nil {
		return nil, errors.Err(err)
	}

	o := s.New()
	o.AuthorID = authorId
	o.Name = name
	o.Hash = hash
	o.Type = http.DetectContentType(header)
	o.Size = size
	o.MD5 = md5.Sum(nil)
	o.SHA256 = sha256.Sum(nil)
	o.CreatedAt = time.Now()

	err = s.Store.Create(table, o)
	return o, errors.Err(err)
}

func (s *Store) Chown(from, to int64) error { return errors.Err(s.Store.Chown(table, from, to)) }

// Delete deletes the object from the database of the given id. The given hash
// is used to remove the object from the underlying fs.Store.
func (s *Store) Delete(id int64, hash string) error {
	if s.blockStore == nil {
		return errors.New("nil block store")
	}

	if err := s.Store.Delete(table, &Object{ID: id}); err != nil {
		return errors.Err(err)
	}
	return errors.Err(s.blockStore.Remove(hash))
}

// Paginate returns the database.Paginator for the objects table for the given
// page. This applies the namespace.WhereCollaborator option to the *user.User
// bound database, and the database.Where option to the *namespace.Namespace bound
// database.
func (s *Store) Paginate(page int64, opts ...query.Option) (database.Paginator, error) {
	opts = append([]query.Option{
		namespace.WhereCollaborator(s.User),
		database.Where(s.Namespace, "namespace_id"),
	}, opts...)

	paginator, err := s.Store.Paginate(table, page, 25, opts...)
	return paginator, errors.Err(err)
}

// New returns a new Object binding any non-nil models to it from the current
// Store.
func (s *Store) New() *Object {
	o := &Object{
		User:      s.User,
		Namespace: s.Namespace,
	}

	if s.User != nil {
		o.UserID = s.User.ID
	}

	if s.Namespace != nil {
		o.NamespaceID = sql.NullInt64{
			Int64: s.Namespace.ID,
			Valid: true,
		}
	}
	return o
}

// All returns a slice of Object models, applying each query.Option that is
// given. The namespace.WhereCollaborator option is applied to the *user.User
// bound database, and the database.Where option is applied to the
// *namespace.Namespace bound database.
func (s *Store) All(opts ...query.Option) ([]*Object, error) {
	oo := make([]*Object, 0)

	opts = append([]query.Option{
		namespace.WhereCollaborator(s.User),
		database.Where(s.Namespace, "namespace_id"),
	}, opts...)

	err := s.Store.All(&oo, table, opts...)

	if err == sql.ErrNoRows {
		err = nil
	}

	for _, o := range oo {
		o.User = s.User
		o.Namespace = s.Namespace
	}
	return oo, errors.Err(err)
}

// Index returns the paginated results from the objects table depending on the
// values that are present in url.Values. Detailed below are the values that
// are used from the given url.Values,
//
// name - This applies the database.Search query.Option using the value of name
func (s *Store) Index(vals url.Values, opts ...query.Option) ([]*Object, database.Paginator, error) {
	page, err := strconv.ParseInt(vals.Get("page"), 10, 64)

	if err != nil {
		page = 1
	}

	opts = append([]query.Option{
		database.Search("name", vals.Get("search")),
	}, opts...)

	paginator, err := s.Paginate(page, opts...)

	if err != nil {
		return []*Object{}, paginator, errors.Err(err)
	}

	oo, err := s.All(append(
		opts,
		query.OrderAsc("name"),
		query.Limit(paginator.Limit),
		query.Offset(paginator.Offset),
	)...)
	return oo, paginator, errors.Err(err)
}

// Load loads in a slice of Object models where the given key is in the list
// of given vals. Each database is loaded individually via a call to the given
// load callback. This method calls Store.All under the hood, so any
// bound models will impact the models being loaded.
func (s *Store) Load(key string, vals []interface{}, load database.LoaderFunc) error {
	oo, err := s.All(query.Where(key, "IN", database.List(vals...)))

	if err != nil {
		return errors.Err(err)
	}

	for i := range vals {
		for _, o := range oo {
			load(i, o)
		}
	}
	return nil
}

// Get returns a single Object database, applying each query.Option that is given.
// The namespace.WhereCollaborator option is applied to the *user.User bound
// database, and the database.Where option is applied to the *namespace.Namespace
// bound database.
func (s *Store) Get(opts ...query.Option) (*Object, error) {
	o := &Object{
		User:      s.User,
		Namespace: s.Namespace,
	}

	opts = append([]query.Option{
		namespace.WhereCollaborator(s.User),
		database.Where(s.Namespace, "namespace_id"),
	}, opts...)

	err := s.Store.Get(o, table, opts...)

	if err == sql.ErrNoRows {
		err = nil
	}
	return o, errors.Err(err)
}
