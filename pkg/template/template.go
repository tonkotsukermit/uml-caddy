package template

import (
	"html/template"
	"io"

	"github.com/Masterminds/sprig/v3"
)

type Template struct {
	Name	string
	Path	string
	Output	io.Writer
}

// Execute executes a templating function with a default sprig implementation to the provided io,Writer
func (t* Template) Execute(data any) error {

	// This example illustrates that the FuncMap *must* be set before the
	// templates themselves are loaded.
	tpl, err := template.New(t.Name).Funcs(sprig.FuncMap()).ParseGlob(t.Path + "*.tmpl")
	if err != nil {
		return err
	}

	return tpl.Execute(t.Output, data)
  }