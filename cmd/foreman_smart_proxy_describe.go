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
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	foreman "github.com/simonfuhrer/nutactl/pkg/foreman"
	"github.com/spf13/cobra"
)

func newForemanSmartProxyDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS]",
		Short:                 "describe a smartproxy",
		Aliases:               []string{"d", "des"},
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  cli.wrap(runForemanSmartProxyDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")

	return cmd
}

func runForemanSmartProxyDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	proxy, err := cli.ForemanClient().GetSmartProxy(cli.Context, idOrName)
	if err != nil {
		return err
	}
	proxyList := foreman.QueryResponseSmartProxy{
		Results: []foreman.SmartProxy{*proxy},
	}

	return outputResponse(displayers.ForemanSmartProxies{QueryResponseSmartProxy: proxyList})
}
