package cmd

import "github.com/spf13/cobra"

func newTaskCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "task",
		Short:                 "Manage tasks",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runTask),
	}

	cmd.AddCommand(
		newTaskListCommand(cli),
		newTaskDescribeCommand(cli),
		newTaskDeleteCommand(cli),
	)
	return cmd
}

func runTask(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
