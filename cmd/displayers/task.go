package displayers

import (
	"io"
	"sort"
	"strconv"

	"github.com/simonfuhrer/nutactl/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Tasks wraps a nutanix TaskListIntent.
type Tasks struct {
	schema.TaskListIntent
}

//var _ Displayable = &Tasks{}

func (o Tasks) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o Tasks) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o Tasks) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o Tasks) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o Tasks) header() []string {
	return []string{
		"UUID",
		"Status",
		"OperationType",
		"PercentageComplete",
		"CreationTime",
		"LastUpdateTime",
	}
}

func (o Tasks) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))

	// Workaround sort
	list := o.Entities
	sort.Slice(list, func(i, j int) bool { return list[i].LastUpdateTime.Before(*list[j].LastUpdateTime) })
	// Workaround End

	for i, t := range o.Entities {
		data[i] = []string{
			utils.StringValue(t.UUID),
			utils.StringValue(t.Status),
			utils.StringValue(t.OperationType),
			strconv.FormatInt(*t.PercentageComplete, 10),
			t.CreationTime.String(),
			t.LastUpdateTime.String(),
		}
	}
	return displayTable(w, data, o.header())
}

func (o Tasks) Text(w io.Writer) error {
	return nil
}
