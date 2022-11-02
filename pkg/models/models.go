package models

import (
	"io"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/exp/maps"
)

type Template struct {
	Name	string
	Base    string
	Model   string
	Output	io.Writer
}

// Execute executes a templating function with a default sprig implementation to the provided io,Writer
func (t *Template) Execute(data any, model string) error {

	// This example illustrates that the FuncMap *must* be set before the
	// templates themselves are loaded.
	tpl, err := template.New(t.Name).Funcs(sprig.FuncMap()).ParseGlob(model)
	if err != nil {
		return err
	}

	return tpl.Execute(t.Output, data)
}

// ExecuteUML executes a templating function with a default sprig implementation and a given base of template.Base to the provided io,Writer
func (t *Template) ExecuteUML(data any) error {

	// This example illustrates that the FuncMap *must* be set before the
	// templates themselves are loaded.
	tpl, err := template.New(t.Name).Funcs(t.FuncMap()).ParseGlob(t.Model)
	if err != nil {
		return err
	}

	return tpl.Execute(t.Output, data)
}

func (t *Template) FuncMap() template.FuncMap {

	s := sprig.FuncMap()
	
	//TODO: needs to run a function that accepts a parameter that can bring down and template the constants
	f := template.FuncMap{
		"k8sVirtual": K8sVirtual,
		"k8sNamespace": NamespaceModel,
		"k8sDeployment": DeploymentModel,
		"k8sPod": PodModel,
	}

	maps.Copy(f, s)

	return f

}