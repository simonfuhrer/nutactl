package displayers

import (
	"io"
	"strconv"
	"strings"

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Categories wraps a nutanix CategoryKeyList.
type Categories struct {
	schema.CategoryKeyList
}

//var _ Displayable = &Categories{}

func (o Categories) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o Categories) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o Categories) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o Categories) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o Categories) header() []string {
	return []string{
		"Name",
		"Description",
		"SystemDefined",
		"Values",
	}
}

func (o Categories) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, category := range o.Entities {
		data[i] = []string{
			category.Name,
			category.Description,
			strconv.FormatBool(category.SystemDefined),
			strings.Join(category.Values, ","),
		}
	}
	return displayTable(w, data, o.header())
}

func (o Categories) Text(w io.Writer) error {
	return nil
}
