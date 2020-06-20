package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newImageCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "image",
		Short:                 "Manage images",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runImage),
	}
	cmd.AddCommand(
		newImageListCommand(cli),
		newImageDescribeCommand(cli),
		newImageDeleteCommand(cli),
		newImageUpdateCommand(cli),
		newImageCreateCommand(cli),
	)
	cmd.Flags().SortFlags = false
	return cmd
}

func runImage(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func createUpdateImageHelper(name string, req *schema.ImageIntent) error {
	description := viper.GetString("description")

	if len(name) != 0 {
		req.Spec.Name = name
	}
	if len(description) != 0 {
		req.Spec.Description = description
	}

	return nil
}

func addImageFlags(flags *pflag.FlagSet) {
	flags.StringP("description", "d", "", "Description")
}
