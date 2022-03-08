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
)

func newFloatingIpDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete [FLAGS] FLOATING IP",
		Short:                 "Delete a Floating IP",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runFloatingIpDelete),
	}

	return cmd
}

func runFloatingIpDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	fip, err := cli.Client().FlotatingIP.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}
	if askForConfirm(fmt.Sprintf("Delete %s ?", fip.Status.Resources.FloatingIP)) == nil {
		err = cli.client.FlotatingIP.Delete(cli.Context, fip.Metadata.UUID)
		if err != nil {
			return err
		}
		fmt.Printf("Floating IP %v deleted\n", idOrName)
		return nil
	}
	return fmt.Errorf("operation aborted")
}
