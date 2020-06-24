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

// ForemanHosts wraps a foreman Hosts.
type ForemanHostgroups struct {
	foreman.QueryResponseHostgroup
}

func (o ForemanHostgroups) JSON(w io.Writer) error {
	return displayJSON(w, o.Results)
}

func (o ForemanHostgroups) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Results)
}

func (o ForemanHostgroups) PP(w io.Writer) error {
	return displayPP(w, o.Results)
}

func (o ForemanHostgroups) YAML(w io.Writer) error {
	return displayYAML(w, o.Results)
}

func (o ForemanHostgroups) header() []string {
	return []string{
		"ID",
		"Title",
		"Name",
		"Environment",
		"CreatedAt",
	}
}
func (o ForemanHostgroups) TableData(w io.Writer) error {
	data := make([][]string, len(o.Results))
	for i, hostgroup := range o.Results {
		data[i] = []string{
			fmt.Sprintf("%v", hostgroup.ID),
			fmt.Sprintf("%v", hostgroup.Title),
			fmt.Sprintf("%v", hostgroup.Name),
			fmt.Sprintf("%v", hostgroup.EnvironmentName),
			fmt.Sprintf("%v", hostgroup.CreatedAt),
		}
	}
	return displayTable(w, data, o.header())
}

func (o ForemanHostgroups) Text(w io.Writer) error {
	return nil
}
