package webutil

import (
	"net/http"

	"github.com/andrewpillar/thrall/template"
)

func HTML(w http.ResponseWriter, content string, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(content))
}

func HTMLError(w http.ResponseWriter, message string, status int) {
	p := &template.Error{
		Code:    status,
		Message: message,
	}

	HTML(w, template.Render(p), status)
}
