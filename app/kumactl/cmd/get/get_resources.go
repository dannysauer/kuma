package get

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	kumactl_cmd "github.com/kumahq/kuma/app/kumactl/pkg/cmd"
	"github.com/kumahq/kuma/app/kumactl/pkg/output"
	"github.com/kumahq/kuma/app/kumactl/pkg/output/printers"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	rest_types "github.com/kumahq/kuma/pkg/core/resources/model/rest"
	core_store "github.com/kumahq/kuma/pkg/core/resources/store"
)

<<<<<<< HEAD
func NewGetResourcesCmd(pctx *kumactl_cmd.RootContext, use string, resourceType model.ResourceType, tablePrinter TablePrinter) *cobra.Command {
=======
func NewGetResourcesCmd(pctx *kumactl_cmd.RootContext, desc model.ResourceTypeDescriptor) *cobra.Command {
>>>>>>> 57212439 (chore(tools): Simplify resource-gen.go by generating`ResourceDescriptor` (#2511))
	cmd := &cobra.Command{
		Use:   desc.KumactlListArg,
		Short: fmt.Sprintf("Show %s", desc.Name),
		Long:  fmt.Sprintf("Show %s entities.", desc.Name),
		RunE: func(cmd *cobra.Command, _ []string) error {
			rs, err := pctx.CurrentResourceStore()
			if err != nil {
				return err
			}

			resources := desc.NewList()
			currentMesh := pctx.CurrentMesh()
<<<<<<< HEAD
			if resources.NewItem().Scope() == model.ScopeGlobal {
=======
			resource := resources.NewItem()
			if resource.Descriptor().Scope == model.ScopeGlobal {
>>>>>>> 57212439 (chore(tools): Simplify resource-gen.go by generating`ResourceDescriptor` (#2511))
				currentMesh = ""
			}
			if err := rs.List(context.Background(), resources, core_store.ListByMesh(currentMesh), core_store.ListByPage(pctx.ListContext.Args.Size, pctx.ListContext.Args.Offset)); err != nil {
				return errors.Wrapf(err, "failed to list "+string(desc.Name))
			}

			switch format := output.Format(pctx.GetContext.Args.OutputFormat); format {
			case output.TableFormat:
<<<<<<< HEAD
				return tablePrinter(pctx.Now(), resources, cmd.OutOrStdout())
=======
				return ResolvePrinter(desc.Name, resource.Descriptor().Scope).Print(pctx.Now(), resources, cmd.OutOrStdout())
>>>>>>> 57212439 (chore(tools): Simplify resource-gen.go by generating`ResourceDescriptor` (#2511))
			default:
				printer, err := printers.NewGenericPrinter(format)
				if err != nil {
					return err
				}
				return printer.Print(rest_types.From.ResourceList(resources), cmd.OutOrStdout())
			}
		},
	}
	return cmd
}
