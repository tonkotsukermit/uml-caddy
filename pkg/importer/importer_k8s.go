package importer

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)


type K8sResources struct {
  Namespaces  []Namespace
  Name      string
  Context   context.Context
  Client    *kubernetes.Clientset

} 

type Namespace struct {
	Namespace	corev1.Namespace
	Deployments	[]appsv1.Deployment
}


func buildConfig(kubeconfig string) (*rest.Config, error) {

	if kubeconfig != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
		return cfg, nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

//New creates a new *K8sResources struct with a provided context and kubeconfig. If kubeconfig is blank, it will assume an in-cluster configuration
func (k *K8sResources) New(ctx context.Context, kubeconfig string) (*K8sResources, error) {

	c, err := buildConfig(kubeconfig)
	if err != nil{
		return nil, err
	}

	
	k.Context = ctx
	k.Name = c.Host

	client, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	k.Client = client

	return k, nil
}

//GetResources retrieves all the pre-configured resources to populate the struct utilizing the built in k8s client
func (k *K8sResources)GetResources() error {


	ns, err := k.Client.CoreV1().Namespaces().List(k.Context, metav1.ListOptions{})
	if err != nil {
		return err
	}

	//append namespaces and deployments to []Namespace
	for _, n := range ns.Items {

		ns := Namespace{Namespace: n}

		deploymentsClient := k.Client.AppsV1().Deployments(n.Name)

		list, err := deploymentsClient.List(k.Context, metav1.ListOptions{})
		if err != nil {
			return err
		}

		ns.Deployments = list.Items
		
		k.Namespaces = append(k.Namespaces, ns)

	}

	return nil
}