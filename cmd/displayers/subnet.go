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
	"strings"

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Subnets wraps a nutanix SubnetListIntent.
type Subnets struct {
	schema.SubnetListIntent
}

// var _ Displayable = &Subnets{}

func (o Subnets) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o Subnets) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o Subnets) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o Subnets) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o Subnets) header() []string {
	return []string{
		"UUID",
		"Name",
		"Description",
		"Cluster",
		"Type",
		"VLAN",
		"SubnetIP",
		"DHCPPOOL",
		"Status",
	}
}

func (o Subnets) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, subnet := range o.Entities {
		subnetIP := ""
		dhcpPool := ""
		if subnet.Spec.Resources.IPConfig != nil {
			subnetIP = fmt.Sprintf("%s/%d", subnet.Spec.Resources.IPConfig.SubnetIP, subnet.Spec.Resources.IPConfig.PrefixLength)
			if subnet.Spec.Resources.IPConfig.PoolList != nil {
				strs := make([]string, len(subnet.Spec.Resources.IPConfig.PoolList))
				for i, v := range subnet.Spec.Resources.IPConfig.PoolList {
					strs[i] = strings.ReplaceAll(v.Range, " ", "-")
				}
				dhcpPool = strings.Join(strs, ",")
			}
		}

		data[i] = []string{
			subnet.Metadata.UUID,
			subnet.Spec.Name,
			subnet.Spec.Description,
			subnet.Spec.ClusterReference.Name,
			subnet.Spec.Resources.SubnetType,
			strconv.FormatInt(*subnet.Spec.Resources.VlanID, 10),
			subnetIP,
			dhcpPool,
			subnet.Status.State,
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o Subnets) Text(w io.Writer) error {
	for _, subnet := range o.Entities {
		fmt.Printf("UUID:\t\t%s\n", subnet.Metadata.UUID)
		fmt.Printf("Name:\t\t%s\n", subnet.Spec.Name)
	}
	return nil
}
