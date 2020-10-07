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

	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newVMRecoveryPointListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "recoverypoint-list [FLAGS] VM",
		Short:                 "List VMs RecoveryPoints",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runVMRecoveryPointList),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")
	return cmd
}

func runVMRecoveryPointList(cli *CLI, cmd *cobra.Command, args []string) error {
	var list *schema.VMRecoveryPointListIntent
	var err error

	if len(args) > 0 {
		vmUUIDOrName := args[0]
		vm, err := cli.Client().VM.Get(cli.Context, vmUUIDOrName)
		if err != nil {
			return err
		}
		opts := &schema.DSMetadata{Filter: fmt.Sprintf("entity_name==%s", vm.Spec.Name)}
		list, err = cli.Client().VMRecoveryPoint.List(cli.Context, opts)
		if err != nil {
			return err
		}
	} else {
		list, err = cli.Client().VMRecoveryPoint.All(cli.Context)
		if err != nil {
			return err
		}
	}

	for _, vol := range list.Entities {
		if vol.Spec.Resources.ParentVMReference.UUID != "" {
			vm, err := cli.Client().VM.GetByUUID(cli.Context, vol.Spec.Resources.ParentVMReference.UUID)
			if err == nil {
				vol.Spec.Resources.ParentVMReference.Name = vm.Spec.Name
			}
		}
	}
	return outputResponse(displayers.VMRecoveryPoints{VMRecoveryPointListIntent: *list})
}
