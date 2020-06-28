// maybe someone find a better solution. For now this is more than a workaround.
//
package foreman

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cheekybits/genny/generic"
)

type Type generic.Type

type Value generic.Number

type Path generic.Number

var TypeEndpointPrefix = fmt.Sprintf("%s", strings.ToLower("Path"))

type QueryResponseType struct {
	QueryResponse
	Results []Type `json:"results"`
}

func (c *Client) GetType(ctx context.Context, idOrName string) (*Type, error) {
	if id, err := strconv.Atoi(idOrName); err == nil {
		return c.GetTypeByID(ctx, int(id))
	}
	return c.GetTypeByName(ctx, idOrName)
}

func (c *Client) GetTypeByID(ctx context.Context, id int) (*Type, error) {
	response := new(Type)
	err := c.requestHelper(ctx, fmt.Sprintf("/%s/%d", TypeEndpointPrefix, id), http.MethodGet, nil, response)
	return response, err
}

func (c *Client) GetTypeByName(ctx context.Context, name string) (*Type, error) {
	response := new(QueryResponseType)
	filter := fmt.Sprintf("%s=\"%s\"", strings.ToLower("Value"), name)
	err := c.requestSearchHelper(ctx, fmt.Sprintf("/%s", TypeEndpointPrefix), http.MethodGet, filter, nil, response)
	if err != nil {
		return nil, err
	}
	if len(response.Results) == 0 {
		return nil, fmt.Errorf("Type not found")

	}
	return &response.Results[0], err
}

func (c *Client) ListType(ctx context.Context) (*QueryResponseType, error) {
	response := new(QueryResponseType)
	err := c.requestHelper(ctx, fmt.Sprintf("/%s", TypeEndpointPrefix), http.MethodGet, nil, response)
	return response, err
}

func (c *Client) SearchType(ctx context.Context, filter string) (*QueryResponseType, error) {
	response := new(QueryResponseType)
	err := c.requestSearchHelper(ctx, fmt.Sprintf("/%s", TypeEndpointPrefix), http.MethodGet, filter, nil, response)
	return response, err
}

func (c *Client) CreateType(ctx context.Context, createRequest interface{}) (*Type, error) {
	response := new(Type)
	err := c.requestHelper(ctx, fmt.Sprintf("/%s", TypeEndpointPrefix), http.MethodPost, createRequest, response)
	return response, err
}

func (c *Client) DeleteType(ctx context.Context, id int) error {
	return c.requestHelper(ctx, fmt.Sprintf("/%s/%d", TypeEndpointPrefix, id), http.MethodDelete, nil, nil)
}
