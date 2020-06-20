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

	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newCategoryListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List categories",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runCategoryList),
	}
	flags := cmd.Flags()
	flags.Bool("with-values", false, "Display all Values")
	addOutputFormatFlags(flags, "table")

	return cmd
}

func runCategoryList(cli *CLI, cmd *cobra.Command, args []string) error {
	withValues := viper.GetBool("with-values")
	var list *schema.CategoryKeyList
	var err error
	if len(args) == 1 {
		list, err = cli.Client().Category.List(
			cli.Context,
			&schema.DSMetadata{Filter: fmt.Sprintf("name==%s", args[0])},
		)
	} else {
		list, err = cli.Client().Category.All(cli.Context)

	}
	if err != nil {
		return err
	}

	if withValues || len(args) == 1 {
		for _, key := range list.Entities {
			listvalues, err := cli.Client().Category.ListValues(cli.Context, key.Name)
			if err != nil {
				return err
			}
			data := make([]string, len(listvalues.Entities))
			for i, val := range listvalues.Entities {
				data[i] = val.Value
			}
			key.Values = data
		}
	}
	return outputResponse(displayers.Categories{CategoryKeyList: *list})
}
