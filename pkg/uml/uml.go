package uml

import (
	"context"
	"io"

	"github.com/chunk-hunkman/uml-caddy/pkg/importer"
	"github.com/chunk-hunkman/uml-caddy/pkg/template"
)


type UML struct {
	Name		 string
	Header		 string
	Title		 string
	TemplatePath string
	Output		 io.Writer
}

type K8sUML struct {
	UML
	K8s	importer.K8sResources
}

func (k *K8sUML) GenerateK8sUML(ctx context.Context, kubeconfig string) error{

	r := importer.K8sResources{}

	kube, err := r.New(ctx, kubeconfig)
	if err != nil{
		return err
	}

	k.K8s = *kube

	err = k.K8s.GetResources()
	if err != nil {
		return err
	}

	t := template.Template{
		Name:   "puml.tmpl",
		Path:   k.TemplatePath,
		Output: k.Output,
	}

	return t.Execute(k)
}