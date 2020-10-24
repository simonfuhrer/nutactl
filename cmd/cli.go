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
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	logrus "github.com/sirupsen/logrus"
	"github.com/tcnksm/go-input"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nutanix "github.com/tecbiz-ch/nutanix-go-sdk"
)

// PollIntervalinSeconds ..
const (
	PollIntervalinSeconds       = 2
	appName                     = "NUTACTL"
	itemsPerPage          int64 = 40
)

// CLI ...
type CLI struct {
	Endpoint            string
	Context             context.Context
	RootCommand         *cobra.Command
	client              *nutanix.Client
	millisecondsPerPoll time.Duration
	config              *Config
	Plugin              bool
}

// NewCLI ...
func NewCLI(plugin bool) *CLI {
	viper.SetEnvPrefix(appName)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.SetConfigType("yaml")
	cli := &CLI{
		Context:             context.Background(),
		millisecondsPerPoll: 1000 * PollIntervalinSeconds,
		Plugin:              plugin,
	}
	cli.RootCommand = NewRootCommand(cli)

	return cli
}

// wrapper func to bind all flags with viper and ensure a logout is perfomed
func (c *CLI) wrap(f func(*CLI, *cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		BindAllFlags(cmd)
		err := f(c, cmd, args)
		return checkErr(err)
	}
}

// Client cli ...
func (c *CLI) Client() *nutanix.Client {
	if c.client == nil {
		context := c.config.ContextByName(c.config.ActiveContext)
		configCreds := nutanix.Credentials{
			Username: context.User,
			Password: context.Password,
		}

		opts := []nutanix.ClientOption{
			nutanix.WithCredentials(&configCreds),
			nutanix.WithEndpoint(context.Endpoint),
		}
		if viper.GetString("log-level") == "trace" {
			opts = append(opts, nutanix.WithDebugWriter(os.Stdout))
		}
		if viper.GetBool("insecure") {
			opts = append(opts, nutanix.WithSkipVerify())
		}

		logrus.Debugf("creating Nutanix Client")

		c.client = nutanix.NewClient(opts...)
	}
	return c.client
}

// WaitTask ...
func (c *CLI) WaitTask(ctx context.Context, taskUUID string, timeoutSeconds int) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	ticker := time.NewTicker(c.millisecondsPerPoll * time.Millisecond)
	defer ticker.Stop()
	s := spinner.New(spinner.CharSets[6], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Waiting Task ID %s", taskUUID)
	s.Start()
	for {
		select {
		case <-ticker.C:
			task, err := c.Client().Task.GetByUUID(ctx, taskUUID)
			if err != nil {
				return err
			}
			logrus.Debugf(*task.Status)
			switch *task.Status {
			case "SUCCEEDED":
				s.Stop()
				return nil
			case "FAILED":
				s.Stop()
				return fmt.Errorf(*task.ErrorDetail)

			}
		case <-ctx.Done():
			s.Stop()
			return fmt.Errorf("error waiting for task to be completed: %s", ctx.Err())
		}
	}
}

func (c *CLI) ensureContext(cmd *cobra.Command, args []string) error {
	context := c.config.ContextByName(c.config.ActiveContext)
	if c.config.ActiveContext == "" || context == nil {
		if context == nil {
			ui := &input.UI{
				Writer: os.Stdout,
				Reader: os.Stdin,
			}

			contextname, err := ui.Ask("Enter context name", &input.Options{
				Default:   "",
				Required:  true,
				Loop:      true,
				HideOrder: true,
			})
			if err != nil {
				logrus.Fatalln(err.Error())
			}

			newcontext := &Context{Name: contextname}
			err = createContext(newcontext, c.config.Contexts)
			if err != nil {
				return err
			}

			context = newcontext
		}
	}

	if context.Endpoint == "" {
		return fmt.Errorf("missing endpoint in config %s", viper.ConfigFileUsed())
	}
	if context.User == "" {
		return fmt.Errorf("missing user in config %s", viper.ConfigFileUsed())
	}
	if context.Password == "" {
		return fmt.Errorf("missing password in config %s", viper.ConfigFileUsed())
	}

	return nil
}

func (c *CLI) ReadConfig() {
	cfgLogJSON := viper.GetBool("log-json")
	if cfgLogJSON {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	cfgLogLevel := viper.GetString("log-level")
	logrusLogLevel, err := logrus.ParseLevel(cfgLogLevel)
	if err == nil {
		logrus.SetLevel(logrusLogLevel)
	}
	logrus.Debugf("logger initialized: loglevel %s", cfgLogLevel)
	cfgFile := viper.GetString("config")
	if cfgFile != "" {
		dir, file := path.Split(cfgFile)
		viper.AddConfigPath(dir)
		viper.SetConfigName(file)
	} else {
		// Search config in home directory with name ".nutactl" (without extension).
		viper.AddConfigPath(DefaultConfigPath)
		viper.SetConfigName(".nutactl")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("Using config file: ", viper.ConfigFileUsed())
	}

	if !fileExists(viper.ConfigFileUsed()) {
		err := viper.SafeWriteConfig()
		if err != nil {
			logrus.Fatalf("%s", err)
		}
	}
	err = viper.Unmarshal(&c.config)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

}
