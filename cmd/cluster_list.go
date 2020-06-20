package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
)

func newClusterListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List clusters",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runClusterList),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")

	return cmd
}

func runClusterList(cli *CLI, cmd *cobra.Command, args []string) error {
	list, err := cli.Client().Cluster.All(cli.Context)
	if err != nil {
		return err
	}

	return outputResponse(displayers.Clusters{ClusterListIntent: *list})
}
