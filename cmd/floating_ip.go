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
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newFloatingIpCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "floatingip",
		Short:                 "Manage vpcs",
		Aliases:               []string{"floating", "fip", ""},
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		RunE:                  cli.wrap(runFloatingIp),
	}
	cmd.AddCommand(
		newFloatingIpListCommand(cli),
		newFloatingIpDescribeCommand(cli),
		newFloatingIpDeleteCommand(cli),
		newFloatingIpCreateCommand(cli),
		newFloatingIpUpdateCommand(cli),
		newFloatingIpAssignCommand(cli),
		newFloatingIpReleaseCommand(cli),
	)
	return cmd
}

func runFloatingIp(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func createUpdateFloatingIpHelper(uuid, vmNameOrId string, cli *CLI) (*schema.FloatingIPIntent, error) {
	if len(vmNameOrId) > 0 {
		vmnicUuid := ""
		vm, err := cli.Client().VM.Get(cli.Context, vmNameOrId)
		if err != nil {
			return nil, err
		}
		for _, v := range vm.Status.Resources.NicList {
			vmnicUuid = v.UUID
			break
		}
		fip, err := cli.Client().FlotatingIP.Get(cli.Context, uuid)
		if err != nil {
			return nil, err
		}
		fip.Spec.Resources.VMNicReference = &schema.Reference{
			Kind: "vm_nic",
			UUID: vmnicUuid,
		}
		fip, err = cli.Client().FlotatingIP.Update(cli.Context, fip)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Floating IP attached to VM with uuid %s\n", vm.Metadata.UUID)
		return fip, nil
	}
	return nil, nil
}
