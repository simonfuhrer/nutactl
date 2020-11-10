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

func newContextDeleteCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "delete [FLAGS] CONTEXT",
		Short:                 "Delete a context",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		RunE:                  cli.wrap(runContextDelete),
	}
	return cmd
}

func runContextDelete(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	context := cli.config.ContextByName(name)
	if context == nil {
		return fmt.Errorf("context not found: %v", name)
	}
	cli.config.RemoveContext(context)
	viper.Set("contexts", cli.config.Contexts)
	activeContext := ""
	if len(cli.config.Contexts) > 0 {
		activeContext = cli.config.Contexts[0].Name
	}
	viper.Set("active_context", activeContext)
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	fmt.Printf("Context %v deleted\n", name)
	return nil
}
