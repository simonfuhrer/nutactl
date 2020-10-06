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
	"github.com/spf13/cobra"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newClusterDescribeCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "describe [FLAGS] CLUSTER",
		Short:                 "Describe an cluster",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runClusterDescribe),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "json")

	return cmd
}

func runClusterDescribe(cli *CLI, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	cluster, err := cli.Client().Cluster.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}

	list := schema.ClusterListIntent{
		Entities: []*schema.ClusterIntent{cluster},
	}

	return outputResponse(displayers.Clusters{ClusterListIntent: list})
}
