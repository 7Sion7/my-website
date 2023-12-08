package main

import (
	"fmt"
	"log"
	"my-website/routes"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

var (
	fileServer = http.FileServer(http.Dir("public/static/"))
)

func main() {

	r := chi.NewRouter()

	routes.Routes(r)

	r.Handle("/static/*", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}
		fileServer.ServeHTTP(w, r)
	})))

	fmt.Println("Serving on Port -> http://localhost:5000")
	if err := http.ListenAndServe(":5000", r); err != nil {
		log.Fatalf("Failure on Listening and Serving: %v", err)
	}
}
