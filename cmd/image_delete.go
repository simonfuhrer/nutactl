package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newImageDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete [FLAGS] Image",
		Short:                 "Delete a Image",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runImageDelete),
	}
	return cmd
}

func runImageDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	image, err := cli.Client().Image.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	if askForConfirm(fmt.Sprintf("Delete %s ?", image.Spec.Name)) == nil {
		err = cli.client.Image.Delete(cli.Context, image.Metadata.UUID)
		if err != nil {
			return err
		}
		fmt.Printf("Image %v deleted\n", idOrName)
		return nil
	}
	return fmt.Errorf("operation aborted")
}
