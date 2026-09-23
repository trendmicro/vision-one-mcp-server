package v1client

import "net/http"

func (c *V1ApiClient) DmmCustomFiltersList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/dmm/customFilters", p)
}

func (c *V1ApiClient) DmmCustomModelsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/dmm/customModels", p)
}

func (c *V1ApiClient) DmmExceptionsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/dmm/exceptions", p)
}

func (c *V1ApiClient) DmmExceptionsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/dmm/exceptions", p)
}

func (c *V1ApiClient) DmmExceptionsDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/dmm/exceptions/delete", p)
}

func (c *V1ApiClient) DmmExceptionGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/dmm/exceptions/%s", id), p)
}

func (c *V1ApiClient) DmmExceptionUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/dmm/exceptions/%s", id), p)
}

func (c *V1ApiClient) DmmModelsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/dmm/models", p)
}

func (c *V1ApiClient) DmmModelUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/dmm/models/%s", id), p)
}
