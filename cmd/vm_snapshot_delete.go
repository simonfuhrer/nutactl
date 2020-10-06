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

func newVMSnapshotDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "snapshot-delete [FLAGS] VM",
		Short:                 "Delete a VM Snapshot",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runVMSnapshotDelete),
	}
	flags := cmd.Flags()
	flags.StringP("name", "n", "", "Snapshot Name")
	MarkFlagsRequired(cmd, "name")

	return cmd
}

func runVMSnapshotDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	snapshotName := viper.GetString("name")
	idOrName := args[0]

	vm, err := cli.Client().VM.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	snapshot, err := cli.Client().Snapshot.Get(cli.Context, snapshotName, vm)
	if err != nil {
		return err
	}
	if askForConfirm(fmt.Sprintf("Delete VM Snapshot %s ?", snapshot.Name)) == nil {
		task, err := cli.Client().Snapshot.Delete(cli.Context, vm, snapshot)
		if err != nil {
			return err
		}
		err = cli.WaitTask(cli.Context, task.TaskUUID, 180)
		if err != nil {
			return err
		}
		fmt.Printf("VM Snapshot %s deleted\n", snapshot.Name)
		return nil
	}

	return fmt.Errorf("operation aborted")
}
