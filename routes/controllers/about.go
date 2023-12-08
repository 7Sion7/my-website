package controllers

import (
	"html/template"
	t "my-website/routes/controllers/templates"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AboutRoute() http.Handler {
	r := chi.NewRouter()

	r.Get("/", About)

	return r
}

func About(w http.ResponseWriter, r *http.Request) {
	pageVariables := t.PageVariables{
		Body:  template.HTML(about),
		Style: template.CSS(``),
	}

	tpl, err := t.RenderTemplate("About")
	if err != nil {
		http.Error(w, "500 Internal Error", 500)
	}
	w.Header().Set("Content-Type", "text/html")
	err = tpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		http.Error(w, "500 Internal Error", 500)
	}
}
