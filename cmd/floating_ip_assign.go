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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newFloatingIpAssignCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "assign [FLAGS] FLOATINGIP",
		Short:                 "Assign a Floating IP to a vm",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runFloatingIpAssign),
	}
	flags := cmd.Flags()
	flags.String("vm", "", "vm name or uuid")
	_ = cmd.MarkFlagRequired("vm")
	return cmd
}

func runFloatingIpAssign(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	vmNameOrId := viper.GetString("vm")
	_, err := createUpdateFloatingIpHelper(idOrName, vmNameOrId, cli)
	return err
}
