package cmd

import (
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newImageDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS] IMAGE",
		Short:                 "Describe an image",
		Aliases:               []string{"d", "des"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runImageDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")
	return cmd
}

func runImageDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]

	image, err := cli.Client().Image.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	imageList := schema.ImageListIntent{
		Entities: []*schema.ImageIntent{image},
	}

	return outputResponse(displayers.Images{ImageListIntent: imageList})
}
