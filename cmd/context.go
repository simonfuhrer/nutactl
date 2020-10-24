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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-input"
)

func newContextCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "context",
		Aliases:               []string{"con", "co"},
		Short:                 "Manage contexts",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runContext),
	}
	cmd.AddCommand(
		newContextListCommand(cli),
		newContextActiveCommand(cli),
		newContextUseCommand(cli),
		newContextCreateCommand(cli),
		newContextDeleteCommand(cli),
	)
	cmd.Flags().SortFlags = false
	return cmd
}

func runContext(cli *CLI, cmd *cobra.Command, args []string) error {
	return cmd.Usage()
}

func createContext(newcontext *Context, contexts []*Context) error {
	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	endpoint, err := ui.Ask("Enter Prismcentral (PC) Endpoint", &input.Options{
		Default:   "",
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}
	newcontext.Endpoint = endpoint

	username, err := ui.Ask("Username", &input.Options{
		Default:   "admin",
		Required:  true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}
	newcontext.User = username

	password, err := ui.Ask("Password", &input.Options{
		Default:   "",
		Required:  true,
		Mask:      true,
		Loop:      true,
		HideOrder: true,
	})
	if err != nil {
		return err
	}
	newcontext.Password = password
	contexts = append(contexts, newcontext)
	viper.Set("contexts", contexts)
	viper.Set("active_context", newcontext.Name)
	return viper.WriteConfig()
}
