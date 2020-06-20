package displayers

import (
	"io"
)

type Displayable interface {
	JSON(w io.Writer) error
	YAML(w io.Writer) error
	PP(w io.Writer) error
	TableData(w io.Writer) error
	JSONPath(w io.Writer, template string) error
	Text(w io.Writer) error
}
