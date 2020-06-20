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
