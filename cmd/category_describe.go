package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newCategoryDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS] CATEGORY",
		Short:                 "Describe an category",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runCategoryDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")

	return cmd
}

func runCategoryDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]

	category, err := cli.Client().Category.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	list := schema.CategoryKeyList{
		Entities: []*schema.CategoryKeyStatus{category},
	}

	return outputResponse(displayers.Categories{CategoryKeyList: list})
}
