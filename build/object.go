package build

import (
	"database/sql"
	"io"
	"os"
	"time"

	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/object"
	"github.com/andrewpillar/thrall/runner"

	"github.com/andrewpillar/query"

	"github.com/jmoiron/sqlx"
)

type Object struct {
	ID        int64
	BuildID   int64         `db:"build_id"`
	ObjectID  sql.NullInt64 `db:"object_id"`
	Source    string        `db:"source"`
	Name      string        `db:"name"`
	Placed    bool          `db:"placed"`
	CreatedAt time.Time     `db:"created_at"`

	Build  *Build         `db:"-"`
	Object *object.Object `db:"-"`
}

type ObjectStore struct {
	model.Store

	placer runner.Placer
	Build  *Build
	Object *object.Object
}

var (
	_ model.Model   = (*Object)(nil)
	_ model.Binder  = (*ObjectStore)(nil)
	_ model.Loader  = (*ObjectStore)(nil)
	_ runner.Placer = (*ObjectStore)(nil)

	objectTable = "build_objects"
)

// NewObjectStore returns a new ObjectStore for querying the build_objects
// table. Each model passed to this function will be bound to the returned
// ObjectStore.
func NewObjectStore(db *sqlx.DB, mm ...model.Model) *ObjectStore {
	s := &ObjectStore{
		Store: model.Store{DB: db},
	}
	s.Bind(mm...)
	return s
}

// NewObjectStoreWithCollector returns a new ObjectStore with the given
// runner.Placer to use for object placement.
func NewObjectStoreWithPlacer(db *sqlx.DB, p runner.Placer, mm ...model.Model) *ObjectStore {
	s := NewObjectStore(db, mm...)
	s.placer = p
	return s
}

// ObjectModel is called along with model.Slice to convert the given slice of
// Object models to a slice of model.Model interfaces.
func ObjectModel(oo []*Object) func(int) model.Model {
	return func(i int) model.Model {
		return oo[i]
	}
}

// SelectObject returns SELECT query that will select the given column from the
// build_objects table with the given query options applied.
func SelectObject(col string, opts ...query.Option) query.Query {
	return query.Select(append([]query.Option{
		query.Columns(col),
		query.From(objectTable),
	}, opts...)...)
}

// Bind the given models to the current Object. This will only bind the model
// if they are one of the following,
//
// - *Build
// - *object.Object
func (o *Object) Bind(mm ...model.Model) {
	for _, m := range mm {
		switch m.(type) {
		case *Build:
			o.Build = m.(*Build)
		case *object.Object:
			o.Object = m.(*object.Object)
		}
	}
}

func (o *Object) SetPrimary(id int64) {
	o.ID = id
}

func (o *Object) Primary() (string, int64) {
	return "id", o.ID
}

// Endpoint is a stub to fulfill the model.Model interface. It returns an empty
// string.
func (*Object) Endpoint(_ ...string) string { return "" }

func (o *Object) IsZero() bool {
	return o == nil || o.ID == 0 &&
		o.BuildID == 0 &&
		!o.ObjectID.Valid &&
		o.Source == "" &&
		o.Name == "" &&
		!o.Placed &&
		o.CreatedAt == time.Time{}
}

func (o *Object) Values() map[string]interface{} {
	return map[string]interface{}{
		"build_id":   o.BuildID,
		"object_id":  o.ObjectID,
		"source":     o.Source,
		"name":       o.Name,
		"placed":     o.Placed,
	}
}

// New returns a new Object binding any non-nil models to it from the current
// ObjectStore.
func (s *ObjectStore) New() *Object {
	o := &Object{
		Build:  s.Build,
		Object: s.Object,
	}

	if s.Build != nil {
		_, id := s.Build.Primary()
		o.BuildID = id
	}

	if s.Object != nil {
		_, id := s.Object.Primary()
		o.ObjectID = sql.NullInt64{
			Int64: id,
			Valid: true,
		}
	}
	return o
}

// Create inserts the given Object models into the build_objects table.
func (s *ObjectStore) Create(oo ...*Object) error {
	models := model.Slice(len(oo), ObjectModel(oo))
	return errors.Err(s.Store.Create(objectTable, models...))
}

// Update updates the given Object models in the build_objects table.
func (s *ObjectStore) Update(oo ...*Object) error {
	models := model.Slice(len(oo), ObjectModel(oo))
	return errors.Err(s.Store.Update(objectTable, models...))
}

// All returns a slice of Object models, applying each query.Option that is
// given. Each model that is bound to the store will be applied to the list of
// query options via model.Where.
func (s *ObjectStore) All(opts ...query.Option) ([]*Object, error) {
	oo := make([]*Object, 0)

	opts = append([]query.Option{
		model.Where(s.Build, "build_id"),
		model.Where(s.Object, "object_id"),
	}, opts...)

	err := s.Store.All(&oo, objectTable, opts...)

	if err == sql.ErrNoRows {
		err = nil
	}

	for _, o := range oo {
		o.Build = s.Build
		o.Object = s.Object
	}
	return oo, errors.Err(err)
}

// Get returns a single Object model, applying each query.Option that is given.
// Each model that is bound to the store will be applied to the list of query
// options via model.Where.
func (s *ObjectStore) Get(opts ...query.Option) (*Object, error) {
	o := &Object{
		Build: s.Build,
		Object: s.Object,
	}

	opts = append([]query.Option{
		model.Where(s.Build, "build_id"),
		model.Where(s.Object, "object_id"),
	}, opts...)

	err := s.Store.Get(o, objectTable, opts...)

	if err == sql.ErrNoRows {
		err = nil
	}
	return o, errors.Err(err)
}

// Bind the given models to the current ObjectStore. This will only bind the
// model if they are one of the following,
//
// - *Build
// - *object.Object
func (s *ObjectStore) Bind(mm ...model.Model) {
	for _, m := range mm {
		switch m.(type) {
		case *Build:
			s.Build = m.(*Build)
		case *object.Object:
			s.Object = m.(*object.Object)
		}
	}
}

// Load loads in a slice of Object models where the given key is in the list of
// given vals. Each model is loaded individually via a call to the given load
// callback. This method calls ObjectStore.All under the hood, so any bound
// models will impact the models being loaded.
func (s *ObjectStore) Load(key string, vals []interface{}, load model.LoaderFunc) error {
	oo, err := s.All(query.Where(key, "IN", vals...))

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

func (s *ObjectStore) getObjectToPlace(name string) (*Object, error) {
	if s.placer == nil {
		return nil, errors.New("cannot place object: nil placer")
	}

	o, err := s.Get(query.Where("source", "=", name))

	if err != nil {
		return o, errors.Err(err)
	}

	if o.IsZero() {
		return o, errors.New("cannot find object: "+name)
	}

	o.Object, err = object.NewStore(s.DB).Get(query.Where("id", "=", o.ObjectID))

	if err != nil {
		return o, errors.Err(err)
	}

	if o.Object.IsZero() {
		return o, errors.New("cannot find object: "+name)
	}
	return o, nil
}

// Place looks up the Object by the given name, and updates it once object
// placement has been done via the underlying runner.Placer on the store.
func (s *ObjectStore) Place(name string, w io.Writer) (int64, error) {
	o, err := s.getObjectToPlace(name)

	if err != nil {
		return 0, errors.Err(err)
	}

	n, errPlace := s.placer.Place(o.Object.Hash, w)

	if errors.Cause(errPlace) == io.EOF {
		errPlace = nil
	}

	o.Placed = errPlace == nil

	if err := s.Update(o); err != nil {
		return n, errors.Err(err)
	}
	return n, errors.Err(errPlace)
}

// Stat returns the os.FileInfo of the Object by the given name. This uses the
// underlying runner.Placer on the store.
func (s *ObjectStore) Stat(name string) (os.FileInfo, error) {
	o, err := s.getObjectToPlace(name)

	if err != nil {
		return nil, errors.Err(err)
	}

	info, err := s.placer.Stat(o.Object.Hash)
	return info, errors.Err(err)
}
