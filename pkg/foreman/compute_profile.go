package foreman

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=ComputeProfile Value=Name Path=compute_profiles"

type ComputeProfile struct {
	// Inherits the base object's attributes
	ForemanObject
	ComputeAttributes []ComputeProfileComputeAttributes `json:"compute_attributes,omitempty"`
}

type ComputeProfileComputeAttributes struct {
	ID int `json:"id"`
	// Human readable name of the API object
	Name                 string               `json:"name"`
	ComputeResourceID    int                  `json:"compute_resource_id"`
	ComputeResourceName  string               `json:"compute_resource_name"`
	ProviderFriendlyName string               `json:"provider_friendly_name"`
	ComputeProfileID     int                  `json:"compute_profile_id"`
	ComputeProfileName   string               `json:"compute_profile_name"`
	VMAttrs              *ComputeProfileAttrs `json:"vm_attrs,omitempty"`
	Attributes           *ComputeProfileAttrs `json:"attributes,omitempty"`
}
type ComputeProfileAttrs struct {
	VCPUsMax  int `json:"vcpus_max"`
	MemoryMin int `json:"memory_min"`
	MemoryMax int `json:"memory_max"`
}
