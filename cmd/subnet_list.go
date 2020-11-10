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

func newSubnetListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List subnets",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runSubnetList),
	}
	flags := cmd.Flags()
	flags.StringP("filter", "f", "", "FIQL filter  (e.g. vlan_id==2711;cluster_name==mycluster)")

	addOutputFormatFlags(flags, "table")
	return cmd
}

func runSubnetList(cli *CLI, cmd *cobra.Command, args []string) error {
	filter := viper.GetString("filter")

	opts := &schema.DSMetadata{Offset: utils.Int64Ptr(0), Length: utils.Int64Ptr(itemsPerPage)}
	var list schema.SubnetListIntent
	if filter != "" {
		opts.Filter = filter
	}

	f := func(opts *schema.DSMetadata) (interface{}, error) {
		list, err := cli.Client().Subnet.List(
			cli.Context,
			opts,
		)
		return list, err
	}
	channelresponse, err := paginateResp(f, opts)
	if err != nil {
		return err
	}
	for response := range channelresponse {
		item := response.(*schema.SubnetListIntent)
		list.Entities = append(list.Entities, item.Entities...)
	}

	return outputResponse(displayers.Subnets{SubnetListIntent: list})
}
