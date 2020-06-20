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
		newCategoryCommand(cli),
		newTaskCommand(cli),
		newVersionCommand(cli),
		newCompletionCommand(cli),
	)

	rootCmd.Flags().SortFlags = false
	flags := rootCmd.PersistentFlags()
	flags.StringP("api-url", "a", "", "Nutanix PC Api URL [NUTACTL_API_URL]")
	flags.StringP("username", "u", "", "Nutanix username [NUTACTL_USERNAME]")
	flags.StringP("password", "p", "", "Nutanix password [NUTACTL_PASSWORD]")
	flags.BoolP("insecure", "", false, "Accept insecure TLS certificates")
	flags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nutactl.yaml)")
	flags.StringP("log-level", "", logrus.InfoLevel.String(), "log level (trace,debug,info,warn/warning,error,fatal,panic)")
	flags.BoolP("log-json", "", false, "log as json")

	BindAllFlags(rootCmd)
	MarkFlagsRequired(rootCmd, "api-url", "username", "password")

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
