package main

import (
	"bytes"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/chunk-hunkman/uml-caddy/pkg/gen"
	"github.com/chunk-hunkman/uml-caddy/pkg/uml"
	"k8s.io/client-go/util/homedir"
)

const purl = "plant-uml:8080"


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

	r.Route("/puml", func(r chi.Router) {

		r.Get("/k8s", generatek8sPUML)

	})

	r.Route("/png", func(r chi.Router){

		r.Get("/k8s", generatek8sPNG)
	})

	http.ListenAndServe(":8080", r)

}

// generatek8sPUML is an http hanlder that generates a .puml from a given request with query params of "name", "header", and "title"
func generatek8sPUML(w http.ResponseWriter, req *http.Request) {

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
		w.Write([]byte(err.Error()))
	}

}

// generatek8sPNG is an http hanlder that generates a .png from plant uml from a given request with query params of "name", "header", and "title"
func generatek8sPNG(w http.ResponseWriter, req *http.Request) {

	buf := new(bytes.Buffer)

	u := uml.K8sUML{
		UML: uml.UML{
			Name:         chi.URLParam(req, "name"),
			Header:       chi.URLParam(req, "header"),
			Title:        chi.URLParam(req, "title"),
			TemplatePath: "templates/uml/virtual/k8s/",
			Output:       buf,
		},
	}

	err := u.GenerateK8sUML(req.Context(), filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	image, err := gen.GetPlantUMLPng(buf.String(), purl + "/png/")
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(image)

}
