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
	"sort"
	"strconv"

	"github.com/tecbiz-ch/nutanix-go-sdk/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Tasks wraps a nutanix TaskListIntent.
type Tasks struct {
	schema.TaskListIntent
}

//var _ Displayable = &Tasks{}

func (o Tasks) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o Tasks) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o Tasks) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o Tasks) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o Tasks) header() []string {
	return []string{
		"UUID",
		"Status",
		"OperationType",
		"PercentageComplete",
		"CreationTime",
		"LastUpdateTime",
	}
}

func (o Tasks) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))

	// Workaround sort
	list := o.Entities
	sort.Slice(list, func(i, j int) bool { return list[i].LastUpdateTime.Before(*list[j].LastUpdateTime) })
	// Workaround End

	for i, t := range o.Entities {
		data[i] = []string{
			utils.StringValue(t.UUID),
			utils.StringValue(t.Status),
			utils.StringValue(t.OperationType),
			strconv.FormatInt(*t.PercentageComplete, 10),
			t.CreationTime.String(),
			t.LastUpdateTime.String(),
		}
	}
	return displayTable(w, data, o.header())
}

func (o Tasks) Text(w io.Writer) error {
	return nil
}
