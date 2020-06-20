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
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newVMCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "vm",
		Short:                 "Manage vms",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runVM),
	}
	cmd.AddCommand(
		newVMListCommand(cli),
		newVMDescribeCommand(cli),
		newVMCreateCommand(cli),
		newVMUpdateCommand(cli),
		newVMDeleteCommand(cli),
		newVMPowerStateOnCommand(cli),
		newVMPowerStateOffCommand(cli),
		newVMPowerStateResetCommand(cli),
		newVMPowerStateRebootCommand(cli),
		newVMPowerStateShutdownCommand(cli),
	)
	cmd.Flags().SortFlags = false
	cmd.PersistentFlags().SortFlags = false
	return cmd
}

func runVM(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func addVMFlags(flags *pflag.FlagSet) {
	flags.StringP("description", "d", "", "Description")
	flags.String("project", "", "Project Name or UUID")
	flags.String("subnet", "", "Subnet Name or UUID")
	flags.Int64("numSockets", 1, "Number of CPU Sockets")
	flags.Int64("numCores", 1, "Number of Cores")
	flags.Int64("root-disk-size", 0, "Root Disk Size in MB")
	flags.Int64("memoryMB", 0, "Memory in MB")
}

func createUpdateVMHelper(name string, req *schema.VMIntent, cli *CLI) error {
	description := viper.GetString("description")
	projectIdorName := viper.GetString("project")
	subnetIdorName := viper.GetString("subnet")
	numSockets := viper.GetInt64("numSockets")
	numCores := viper.GetInt64("numCores")
	memoryMB := viper.GetInt64("memoryMB")
	rootDiskSize := viper.GetInt64("root-disk-size")

	if len(name) != 0 {
		req.Spec.Name = name
	}

	if len(description) != 0 {
		req.Spec.Description = description
	}

	if memoryMB > 0 {
		req.Spec.Resources.MemorySizeMib = memoryMB
	}

	if numSockets > 0 {
		req.Spec.Resources.NumSockets = numSockets
	}

	if numCores > 0 {
		req.Spec.Resources.NumVcpusPerSocket = numCores
	}

	if rootDiskSize > 0 {
		req.Spec.Resources.DiskList[0].DiskSizeMib = rootDiskSize
	}

	if len(projectIdorName) != 0 {
		project, err := cli.Client().Project.Get(cli.Context, projectIdorName)
		if err != nil {
			return fmt.Errorf("project not found %s", projectIdorName)
		}
		req.Metadata.ProjectReference = &schema.Reference{
			Kind: "project",
			UUID: project.Metadata.UUID,
		}
	}

	if len(subnetIdorName) != 0 {
		subnet, err := cli.Client().Subnet.Get(cli.Context, subnetIdorName)
		if err != nil {
			return fmt.Errorf("subnet not found %s", subnetIdorName)
		}
		if subnet.Spec.ClusterReference.UUID != req.Spec.ClusterReference.UUID {
			return fmt.Errorf("subnet not available on provided cluster %s", req.Spec.ClusterReference.UUID)
		}

		if len(req.Spec.Resources.NicList) == 0 {
			req.Spec.Resources.NicList = []*schema.VMNic{
				{
					IsConnected: true,
					SubnetReference: &schema.Reference{
						Kind: "subnet",
					},
				},
			}
		}
		req.Spec.Resources.NicList[0].SubnetReference.UUID = subnet.Metadata.UUID
	}

	return nil
}
