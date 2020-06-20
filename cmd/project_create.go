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
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newProjectCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] projectname",
		Short:                 "Create an project",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runProjectCreate),
	}
	flags := cmd.Flags()
	addProjectFlags(flags)

	return cmd
}

func runProjectCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]

	req := &schema.ProjectIntent{
		Spec: &schema.Project{
			Resources: &schema.ProjectResources{},
		},
		Metadata: &schema.Metadata{
			Kind: "project",
		},
	}

	err := createUpdateProjectHelper(name, req)
	if err != nil {
		return err
	}

	result, err := cli.Client().Project.Create(cli.Context, req)
	if err != nil {
		return err
	}
	fmt.Printf("Project %s with uuid %s created\n", result.Spec.Name, result.Metadata.UUID)

	return nil
}
