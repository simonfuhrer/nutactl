package foreman

import (
	"context"
	"fmt"
	"net/http"
)

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=OperatingSystem Value=Title Path=operatingsystems"
type OperatingSystem struct {
	// Inherits the base object's attributes
	ForemanObject
	Title           string                `json:"title,omitempty"`
	Family          string                `json:"family,omitempty"`
	Description     string                `json:"description,omitempty"`
	ReleaseName     string                `json:"release_name,omitempty"`
	Media           []ForemanKeyValuePair `json:"media,omitempty"`
	Architectures   []ForemanKeyValuePair `json:"architectures,omitempty"`
	PartitionTables []ForemanKeyValuePair `json:"ptables,omitempty"`
}

type ForemanKeyValuePair struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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

type QueryResponseOperatingSystemImages struct {
	QueryResponse
	Results []OperatingSystemImage `json:"results"`
}

type OperatingSystemImage struct {
	ForemanObject
	OperatingsystemID   int    `json:"operatingsystem_id,omitempty"`
	OperatingsystemName string `json:"operatingsystem_name,omitempty"`
	ComputeResourceID   int    `json:"compute_resource_id,omitempty"`
	ComputeResourceName string `json:"compute_resource_name,omitempty"`
	ArchitectureID      int    `json:"architecture_id,omitempty"`
	ArchitectureName    string `json:"architecture_name,omitempty"`
	UUID                string `json:"uuid,omitempty"`
	Username            string `json:"username,omitempty"`
}

func (c *Client) SearchOperatingSystemImages(ctx context.Context, os *OperatingSystem, filter string) (*QueryResponseOperatingSystemImages, error) {
	response := new(QueryResponseOperatingSystemImages)
	err := c.requestSearchHelper(ctx, fmt.Sprintf("/%s/%d/images", OperatingSystemEndpointPrefix, os.ID), http.MethodGet, filter, nil, response)
	return response, err
}
