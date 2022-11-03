package models

import (
	"bytes"
	"errors"
	"io"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"golang.org/x/exp/maps"
)

type Template struct {
	Name	string
	Model   string
	Output	io.Writer
}

// Execute executes a templating function with a default sprig implementation to the provided io,Writer
func (t *Template) Execute(data any) error {

	tpl, err := template.New(t.Name).Funcs(t.funcMap()).Parse(t.Model)
	if err != nil {
		return err
	}

	return tpl.Execute(t.Output, data)
}

func (t *Template) retrieveModel(s string) (string, error){

	i := loadIndex()

	if i[s] != ""{
		return i[s], nil
	}

	return "", errors.New("Could not find model " + s)

}

func (t *Template) funcMap() template.FuncMap {

	s := sprig.FuncMap()
	
	//TODO: needs to run a function that accepts a parameter that can bring down and template the constants
	f := template.FuncMap{
		"partial": t.partial,
	}

	maps.Copy(f, s)

	return f

}

func (t *Template) partial(modelName string, data any) (string, error) {

	buf := new(bytes.Buffer)

	m, err := t.retrieveModel(modelName)
	if err != nil {
		return "", err
	}

	tpl, err := template.New(modelName).Funcs(t.funcMap()).Parse(m)
	if err != nil {
		return "", err
	}

	err = tpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil

}
