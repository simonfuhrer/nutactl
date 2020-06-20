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

	"github.com/simonfuhrer/nutactl/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// AvailabilityZones wraps a nutanix AvailabilityZoneListIntent.
type AvailabilityZones struct {
	schema.AvailabilityZoneListIntent
}

//var _ Displayable = &AvailabilityZones{}

func (o AvailabilityZones) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o AvailabilityZones) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o AvailabilityZones) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o AvailabilityZones) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o AvailabilityZones) header() []string {
	return []string{
		"Name",
		"Type",
	}
}

func (o AvailabilityZones) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, av := range o.Entities {
		data[i] = []string{
			utils.StringValue(av.Spec.Name),
			utils.StringValue(av.Spec.Resources.ManagementPlaneType),
		}
	}
	return displayTable(w, data, o.header())
}

func (o AvailabilityZones) Text(w io.Writer) error {
	return nil
}
