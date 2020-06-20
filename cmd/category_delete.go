package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCategoryyDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete [FLAGS] CATEGORY",
		Short:                 "Delete a Category",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runCategoryDelete),
	}

	return cmd
}

func runCategoryDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	categoryKey, err := cli.Client().Category.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}
	if askForConfirm(fmt.Sprintf("Delete %s ?", categoryKey.Name)) == nil {
		if err := cli.Client().Category.Delete(cli.Context, categoryKey.Name); err != nil {
			return err
		}
		fmt.Printf("Category %v deleted\n", idOrName)
		return nil
	}
	return fmt.Errorf("operation aborted")
}
