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
	"git.atilf.fr/atilf/portainer-cli/cmd/util"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:                   "nutactl",
	Short:                 "nutanix prism central CLI",
	Long:                  "A command-line interface for nutanix prism central",
	TraverseChildren:      false,
	SilenceUsage:          false,
	SilenceErrors:         true,
	DisableFlagsInUseLine: true,
}

// NewRootCommand ...
func NewRootCommand(cli *CLI) *cobra.Command {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(
		newVMCommand(cli),
		newImageCommand(cli),
		newClusterCommand(cli),
		newProjectCommand(cli),
		newSubnetCommand(cli),
		newAvailabilityZoneCommand(cli),
		newCategoryCommand(cli),
		newTaskCommand(cli),
		newForemanCommand(cli),
		newVersionCommand(cli),
		newCompletionCommand(cli),
	)

	rootCmd.Flags().SortFlags = false
	flags := rootCmd.PersistentFlags()
	flags.StringP("api-url", "a", "", "Nutanix PC API URL [NUTACTL_API_URL]")
	flags.StringP("foreman-api-url", "f", "", "Foreman API URL [NUTACTL_FOREMAN_API_URL]")
	flags.StringP("username", "u", "", "Nutanix username [NUTACTL_USERNAME]")
	flags.StringP("password", "p", "", "Nutanix password [NUTACTL_PASSWORD]")
	flags.BoolP("insecure", "", false, "Accept insecure TLS certificates")
	flags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nutactl.yaml)")
	flags.StringP("log-level", "", logrus.InfoLevel.String(), "log level (trace,debug,info,warn/warning,error,fatal,panic)")
	flags.BoolP("log-json", "", false, "log as json")

	BindAllFlags(rootCmd)
	MarkFlagsRequired(rootCmd, "api-url", "username", "password", "foreman-api-url")

	return rootCmd
}
func initConfig() {
	if viper.GetBool("log-json") {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	logLevel := viper.GetString("log-level")
	logrusLogLevel, err := logrus.ParseLevel(logLevel)
	if err == nil {
		logrus.SetLevel(logrusLogLevel)
	}
	logrus.Debugf("logger initialized: loglevel %s", logLevel)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		util.HandleError(err)

		// Search config in home directory with name ".nutactl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".nutactl")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("Using config file:", viper.ConfigFileUsed())
	}
}
