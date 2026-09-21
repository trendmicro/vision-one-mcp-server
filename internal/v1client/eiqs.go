package v1client

import "net/http"

func (c *V1ApiClient) EiqsEndpointsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/eiqs/endpoints", p)
}
