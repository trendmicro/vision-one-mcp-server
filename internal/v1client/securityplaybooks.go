package v1client

import "net/http"

func (c *V1ApiClient) SecurityPlaybooksPlaybooksList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/securityPlaybooks/playbooks", p)
}

func (c *V1ApiClient) SecurityPlaybooksPlaybooksRun(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/securityPlaybooks/playbooks/run", p)
}

func (c *V1ApiClient) SecurityPlaybooksTasksList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/securityPlaybooks/tasks", p)
}

func (c *V1ApiClient) SecurityPlaybooksTaskGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/securityPlaybooks/tasks/%s", id), p)
}

func (c *V1ApiClient) SecurityPlaybooksTaskActionsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/securityPlaybooks/tasks/%s/actions", id), p)
}
