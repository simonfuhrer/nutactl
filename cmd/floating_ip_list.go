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
	"github.com/spf13/viper"
	"github.com/tecbiz-ch/nutanix-go-sdk/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newFloatingIpListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List Floating IPs",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runFloatingIpList),
	}
	flags := cmd.Flags()
	flags.StringP("filter", "f", "", "FIQL filter  (e.g.)")

	addOutputFormatFlags(flags, "table")

	return cmd
}

func runFloatingIpList(cli *CLI, cmd *cobra.Command, args []string) error {
	filter := viper.GetString("filter")

	opts := &schema.DSMetadata{Offset: utils.Int64Ptr(0), Length: utils.Int64Ptr(itemsPerPage)}
	var list schema.FloatingIPListIntent

	if filter != "" {
		opts.Filter = filter
	}

	f := func(opts *schema.DSMetadata) (interface{}, error) {
		list, err := cli.Client().FlotatingIP.List(
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
		item := response.(*schema.FloatingIPListIntent)
		list.Entities = append(list.Entities, item.Entities...)

	}

	cache := make(map[string]string)
	for _, f := range list.Entities {
		_, ok := cache[f.Status.Resources.ExternalSubnetReference.UUID]
		if !ok {
			subnet, err := cli.Client().Subnet.GetByUUID(cli.Context, f.Status.Resources.ExternalSubnetReference.UUID)
			if err != nil {
				return err
			}
			cache[f.Status.Resources.ExternalSubnetReference.UUID] = subnet.Spec.Name
		}
		f.Status.Resources.ExternalSubnetReference.Name = cache[f.Status.Resources.ExternalSubnetReference.UUID]
	}

	return outputResponse(displayers.FloatingIps{FloatingIPListIntent: list})
}
