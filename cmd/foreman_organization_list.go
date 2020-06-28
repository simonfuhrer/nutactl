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
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/simonfuhrer/nutactl/pkg/foreman"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newForemanOrganizationListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Aliases:               []string{"l", "li"},
		Short:                 "List organizations",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runForemanOrganizationList),
	}
	flags := cmd.Flags()
	flags.StringP("filter", "f", "", "Foreman search filter (e.g. os ~ windows and environment == test_win)")
	addOutputFormatFlags(flags, "table")

	return cmd
}

func runForemanOrganizationList(cli *CLI, cmd *cobra.Command, args []string) error {
	filter := viper.GetString("filter")
	var organizations *foreman.QueryResponseOrganization
	var err error
	if filter != "" {
		organizations, err = cli.ForemanClient().SearchOrganization(cli.Context, filter)
		if err != nil {
			return err
		}
	} else {
		organizations, err = cli.ForemanClient().ListOrganization(cli.Context)
		if err != nil {
			return err
		}
	}
	return outputResponse(displayers.ForemanOrganizations{QueryResponseOrganization: *organizations})
}
