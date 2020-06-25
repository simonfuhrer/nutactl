package foreman

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

const (
	// Every Foreman API call has the following prefix to the path component
	// of the URL.  The client hepler functions utilize this to automatically
	// create endpoint URLs.
	FOREMAN_API_URL_PREFIX = "/api"
	// The Foreman API allows you to request a specific API version in the
	// Accept header of the HTTP request.  The two supported versions (at
	// the time of writing) are 1 and 2, which version 1 planning on being
	// deprecated after version 1.17.
	FOREMAN_API_VERSION = "2"
)

// ----------------------------------------------------------------------------
// Client / Server Configuration
// ----------------------------------------------------------------------------

type Server struct {
	// The URL of the API gateway
	URL url.URL
}

// Credentials used to authenticate the client against the remote server - in
// this case, the Foreman API
type ClientCredentials struct {
	Username string
	Password string
}

// Configurable features to apply the REST client
type ClientConfig struct {
	// Whether or not to verify the server's certificate/hostname.  This flag
	// is passed to the TLS config when initializing the REST client for API
	// communication.
	//
	// See 'pkg/crypto/tls/#Config.InsecureSkipVerify' for more information
	TLSInsecureEnabled bool
}

type Client struct {
	// Foreman URL used to communicate and interact with the API.
	server Server
	// Set of credentials to authenticate the client
	credentials ClientCredentials
	// Instance of the HTTP client used to communicate with the webservice.  After
	// the intial setup, the client should never modify or interact directly with
	// the underlying HTTP client and should instead use the helper functions.
	httpClient  *http.Client
	validateobj *validator.Validate
}

// NewClient creates a new instance of the REST client for communication with
// the API gateway.
func NewClient(s Server, c ClientCredentials, cfg ClientConfig) *Client {

	// Initialize the HTTP client for use by the provider.  The insecure flag
	// from the provider config is used when configuring the TLS settings of
	// the HTTP client.

	transCfg := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: cfg.TLSInsecureEnabled,
		},
	}

	httpClient := &http.Client{
		Transport: transCfg,
	}

	// Initialize and return the unauthenticated client.
	client := Client{
		httpClient:  httpClient,
		server:      s,
		credentials: c,
		validateobj: validator.New(),
	}
	return &client
}

func (client *Client) validate(o interface{}) error {
	err := client.validateobj.Struct(o)
	if err != nil {
		return err
	}
	return nil
}

func (client *Client) NewRequest(ctx context.Context, method string, endpoint string, body interface{}) (*http.Request, error) {
	if !isValidRequestMethod(method) {
		return nil, fmt.Errorf("Invalid HTTP request method: [%s]", method)
	}

	var contentBody io.Reader

	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	//var prettyJSON bytes.Buffer
	//_ = json.Indent(&prettyJSON, buf, "", "\t")
	//fmt.Println(string(prettyJSON.Bytes()))
	contentBody = bytes.NewReader(buf)

	// Build the URL for the request
	reqURL := client.server.URL
	if strings.HasPrefix(endpoint, "/") {
		reqURL.Path = FOREMAN_API_URL_PREFIX + endpoint
	} else {
		reqURL.Path = FOREMAN_API_URL_PREFIX + "/" + endpoint
	}

	// Create the request object, bubble up errors if any were encountered
	req, reqErr := http.NewRequest(
		strings.ToUpper(method),
		reqURL.String(),
		contentBody,
	)
	if reqErr != nil {

		return req, reqErr
	}
	// Add common meta-data and header information for the request
	req.Header.Add("Accept", "application/json,version="+FOREMAN_API_VERSION)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(client.credentials.Username, client.credentials.Password)

	return req.WithContext(ctx), nil
}

func (client *Client) requestSearchHelper(ctx context.Context, path, method, key, name string, request interface{}, output interface{}) error {
	req, err := client.NewRequest(ctx, method, path, request)
	if err != nil {
		return err
	}
	reqQuery := req.URL.Query()
	reqQuery.Add("search", fmt.Sprintf("%s=\"%s\"", key, name))

	req.URL.RawQuery = reqQuery.Encode()
	err = client.Do(req, &output)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) requestHelper(ctx context.Context, path, method string, request interface{}, output interface{}) error {
	if method == http.MethodPost || method == http.MethodPut {
		err := client.validate(request)
		if err != nil {
			return err
		}
	}

	req, err := client.NewRequest(ctx, method, path, request)
	if err != nil {
		return err
	}

	err = client.Do(req, &output)
	if err != nil {
		return err
	}

	return nil
}

// isValidRequestMethod is a helper function used to determine if an HTTP
// request method is valid.
//
// NOTE(ALL): Go's HTTP client does not support sending a request with
//   the 'CONNECT' method and therefore is not counted as a valid request
//   method. See http.Transport, http.Client for more information.
func isValidRequestMethod(method string) bool {
	// Slice of valid HTTP methods for sending and creating requests
	validHTTPMethods := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodTrace,
	}
	// list isn't large - use linear search to validate the method.  Use
	// strings.EqualFold to perform case-insensitive comparisons
	for _, value := range validHTTPMethods {
		if strings.EqualFold(value, method) {
			return true
		}
	}
	return false
}

func (client *Client) Do(request *http.Request, v interface{}) error {
	// Send the request to the server
	resp, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if resp.Request.Method == http.MethodDelete {
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("HTTP Error, endpoint: [%s], statusCode: [%d], error [%v]", request.URL, resp.StatusCode, string(body))
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return fmt.Errorf("error unmarshalling json: %s", err)
		}
	}

	return err
}
