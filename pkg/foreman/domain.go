package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=Domain Value=name Path=domains"

type Domain struct {
	// Inherits the base object's attributes
	ForemanObject
	Fullname   string                `json:"fullname",omitempty`
	DNS        interface{}           `json:"dns"`
	DNSID      int                   `json:"dns_id,omitempty"`
	Parameters []Parameters          `json:"parameters,omitempty"`
	Interfaces []ForemanNetInterface `json:"interfaces,omitempty"`
}

type ForemanNetInterface struct {
	ForemanObject
	IP  string `json:"ip,omitempty"`
	IP6 string `json:"ip6,omitempty"`
	Mac string `json:"mac,omitempty"`
	//	Mtu        int `json:"mtu,omitempty"`
	FQDN       string `json:"fqdn,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Primary    bool   `json:"primary,omitempty"`
	Provision  bool   `json:"provision,omitempty"`
	Type       string `json:"type,omitempty"`
}

type NewDomainData struct {
	ForemanObject
	Fullname             string                 `json:"fullname",omitempty`
	ParametersAttributes []ParametersAttributes `json:"domain_parameters_attributes,omitempty"`
}

// Structures used to create a new domain
type DomainRequest struct {
	OrganizationID int           `json:"organization_id,omitempty"`
	LocationID     int           `json:"location_id,omitempty"`
	Domain         NewDomainData `json:"domain"`
}
