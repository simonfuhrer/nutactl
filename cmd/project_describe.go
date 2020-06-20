package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newProjectDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS] PROJECT",
		Short:                 "Describe an project",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runProjectDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")
	return cmd
}

func runProjectDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]

	project, err := cli.Client().Project.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	list := schema.ProjectListIntent{
		Entities: []*schema.ProjectIntent{project},
	}

	return outputResponse(displayers.Projects{ProjectListIntent: list})
}
