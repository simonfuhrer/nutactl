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

package displayers

import (
	"fmt"
	"io"
	"strconv"

	humanize "github.com/dustin/go-humanize"
	"github.com/tecbiz-ch/nutanix-go-sdk/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Clusters wraps a nutanix ImageListIntent.
type VMs struct {
	schema.VMListIntent
}

//var _ Displayable = &VMs{}

func (o VMs) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o VMs) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o VMs) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o VMs) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o VMs) header() []string {
	return []string{
		"UUID",
		"Name",
		"Power",
		"Project",
		"Subnet",
		"IP",
		"Cluster",
		"MiB",
		"CPU",
		"Disks",
		"Status",
		"UpdatedAt",
		"CreatedAt",
	}
}

func (o VMs) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, vm := range o.Entities {
		subnetName := ""
		ip := ""
		if len(vm.Spec.Resources.NicList) > 0 {
			subnetName = vm.Spec.Resources.NicList[0].SubnetReference.Name
			ip = vm.Status.Resources.NicList[0].IPEndpointList[0].IP
		}
		project := ""
		if vm.Metadata.ProjectReference != nil {
			project = vm.Metadata.ProjectReference.Name
		}
		data[i] = []string{
			vm.Metadata.UUID,
			vm.Spec.Name,
			vm.Spec.Resources.PowerState,
			project,
			subnetName,
			ip,
			vm.Spec.ClusterReference.Name,
			strconv.FormatInt(vm.Spec.Resources.MemorySizeMib, 10),
			fmt.Sprintf("%d/%d", vm.Spec.Resources.NumSockets, vm.Spec.Resources.NumVcpusPerSocket),
			fmt.Sprintf("%d", len(vm.Spec.Resources.DiskList)),
			utils.StringValue(vm.Status.State),
			humanize.Time(*vm.Metadata.LastUpdateTime),
			humanize.Time(*vm.Metadata.CreationTime),
		}
	}
	return displayTable(w, data, o.header())
}

func (o VMs) Text(w io.Writer) error {
	return nil
}
