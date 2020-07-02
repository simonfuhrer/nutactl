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
type ForemanSmartClassParameters struct {
	foreman.QueryResponseSmartClassParameter
}

func (o ForemanSmartClassParameters) JSON(w io.Writer) error {
	return displayJSON(w, o.Results)
}

func (o ForemanSmartClassParameters) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Results)
}

func (o ForemanSmartClassParameters) PP(w io.Writer) error {
	return displayPP(w, o.Results)
}

func (o ForemanSmartClassParameters) YAML(w io.Writer) error {
	return displayYAML(w, o.Results)
}

func (o ForemanSmartClassParameters) header() []string {
	return []string{
		"ID",
		"Class Name",
		"Required",
		"Parameter",
		"Type",
	}
}
func (o ForemanSmartClassParameters) TableData(w io.Writer) error {
	data := make([][]string, len(o.Results))
	for i, p := range o.Results {
		data[i] = []string{
			fmt.Sprintf("%v", p.ID),
			p.PuppetclassName,
			fmt.Sprintf("%v", p.Required),
			p.Parameter,
			p.ParameterType,
		}
	}
	return displayTable(w, data, o.header())
}

func (o ForemanSmartClassParameters) Text(w io.Writer) error {
	return nil
}
