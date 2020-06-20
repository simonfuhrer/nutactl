package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	colorWarn = color.YellowString("Warning")
)

func newTaskDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete [FLAGS] Task",
		Short:                 "Delete a Task",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runTaskDelete),
	}

	return cmd
}

func runTaskDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	task, err := cli.Client().Task.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	if err := cli.Client().Task.Delete(cli.Context, task); err != nil {
		return err
	}
	fmt.Printf("Task %v deleted\n", idOrName)
	return nil
}
