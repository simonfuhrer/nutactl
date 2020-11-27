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
	"github.com/tecbiz-ch/nutanix-go-sdk/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newHostListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List hosts",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runHostList),
	}
	flags := cmd.Flags()
	addOutputFormatFlags(flags, "table")

	return cmd
}

func runHostList(cli *CLI, cmd *cobra.Command, args []string) error {
	opts := &schema.DSMetadata{
		Offset:        utils.Int64Ptr(0),
		Length:        utils.Int64Ptr(itemsPerPage),
		SortOrder:     "ASCENDING",
		SortAttribute: "Name",
	}
	var list schema.HostListIntent
	f := func(opts *schema.DSMetadata) (interface{}, error) {
		list, err := cli.Client().Host.List(
			cli.Context,
			opts,
		)
		return list, err
	}
	responses, err := paginateResp(f, opts)
	if err != nil {
		return err
	}
	for _, response := range responses {
		item := response.(*schema.HostListIntent)
		list.Entities = append(list.Entities, item.Entities...)
	}

	clusters, err := cli.client.Cluster.All(cli.Context)
	if err != nil {
		return err
	}
	data := make(map[string]string)
	for _, cluster := range clusters.Entities {
		data[cluster.Metadata.UUID] = cluster.Spec.Name
	}

	for _, host := range list.Entities {
		if host.Status.ClusterReference.UUID != "" {
			host.Status.ClusterReference.Name = data[host.Status.ClusterReference.UUID]
		}
	}

	return outputResponse(displayers.Hosts{HostListIntent: list})
}
