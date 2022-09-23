package main

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/chunk-hunkman/uml-caddy/pkg/uml"
	"k8s.io/client-go/util/homedir"
)


func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("howdy"))
	})

	r.Route("/uml", func(r chi.Router) {

		r.Get("/k8s", generatek8sUML)

	})

	http.ListenAndServe(":8080", r)

}

// generatek8sUML is an http hanlder that generates a .puml from a given request with query params of "name", "header", and "title"
func generatek8sUML(w http.ResponseWriter, req *http.Request) {

	u := uml.K8sUML{
		UML: uml.UML{
			Name:         chi.URLParam(req, "name"),
			Header:       chi.URLParam(req, "header"),
			Title:        chi.URLParam(req, "title"),
			TemplatePath: "templates/uml/virtual/k8s/",
			Output:       w,
		},
	}

	err := u.GenerateK8sUML(req.Context(), filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		panic(err)
	}

}
