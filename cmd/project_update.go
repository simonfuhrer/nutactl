package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newProjectUpdateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "update [FLAGS] PROJECT",
		Short:                 "Describe an project",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runProjectUpdate),
	}
	flags := cmd.Flags()
	flags.StringP("name", "n", "", "New project name")
	addProjectFlags(flags)

	return cmd
}

func runProjectUpdate(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	newName := viper.GetString("name")

	project, err := cli.Client().Project.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	err = createUpdateProjectHelper(newName, project)
	if err != nil {
		return err
	}

	result, err := cli.Client().Project.Update(cli.Context, project)
	if err != nil {
		return err
	}
	fmt.Printf("Project %s with uuid %s updated\n", result.Spec.Name, result.Metadata.UUID)

	return nil
}
