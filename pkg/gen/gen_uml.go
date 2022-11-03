package gen

import (
	"context"
	"io"

	"github.com/chunk-hunkman/uml-caddy/pkg/importer"
	"github.com/chunk-hunkman/uml-caddy/pkg/models"
)


type UML struct {
	Name		 string
	Header		 string
	Title		 string
	Model        string 
	Output		 io.Writer
}

type K8sUML struct {
	UML
	K8s	importer.K8sResources
}

func (k *K8sUML) GenerateVirtualK8sUML(ctx context.Context, kubeconfig string) error{

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

	t := models.Template{
		Model:  models.K8sUMLVirtualBase,
		Output: k.Output,
	}

	return t.Execute(k)
}

func (k *K8sUML) GenerateInfraK8sUML(ctx context.Context, kubeconfig string) error{

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

	t := models.Template{
		Model:  models.K8sUMLInfraBase,
		Output: k.Output,
	}

	return t.Execute(k)
}
