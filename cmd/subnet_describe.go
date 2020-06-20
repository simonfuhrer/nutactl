package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newSubnetDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS] SUBNET",
		Short:                 "Describe an subnet",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runSubnetDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")
	return cmd
}

func runSubnetDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	subnet, err := cli.Client().Subnet.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	list := schema.SubnetListIntent{
		Entities: []*schema.SubnetIntent{subnet},
	}

	return outputResponse(displayers.Subnets{SubnetListIntent: list})
}
