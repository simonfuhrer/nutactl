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
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newImageCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "image",
		Short:                 "Manage images",
		Aliases:               []string{"i", "img", "ima"},
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		RunE:                  cli.wrap(runImage),
	}
	cmd.AddCommand(
		newImageListCommand(cli),
		newImageDescribeCommand(cli),
		newImageDeleteCommand(cli),
		newImageUpdateCommand(cli),
		newImageCreateCommand(cli),
	)
	cmd.Flags().SortFlags = false
	return cmd
}

func runImage(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func createUpdateImageHelper(name string, req *schema.ImageIntent) error {
	description := viper.GetString("description")

	if len(name) != 0 {
		req.Spec.Name = name
	}
	if len(description) != 0 {
		req.Spec.Description = description
	}

	return nil
}

func addImageFlags(flags *pflag.FlagSet) {
	flags.StringP("description", "d", "", "Description")
}
