package displayers

import (
	"io"
	"strconv"

	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Clusters wraps a nutanix ImageListIntent.
type Clusters struct {
	schema.ClusterListIntent
}

//var _ Displayable = &Clusters{}

func (o Clusters) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o Clusters) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o Clusters) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o Clusters) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o Clusters) header() []string {
	return []string{
		"UUID",
		"Name",
		"nos_version",
		"ncc_version",
		"OperationMode",
		"ExternalIP",
		"Nodes",
	}
}

func (o Clusters) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, cluster := range o.Entities {
		if cluster.Spec.Name == "Unnamed" {
			continue
		}
		var hosts = 0
		if cluster.Status.Resources.Nodes != nil {
			hosts = len(cluster.Status.Resources.Nodes.HypervisorServerList)

		}
		data[i] = []string{
			cluster.Metadata.UUID,
			cluster.Spec.Name,
			cluster.Spec.Resources.Config.SoftwareMap["NOS"].Version,
			cluster.Spec.Resources.Config.SoftwareMap["NCC"].Version,
			cluster.Spec.Resources.Config.OperationMode,
			cluster.Spec.Resources.Network.ExternalIP,
			strconv.Itoa(hosts),
		}
	}
	return displayTable(w, data, o.header())
}

func (o Clusters) Text(w io.Writer) error {
	return nil
}
