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
	"io"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Hosts wraps a nutanix HostListItnent.
type Hosts struct {
	schema.HostListIntent
}

// var _ Displayable = &Projects{}

func (o Hosts) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o Hosts) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o Hosts) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o Hosts) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o Hosts) header() []string {
	return []string{
		"UUID",
		"Name",
		"Model",
		"Serial",
		"Hypervisor IP",
		"CVM IP",
		"Num Cores",
		"Memory",
		"VMs",
		"Cluster",
		"UpdatedAt",
		"CreatedAt",
	}
}

func (o Hosts) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, host := range o.Entities {
		if host.Spec.Name == "" {
			continue
		}
		data[i] = []string{
			host.Metadata.UUID,
			host.Spec.Name,
			host.Status.Resources.Block.BlockModel,
			host.Status.Resources.Block.BlockSerialNumber,
			host.Status.Resources.Hypervisor.IP,
			host.Status.Resources.ControllerVM.IP,
			strconv.Itoa(host.Status.Resources.NumCPUCores),
			humanize.Bytes(host.Status.Resources.MemoryCapacityMib * humanize.MByte),
			strconv.Itoa(host.Status.Resources.Hypervisor.NumVMs),
			host.Status.ClusterReference.Name,
			RenderTime(host.Metadata.LastUpdateTime),
			RenderTime(host.Metadata.CreationTime),
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o Hosts) Text(w io.Writer) error {
	return nil
}
