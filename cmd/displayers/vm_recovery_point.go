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
	return displayJSON(w, o.Entities)
}

func (o VMRecoveryPoints) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o VMRecoveryPoints) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o VMRecoveryPoints) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
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
	return displayTable(w, data, o.header())
}

func (o VMRecoveryPoints) Text(w io.Writer) error {
	return nil
}
