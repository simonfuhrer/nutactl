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

	"github.com/simonfuhrer/nutactl/pkg/foreman"
	"github.com/spf13/cobra"
)

func newForemanHostCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "host",
		Short:                 "Manage hosts",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runForemanHost),
	}
	cmd.AddCommand(
		newForemanHostListCommand(cli),
		newForemanHostDescribeCommand(cli),
		newForemanHostCreateCommand(cli),
		newForemanHostUpdateCommand(cli),
		newForemanHostDeleteCommand(cli),
	)
	cmd.Flags().SortFlags = false
	return cmd
}

func runForemanHost(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func enableNetbackup(cli *CLI, os *foreman.OperatingSystem, hostID int) error {
	filterClass := "puppetclass_name==bi_backup::nbu"
	if os.Family == "Windows" {
		filterClass = "puppetclass_name==bi_databackup"
	}
	smartClasses, err := cli.foremanclient.SearchSmartClassParameter(cli.Context, filterClass)
	if err != nil {
		return err
	}
	if len(smartClasses.Results) == 0 {
		return fmt.Errorf("PuppetSmartClass not found: %s", filterClass)
	}
	_, err = cli.foremanclient.AddPuppetClassToHost(cli.Context, hostID, smartClasses.Results[0].PuppetclassID)
	if err != nil {
		return err
	}
	return nil
}
