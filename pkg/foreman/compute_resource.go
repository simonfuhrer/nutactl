package foreman

import (
	"context"
	"fmt"
	"net/http"
)

//go:generate genny -in=template.go -out=$GOFILE-gen.go gen "Type=ComputeResource Value=Name Path=compute_resources"

type ComputeResource struct {
	// Inherits the base object's attributes
	ForemanObject
	Description           string                 `json:"description"`
	Provider              string                 `json:"provider"`
	ProviderFriendlyName  string                 `json:"provider_friendly_name"`
	User                  string                 `json:"user"`
	ISOLibraryMsountpoint string                 `json:"iso_library_mountpoint"`
	Images                []ComputeResourceImage `json:"images,omitempty"`
	ComputeAttributes     interface{}            `json:"compute_attributes,omitempty"`
}

type ComputeResourceImage struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type QueryResponseComputeResourceStorageDomains struct {
	QueryResponse
	Results []ComputeResourceStorageDomain `json:"results"`
}

type QueryResponseComputeResourceAvailableNetworks struct {
	QueryResponse
	Results []ComputeResourceNetwork `json:"results"`
}

type ComputeResourceStorageDomain struct {
	Name      string `json:"name"`
	UUID      string `json:"id"`
	Capacity  int    `json:"capacity"`
	Freespace int    `json:"freespace"`
}

type ComputeResourceNetwork struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}

func (c *Client) GetComputeResourceStorageDomains(ctx context.Context, computeResource *ComputeResource, filter string) (*QueryResponseComputeResourceStorageDomains, error) {
	response := new(QueryResponseComputeResourceStorageDomains)
	err := c.requestSearchHelper(ctx, fmt.Sprintf("/%s/%d/available_storage_domains", ComputeResourceEndpointPrefix, computeResource.ID), http.MethodGet, filter, nil, response)
	return response, err
}

func (c *Client) GetComputeResourceAvailableNetworks(ctx context.Context, computeResource *ComputeResource, filter string) (*QueryResponseComputeResourceStorageDomains, error) {
	response := new(QueryResponseComputeResourceStorageDomains)
	err := c.requestSearchHelper(ctx, fmt.Sprintf("/%s/%d/available_networks", ComputeResourceEndpointPrefix, computeResource.ID), http.MethodGet, filter, nil, response)
	return response, err
}
