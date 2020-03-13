// Package wazix provides wazirx.com's Public Rest API client.
// You can read the API server documentation at https://github.com/WazirX/wazirx-api
package wazix

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	// DefaultBaseURL is the default server URL.
	DefaultBaseURL = "https://api.wazirx.com"
)

// Requester is implemented by any value that has a Request method.
type Requester interface {
	Request() (*http.Request, error)
}

// RequesterFunc implements Requester
type RequesterFunc func() (*http.Request, error)

// Request invokes 'f'
func (f RequesterFunc) Request() (*http.Request, error) {
	return f()
}

// WithCtx applies 'ctx' to the the http.Request and returns *http.Request
// The provided ctx and req must be non-nil
func WithCtx(ctx context.Context, req *http.Request) *http.Request {
	if req == nil {
		panic("nil http.Request")
	}
	return req.WithContext(ctx)
}

// Client parameters
type Client struct {
	// HTTPClient is a reusable http client instance.
	HTTP *http.Client
	// BaseURL is the REST endpoints URL of the api server
	BaseURL *url.URL
}

// Do sends the http.Request and unmarshalls the JSON response into 'target'
func (c Client) Do(r Requester, target interface{}) (*http.Response, error) {
	if r == nil {
		return nil, errors.New("invalid Requester")
	}

	req, err := r.Request()
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	if c.HTTP == nil {
		c.HTTP = http.DefaultClient
		c.HTTP.Transport = http.DefaultTransport
		c.HTTP.Timeout = 15 * time.Second
	}

	if c.BaseURL != nil {
		req.URL.Scheme = c.BaseURL.Scheme
		req.URL.Host = c.BaseURL.Host
	}

	// make request to the api and read the response
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}

	var buf bytes.Buffer
	return resp, json.NewDecoder(io.TeeReader(resp.Body, &buf)).Decode(target)
}
