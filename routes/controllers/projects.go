package controllers

import (
	"fmt"
	"html/template"
	t "my-website/routes/controllers/templates"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProjectsRoute() http.Handler {
	r := chi.NewRouter()

	r.Get("/", Projects)
	r.Get("/{project}", Project)

	return r
}

func Projects(w http.ResponseWriter, r *http.Request) {
	pageVariables := t.PageVariables{
		Title: "Projects",
		Body:  template.HTML(projects),
		Style: template.CSS(``),
	}
	tpl, err := t.RenderTemplate("Projects")
	if err != nil {
		http.Error(w, "500 Internal Error", 500)
	}

	w.Header().Set("Content-Type", "text/html")
	err = tpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, "500 Internal Error", 500)
	}
}

func Project(w http.ResponseWriter, r *http.Request) {
	project := chi.URLParam(r, "project")

	var pageVariables t.PageVariables

	switch project {
	case "recipes":
		project = recipes
	case "forum":
		project = forum
	case "currency-converter":
		project = currencyConverter
	case "artists-tracker":
		project = artistTracker
	}

	pageVariables = t.PageVariables{
		Body:  template.HTML(project),
		Style: template.CSS(``),
	}

	tpl, err := t.RenderTemplate("Project")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Error", 500)
	}

	w.Header().Set("Content-Type", "text/html")

	err = tpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		http.Error(w, "500 Internal Error", 500)
	}
}
