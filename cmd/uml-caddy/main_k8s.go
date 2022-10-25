package main

import (
	"bytes"
	"image/png"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"

	"github.com/chunk-hunkman/uml-caddy/pkg/gen"
	"github.com/chunk-hunkman/uml-caddy/pkg/uml"
	"k8s.io/client-go/util/homedir"
)

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
		return
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
		return
	}

	image, err := gen.GetPlantUMLPng(buf.String(), purl + "/png/")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, image)

}
