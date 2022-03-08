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

func newSubnetCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] subnetname",
		Short:                 "Create an subnet",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runSubnetCreate),
	}
	flags := cmd.Flags()
	flags.String("cluster", "", "Cluster (UUID or name)")
	//MarkFlagsRequired(cmd, "cluster")
	addSubnetFlags(flags)

	return cmd
}

func runSubnetCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	clusterIdorName := viper.GetString("cluster")
	vpcIdorName := viper.GetString("vpc")
	subnetType := viper.GetString("type")
	var cluster *schema.ClusterIntent
	if subnetType != "VLAN" && subnetType != "OVERLAY" {
		return fmt.Errorf("type should be VLAN or OVERLAY not %s", subnetType)

	}

	subnet, _ := cli.Client().Subnet.Get(cli.Context, name)
	if subnet != nil {
		return fmt.Errorf("subnet %s already exists with uuid %s", name, subnet.Metadata.UUID)
	}

	req := &schema.SubnetIntent{
		Spec: &schema.Subnet{
			Resources: &schema.SubnetResources{
				SubnetType: subnetType,
			},
		},
		Metadata: &schema.Metadata{
			Kind: "subnet",
		},
	}

	if subnetType == "VLAN" {
		var err error
		cluster, err = cli.Client().Cluster.Get(cli.Context, clusterIdorName)
		if err != nil {
			return fmt.Errorf("cluster not found %s", clusterIdorName)
		}
		req.Spec.ClusterReference = &schema.Reference{
			Kind: "cluster",
			UUID: cluster.Metadata.UUID,
		}
	}

	if subnetType == "OVERLAY" {
		if len(vpcIdorName) == 0 {
			return fmt.Errorf("vpc id or name is required")
		}
		vpc, err := cli.Client().VPC.Get(cli.Context, vpcIdorName)
		if err != nil {
			return fmt.Errorf("vpc not found %s", vpcIdorName)
		}
		req.Spec.Resources.VPCReference = &schema.Reference{
			Kind: "vpc",
			UUID: vpc.Metadata.UUID,
		}
	}

	err := createUpdateSubnetHelper(name, req)
	if err != nil {
		return err
	}

	result, err := cli.Client().Subnet.Create(cli.Context, req)
	if err != nil {
		return err
	}

	fmt.Printf("Subnet %s with uuid %s created\n", result.Spec.Name, result.Metadata.UUID)
	return nil
}
