package routes

import (
	"my-website/routes/controllers"

	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	r.Mount("/", controllers.HomeRoute())
	r.Mount("/projects", controllers.ProjectsRoute())
	r.Mount("/about", controllers.AboutRoute())
}
