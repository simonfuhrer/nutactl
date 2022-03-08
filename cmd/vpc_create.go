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

func newVpcCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] vpcname",
		Short:                 "Create a vpc",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runVpcCreate),
	}
	flags := cmd.Flags()
	flags.StringP("external-subnet", "e", "", "external subnet uuid or name")

	return cmd
}

func runVpcCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	subnetNameorUuid := viper.GetString("external-subnet")

	req := &schema.VpcIntent{
		Spec: &schema.Vpc{
			Resources: &schema.VpcResources{},
		},
		Metadata: &schema.Metadata{
			Kind: "vpc",
		},
	}

	req.Spec.Name = name

	if len(subnetNameorUuid) > 0 {
		subnet, err := cli.Client().Subnet.Get(cli.Context, subnetNameorUuid)
		if err != nil {
			return err
		}

		req.Spec.Resources.ExternalSubnetList = []*schema.ExternalSubnet{
			{
				ExternalSubnetReference: &schema.Reference{Kind: "subnet", UUID: subnet.Metadata.UUID},
			},
		}
	}

	result, err := cli.Client().VPC.Create(cli.Context, req)
	if err != nil {
		return err
	}
	fmt.Printf("VPC %s with uuid %s created\n", result.Spec.Name, result.Metadata.UUID)

	return nil
}
