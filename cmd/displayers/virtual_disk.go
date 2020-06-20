package displayers

import (
	"io"
	"strconv"

	v2 "github.com/tecbiz-ch/nutanix-go-sdk/schema/v2"
)

// VirtualDisks wraps a nutanix VirtualDiskList.
type VirtualDisks struct {
	v2.VirtualDiskList
}

//var _ Displayable = &VirtualDisks{}

func (o VirtualDisks) JSON(w io.Writer) error {
	return displayJSON(w, o.Entities)
}

func (o VirtualDisks) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, o.Entities)
}

func (o VirtualDisks) PP(w io.Writer) error {
	return displayPP(w, o.Entities)
}

func (o VirtualDisks) YAML(w io.Writer) error {
	return displayYAML(w, o.Entities)
}

func (o VirtualDisks) header() []string {
	return []string{
		"UUID",
		"VM",
		"SizeinMB",
		"Cluster",
	}
}

func (o VirtualDisks) TableData(w io.Writer) error {
	data := make([][]string, len(o.Entities))
	for i, vmdisk := range o.Entities {
		data[i] = []string{
			vmdisk.UUID,
			vmdisk.AttachedVMName,
			strconv.FormatInt(vmdisk.DiskCapacityInBytes/1024/1024, 10),
			vmdisk.ClusterUUID,
		}
	}
	return displayTable(w, data, o.header())
}

func (o VirtualDisks) Text(w io.Writer) error {
	return nil
}
