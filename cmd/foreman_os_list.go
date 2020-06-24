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
	"log"

	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
)

func newForemanOSListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Aliases:               []string{"l", "li"},
		Short:                 "List operating systems",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runForemanOSList),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")

	return cmd
}

func runForemanOSList(cli *CLI, cmd *cobra.Command, args []string) error {
	lists, err := cli.ForemanClient().ListOperatingSystem(cli.Context)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return outputResponse(displayers.ForemanOperatingSystems{QueryResponseOperatingSystem: *lists})
}
