package models

func loadIndex() map[string]string {
	models := map[string]string{
		"K8sUMLVirtualBase":  K8sUMLVirtualBase,
		"K8sVirtual":         K8sVirtual,
		"NamespaceModel":     NamespaceModel,
		"DeploymentModel":    DeploymentModel,
		"PodModel":           PodModel,
	}

	return models
}