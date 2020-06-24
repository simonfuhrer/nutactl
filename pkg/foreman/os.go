package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=OperatingSystem"
type OperatingSystem struct {
	// Inherits the base object's attributes
	ForemanObject
	Title       string `json:"title,omitempty"`
	Family      string `json:"family,omitempty"`
	Description string `json:"description,omitempty"`
	ReleaseName string `json:"release_name,omitempty"`
}

type NewOperatingSystemData struct {
	ForemanObject
	Fullname             string                 `json:"fullname,omitempty"`
	ParametersAttributes []ParametersAttributes `json:"domain_parameters_attributes,omitempty"`
}

// Structures used to create a new domain
type OperatingSystemRequest struct {
	OrganizationID  int                    `json:"organization_id,omitempty"`
	LocationID      int                    `json:"location_id,omitempty"`
	OperatingSystem NewOperatingSystemData `json:"operatingsystem"`
}
