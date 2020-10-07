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
	"fmt"
	"strings"

	"github.com/prometheus/common/log"
	"github.com/simonfuhrer/nutactl/cmd/displayers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tecbiz-ch/nutanix-go-sdk/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

func newVMListCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "list [FLAGS]",
		Short:                 "List all VM",
		Aliases:               []string{"l", "li"},
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runVMList),
	}
	addOutputFormatFlags(cmd.Flags(), "table")
	flags := cmd.Flags()
	flags.StringP("filter", "f", "", "FIQL filter (e.g. vm_name==x3012.*)")
	flags.StringP("cluster", "c", "", "filter vms by cluster")
	flags.Int64P("limits", "L", 100, "limit objects")
	return cmd
}

func runVMList(cli *CLI, cmd *cobra.Command, args []string) error {
	filter := viper.GetString("filter")
	cluster := viper.GetString("cluster")
	limits := viper.GetInt64("limits")

	var list *schema.VMListIntent
	var err error

	var finalfilter []string
	if filter != "" {
		finalfilter = append(finalfilter, filter)
	}

	if cluster != "" {
		clustername, err := cli.Client().Cluster.Get(cli.Context, cluster)
		if err != nil {
			return err
		}
		finalfilter = append(finalfilter, fmt.Sprintf("cluster=in=%s", clustername.Metadata.UUID))
	}

	if len(finalfilter) > 0 {
		list, err = cli.Client().VM.List(
			cli.Context,
			&schema.DSMetadata{Length: utils.Int64Ptr(limits), Offset: utils.Int64Ptr(0), Filter: strings.Join(finalfilter, ";")},
		)
		if err != nil {
			return err
		}
	} else {
		list, err = cli.Client().VM.List(
			cli.Context,
			&schema.DSMetadata{Length: utils.Int64Ptr(limits), Offset: utils.Int64Ptr(0), SortAttribute: "_created_timestamp_usecs_", SortOrder: "DESCENDING"},
		)
		if err != nil {
			return err
		}

	}
	err = outputResponse(displayers.VMs{VMListIntent: *list})
	if list.Metadata.TotalMatches > list.Metadata.Length {
		log.Warnf("Entities found: %d, Total: %d. use --limits=<Integer> to show more Objects or specified a filtercriteria with --filter", list.Metadata.Length, list.Metadata.TotalMatches)
	}
	return err
}
