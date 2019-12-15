package form

import (
	"regexp"
	"strings"

	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/model/types"

	"github.com/andrewpillar/query"
)

var reNamespace = regexp.MustCompile("^[a-zA-Z0-9]+$")

type Namespace struct {
	Namespaces model.NamespaceStore
	Namespace  *model.Namespace

	UserID      int64
	Parent      string           `schema:"parent"`
	Name        string           `schema:"name"`
	Description string           `schema:"description"`
	Visibility  types.Visibility `schema:"visibility"`
}

func (f Namespace) Fields() map[string]string {
	m := make(map[string]string)
	m["name"] = f.Name
	m["description"] = f.Description

	return m
}

func (f Namespace) Validate() error {
	errs := NewErrors()

	if f.Name == "" {
		errs.Put("name", ErrFieldRequired("Name"))
	}

	if len(f.Name) < 3 || len(f.Name) > 32 {
		errs.Put("name", ErrFieldInvalid("Name", "must be between 3 and 32 characters in length"))
	}

	if !reNamespace.Match([]byte(f.Name)) {
		errs.Put("name", ErrFieldInvalid("Name", "can only contain numbers and letters"))
	}

	checkUnique := true

	if f.Namespace != nil && !f.Namespace.IsZero() {
		parts := strings.Split(f.Namespace.Path, "/")
		parts[len(parts) - 1] = f.Name

		if f.Namespace.Name == f.Name {
			checkUnique = false
		}

		f.Name = strings.Join(parts, "/")
	} else if f.Parent != "" {
		f.Name = strings.Join([]string{f.Parent, f.Name}, "/")
	}

	if checkUnique {
		n, err := f.Namespaces.Get(query.Where("path", "=", f.Name))

		if err != nil {
			return errors.Err(err)
		}

		if !n.IsZero() {
			errs.Put("name", ErrFieldExists("Name"))
		}
	}

	if len(f.Description) > 255 {
		errs.Put("description", ErrFieldInvalid("Description", "must be shorted than 255 characters in length"))
	}

	return errs.Err()
}
