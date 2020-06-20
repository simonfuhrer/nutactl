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
