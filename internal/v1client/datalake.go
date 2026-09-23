package v1client

import "net/http"

func (c *V1ApiClient) DatalakeDataPipelinesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/datalake/dataPipelines", p)
}

func (c *V1ApiClient) DatalakeDataPipelinesCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/datalake/dataPipelines", p)
}

func (c *V1ApiClient) DatalakeDataPipelinesDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/datalake/dataPipelines/delete", p)
}

func (c *V1ApiClient) DatalakeDataPipelineGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/datalake/dataPipelines/%s", id), p)
}

func (c *V1ApiClient) DatalakeDataPipelineUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/datalake/dataPipelines/%s", id), p)
}

func (c *V1ApiClient) DatalakeDataPipelinePackagesList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/datalake/dataPipelines/%s/packages", id), p)
}

func (c *V1ApiClient) DatalakeDataPipelinePackageGet(id string, packageId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/datalake/dataPipelines/%s/packages/%s", id, packageId), p)
}
