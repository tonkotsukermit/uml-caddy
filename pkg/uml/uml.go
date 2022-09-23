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
	importer.K8sResources
}

func (k *K8sUML) GenerateK8sUML(ctx context.Context, kubeconfig string) error{

	r := importer.K8sResources{}

	kube, err := r.New(ctx, kubeconfig)
	if err != nil{
		return err
	}

	k.K8sResources = *kube

	t := template.Template{
		Name:   "k8sUML",
		Path:   k.TemplatePath,
		Output: k.Output,
	}

	return t.Execute(k)
}