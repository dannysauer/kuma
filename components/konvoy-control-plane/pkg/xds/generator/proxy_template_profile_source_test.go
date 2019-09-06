package generator_test

import (
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/config/xds"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"path/filepath"

	mesh_proto "github.com/Kong/konvoy/components/konvoy-control-plane/api/mesh/v1alpha1"
	mesh_core "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/apis/mesh"
	model "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/xds"
	util_proto "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/util/proto"
	xds_envoy "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/xds/envoy"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/xds/generator"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/xds/template"

	test_model "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/test/resources/model"
)

var _ = Describe("ProxyTemplateProfileSource", func() {

	type testCaseFile struct {
		dataplaneFile   string
		profile         *mesh_proto.ProxyTemplateProfileSource
		envoyConfigFile string
	}

	DescribeTable("Generate Envoy xDS resources",
		func(given testCaseFile) {
			// setup
			gen := &generator.ProxyTemplateProfileSource{
				Profile: given.profile,
			}

			// given
			ctx := xds_envoy.Context{
				ControlPlane: &xds_envoy.ControlPlaneContext{
					Config: xds.SnapshotConfig{
						SdsLocation: "konvoy-system:5677",
					},
					SdsTlsCert: []byte("12345"),
				},
			}

			dataplane := mesh_proto.Dataplane{}
			dpBytes, err := ioutil.ReadFile(filepath.Join("testdata", "profile-source", given.dataplaneFile))
			Expect(err).ToNot(HaveOccurred())
			Expect(util_proto.FromYAML(dpBytes, &dataplane)).To(Succeed())
			proxy := &model.Proxy{
				Id: model.ProxyId{Name: "side-car", Namespace: "default"},
				Dataplane: &mesh_core.DataplaneResource{
					Meta: &test_model.ResourceMeta{
						Version: "1",
					},
					Spec: dataplane,
				},
			}

			// when
			rs, err := gen.Generate(ctx, proxy)

			// then
			Expect(err).ToNot(HaveOccurred())

			// then
			resp := generator.ResourceList(rs).ToDeltaDiscoveryResponse()
			actual, err := util_proto.ToYAML(resp)

			expected, err := ioutil.ReadFile(filepath.Join("testdata", "profile-source", given.envoyConfigFile))
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(MatchYAML(expected))
		},
		Entry("should support pre-defined `default-proxy` profile; transparent_proxying=false", testCaseFile{
			dataplaneFile: "1-dataplane.yaml",
			profile: &mesh_proto.ProxyTemplateProfileSource{
				Name: template.ProfileDefaultProxy,
			},
			envoyConfigFile: "1-envoy-config.yaml",
		}),
		Entry("should support pre-defined `default-proxy` profile; transparent_proxying=true", testCaseFile{
			dataplaneFile: "2-dataplane.yaml",
			profile: &mesh_proto.ProxyTemplateProfileSource{
				Name: template.ProfileDefaultProxy,
			},
			envoyConfigFile: "2-envoy-config.yaml",
		}),
	)
})
