package object

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/form"
	"github.com/andrewpillar/thrall/namespace"
	"github.com/andrewpillar/thrall/user"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/jmoiron/sqlx"
)

var (
	width  = 200
	height = 100

	namespaceCols = []string{
		"id",
		"user_id",
		"root_id",
		"parent_id",
		"name",
		"path",
		"description",
		"level",
		"visibility",
		"created_at",
	}

	collabCols = []string{
		"namespace_id",
		"user_id",
	}
)

func createImage(t *testing.T) image.Image {
	img := image.NewRGBA(image.Rectangle{
		image.Point{0, 0},
		image.Point{width, height},
	})

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			r := uint8(rand.Intn(255))
			g := uint8(rand.Intn(255))
			b := uint8(rand.Intn(255))

			img.Set(i, j, color.RGBA{r, g, b, 0xFF})
		}
	}
	return img
}

func namespaceStore(t *testing.T) (namespace.Store, sqlmock.Sqlmock, func() error) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}
	return namespace.NewStore(sqlx.NewDb(db, "sqlmock")), mock, db.Close
}

func Test_FormValidate(t *testing.T) {
	objectStore, objectMock, objectClose := store(t)
	defer objectClose()

	namespaceStore, namespaceMock, namespaceClose := namespaceStore(t)
	defer namespaceClose()

	tests := []struct {
		form        Form
		errs        []string
		shouldError bool
	}{
		{
			Form{
				ResourceForm: namespace.ResourceForm{
					Namespaces: namespaceStore,
				},
				Objects:      objectStore,
				Name:         "rand.png",
			},
			[]string{},
			false,
		},
		{
			Form{
				ResourceForm: namespace.ResourceForm{
					Namespaces: namespaceStore,
					Namespace:  "blackmesa",
					User:       &user.User{ID: 10},
				},
				Objects:      objectStore,
				Name:         "rand.png",
			},
			[]string{},
			false,
		},
		{
			Form{
				ResourceForm: namespace.ResourceForm{
					Namespaces: namespaceStore,
					Namespace:  "blackmesa",
					User:       &user.User{ID: 10},
				},
				Objects:      objectStore,
			},
			[]string{"name"},
			true,
		},
	}

	for _, test := range tests {
		if test.form.Namespace != "" {
			namespaceMock.ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM namespaces WHERE (path = $1)"),
			).WithArgs(test.form.Namespace).WillReturnRows(
				sqlmock.NewRows(namespaceCols).AddRow(1, 10, 1, 0, "blackmesa", "blackmesa", "", 1, namespace.Internal, time.Now()),
			)

			namespaceMock.ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM namespace_collaborators WHERE (namespace_id = $1)"),
			).WithArgs(1).WillReturnRows(
				sqlmock.NewRows(collabCols).AddRow(1, test.form.User.ID),
			)
		}

		objectMock.ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM objects WHERE (name = $1)"),
		).WithArgs(test.form.Name).WillReturnRows(sqlmock.NewRows(objectCols))

		pr, pw := io.Pipe()

		mw := multipart.NewWriter(pw)

		go func() {
			defer mw.Close()

			w, err := mw.CreateFormFile("file", "rand.png")

			if err != nil {
				t.Fatal(err)
			}

			if err := png.Encode(w, createImage(t)); err != nil {
				t.Fatal(err)
			}
		}()

		r := httptest.NewRequest("POST", "/", pr)
		r.Header.Add("Content-Type", mw.FormDataContentType())

		w := httptest.NewRecorder()

		test.form.File = form.File{
			Writer:  w,
			Request: r,
		}

		if err := test.form.Validate(); err != nil {
			if test.shouldError {
				if len(test.errs) == 0 {
					continue
				}

				ferrs, ok := err.(form.Errors)

				if !ok {
					t.Fatalf("expected error to be form.Errors, it was not\n%s\n", errors.Cause(err))
				}

				for _, err := range test.errs {
					if _, ok := ferrs[err]; !ok {
						t.Fatalf("expected field '%s' to be in form.Errors, it was not\n", err)
					}
				}
				continue
			}
			t.Fatal(errors.Cause(err))
		}
	}
}
