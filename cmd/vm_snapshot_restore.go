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

func newVMSnapshotRestoreCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "snapshot-restore [FLAGS] VM",
		Short:                 "Restore a VM Snapshot",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMSnapshotRestore),
	}
	flags := cmd.Flags()
	flags.StringP("name", "n", "", "Snapshot Name")
	markFlagsRequired(cmd, "name")

	return cmd
}

func runVMSnapshotRestore(cli *CLI, cmd *cobra.Command, args []string) error {
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

	task, err := cli.Client().Snapshot.Restore(cli.Context, snapshot, vm)
	if err != nil {
		return err
	}

	err = cli.WaitTask(cli.Context, task.TaskUUID, 180)
	if err != nil {
		return err
	}
	fmt.Printf("VM Snapshot %s restored\n", snapshotName)
	return nil
}
