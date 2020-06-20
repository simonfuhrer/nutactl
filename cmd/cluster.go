package cmd

import "github.com/spf13/cobra"

func newClusterCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "cluster",
		Short:                 "Manage cluster",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runCluster),
	}
	cmd.AddCommand(
		newClusterListCommand(cli),
		newClusterDescribeCommand(cli),
	)
	return cmd
}

func runCluster(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
