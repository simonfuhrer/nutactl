package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newImageUpdateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "update [FLAGS] IMAGE",
		Short:                 "Update a image",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runImageUpdate),
	}
	flags := cmd.Flags()
	flags.StringP("name", "n", "", "New image name")
	addImageFlags(flags)

	return cmd
}

func runImageUpdate(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	newName := viper.GetString("name")

	image, err := cli.Client().Image.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	err = createUpdateImageHelper(newName, image)
	if err != nil {
		return err
	}

	result, err := cli.Client().Image.Update(cli.Context, image)
	if err != nil {
		return err
	}

	fmt.Printf("Image %s with uuid %s updated\n", result.Spec.Name, result.Metadata.UUID)

	return nil
}
