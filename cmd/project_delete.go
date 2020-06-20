package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newProjectDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete [FLAGS] Project",
		Short:                 "Delete a Project",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runProjectDelete),
	}

	return cmd
}

func runProjectDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	project, err := cli.Client().Project.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}
	if askForConfirm(fmt.Sprintf("Delete %s ?", project.Spec.Name)) == nil {
		err = cli.client.Project.Delete(cli.Context, project.Metadata.UUID)
		if err != nil {
			return err
		}
		fmt.Printf("Project %v deleted\n", idOrName)
		return nil
	}
	return fmt.Errorf("operation aborted")
}
