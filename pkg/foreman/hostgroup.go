package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Hostgroup"

type Hostgroup struct {
	// Inherits the base object's attributes
	ForemanObject

	// The title is a computed property representing the fullname of the
	// hostgroup.  A hostgroup's title is a path-like string from the head
	// of the hostgroup tree down to this hostgroup.  The title will be
	// in the form of: "<parent 1>/<parent 2>/.../<name>"
	Title string `json:"title"`
	// ID of the architecture associated with this hostgroup
	ArchitectureID   int    `json:"architecture_id"`
	ArchitectureName string `json:"architecture_name"`
	// ID of the compute profile associated with this hostgroup
	ComputeProfileID   int    `json:"compute_profile_id"`
	ComputeProfileName string `json:"compute_profile_name"`
	// ID of the domain associated with this hostgroup
	DomainID   int    `json:"domain_id"`
	DomainName string `json:"domain_name"`
	// ID of the environment associated with this hostgroup
	EnvironmentID   int    `json:"environment_id"`
	EnvironmentName string `json:"environment_name"`

	// ID of the media associated with this hostgroup
	MediaID int `json:"medium_id"`
	// ID of the operating system associated with this hostgroup
	OperatingSystemID   int    `json:"operatingsystem_id"`
	OperatingSystemName string `json:"operatingsystem_name"`

	// ID of this hostgroup's parent hostgroup
	ParentID int `json:"parent_id"`
	// ID of the partition table to use with this hostgroup
	PartitionTableID int `json:"ptable_id"`
	// ID of the smart proxy acting as the puppet certificate authority
	// server for the hostgroup
	PuppetCAProxyID int `json:"puppet_ca_proxy_id"`
	// ID of the smart proxy acting as the puppet proxy server for the
	// hostgroup
	PuppetProxyID int `json:"puppet_proxy_id"`
	// ID of the realm associated with the hostgroup
	RealmID int `json:"realm_id"`
	// ID of the subnet associated with the hostgroup
	SubnetID   int    `json:"subnet_id"`
	SubnetName string `json:"subnet_name"`
}
