package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newTaskDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS] TASK",
		Short:                 "Describe an task",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runTaskDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")
	return cmd
}

func runTaskDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]

	task, err := cli.Client().Task.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	list := schema.TaskListIntent{
		Entities: []*schema.Task{task},
	}

	return outputResponse(displayers.Tasks{TaskListIntent: list})
}
