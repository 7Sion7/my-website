package controllers

import (
	"fmt"
	"html/template"
	t "my-website/routes/controllers/templates"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HomeRoute() http.Handler {
	r := chi.NewRouter()

	r.Get("/", Home)

	return r
}

func Home(w http.ResponseWriter, r *http.Request) {
	pageVariables := t.PageVariables{
		Body:   template.HTML(home),
		Style:  template.CSS(``),
	}
	tpl, err := t.RenderTemplate("Home")
	if err != nil {
		http.Error(w, "500 Internal Error", 500)
	}
	w.Header().Set("Content-Type", "text/html")
	err = tpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Error", 500)
	}
}
