package v1client

import (
	"bytes"
	"encoding/json"
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

type ClientOptions struct {
	ApiKey string
	// The Vision One service region.
	Region string
	// The Host to use. Use for pre-prod environments. If specified will be used instead of [Region]
	Host      string
	UserAgent string
}

func NewV1ApiClient(co ClientOptions) (*V1ApiClient, error) {
	if co.Host == "" && co.Region == "" {
		panic("must specify either host or region when creating a client")
	}

	if co.Host != "" {
		hostUrl := fmt.Sprintf("https://%s/", co.Host)
		baseUrl, err := url.Parse(hostUrl)
		if err != nil {
			return nil, err
		}

		return &V1ApiClient{
			client:  &http.Client{},
			apiKey:  co.ApiKey,
			baseUrl: baseUrl,
		}, nil
	}

	baseUrl, err := getRegionURL(co.Region)
	if err != nil {
		return nil, err
	}

	return &V1ApiClient{
		client:  &http.Client{},
		apiKey:  co.ApiKey,
		baseUrl: baseUrl,
	}, nil
}

func getRegionURL(region string) (*url.URL, error) {
	if region == "us" {
		u, err := url.Parse("https://api.xdr.trendmicro.com/")
		if err != nil {
			return nil, err
		}
		return u, nil
	}

	if region == "jp" {
		u, err := url.Parse("https://api.xdr.trendmicro.co.jp/")
		if err != nil {
			return nil, err
		}
		return u, nil
	}

	u, err := url.Parse(fmt.Sprintf("https://api.%s.xdr.trendmicro.com/", region))
	if err != nil {
		return nil, err
	}

	return u, nil
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

// withHeader adds a custom header to the request.
func withHeader(name, value string) requestOptionFunc {
	return func(r *http.Request) {
		if value == "" {
			return
		}
		r.Header.Add(name, value)
	}
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

func (c *V1ApiClient) genericJSONPost(path string, body any, options ...requestOptionFunc) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	opts := append([]requestOptionFunc{withContentTypeJSON()}, options...)
	r, err := c.newRequest(
		http.MethodPost,
		path,
		bytes.NewReader(b),
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) genericPost(path string) (*http.Response, error) {
	r, err := c.newRequest(
		http.MethodPost,
		path,
		http.NoBody,
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) genericJSONPatch(path string, body any) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPatch,
		path,
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) genericDelete(path string) (*http.Response, error) {
	r, err := c.newRequest(
		http.MethodDelete,
		path,
		http.NoBody,
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}
