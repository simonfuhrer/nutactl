package displayers

import (
	"fmt"
	"io"
	"strconv"

	"github.com/simonfuhrer/nutactl/pkg/utils"
	"github.com/tecbiz-ch/nutanix-go-sdk/schema"
)

// Clusters wraps a nutanix ImageListIntent.
type VMs struct {
	schema.VMListIntent
}

//var _ Displayable = &VMs{}

func (o VMs) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o VMs) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o VMs) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o VMs) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o VMs) header() []string {
	return []string{
		"UUID",
		"Name",
		"Power",
		"Project",
		"Subnet",
		"IP",
		"Cluster",
		"MiB",
		"CPU",
		"Status",
	}
}

func (o VMs) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, vm := range o.Entities {
		subnet := ""
		ip := ""
		if len(vm.Spec.Resources.NicList) > 0 {
			subnet = vm.Spec.Resources.NicList[0].SubnetReference.Name
			ip = vm.Status.Resources.NicList[0].IPEndpointList[0].IP
		}
		state := ""
		if vm.Metadata.ProjectReference != nil {
			state = vm.Metadata.ProjectReference.Name
		}
		data[i] = []string{
			vm.Metadata.UUID,
			vm.Spec.Name,
			vm.Spec.Resources.PowerState,
			state,
			subnet,
			ip,
			vm.Spec.ClusterReference.Name,
			strconv.FormatInt(vm.Spec.Resources.MemorySizeMib, 10),
			fmt.Sprintf("%d/%d", vm.Spec.Resources.NumSockets, vm.Spec.Resources.NumVcpusPerSocket),
			utils.StringValue(vm.Status.State),
		}
	}
	return displayTable(w, data, o.header())
}

func (o VMs) Text(w io.Writer) error {
	return nil
}
