package displayers

import (
	"io"

	"github.com/simonfuhrer/nutactl/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// AvailabilityZones wraps a nutanix AvailabilityZoneListIntent.
type AvailabilityZones struct {
	schema.AvailabilityZoneListIntent
}

//var _ Displayable = &AvailabilityZones{}

func (o AvailabilityZones) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o AvailabilityZones) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o AvailabilityZones) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o AvailabilityZones) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o AvailabilityZones) header() []string {
	return []string{
		"Name",
		"Type",
	}
}

func (o AvailabilityZones) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, av := range o.Entities {
		data[i] = []string{
			utils.StringValue(av.Spec.Name),
			utils.StringValue(av.Spec.Resources.ManagementPlaneType),
		}
	}
	return displayTable(w, data, o.header())
}

func (o AvailabilityZones) Text(w io.Writer) error {
	return nil
}
