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

func newContextActiveCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "active [FLAGS]",
		Short:                 "Show active context",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		RunE:                  cli.wrap(runContextActive),
	}
	return cmd
}

func runContextActive(cli *CLI, cmd *cobra.Command, args []string) error {
	fmt.Println(cli.config.ActiveContext)
	return nil
}
