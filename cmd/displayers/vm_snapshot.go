package displayers

import (
	"io"

	v2 "github.com/tecbiz-ch/nutanix-go-sdk/schema/v2"
)

// VMSnapshots wraps a nutanix SnapshotList.
type VMSnapshots struct {
	v2.SnapshotList
}

//var _ Displayable = &VMSnapshots{}

func (o VMSnapshots) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o VMSnapshots) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o VMSnapshots) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o VMSnapshots) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o VMSnapshots) header() []string {
	return []string{
		"UUID",
		"Name",
		"CreatedTime",
	}
}

func (o VMSnapshots) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, vmdisk := range o.Entities {
		data[i] = []string{
			vmdisk.UUID,
			vmdisk.Name,
			vmdisk.CreatedTime.String(),
		}
	}
	return displayTable(w, data, o.header())
}

func (o VMSnapshots) Text(w io.Writer) error {
	return nil
}
