package v1client

import "net/http"

func (c *V1ApiClient) FileSecurityStoragesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/fileSecurity/storages", p)
}

func (c *V1ApiClient) FileSecurityStoragesUpdate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/fileSecurity/storages/update", p)
}
