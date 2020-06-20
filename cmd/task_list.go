package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
)

func newTaskListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List tasks",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runTaskList),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")
	return cmd
}

func runTaskList(cli *CLI, cmd *cobra.Command, args []string) error {
	list, err := cli.Client().Task.All(cli.Context)
	if err != nil {
		return err
	}

	return outputResponse(displayers.Tasks{TaskListIntent: *list})
}
