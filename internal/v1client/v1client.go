package v1client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

var version = "1.0.0"

type V1ApiClient struct {
	client    *http.Client
	apiKey    string
	baseUrl   *url.URL
	UserAgent string
}

func NewV1ApiClient(apiKey, v1Region string) *V1ApiClient {
	var u string
	if v1Region == "us" {
		u = "https://api.xdr.trendmicro.com/"
	} else {
		u = fmt.Sprintf("https://api.%s.xdr.trendmicro.com/", v1Region)
	}
	// We ignore the error below. We know the URL will be valid due to previous validation.
	baseUrl, _ := url.Parse(u)
	return &V1ApiClient{
		client:  &http.Client{},
		apiKey:  apiKey,
		baseUrl: baseUrl,
	}
}

// requestOptionFunc is a function that modifies an HTTP request.
// For example, adding headers.
type requestOptionFunc func(*http.Request)

// Creates new requests.
// Path MUST NOT start with a "/". E.g. "v3.0/service/call".
// Path is based on the client base URL. "v3.0/service/call" -> https://api.xdr.trendmicro.com/v3.0/service/call
// Options allow the caller to specify methods to modify the request.
func (c *V1ApiClient) newRequest(method string, path string, body io.Reader, options ...requestOptionFunc) (*http.Request, error) {
	u, err := c.baseUrl.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	// Add headers needed for all requests
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	if c.UserAgent == "" {
		req.Header.Add("User-Agent", fmt.Sprintf("vision-one-client/%s", version))
	}
	req.Header.Add("User-Agent", c.UserAgent)

	for _, opt := range options {
		opt(req)
	}

	return req, err
}

// Adds the TMV1-Filter headers with the given filter
func withFilter(filter string) requestOptionFunc {
	return func(r *http.Request) {
		if filter == "" {
			return
		}
		r.Header.Add("TMV1-Filter", filter)
	}
}

// Replaces request RawQuery with params.
func withUrlParameters(params url.Values) requestOptionFunc {
	return func(r *http.Request) {
		r.URL.RawQuery = params.Encode()
	}
}

func withContentTypeJSON() requestOptionFunc {
	return func(r *http.Request) {
		r.Header.Add("content-type", "application/json")
	}
}

func contentTypeJSON(r *http.Request) {
	r.Header.Add("content-type", "application/json")
}

func (c *V1ApiClient) searchAndFilter(path, filter string, queryParams any) (*http.Response, error) {
	p, err := query.Values(queryParams)
	if err != nil {
		return nil, err
	}
	r, err := c.newRequest(
		http.MethodGet,
		path,
		http.NoBody,
		withFilter(filter),
		withUrlParameters(p),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) genericGet(path string) (*http.Response, error) {
	r, err := c.newRequest(
		http.MethodGet,
		path,
		http.NoBody,
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}
