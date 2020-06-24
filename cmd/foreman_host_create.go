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
	"github.com/k0kubun/pp"
	foreman "github.com/simonfuhrer/nutactl/pkg/foreman"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newForemanHostCreateCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "create [FLAGS]",
		Short:                 "create a host",
		Aliases:               []string{"cre", "c"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runForemanHostCreate),
	}
	flags := cmd.Flags()
	flags.String("mac", "", "mac address")
	flags.String("domain", "", "domain name or id")
	flags.String("os", "", "os name or id")
	MarkFlagsRequired(cmd, "domain", "os")
	return cmd
}

func runForemanHostCreate(cli *CLI, cmd *cobra.Command, args []string) error {
	name := args[0]
	mac := viper.GetString("mac")
	domainIDOrName := viper.GetString("domain")
	osIDOrName := viper.GetString("os")

	domain, err := cli.ForemanClient().GetDomain(cli.Context, domainIDOrName)
	if err != nil {
		return err
	}

	os, err := cli.ForemanClient().GetOperatingSystem(cli.Context, osIDOrName)
	if err != nil {
		return err
	}

	if len(mac) == 0 {
		mac = GenerateMac().String()
	}

	request := foreman.HostRequest{
		Host: foreman.NewHostData{
			ForemanObject: foreman.ForemanObject{
				Name: name,
			},
			Mac:               mac,
			DomainID:          domain.ID,
			OperatingsystemID: os.ID,
		},
	}

	host, err := cli.ForemanClient().CreateHost(cli.Context, &request)
	if err != nil {
		return err
	}
	pp.Println(host)
	return nil
}
