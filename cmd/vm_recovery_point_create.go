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
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newVMRecoveryPointCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "recoverypoint-create [FLAGS] VM",
		Short:                 "Create a VM RecoveryPoint",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      false,
		DisableFlagsInUseLine: true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runVMRecoveryPointCreate),
	}

	cmd.Flags().String("name", "", "VM RecoveryPoint Name ")
	return cmd
}

func runVMRecoveryPointCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := viper.GetString("name")
	idOrName := args[0]

	vm, err := cli.Client().VM.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	recoveryPointRequest := &schema.VMRecoveryPointRequest{
		Metadata: &schema.Metadata{
			Kind: "vm_recovery_point",
		},
		Spec: &schema.VMRecoveryPoint{
			ClusterReference: vm.Spec.ClusterReference,
			Name:             name,
			Resources: &schema.VMRecoveryPointResources{
				ParentVMReference: &schema.Reference{
					UUID: vm.Metadata.UUID,
					Kind: vm.Metadata.Kind,
				},
			},
		},
	}

	recoveryPoint, err := cli.Client().VMRecoveryPoint.Create(cli.Context, recoveryPointRequest)
	if err != nil {
		return err
	}

	fmt.Printf("VM RecoveryPoint %s with UUID %s created\n", name, recoveryPoint.Metadata.UUID)
	return nil
}
