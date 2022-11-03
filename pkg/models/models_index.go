package models

func loadIndex() map[string]string {
	models := map[string]string{
		"K8sUMLVirtualBase":  K8sUMLVirtualBase,
		"K8sVirtual":         K8sVirtual,
		"NamespaceModel":     NamespaceModel,
		"DeploymentModel":    DeploymentModel,
		"PodModel":           PodModel,
		"K8sUMLInfraBase":    K8sUMLInfraBase,
		"K8sInfra":           K8sInfra,
		"K8sNode":            K8sNode,
		"K8sTaints":          K8sTaints,

	}

	return models
}