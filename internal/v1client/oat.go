package v1client

import "net/http"

func (c *V1ApiClient) ObservedAttackTechniquesList(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/oat/detections",
		filter,
		qp,
	)
}

func (c *V1ApiClient) OatDataPipelinesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/oat/dataPipelines", p)
}

func (c *V1ApiClient) OatDataPipelinesCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/oat/dataPipelines", p)
}

func (c *V1ApiClient) OatDataPipelinesDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/oat/dataPipelines/delete", p)
}

func (c *V1ApiClient) OatDataPipelineGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/oat/dataPipelines/%s", id), p)
}

func (c *V1ApiClient) OatDataPipelineUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/oat/dataPipelines/%s", id), p)
}

func (c *V1ApiClient) OatDataPipelinePackagesList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/oat/dataPipelines/%s/packages", id), p)
}

func (c *V1ApiClient) OatDataPipelinePackageGet(id string, packageId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/oat/dataPipelines/%s/packages/%s", id, packageId), p)
}
