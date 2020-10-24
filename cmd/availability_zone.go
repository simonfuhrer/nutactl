package cmd

import "github.com/spf13/cobra"

func newAvailabilityZoneCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "availabilityzone",
		Short:                 "Manage availability zones",
		Aliases:               []string{"a", "avz"},
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runAvailabilityZone),
	}
	cmd.AddCommand(
		newAvailabilityZoneListCommand(cli),
	)
	return cmd
}

func runAvailabilityZone(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
