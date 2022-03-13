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

func newFloatingIpCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] External Subnet Name or Uuid",
		Short:                 "Create a Floating IP",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runFloatingIpCreate),
	}
	flags := cmd.Flags()
	flags.String("vm", "", "vm name or uuid to attach")

	return cmd
}

func runFloatingIpCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	vmNameOrId := viper.GetString("vm")
	subnet, err := cli.Client().Subnet.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	req := &schema.FloatingIPIntent{
		Spec: &schema.FloatingIP{
			Resources: &schema.FloatingIPResources{
				ExternalSubnetReference: &schema.Reference{
					Kind: "subnet",
					UUID: subnet.Metadata.UUID,
				},
			},
		},
		Metadata: &schema.Metadata{
			Kind: "floating_ip",
		},
	}

	result, err := cli.Client().FlotatingIP.Create(cli.Context, req)
	if err != nil {
		return err
	}

	fmt.Printf("Flotating IP with uuid %s created\n", result.Metadata.UUID)

	_, err = createUpdateFloatingIpHelper(result.Metadata.UUID, vmNameOrId, cli)
	if err != nil {
		return err
	}
	return nil
}
