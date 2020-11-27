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
		DisableAutoGenTag:     true,
		PreRunE:               cli.ensureContext,
		RunE:                  cli.wrap(runVMList),
	}
	addOutputFormatFlags(cmd.Flags(), "table")
	flags := cmd.Flags()
	flags.StringP("filter", "f", "", "FIQL filter (e.g. vm_name==srv.*)")
	flags.StringP("cluster", "c", "", "filter vms by cluster")
	return cmd
}

func runVMList(cli *CLI, cmd *cobra.Command, args []string) error {
	filter := viper.GetString("filter")
	cluster := viper.GetString("cluster")

	var list schema.VMListIntent
	var err error
	opts := &schema.DSMetadata{Offset: utils.Int64Ptr(0), Length: utils.Int64Ptr(itemsPerPage)}

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
		opts.Filter = strings.Join(finalfilter, ";")
	}

	f := func(opts *schema.DSMetadata) (interface{}, error) {
		list, err := cli.Client().VM.List(
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
		item := response.(*schema.VMListIntent)
		list.Entities = append(list.Entities, item.Entities...)
	}

	hosts, err := cli.Client().Host.All(cli.Context)
	if err != nil {
		return err
	}

	m := make(map[string]string)
	for _, h := range hosts.Entities {
		m[h.Metadata.UUID] = h.Spec.Name
	}
	for _, vm := range list.Entities {
		if vm.Status.Resources.HostReference != nil {
			vm.Status.Resources.HostReference.Name = m[vm.Status.Resources.HostReference.UUID]
		}
	}

	return outputResponse(displayers.VMs{VMListIntent: list})
}
