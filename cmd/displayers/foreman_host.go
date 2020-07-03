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

// ForemanHosts wraps a foreman Hosts.
type ForemanHosts struct {
	foreman.QueryResponseHost
}

func (o ForemanHosts) JSON(w io.Writer) error {
	return displayJSON(w, o.Results)
}

func (o ForemanHosts) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Results)
}

func (o ForemanHosts) PP(w io.Writer) error {
	return displayPP(w, o.Results)
}

func (o ForemanHosts) YAML(w io.Writer) error {
	return displayYAML(w, o.Results)
}

func (o ForemanHosts) header() []string {
	return []string{
		"ID",
		"Name",
		"IP",
		"OS",
		"Environment",
		"Model",
		"Build Status",
		"Last Report",
		"Status",
		"UpdatedAt",
		"CreatedAt",
	}
}
func (o ForemanHosts) TableData(w io.Writer) error {
	data := make([][]string, len(o.Results))
	for i, host := range o.Results {
		data[i] = []string{
			fmt.Sprintf("%v", host.ID),
			host.Name,
			host.IP,
			host.OperatingsystemName,
			host.EnvironmentName,
			host.ModelName,
			host.BuildStatusLabel,
			humanize.Time(host.LastReport),
			host.GlobalStatusLabel,
			humanize.Time(host.UpdatedAt),
			humanize.Time(host.CreatedAt),
		}
	}
	return displayTable(w, data, o.header())
}

func (o ForemanHosts) Text(w io.Writer) error {
	return nil
}
