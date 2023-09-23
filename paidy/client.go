package paidy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

// API endpoint base constants
const (
	APIEndpoint = "https://api.paidy.com"
)

const (
	APIVersion = "2018-04-10"
)

// Client type
type Client struct {
	secretKey  string
	endpoint   *url.URL
	version    string
	httpClient *http.Client
}

// ClientOption type
type ClientOption func(*Client) error

// New returns a new pay client instance.
func New(secretKey string, options ...ClientOption) (*Client, error) {
	if secretKey == "" {
		return nil, errors.New("missing api secret key")
	}
	c := &Client{
		secretKey:  secretKey,
		httpClient: http.DefaultClient,
	}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	if c.version == "" {
		c.version = APIVersion
	}
	if c.endpoint == nil {
		u, err := url.Parse(APIEndpoint)
		if err != nil {
			return nil, err
		}
		c.endpoint = u
	}
	return c, nil
}

// WithHTTPClient function
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}

// mergeQuery method
func (c *Client) mergeQuery(path string, q any) (string, error) {
	v := reflect.ValueOf(q)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return path, nil
	}

	u, err := url.Parse(path)
	if err != nil {
		return path, err
	}

	qs, err := query.Values(q)
	if err != nil {
		return path, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewRequest method
func (c *Client) NewRequest(method, path string, body any) (*http.Request, error) {
	switch method {
	case http.MethodGet, http.MethodDelete:
		if body != nil {
			merged, err := c.mergeQuery(path, body)
			if err != nil {
				return nil, err
			}
			path = merged
		}
	}
	u, err := c.endpoint.Parse(path)
	if err != nil {
		return nil, err
	}

	var reqBody io.ReadWriter
	switch method {
	case http.MethodPost, http.MethodPut:
		if body != nil {
			b, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			reqBody = bytes.NewBuffer(b)
		}
	}

	req, err := http.NewRequest(method, u.String(), reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Paidy-Version", c.version)
	req.Header.Set("Authorization", "Bearer "+c.secretKey)
	return req, nil
}

// Do method
func (c *Client) Do(ctx context.Context, req *http.Request, v any) (*http.Response, error) {
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer resp.Body.Close()

	// error response
	if !isSuccess(resp) {
		errResp := new(APIError)
		if err := json.NewDecoder(resp.Body).Decode(errResp); err != nil {
			return resp, err
		}
		return resp, errResp
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
				return resp, err
			}
		}
	}
	return resp, err
}

func isSuccess(resp *http.Response) bool {
	code := resp.StatusCode
	return code/100 == 2
}

type APIError struct {
	Reference   string `json:"reference,omitempty"`
	Status      string `json:"status,omitempty"`
	Code        string `json:"code,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// Error method
func (e *APIError) Error() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "paidy: Error %s ", e.Code)
	return buf.String()
}
