package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newSubnetDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete [FLAGS] Subnet",
		Short:                 "Delete a Subnet",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runSubnetDelete),
	}

	return cmd
}

func runSubnetDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	subnet, err := cli.Client().Subnet.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	if askForConfirm(fmt.Sprintf("Delete %s ?", subnet.Spec.Name)) == nil {
		if err := cli.Client().Subnet.Delete(cli.Context, subnet.Metadata.UUID); err != nil {
			return err
		}
		fmt.Printf("Subnet %v deleted\n", idOrName)
		return nil
	}
	return fmt.Errorf("operation aborted")
}
