package k8s

import (
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
)

const (
	// k8sNamespaceComponent identifies the namespace component of a resource name on Kubernetes.
	// The value is considered a part of user-facing Kuma API and should not be changed lightly.
	// The value has a format of a Kubernetes label name.
	k8sNamespaceComponent = "k8s.kuma.io/namespace"

	// k8sNameComponent identifies the name component of a resource name on Kubernetes.
	// The value is considered a part of user-facing Kuma API and should not be changed lightly.
	// The value has a format of a Kubernetes label name.
	k8sNameComponent = "k8s.kuma.io/name"

	// k8sSynced identifies that resource was synced
	K8sSynced = "k8s.kuma.io/synced"

	// K8sProcessed identifies that resource was already processed on creation. Specifically it's used to avoid
	// creation of default resources twice when Kuma CP restarts
	K8sProcessed = "k8s.kuma.io/processed"

	// Kubernetes secret type to differentiate Kuma System secrets. Secret is bound to a mesh
	MeshSecretType = "system.kuma.io/secret"

	// Kubernetes secret type to differentiate Kuma System secrets. Secret is bound to a control plane
	GlobalSecretType = "system.kuma.io/global-secret"
)

func ResourceNameExtensions(namespace, name string) core_model.ResourceNameExtensions {
	return core_model.ResourceNameExtensions{
		k8sNamespaceComponent: namespace,
		k8sNameComponent:      name,
	}
}
