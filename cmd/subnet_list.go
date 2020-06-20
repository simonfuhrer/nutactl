package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
)

func newSubnetListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List subnets",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runSubnetList),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")
	return cmd
}

func runSubnetList(cli *CLI, cmd *cobra.Command, args []string) error {
	list, err := cli.Client().Subnet.All(cli.Context)
	if err != nil {
		return err
	}

	return outputResponse(displayers.Subnets{SubnetListIntent: *list})
}
