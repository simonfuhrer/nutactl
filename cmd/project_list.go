package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
)

func newProjectListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List projects",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runProjectList),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")

	return cmd
}

func runProjectList(cli *CLI, cmd *cobra.Command, args []string) error {
	list, err := cli.Client().Project.All(cli.Context)
	if err != nil {
		return err
	}

	return outputResponse(displayers.Projects{ProjectListIntent: *list})
}
