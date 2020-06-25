package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Host"

// Structures used to hold the returned host value
type Host struct {
	ForemanObject
	Name                     string            `json:"name"`
	IP                       string            `json:"ip,omitempty"`
	IP6                      string            `json:"ip6,omitempty"`
	EnvironmentID            int               `json:"environment_id,omitempty"`
	EnvironmentName          string            `json:"environment_name,omitempty"`
	Mac                      string            `json:"mac,omitempty"`
	RealmID                  int               `json:"realm_id,omitempty"`
	RealmName                string            `json:"realm_name,omitempty"`
	SpMac                    string            `json:"sp_mac,omitempty"`
	SpIP                     string            `json:"sp_ip,omitempty"`
	SpName                   string            `json:"sp_name,omitempty"`
	DomainID                 int               `json:"domain_id,omitempty"`
	DomainName               string            `json:"domain_name,omitempty"`
	ArchitectureID           int               `json:"architecture_id,omitempty"`
	ArchitectureName         string            `json:"architecture_name,omitempty"`
	OperatingsystemID        int               `json:"operatingsystem_id,omitempty"`
	OperatingsystemName      string            `json:"operatingsystem_name,omitempty"`
	SubnetID                 int               `json:"subnet_id,omitempty"`
	SubnetName               string            `json:"subnet_name,omitempty"`
	Subnet6ID                int               `json:"subnet6_id,omitempty"`
	Subnet6Name              string            `json:"subnet6_name,omitempty"`
	SpSubnetID               int               `json:"sp_subnet_id,omitempty"`
	PtableID                 int               `json:"ptable_id,omitempty"`
	PtableName               string            `json:"ptable_name,omitempty"`
	MediumID                 int               `json:"medium_id,omitempty"`
	MediumName               string            `json:"medium_name,omitempty"`
	PXELoader                string            `json:"pxe_loader,omitempty"`
	Build                    bool              `json:"build,omitempty"`
	Comment                  string            `json:"comment,omitempty"`
	Disk                     interface{}       `json:"disk,omitempty"`
	InstalledAt              interface{}       `json:"instanlled_at,omitempty"`
	ModelID                  int               `json:"model_id,omitempty"`
	HostgroupID              int               `json:"hostgroup_id,omitempty"`
	OwnerID                  int               `json:"owner_id,omitempty"`
	OwnerName                string            `json:"owner_name,omitempty"`
	OwnerType                string            `json:"owner_type,omitempty"`
	Enabled                  bool              `json:"enabled,omitempty"`
	Managed                  bool              `json:"managed,omitempty"`
	UseImage                 interface{}       `json:"use_image,omitempty"`
	ImageFile                string            `json:"image_file,omitempty"`
	UUID                     string            `json:"uuid,omitempty"`
	ComputeResourceID        int               `json:"compute_resource_id,omitempty"`
	ComputeResourceName      string            `json:"compute_resource_name,omitempty"`
	ComputeProfileID         int               `json:"compute_profile_id,omitempty"`
	ComputeProfileName       string            `json:"compute_profile_name,omitempty"`
	Capabilities             []string          `json:"capabilities,omitempty"`
	CertName                 string            `json:"certname,omitempty"`
	ImageID                  int               `json:"image_id,omitempty"`
	ImageName                string            `json:"image_name,omitempty"`
	CreatedAt                string            `json:"created_at,omitempty"`
	UpdatedAt                string            `json:"updated_at,omitempty"`
	LastCompile              interface{}       `json:"last_compile,omitempty"`
	GlobalStatus             int               `json:"global_status,omitempty"`
	GlobalStatusLabel        string            `json:"global_status_label,omitempty"`
	UptimeSeconds            interface{}       `json:"uptime_seconds,omitempty"`
	OrganizationID           int               `json:"organization_id,omitempty"`
	OrganizationName         string            `json:"organization_name,omitempty"`
	LocationID               int               `json:"location_id,omitempty"`
	LocationName             string            `json:"location_name,omitempty"`
	PuppetStatus             int               `json:"puppet_proxy_status,omitempty"`
	ModelName                string            `json:"model_name,omitempty"`
	ConfigurationStatus      int               `json:"configuration_status,omitempty"`
	ConfigurationStatusLabel string            `json:"configuration_status_label,omitempty"`
	BuildStatus              int               `json:"build_status,omitempty"`
	BuildStatusLabel         string            `json:"build_status_label,omitempty"`
	PuppetProxyID            int               `json:"puppet_proxy_id,omitempty"`
	PuppetProxyName          string            `json:"puppet_proxy_name,omitempty"`
	PuppetCAProxyID          int               `json:"puppet_ca_proxy_id,omitempty"`
	PuppetCAProxyName        string            `json:"puppet_ca_proxy_name,omitempty"`
	PuppetProxy              PuppetProxy       `json:"puppet_proxy,omitempty"`
	PuppetCaProxy            PuppetCaProxy     `json:"puppet_ca_proxy,omitempty"`
	Parameters               []interface{}     `json:"parameters,omitempty"`
	HostGroupName            string            `json:"host_group_name,omitempty"`
	HostGroupTitle           string            `json:"host_group_title,omitempty"`
	AllParameters            []Parameters      `json:"all_parameters,omitempty"`
	Interfaces               []NetInterfaceRet `json:"interfaces,omitempty"`
	PuppeClasses             []interface{}     `json:"puppetclasses,omitempty"`
	ConfigGroups             []interface{}     `json:"config_groups,omitempty"`
	AllPuppeClasses          []interface{}     `json:"all_puppetclasses,omitempty"`
}
type PuppetProxy struct {
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
}
type PuppetCaProxy struct {
	Name string `json:"name,omitempty"`
	ID   int    `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Parameters struct {
	Name          string      `json:"name,omitempty"`
	Priority      interface{} `json:"priority,omitempty"`
	ID            int         `json:"id,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	ParameterType string      `json:"parameter_type,omitempty"`
	CreatedAt     string      `json:"created_at,omitempty"`
	UpdatedAt     string      `json:"updated_at,omitempty"`
	Permissions   Permissions `json:"permissions,omitempty"`
}
type NetInterfaceRet struct {
	SubnetID        int         `json:"subnet_id,omitempty"`
	SubnetName      string      `json:"subnet_name,omitempty"`
	Subnet6ID       int         `json:"subnet6_id,omitempty"`
	Subnet6Name     string      `json:"subnet6_name,omitempty"`
	DomainID        int         `json:"domain_id,omitempty"`
	DomainName      string      `json:"domain_name,omitempty"`
	CreatedAt       string      `json:"created_at,omitempty"`
	UpdatedAt       string      `json:"updated_at,omitempty"`
	Managed         bool        `json:"managed,omitempty"`
	Identifier      string      `json:"identifier,omitempty"`
	ID              int         `json:"id,omitempty"`
	Name            string      `json:"name,omitempty"`
	IP              string      `json:"ip,omitempty"`
	IP6             string      `json:"ip6,omitempty"`
	Mac             string      `json:"mac,omitempty"`
	Mtu             int         `json:"mtu,omitempty"`
	Fqdn            string      `json:"fqdn,omitempty"`
	Primary         bool        `json:"primary,omitempty"`
	Provision       bool        `json:"provision,omitempty"`
	Type            string      `json:"type,omitempty"`
	Virtual         bool        `json:"virtual,omitempty"`
	Username        string      `json:"username,omitempty"`
	Password        string      `json:"password,omitempty"`
	Provider        string      `json:"provider,omitempty"`
	Tag             string      `json:"tag,omitempty"`
	AttachedTo      string      `json:"attached_to,omitempty"`
	Mode            string      `json:"mode,omitempty"`
	AttachedDevices interface{} `json:"attached_devices,omitempty"`
	BondOptions     string      `json:"bond_options,omitempty"`
}
type Permissions struct {
	ViewHosts      bool `json:"view_hosts,omitempty"`
	CreateHosts    bool `json:"create_hosts,omitempty"`
	EditHosts      bool `json:"edit_hosts,omitempty"`
	DestroyHosts   bool `json:"destroy_hosts,omitempty"`
	BuildHosts     bool `json:"build_hosts,omitempty"`
	PowerHosts     bool `json:"power_hosts,omitempty"`
	ConsoleHosts   bool `json:"console_hosts,omitempty"`
	IpmiBootHosts  bool `json:"ipmi_boot_hosts,omitempty"`
	PuppetrunHosts bool `json:"puppetrun_hosts,omitempty"`
}

// Structures used to create a new host
type HostRequest struct {
	OrganizationID int         `json:"organization_id,omitempty"`
	LocationID     int         `json:"location_id,omitempty"`
	Host           NewHostData `json:"host"`
}

type NewHostData struct {
	ForemanObject
	EnvironmentID            string                 `json:"environment_id,omitempty"`
	IP                       string                 `json:"ip,omitempty"`
	Mac                      string                 `json:"mac" validate:"required"`
	ArchitectureID           int                    `json:"architecture_id,omitempty"`
	DomainID                 int                    `json:"domain_id" validate:"required"`
	RealmID                  int                    `json:"realm_id,omitempty"`
	PuppetProxyID            int                    `json:"puppet_proxy_id,omitempty"`
	PuppetCAProxyID          int                    `json:"puppet_ca_proxy_id,omitempty"`
	PuppeClaassIDs           []interface{}          `json:"puppet_class_ids,omitempty"`
	ConfigGroupIDs           []interface{}          `json:"config_group_ids,omitempty"`
	OperatingsystemID        int                    `json:"operatingsystem_id,omitempty"`
	MediumID                 int                    `json:"medium_id,omitempty"`
	PXELoader                string                 `json:"pxe_loader,omitempty"`
	PtableID                 int                    `json:"ptable_id,omitempty"`
	SubnetID                 int                    `json:"subnet_id,omitempty"`
	ComputeResourceID        int                    `json:"compute_resource_id,omitempty"`
	RootPass                 string                 `json:"root_pass,omitempty"`
	ModelID                  int                    `json:"model_id,omitempty"`
	HostgroupID              int                    `json:"hostgroup_id,omitempty"`
	OwnerID                  int                    `json:"owner_id,omitempty"`
	OwnerType                string                 `json:"owner_type,omitempty"`
	ImageID                  int                    `json:"image_id,omitempty"`
	HostParametersAttributes []ParametersAttributes `json:"host_parameters_attributes,omitempty"`
	Build                    bool                   `json:"build,omitempty"`
	Enabled                  bool                   `json:"enabled,omitempty"`
	ProvisionMethod          string                 `json:"provision_method,omitempty"`
	Managed                  bool                   `json:"managed,omitempty"`
	ProgressReportID         string                 `json:"progress_report_id,omitempty"`
	Comment                  string                 `json:"comment,omitempty"`
	Capabilities             string                 `json:"capabilities,omitempty"`
	ComputeProfileID         int                    `json:"compute_profile_id,omitempty"`
	InterfacesAttributes     *InterfacesAttributes  `json:"interfaces_attributes,omitempty"`
}

type ParametersAttributes struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type InterfacesAttributes struct {
	Primary           NetInterface `json:"1,omitempty"`
	Management        NetInterface `json:"2,omitempty"`
	Storage           NetInterface `json:"3,omitempty"`
	StorageManagement NetInterface `json:"4,omitempty"`
	Tenant            NetInterface `json:"5,omitempty"`
	LBAAS             NetInterface `json:"6,omitempty"`
	InsideNet         NetInterface `json:"7,omitempty"`
	GatewayNet        NetInterface `json:"8,omitempty"`
}

type NetInterface struct {
	Name            string   `json:"name,omitempty"`
	Primary         bool     `json:"primary,omitempty"`
	IP              string   `json:"ip,omitempty"`
	IP6             string   `json:"ip6,omitempty"`
	Mac             string   `json:"mac,omitempty"`
	Type            string   `json:"type,omitempty"`
	SubnetID        int      `json:"subnet_id,omitempty"`
	Subnet6ID       int      `json:"subnet6_id,omitempty"`
	DomainID        int      `json:"domain_id,omitempty"`
	Identifier      string   `json:"identifier,omitempty"`
	Managed         bool     `json:"managed,omitempty"`
	Provision       bool     `json:"provision,omitempty"`
	Username        string   `json:"username,omitempty"`
	Password        string   `json:"password,omitempty"`
	Provider        string   `json:"provider,omitempty"`
	Virtual         bool     `json:"virtual,omitempty"`
	Tag             string   `json:"tag,omitempty"`
	Mtu             int      `json:"mtu,omitempty"`
	AttachedTo      string   `json:"attached_to,omitempty"`
	Mode            string   `json:"mode,omitempty"`
	AttachedDevices []string `json:"attached_devices,omitempty"`
	BondOptions     string   `json:"bond_options,omitempty"`
}

// Structures used to create a new interface
type InterfacePostData struct {
	OrganizationID int          `json:"organization_id,omitempty"`
	LocationID     int          `json:"location_id,omitempty"`
	HostID         string       `json:"host_id,omitempty"`
	Interface      NetInterface `json:"interface"`
}