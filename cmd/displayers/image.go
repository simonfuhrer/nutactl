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

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Images wraps a nutanix ImageListIntent.
type Images struct {
	schema.ImageListIntent
}

// var _ Displayable = &Images{}

func (o Images) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o Images) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o Images) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o Images) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o Images) header() []string {
	return []string{
		"UUID",
		"Name",
		"Description",
		"Arch",
		"ImageType",
		"MB",
		"Status",
	}
}
func (o Images) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, image := range o.Entities {
		var size string
		if image.Status.Resources.SizeBytes != 0 {
			size = strconv.FormatInt(image.Status.Resources.SizeBytes/1024/1024, 10)
		}
		data[i] = []string{
			image.Metadata.UUID,
			image.Spec.Name,
			image.Spec.Description,
			image.Spec.Resources.Architecture,
			image.Spec.Resources.ImageType,
			size,
			image.Status.State,
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o Images) Text(w io.Writer) error {
	return nil
}
