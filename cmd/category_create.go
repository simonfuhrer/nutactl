package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newCategoryCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] CATEGORY",
		Short:                 "Create an category",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runCategoryCreate),
	}
	cmd.Flags().StringP("description", "d", "", "Description")

	return cmd
}

func runCategoryCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	description := viper.GetString("description")

	req := &schema.CategoryKey{
		Name: name,
	}
	if description != "" {
		req.Description = description
	}
	result, err := cli.Client().Category.Create(cli.Context, req)
	if err != nil {
		return err
	}
	fmt.Printf("Category %s created\n", result.Name)

	return nil
}
