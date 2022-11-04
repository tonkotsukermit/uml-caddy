package main

import (
	"bytes"
	"net/http"
	"os"
	"path/filepath"

	"github.com/chunk-hunkman/uml-caddy/pkg/gen"
	"k8s.io/client-go/util/homedir"
)

var kubeconfig string

func init(){
	if _, err := os.Stat(filepath.Join(homedir.HomeDir(), ".kube", "config")); err == nil {
		//kubeconfig exists
		kubeconfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
	} else {
		//kubeconfig not present
		kubeconfig = ""
	}
}

// generateK8sPUML is an http handler that generates a .puml from a given request with query params of "name", "header", and "title"
// @Summary      Generate a k8s virtual puml document
// @Param        name   query     string  false  "name"
// @Param        header query     string  false  "header"
// @Param        title  query     string  false  "title"
// @Success      200
// @Router       /puml/k8s [get]
func generateK8sPUML(w http.ResponseWriter, req *http.Request) {

	u := newK8sPUML(w, req)

	err := u.GenerateVirtualK8sUML(req.Context(), filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

// generateK8sPNG is an http handler that generates a .png from plant uml from a given request with query params of "name", "header", and "title"
// @Summary      Generate a k8s infrastructure uml png
// @Produce      image/png
// @Param        name   query     string  false  "name"
// @Param        header query     string  false  "header"
// @Param        title  query     string  false  "title"
// @Success      200
// @Router       /png/k8s [get]
func generateK8sPNG(w http.ResponseWriter, req *http.Request) {

	buf := new(bytes.Buffer)

	u := newK8sPUML(w, req)

	err := u.GenerateVirtualK8sUML(req.Context(), kubeconfig)
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
// @Summary      Generate a k8s infrastructure puml document
// @Param        name   query     string  false  "name"
// @Param        header query     string  false  "header"
// @Param        title  query     string  false  "title"
// @Success      200
// @Router       /puml/k8sInfra [get]
func generateK8sInfraPUML(w http.ResponseWriter, req *http.Request) {

	u := newK8sPUML(w, req)

	err := u.GenerateInfraK8sUML(req.Context(), kubeconfig)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

// generateK8sInfraPNG is an http handler that generates a .png from plant uml from a given request with query params of "name", "header", and "title"
// @Summary      Generate a k8s infrastructure uml png
// @Produce      image/png
// @Param        name   query     string  false  "name"
// @Param        header query     string  false  "header"
// @Param        title  query     string  false  "title"
// @Success      200
// @Router       /png/k8sInfra [get]
func generateK8sInfraPNG(w http.ResponseWriter, req *http.Request) {

	buf := new(bytes.Buffer)

	u := newK8sPUML(w, req)

	err := u.GenerateInfraK8sUML(req.Context(), kubeconfig)
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

//newK8sPUML returns a standardized gen.K8sUML struct
func newK8sPUML(w http.ResponseWriter, req *http.Request) gen.K8sUML {

	return gen.K8sUML{
		UML: gen.UML{
			Name:         checkParam(req.URL.Query().Get("name")),
			Header:       checkParam(req.URL.Query().Get("header")),
			Title:        checkParam(req.URL.Query().Get("title")),
			Output:       w,
		},
	}

}