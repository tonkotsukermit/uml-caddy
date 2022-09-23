package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/chunk-hunkman/uml-caddy/pkg/uml"
	"k8s.io/client-go/util/homedir"
)

const (
	name = "myUML"
	header = "my header"
	title = "my title"
	
)

func main(){

	f, err := os.Create("tmp.puml")
	if err != nil {
		panic(err)
	}
	
	ctx := context.Background()

	u := uml.K8sUML{
		UML:          uml.UML{
			Name:         name,
			Header:       header,
			Title:        title,
			TemplatePath: "templates/uml/virtual/k8s/",
			Output:       f,
		},
	}

	err = u.GenerateK8sUML(ctx, filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil{
		panic(err)
	}

}