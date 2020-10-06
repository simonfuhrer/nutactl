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

	v2 "github.com/tecbiz-ch/nutanix-go-sdk/schema/v2"
)

// VirtualDisks wraps a nutanix VirtualDiskList.
type VirtualDisks struct {
	v2.VirtualDiskList
}

// var _ Displayable = &VirtualDisks{}

func (o VirtualDisks) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o VirtualDisks) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o VirtualDisks) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o VirtualDisks) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o VirtualDisks) header() []string {
	return []string{
		"UUID",
		"VM",
		"SizeinMB",
		"Cluster",
	}
}

func (o VirtualDisks) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, vmdisk := range o.Entities {
		data[i] = []string{
			vmdisk.UUID,
			vmdisk.AttachedVMName,
			strconv.FormatInt(vmdisk.DiskCapacityInBytes/1024/1024, 10),
			vmdisk.ClusterUUID,
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o VirtualDisks) Text(w io.Writer) error {
	return nil
}
