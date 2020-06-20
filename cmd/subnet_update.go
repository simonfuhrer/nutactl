package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newSubnetUpdateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "update [FLAGS] SUBNET",
		Short:                 "Update an subnet",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runSubnetUpdate),
	}
	flags := cmd.Flags()
	flags.StringP("name", "n", "", "New project name")
	addSubnetFlags(flags)

	return cmd
}

func runSubnetUpdate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	newName := viper.GetString("name")
	req, err := cli.Client().Subnet.Get(cli.Context, name)
	if err != nil {
		return err
	}
	err = createUpdateSubnetHelper(newName, req)
	if err != nil {
		return err
	}

	result, err := cli.Client().Subnet.Update(cli.Context, req)
	if err != nil {
		return err
	}
	fmt.Printf("Subnet %s with uuid %s updated\n", result.Spec.Name, result.Metadata.UUID)
	return nil
}
