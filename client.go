// Package wazirx provides wazirx.com's Public Rest API client.
// You can read the API server documentation at https://github.com/WazirX/wazirx-api
package wazirx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	// DefaultBaseURL is the default server URL.
	DefaultBaseURL = "https://api.wazirx.com"
)

// WithCtx applies 'ctx' to the the http.Request and returns *http.Request
// The provided ctx and req must be non-nil
func WithCtx(ctx context.Context, req *http.Request) *http.Request {
	if req == nil {
		panic("nil http.Request")
	}
	return req.WithContext(ctx)
}

// Client for accessing different endpoints of the API
type Client struct {
	// HTTPClient is a reusable http client instance.
	HTTP *http.Client
	// BaseURL is the REST endpoints URL of the api server
	BaseURL *url.URL
}

// Do sends the http.Request and unmarshalls the JSON response into 'target'
func (c Client) Do(req *http.Request, target interface{}) error {
	if req == nil {
		return errors.New("invalid Request")
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
		return errors.Wrap(err, "request failed")
	}

	if resp.StatusCode != http.StatusOK {
		return ErrAPI{resp}
	}

	defer func() {
		// Ensure the response body is fully read and closed
		// before we reconnect, so that we reuse the same TCPconnection.
		const maxBodySlurpSize = 2 << 10
		if resp.ContentLength == -1 || resp.ContentLength <= maxBodySlurpSize {
			io.CopyN(ioutil.Discard, resp.Body, maxBodySlurpSize)
		}
		resp.Body.Close()
	}()

	var buf bytes.Buffer
	return json.NewDecoder(io.TeeReader(resp.Body, &buf)).Decode(target)
}

// ErrAPI is returned by API calls when the response status code isn't 200.
type ErrAPI struct {
	// Response from the request which returned error.
	Response *http.Response
}

// Error implements the error interface.
func (err ErrAPI) Error() (errStr string) {
	if err.Response != nil {
		errStr += fmt.Sprintf(
			"request to %s returned %d (%s)",
			err.Response.Request.URL,
			err.Response.StatusCode,
			http.StatusText(err.Response.StatusCode),
		)
	}
	return errStr
}
