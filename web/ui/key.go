package ui

import (
	"net/http"
	"strconv"

	"github.com/andrewpillar/thrall/crypto"
	"github.com/andrewpillar/thrall/errors"
	"github.com/andrewpillar/thrall/form"
	"github.com/andrewpillar/thrall/log"
	"github.com/andrewpillar/thrall/model"
	"github.com/andrewpillar/thrall/template"
	"github.com/andrewpillar/thrall/template/key"
	"github.com/andrewpillar/thrall/web"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

type Key struct {
	web.Handler
}

func (h Key) Index(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	search := r.URL.Query().Get("search")

	kk, err := u.KeyStore().All(model.Search("name", search))

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	p := &key.IndexPage{
		Page: template.Page{
			URI: r.URL.Path,
		},
		CSRF:   csrf.TemplateField(r),
		Search: search,
		Keys:   kk,
	}

	d := template.NewDashboard(p, r.URL.Path, h.Alert(w, r))

	web.HTML(w, template.Render(d), http.StatusOK)
}

func (h Key) Create(w http.ResponseWriter, r *http.Request) {
	p := &key.Form{
		Form: template.Form{
			CSRF:   csrf.TemplateField(r),
			Errors: h.Errors(w, r),
			Form:   h.Form(w, r),
		},
	}

	d := template.NewDashboard(p, r.URL.Path, h.Alert(w, r))

	web.HTML(w, template.Render(d), http.StatusOK)
}

func (h Key) Store(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to create SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	keys := u.KeyStore()

	f := &form.Key{
		Keys: keys,
	}

	if err := h.ValidateForm(f, w, r); err != nil {
		if _, ok := err.(form.Errors); ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to create SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	enc, err := crypto.Encrypt([]byte(f.Key))

	if err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to create SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	k := keys.New()
	k.Name = f.Name
	k.Key = []byte(enc)
	k.Config = f.Config

	if err := k.Create(); err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to create SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/keys", http.StatusSeeOther)
}

func (h Key) Edit(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)

	id, _ := strconv.ParseInt(vars["key"], 10, 64)

	k, err := u.KeyStore().Find(id)

	if err != nil {
		log.Error.Println(errors.Err(err))
		web.HTMLError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	if k.IsZero() {
		web.HTMLError(w, "Not found", http.StatusNotFound)
		return
	}

	p := &key.Form{
		Form: template.Form{
			CSRF:   csrf.TemplateField(r),
			Errors: h.Errors(w, r),
			Form:   h.Form(w, r),
		},
		Key: k,
	}

	d := template.NewDashboard(p, r.URL.Path, h.Alert(w, r))

	web.HTML(w, template.Render(d), http.StatusOK)
}

func (h Key) Update(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to update SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)

	id, _ := strconv.ParseInt(vars["key"], 10, 64)

	keys := u.KeyStore()

	k, err := keys.Find(id)

	if err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to update SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	if k.IsZero() {
		web.HTMLError(w, "Not found", http.StatusNotFound)
		return
	}

	f := &form.Key{}

	if err := form.Unmarshal(f, r); err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to update SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	if f.Config == "" {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	k.Config = f.Config

	if err := k.Update(); err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to update SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	h.FlashAlert(w, r, template.Success("Key changes saved for key: " + k.Name))

	http.Redirect(w, r, "/keys", http.StatusSeeOther)
}

func (h Key) Destroy(w http.ResponseWriter, r *http.Request) {
	u, err := h.User(r)

	if err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to delete SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)

	id, _ := strconv.ParseInt(vars["key"], 10, 64)

	k, err := u.KeyStore().Find(id)

	if err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to delete SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	if k.IsZero() {
		web.HTMLError(w, "Not found", http.StatusNotFound)
		return
	}

	if err := k.Destroy(); err != nil {
		log.Error.Println(errors.Err(err))
		h.FlashAlert(w, r, template.Danger("Failed to delete SSH key: " + errors.Cause(err).Error()))
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	h.FlashAlert(w, r, template.Success("Key has been deleted: " + k.Name))

	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
}