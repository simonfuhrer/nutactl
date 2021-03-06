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
	"strings"

	"github.com/spf13/cobra"
)

func newContextCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS] NAME",
		Short:                 "create a context",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		RunE:                  cli.wrap(runContextCreate),
	}
	return cmd
}

func runContextCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := strings.TrimSpace(args[0])
	if name == "" {
		return fmt.Errorf("emtpy name not allowed")
	}
	if cli.config.ContextByName(name) != nil {
		return fmt.Errorf("name %s already used", name)
	}
	context := &Context{Name: name}
	err := createContext(context, cli.config.Contexts)
	if err != nil {
		return err
	}
	fmt.Printf("Context %s created and activated\n", name)
	return nil
}
