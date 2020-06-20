package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVMDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete [FLAGS] VM",
		Short:                 "Delete a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMDelete),
	}
	return cmd
}

func runVMDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	vm, err := cli.Client().VM.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	if askForConfirm(fmt.Sprintf("Delete %s ?", vm.Spec.Name)) == nil {
		if err := cli.Client().VM.Delete(cli.Context, vm.Metadata.UUID); err != nil {
			return err
		}
		fmt.Printf("VM %v deleted\n", idOrName)
		return nil
	}
	return fmt.Errorf("operation aborted")

}
