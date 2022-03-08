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

	"github.com/dustin/go-humanize"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Vpcs wraps a nutanix VpcListIntent.
type FloatingIps struct {
	schema.FloatingIPListIntent
}

// var _ Displayable = &Projects{}

func (o FloatingIps) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o FloatingIps) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o FloatingIps) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o FloatingIps) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o FloatingIps) header() []string {
	return []string{
		"UUID",
		"IP",
		"Subnet",
		"VM Nic",
		"UpdatedAt",
		"CreatedAt",
	}
}

func (o FloatingIps) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, fip := range o.Entities {
		vm := ""
		if fip.Status.Resources.VMNicReference != nil {
			vm = fip.Status.Resources.VMNicReference.UUID
		}
		data[i] = []string{
			fip.Metadata.UUID,
			fip.Status.Resources.FloatingIP,
			fip.Status.Resources.ExternalSubnetReference.Name,
			vm,
			humanize.Time(*fip.Metadata.LastUpdateTime),
			humanize.Time(*fip.Metadata.CreationTime),
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o FloatingIps) Text(w io.Writer) error {
	return nil
}
