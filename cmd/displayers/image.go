package displayers

import (
	"io"
	"strconv"

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Images wraps a nutanix ImageListIntent.
type Images struct {
	schema.ImageListIntent
}

//var _ Displayable = &Images{}

func (o Images) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o Images) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o Images) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o Images) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o Images) header() []string {
	return []string{
		"UUID",
		"Name",
		"Description",
		"Arch",
		"ImageType",
		"MB",
		"Status",
	}
}
func (o Images) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, image := range o.Entities {
		var size string
		if image.Status.Resources.SizeBytes != 0 {
			size = strconv.FormatInt(image.Status.Resources.SizeBytes/1024/1024, 10)
		}
		data[i] = []string{
			image.Metadata.UUID,
			image.Spec.Name,
			image.Spec.Description,
			image.Spec.Resources.Architecture,
			image.Spec.Resources.ImageType,
			size,
			image.Status.State,
		}
	}
	return displayTable(w, data, o.header())
}

func (o Images) Text(w io.Writer) error {
	return nil
}
