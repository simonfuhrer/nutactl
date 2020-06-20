package cmd

import "github.com/spf13/cobra"

func newCategoryCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "category",
		Short:                 "Manage categories",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runCategory),
	}
	cmd.AddCommand(
		newCategoryListCommand(cli),
		newCategoryDescribeCommand(cli),
		newCategoryCreateCommand(cli),
		newCategoryyDeleteCommand(cli),
	)
	return cmd
}

func runCategory(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}
