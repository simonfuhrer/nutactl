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

func newForemanDomainCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS]",
		Short:                 "create a domain",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runForemanDomainCreate),
	}
	return cmd
}

func runForemanDomainCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]

	request := foreman.DomainRequest{
		Domain: foreman.NewDomainData{
			ForemanObject: foreman.ForemanObject{
				Name: name,
			},
		},
	}

	domain, err := cli.ForemanClient().CreateDomain(cli.Context, &request)
	if err != nil {
		return err
	}

	fmt.Printf("Domain %s with ID %d created\n", domain.Name, domain.ID)

	return nil
}
