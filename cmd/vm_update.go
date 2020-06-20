package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newVMUpdateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "update [FLAGS] VM",
		Short:                 "Update a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMUpdate),
	}
	flags := cmd.Flags()
	flags.StringP("name", "n", "", "New VM name")

	addVMFlags(flags)

	return cmd
}

func runVMUpdate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	newName := viper.GetString("name")
	req, err := cli.Client().VM.Get(cli.Context, name)
	if err != nil {
		return err
	}
	err = createUpdateVMHelper(newName, req, cli)
	if err != nil {
		return err
	}

	result, err := cli.Client().VM.Update(cli.Context, req)
	if err != nil {
		return err
	}
	fmt.Printf("VM %s with uuid %s updated\n", result.Spec.Name, result.Metadata.UUID)
	return nil
}
