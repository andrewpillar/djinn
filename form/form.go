package form

import (
	"net/http"
	"regexp"

	"github.com/andrewpillar/thrall/errors"

	"github.com/gorilla/schema"
)

var	reAlphaNumDotDash = regexp.MustCompile("^[a-zA-Z0-9\\._\\-]+$")

type Form interface {
	Validate() error

	Fields() map[string]string
}

func ErrField(field string, err error) error {
	err = errors.Cause(err)

	return errors.New(field + " unexpected error: " + err.Error())
}

func ErrFieldExists(field string) error {
	return errors.New(field + " already exists")
}

func ErrFieldInvalid(field, req string) error {
	msg := field + " is invalid"

	if req != "" {
		msg += ", " + req
	}

	return errors.New(msg)
}

func ErrFieldRequired(field string) error {
	return errors.New(field + " can't be blank")
}

func Unmarshal(f Form, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return errors.Err(err)
	}

	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)

	return errors.Err(dec.Decode(f, r.Form))
}

func UnmarshalAndValidate(f Form, r *http.Request) error {
	if err := Unmarshal(f, r); err != nil {
		return errors.Err(err)
	}

	return f.Validate()
}
