// Copyright © 2020 Simon Fuhrer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
