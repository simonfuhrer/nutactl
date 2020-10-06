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

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Projects wraps a nutanix ProjectListItnent.
type Projects struct {
	schema.ProjectListIntent
}

//var _ Displayable = &Projects{}

func (o Projects) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o Projects) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o Projects) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o Projects) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o Projects) header() []string {
	return []string{
		"UUID",
		"Name",
		"Description",
	}
}

func (o Projects) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, project := range o.Entities {
		data[i] = []string{
			project.Metadata.UUID,
			project.Spec.Name,
			project.Spec.Description,
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o Projects) Text(w io.Writer) error {
	return nil
}
