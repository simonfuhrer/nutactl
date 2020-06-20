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

	v2 "github.com/tecbiz-ch/nutanix-go-sdk/schema/v2"
)

// VMSnapshots wraps a nutanix SnapshotList.
type VMSnapshots struct {
	v2.SnapshotList
}

//var _ Displayable = &VMSnapshots{}

func (o VMSnapshots) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o VMSnapshots) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o VMSnapshots) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o VMSnapshots) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o VMSnapshots) header() []string {
	return []string{
		"UUID",
		"Name",
		"CreatedTime",
	}
}

func (o VMSnapshots) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, vmdisk := range o.Entities {
		data[i] = []string{
			vmdisk.UUID,
			vmdisk.Name,
			vmdisk.CreatedTime.String(),
		}
	}
	return displayTable(w, data, o.header())
}

func (o VMSnapshots) Text(w io.Writer) error {
	return nil
}
