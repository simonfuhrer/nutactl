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
	"strings"

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Vpcs wraps a nutanix VpcListIntent.
type Vpcs struct {
	schema.VpcListIntent
}

// var _ Displayable = &Projects{}

func (o Vpcs) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o Vpcs) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o Vpcs) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o Vpcs) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o Vpcs) header() []string {
	return []string{
		"UUID",
		"Name",
		"Gateway Address",
		"External Subnet",
		"UpdatedAt",
		"CreatedAt",
		"Status",
	}
}

func (o Vpcs) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, vpc := range o.Entities {
		subnet := []string{}
		ip := []string{}
		for _, v := range vpc.Status.Resources.ExternalSubnetList {
			ip = append(ip, v.ExternalIPList...)
			subnet = append(subnet, v.ExternalSubnetReference.Name)
		}

		data[i] = []string{
			vpc.Metadata.UUID,
			vpc.Spec.Name,
			strings.Join(ip, ","),
			strings.Join(subnet, ","),
			RenderTime(vpc.Metadata.LastUpdateTime),
			RenderTime(vpc.Metadata.CreationTime),
			vpc.Status.State,
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o Vpcs) Text(w io.Writer) error {
	return nil
}
