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

	humanize "github.com/dustin/go-humanize"
	"github.com/simonfuhrer/nutactl/pkg/foreman"
)

// ForemanHosts wraps a foreman Organisation.
type ForemanOrganizations struct {
	foreman.QueryResponseOrganization
}

func (o ForemanOrganizations) JSON(w io.Writer) error {
	return displayJSON(w, o.Results)
}

func (o ForemanOrganizations) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Results)
}

func (o ForemanOrganizations) PP(w io.Writer) error {
	return displayPP(w, o.Results)
}

func (o ForemanOrganizations) YAML(w io.Writer) error {
	return displayYAML(w, o.Results)
}

func (o ForemanOrganizations) header() []string {
	return []string{
		"ID",
		"Name",
		"UpdatedAt",
		"CreatedAt",
	}
}
func (o ForemanOrganizations) TableData(w io.Writer) error {
	data := make([][]string, len(o.Results))
	for i, org := range o.Results {
		data[i] = []string{
			fmt.Sprintf("%v", org.ID),
			org.Name,
			humanize.Time(org.UpdatedAt),
			humanize.Time(org.CreatedAt),
		}
	}
	return displayTable(w, data, o.header())
}

func (o ForemanOrganizations) Text(w io.Writer) error {
	return nil
}
