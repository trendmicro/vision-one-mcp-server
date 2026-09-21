package v1client

import "net/http"

func (c *V1ApiClient) BusinessProfileGet(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/business/profile", p)
}

func (c *V1ApiClient) BusinessProfileUpdate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, "v3.0/business/profile", p)
}
