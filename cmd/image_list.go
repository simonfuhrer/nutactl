package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
)

func newImageListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Aliases:               []string{"l", "li"},
		Short:                 "List images",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runImageList),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")
	return cmd
}

func runImageList(cli *CLI, cmd *cobra.Command, args []string) error {
	list, err := cli.Client().Image.All(cli.Context)
	if err != nil {
		return err
	}
	return outputResponse(displayers.Images{ImageListIntent: *list})
}
