package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Subnet"

type Subnet struct {
	ForemanObject
	NetworkType          string                `json:"network_type"`
	CIDR                 int                   `json:"cidr"`
	Network              string                `json:"network"`
	NetworkAddress       string                `json:"network_address"`
	Mask                 string                `json:"mask"`
	Priority             string                `json:"priority"`
	Gateway              string                `json:"gateway"`
	From                 string                `json:"from"`
	To                   string                `json:"to"`
	Ipam                 string                `json:"ipam"`
	BootMode             string                `json:"boot-mode"`
	DomainIds            []int                 `json:"domain_ids"`
	Vlanid               int                   `json:"vlanid"`
	Mtu                  int                   `json:"mtu"`
	DNSPrimary           string                `json:"dns_primary"`
	DNSSecondary         string                `json:"dns_secondary"`
	Description          string                `json:"description"`
	DHCPID               int                   `json:"dhcp_id"`
	DHCPName             string                `json:"dhcp_name"`
	Interfaces           []ForemanNetInterface `json:"interfaces,omitempty"`
	ParametersAttributes []Parameters          `json:"parameters"`
}

type NewSubnetData struct {
	Subnet
}

type SubnetRequest struct {
	OrganizationID int           `json:"organization_id,omitempty"`
	LocationID     int           `json:"location_id,omitempty"`
	Subnet         NewSubnetData `json:"subnet"`
}
