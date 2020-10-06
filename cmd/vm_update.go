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

func newVMUpdateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "update [FLAGS] VM",
		Short:                 "Update a VM",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               cli.ensureContext,
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
