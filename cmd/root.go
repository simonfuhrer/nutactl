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
	logrus "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "nutactl",
	Short:                 "nutanix prism central CLI",
	Long:                  "A command-line interface for nutanix prism central",
	TraverseChildren:      false,
	SilenceUsage:          false,
	SilenceErrors:         true,
	DisableAutoGenTag:     true,
	DisableFlagsInUseLine: true,
}

// NewRootCommand ...
func NewRootCommand(cli *CLI) *cobra.Command {
	cobra.OnInitialize(cli.readConfig)

	rootCmd.AddCommand(
		newVMCommand(cli),
		newImageCommand(cli),
		newClusterCommand(cli),
		newHostCommand(cli),
		newProjectCommand(cli),
		newSubnetCommand(cli),
		newAvailabilityZoneCommand(cli),
		newCategoryCommand(cli),
		newContextCommand(cli),
		newTaskCommand(cli),
		newVersionCommand(cli),
		newCompletionCommand(cli),
		newDocCommand(cli),
	)

	rootCmd.Flags().SortFlags = false
	if !cli.Plugin {
		flags := rootCmd.PersistentFlags()
		flags.BoolP("insecure", "", false, "Accept insecure TLS certificates")
		flags.StringP("config", "", "", "config file to use (default $HOME/.nutactl.yaml)")
		flags.StringP("log-level", "", logrus.InfoLevel.String(), "log level (trace,debug,info,warn/warning,error,fatal,panic)")
		flags.BoolP("log-json", "", false, "log as json")
		BindAllFlags(rootCmd)
	}
	return rootCmd
}
