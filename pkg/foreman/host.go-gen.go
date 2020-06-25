// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package foreman

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var HostEndpointPrefix = fmt.Sprintf("%ss", strings.ToLower("Host"))

type QueryResponseHost struct {
	QueryResponse
	Results []Host `json:"results"`
}

func (c *Client) GetHost(ctx context.Context, idOrName string) (*Host, error) {
	if id, err := strconv.Atoi(idOrName); err == nil {
		return c.GetHostByID(ctx, int(id))
	}
	return c.GetHostByName(ctx, idOrName)
}

func (c *Client) GetHostByID(ctx context.Context, id int) (*Host, error) {
	response := new(Host)
	err := c.requestHelper(ctx, fmt.Sprintf("/%s/%d", HostEndpointPrefix, id), http.MethodGet, nil, response)
	return response, err
}

func (c *Client) GetHostByName(ctx context.Context, name string) (*Host, error) {
	response := new(QueryResponseHost)
	err := c.requestSearchHelper(ctx, fmt.Sprintf("/%s", HostEndpointPrefix), http.MethodGet, "name", name, nil, response)
	if err != nil {
		return nil, err
	}
	if len(response.Results) == 0 {
		return nil, fmt.Errorf("Host not found: %s", name)

	}
	return &response.Results[0], err
}

func (c *Client) ListHost(ctx context.Context) (*QueryResponseHost, error) {
	response := new(QueryResponseHost)
	err := c.requestHelper(ctx, fmt.Sprintf("/%s", HostEndpointPrefix), http.MethodGet, nil, response)
	return response, err
}

func (c *Client) CreateHost(ctx context.Context, createRequest interface{}) (*Host, error) {
	response := new(Host)
	err := c.requestHelper(ctx, fmt.Sprintf("/%s", HostEndpointPrefix), http.MethodPost, createRequest, response)
	return response, err
}

func (c *Client) DeleteHost(ctx context.Context, id int) error {
	return c.requestHelper(ctx, fmt.Sprintf("/%s/%d", HostEndpointPrefix, id), http.MethodDelete, nil, nil)
}