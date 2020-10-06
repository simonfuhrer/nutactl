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
	"strings"
	"time"

	"github.com/briandowns/spinner"
	logrus "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	nutanix "github.com/tecbiz-ch/nutanix-go-sdk"
)

// PollIntervalinSeconds ..
const (
	PollIntervalinSeconds = 2
	appName               = "NUTACTL"
)

// CLI ...
type CLI struct {
	Endpoint            string
	Context             context.Context
	RootCommand         *cobra.Command
	client              *nutanix.Client
	millisecondsPerPoll time.Duration
	clusters            map[string]string
	config              *Config
}

// NewCLI ...
func NewCLI() *CLI {
	viper.SetEnvPrefix(appName)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.SetConfigType("yaml")
	initConfig()
	cli := &CLI{
		Context:             context.Background(),
		millisecondsPerPoll: 1000 * PollIntervalinSeconds,
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
		if viper.GetBool("insecure") {
			opts = append(opts, nutanix.WithSkipVerify())
		}

		logrus.Debugf("creating Nutanix Client")
		c.client = nutanix.NewClient(opts...)
	}
	return c.client
}

func (c *CLI) ensureContext(cmd *cobra.Command, args []string) error {
	if c.config.ActiveContext == "" || c.config.ContextByName(c.config.ActiveContext) == nil {
		return fmt.Errorf("no active context or context does not exists")
	}
	return nil
}

// InitAllClusters ...
func (c *CLI) InitAllClusters() error {
	logrus.Debugf("init Nutanix Clusters")
	if c.clusters == nil {
		clusters, err := c.client.Cluster.All(context.Background())
		if err != nil {
			return err
		}
		data := make(map[string]string)
		for _, cluster := range clusters.Entities {
			data[cluster.Metadata.UUID] = cluster.Spec.Name
		}
		c.clusters = data
	}
	return nil
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
