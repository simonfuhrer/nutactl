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

func newForemanHostUpdateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "update [FLAGS] HOST",
		Short:                 "Update an host",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runForemanHostUpdate),
	}
	flags := cmd.Flags()
	flags.StringP("name", "n", "", "New host name")
	flags.Bool("enable-netbackup", false, "link netbackup puppetclass to host")
	return cmd
}

func runForemanHostUpdate(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	_ = viper.GetString("name")
	netbackup := viper.GetBool("enable-netbackup")

	host, err := cli.ForemanClient().GetHost(cli.Context, idOrName)
	if err != nil {
		return err
	}

	os, err := cli.ForemanClient().GetOperatingSystemByID(cli.Context, host.OperatingsystemID)
	if err != nil {
		return err
	}

	if netbackup {
		err := enableNetbackup(cli, os, host.ID)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Host %s with ID %d updated\n", host.Name, host.ID)

	return nil
}
