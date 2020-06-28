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

	"github.com/simonfuhrer/nutactl/pkg/foreman"
)

// ForemanHosts wraps a foreman ComputeProfiles.
type ForemanComputeProfiles struct {
	foreman.QueryResponseComputeProfile
}

func (o ForemanComputeProfiles) JSON(w io.Writer) error {
	return displayJSON(w, o.Results)
}

func (o ForemanComputeProfiles) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Results)
}

func (o ForemanComputeProfiles) PP(w io.Writer) error {
	return displayPP(w, o.Results)
}

func (o ForemanComputeProfiles) YAML(w io.Writer) error {
	return displayYAML(w, o.Results)
}

func (o ForemanComputeProfiles) header() []string {
	return []string{
		"ID",
		"Name",
		"UpdatedAt",
		"CreatedAt",
	}
}
func (o ForemanComputeProfiles) TableData(w io.Writer) error {
	data := make([][]string, len(o.Results))
	for i, com := range o.Results {
		data[i] = []string{
			fmt.Sprintf("%v", com.ID),
			com.Name,
			com.UpdatedAt,
			com.CreatedAt,
		}
	}
	return displayTable(w, data, o.header())
}

func (o ForemanComputeProfiles) Text(w io.Writer) error {
	return nil
}
