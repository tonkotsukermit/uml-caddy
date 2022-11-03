package main

import (
	"bytes"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"

	"github.com/chunk-hunkman/uml-caddy/pkg/gen"
	"k8s.io/client-go/util/homedir"
)


// generateK8sPUML is an http handler that generates a .puml from a given request with query params of "name", "header", and "title"
func generateK8sPUML(w http.ResponseWriter, req *http.Request) {

	u := gen.K8sUML{
		UML: gen.UML{
			Name:         checkParam(chi.URLParam(req, "name")),
			Header:       checkParam(chi.URLParam(req, "header")),
			Title:        checkParam(chi.URLParam(req, "title")),
			Output:       w,
		},
	}

	err := u.GenerateVirtualK8sUML(req.Context(), filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

// generateK8sPNG is an http handler that generates a .png from plant uml from a given request with query params of "name", "header", and "title"
func generateK8sPNG(w http.ResponseWriter, req *http.Request) {

	buf := new(bytes.Buffer)

	u := gen.K8sUML{
		UML: gen.UML{
			Name:         checkParam(chi.URLParam(req, "name")),
			Header:       checkParam(chi.URLParam(req, "header")),
			Title:        checkParam(chi.URLParam(req, "title")),
			Output:       w,
		},
	}

	err := u.GenerateVirtualK8sUML(req.Context(), filepath.Join(homedir.HomeDir(), ".kube", "config"))
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
	w.Write(image)

}


// generateK8sInfraPUML is an http handler that generates a .puml from a given request with query params of "name", "header", and "title"
func generateK8sInfraPUML(w http.ResponseWriter, req *http.Request) {

	u := gen.K8sUML{
		UML: gen.UML{
			Name:         checkParam(chi.URLParam(req, "name")),
			Header:       checkParam(chi.URLParam(req, "header")),
			Title:        checkParam(chi.URLParam(req, "title")),
			Output:       w,
		},
	}

	err := u.GenerateInfraK8sUML(req.Context(), filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

// generateK8sInfraPNG is an http handler that generates a .png from plant uml from a given request with query params of "name", "header", and "title"
func generateK8sInfraPNG(w http.ResponseWriter, req *http.Request) {

	buf := new(bytes.Buffer)

	u := gen.K8sUML{
		UML: gen.UML{
			Name:         checkParam(chi.URLParam(req, "name")),
			Header:       checkParam(chi.URLParam(req, "header")),
			Title:        checkParam(chi.URLParam(req, "title")),
			Output:       w,
		},
	}

	err := u.GenerateInfraK8sUML(req.Context(), filepath.Join(homedir.HomeDir(), ".kube", "config"))
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
	w.Write(image)

}
