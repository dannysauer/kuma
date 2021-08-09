// +build gateway

package gateway

<<<<<<< HEAD
import core_plugins "github.com/kumahq/kuma/pkg/core/plugins"
=======
import (
	"github.com/kumahq/kuma/pkg/core/plugins"
	_ "github.com/kumahq/kuma/pkg/plugins/runtime/gateway/register"
)
>>>>>>> 57212439 (chore(tools): Simplify resource-gen.go by generating`ResourceDescriptor` (#2511))

func init() {
	plugins.Register("gateway", &plugin{})
}
