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

// VMRecoveryPoints wraps a nutanix VMRecoveryPointListIntent.
type VMRecoveryPoints struct {
	schema.VMRecoveryPointListIntent
}

//var _ Displayable = &VMRecoveryPoints{}

func (o VMRecoveryPoints) JSON(w io.Writer) error {
	return DisplayJSON(w, o.Entities)
}

func (o VMRecoveryPoints) JSONPath(w io.Writer, template string) error {
	return DisplayJSONPath(w, template, o.Entities)
}

func (o VMRecoveryPoints) PP(w io.Writer) error {
	return DisplayPP(w, o.Entities)
}

func (o VMRecoveryPoints) YAML(w io.Writer) error {
	return DisplayYAML(w, o.Entities)
}

func (o VMRecoveryPoints) header() []string {
	return []string{
		"UUID",
		"Name",
		"CreationTime",
		"ExpirationTime",
		"Type",
		"VM",
		"State",
	}
}

func (o VMRecoveryPoints) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, vmrecoveryPoint := range o.Entities {

		data[i] = []string{
			vmrecoveryPoint.Metadata.UUID,
			vmrecoveryPoint.Spec.Name,
			vmrecoveryPoint.Metadata.CreationTime.String(),
			vmrecoveryPoint.Spec.Resources.ExpirationTime.String(),
			vmrecoveryPoint.Spec.Resources.RecoveryPointType,
			vmrecoveryPoint.Spec.Resources.ParentVMReference.Name,
			vmrecoveryPoint.Status.State,
		}
	}
	return DisplayTable(w, data, o.header())
}

func (o VMRecoveryPoints) Text(w io.Writer) error {
	return nil
}
