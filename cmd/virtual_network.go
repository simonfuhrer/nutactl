package cmd

import "github.com/spf13/cobra"

func newVirtualNetworkCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "virtualnetwork",
		Short:                 "Manage virtual networks",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVirtualNetwork),
	}
	cmd.AddCommand()
	return cmd
}

func runVirtualNetwork(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
