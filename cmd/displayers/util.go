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
	"encoding/json"
	"fmt"
	"io"

	"github.com/k0kubun/pp"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/util/jsonpath"
)

func JSONToYAML(j []byte) ([]byte, error) {
	var jsonObj interface{}

	err := yaml.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}

	return yaml.Marshal(jsonObj)
}

func DisplayYAML(w io.Writer, data interface{}) error {
	j, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "marshaling to JSON")
	}

	y, err := JSONToYAML(j)
	if err != nil {
		return errors.Wrap(err, "converting JSON to YAML")
	}

	fmt.Fprintln(w, string(y))

	return nil
}

func DisplayJSON(w io.Writer, data interface{}) error {
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshaling to JSON")
	}
	fmt.Fprintln(w, string(j))
	return nil
}

func DisplayPP(w io.Writer, data interface{}) error {
	_, err := pp.Println(w, data)
	if err != nil {
		return errors.Wrap(err, "marshaling to PP")
	}
	return nil
}

func DisplayText(w io.Writer, data []string) error {
	for _, line := range data {
		fmt.Print(line)
	}
	return nil
}

func DisplayTable(w io.Writer, data [][]string, header []string) error {
	table := tablewriter.NewWriter(w)
	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
	return nil
}

func DisplayJSONPath(w io.Writer, template string, data interface{}) error {
	jp := jsonpath.New("")
	err := jp.Parse(template)
	if err != nil {
		return errors.Wrap(err, "parsing jsonpath")
	}

	err = jp.Execute(w, data)
	if err != nil {
		return errors.Wrap(err, "executing jsonpath")
	}

	return nil
}
