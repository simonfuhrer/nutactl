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
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v2 "github.com/tecbiz-ch/nutanix-go-sdk/schema/v2"
)

func newVMSnapshotDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "snapshot-describe [FLAGS] VM",
		Short:                 "Describe a VM Snapshot",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVMSnapshotDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")
	cmd.Flags().StringP("name", "n", "", "Snapshot Name")
	markFlagsRequired(cmd, "name")
	return cmd
}

func runVMSnapshotDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
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

	list := v2.SnapshotList{
		Entities: []*v2.SnapshotSpec{snapshot},
	}

	return outputResponse(displayers.VMSnapshots{SnapshotList: list})
}
