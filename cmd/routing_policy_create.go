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

func newRoutingPolicyCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] routing policy",
		Short:                 "Create a routing policy",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runRoutingPolicyCreate),
	}
	flags := cmd.Flags()
	flags.String("vpc", "", "vpc uuid or name")
	flags.Bool("isbidirectional", false, "Additionally Create Policy in reverse direction")
	flags.Int32("priority", 0, "priority of rule (between 10-1000")
	flags.String("protocol-type", "ALL", "any of 'ALL', 'TCP', 'UDP', 'ICMP', 'PROTOCOL_NUMBER' ")
	flags.String("source", "", "ALL, INTERNET or CIDR")
	flags.String("destination", "", "ALL, INTERNET or CIDR")
	flags.String("action", "PERMIT", "PERMIT or DENY")
	MarkFlagsRequired(cmd, "vpc", "priority", "action", "protocol-type", "source", "destination")
	return cmd
}

func runRoutingPolicyCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	vpcIdOrName := viper.GetString("vpc")
	isbidirectional := viper.GetBool("isbidirectional")
	priority := viper.GetInt32("priority")
	action := viper.GetString("action")
	protocolType := viper.GetString("protocol-type")
	source := viper.GetString("source")
	destination := viper.GetString("destination")
	vpc, err := cli.Client().VPC.Get(cli.Context, vpcIdOrName)
	if err != nil {
		return err
	}

	req := &schema.RoutingPolicyIntent{
		Metadata: &schema.Metadata{Kind: "routing_policy"},
		Spec: &schema.RoutingPolicy{Name: name, Resources: &schema.RoutingPolicyResources{
			VpcReference: &schema.Reference{
				Kind: "vpc",
				UUID: vpc.Metadata.UUID,
			},
			ProtocolType:    protocolType,
			IsBidirectional: isbidirectional,
			Priority:        int16(priority),
			Action: &schema.RoutingPolicyAction{
				Action: action,
			},
			Source: &schema.NetworkAddress{
				AddressType: source,
			},
			Destination: &schema.NetworkAddress{
				AddressType: destination,
			},
		}},
	}

	result, err := cli.Client().RoutingPolicy.Create(cli.Context, req)
	if err != nil {
		return err
	}
	fmt.Printf("Routing Policy %s with uuid %s created\n", result.Spec.Name, result.Metadata.UUID)

	return nil
}
