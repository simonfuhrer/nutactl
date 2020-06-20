package displayers

import (
	"io"

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Projects wraps a nutanix ProjectListItnent.
type Projects struct {
	schema.ProjectListIntent
}

//var _ Displayable = &Projects{}

func (o Projects) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o Projects) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o Projects) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o Projects) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o Projects) header() []string {
	return []string{
		"UUID",
		"Name",
		"Description",
	}
}

func (o Projects) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, project := range o.Entities {
		data[i] = []string{
			project.Metadata.UUID,
			project.Spec.Name,
			project.Spec.Description,
		}
	}
	return displayTable(w, data, o.header())
}

func (o Projects) Text(w io.Writer) error {
	return nil
}
