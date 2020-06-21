package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
)

func newAvailabilityZoneListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List availability zones",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runAvailabilityZoneList),
	}

	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")

	return cmd
}

func runAvailabilityZoneList(cli *CLI, cmd *cobra.Command, args []string) error {
	list, err := cli.Client().AvailabilityZone.All(cli.Context)
	if err != nil {
		return err
	}
	return outputResponse(displayers.AvailabilityZones{AvailabilityZoneListIntent: *list})
}
