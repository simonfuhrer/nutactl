package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newClusterDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS] CLUSTER",
		Short:                 "Describe an cluster",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runClusterDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")

	return cmd
}

func runClusterDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	cluster, err := cli.Client().Cluster.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	list := schema.ClusterListIntent{
		Entities: []*schema.ClusterIntent{cluster},
	}

	return outputResponse(displayers.Clusters{ClusterListIntent: list})
}
