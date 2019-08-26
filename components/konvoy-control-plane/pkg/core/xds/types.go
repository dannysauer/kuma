package xds

import (
	"fmt"
	"strings"

	mesh_core "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/apis/mesh"
	core_model "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/model"

	envoy_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"

	"github.com/pkg/errors"
)

type ProxyId struct {
	Mesh      string
	Namespace string
	Name      string
}

func (id *ProxyId) String() string {
	return fmt.Sprintf("%s.%s.%s", id.Name, id.Namespace, id.Mesh)
}

type Proxy struct {
	Id        ProxyId
	Dataplane *mesh_core.DataplaneResource
}

func ParseProxyId(node *envoy_core.Node) (*ProxyId, error) {
	if node == nil {
		return nil, errors.Errorf("Envoy node must not be nil")
	}
	return ParseProxyIdFromString(node.Id)
}

func ParseProxyIdFromString(id string) (*ProxyId, error) {
	parts := strings.Split(id, ".")
	name := parts[0]
	if id == "" {
		return nil, errors.New("name must not be empty")
	}
	ns := core_model.DefaultNamespace
	if 1 < len(parts) {
		ns = parts[1]
	}
	mesh := core_model.DefaultMesh
	if 2 < len(parts) {
		mesh = parts[2]
	}
	return &ProxyId{
		Mesh:      mesh,
		Namespace: ns,
		Name:      name,
	}, nil
}

func (id *ProxyId) ToResourceKey() core_model.ResourceKey {
	return core_model.ResourceKey{
		Name:      id.Name,
		Namespace: id.Namespace,
		Mesh:      id.Mesh,
	}
}
