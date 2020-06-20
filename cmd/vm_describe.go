package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"

	"github.com/simonfuhrer/nutactl/cmd/displayers"
)

func newVMDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS] VM",
		Short:                 "Describe a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")
	return cmd
}

func runVMDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]

	vm, err := cli.Client().VM.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	list := schema.VMListIntent{
		Entities: []*schema.VMIntent{vm},
	}

	return outputResponse(displayers.VMs{VMListIntent: list})
}
