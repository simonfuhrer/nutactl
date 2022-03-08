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

func newVpcUpdateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "update [FLAGS] VPC",
		Short:                 "Update a VPC",
		Aliases:               []string{"upd", "up"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runVpcUpdate),
	}
	flags := cmd.Flags()
	flags.StringP("external-subnet", "e", "", "external subnet uuid or name")
	flags.StringP("name", "n", "", "new vpc name")

	return cmd
}

func runVpcUpdate(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	subnetNameorUuid := viper.GetString("external-subnet")
	newName := viper.GetString("name")

	vpc, err := cli.Client().VPC.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	if len(subnetNameorUuid) > 0 {
		subnet, err := cli.Client().Subnet.Get(cli.Context, subnetNameorUuid)
		if err != nil {
			return err
		}

		vpc.Spec.Resources.ExternalSubnetList[0].ExternalSubnetReference.UUID = subnet.Metadata.UUID

	}

	if len(newName) > 0 {
		vpc.Spec.Name = newName
	}
	vpc.Metadata.Kind = "vpc"
	result, err := cli.Client().VPC.Update(cli.Context, vpc)
	if err != nil {
		return err
	}
	fmt.Printf("VPC %s with uuid %s updated\n", result.Spec.Name, result.Metadata.UUID)

	return nil
}
