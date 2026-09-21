package v1client

import "net/http"

func (c *V1ApiClient) SandboxAnalysisResultsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/sandbox/analysisResults", p)
}

func (c *V1ApiClient) SandboxAnalysisResultGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/sandbox/analysisResults/%s", id), p)
}

func (c *V1ApiClient) SandboxAnalysisResultInvestigationPackageGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/sandbox/analysisResults/%s/investigationPackage", id), p)
}

func (c *V1ApiClient) SandboxAnalysisResultReportGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/sandbox/analysisResults/%s/report", id), p)
}

func (c *V1ApiClient) SandboxAnalysisResultSuspiciousObjectsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/sandbox/analysisResults/%s/suspiciousObjects", id), p)
}

func (c *V1ApiClient) SandboxFilesAnalyze(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/sandbox/files/analyze", p)
}

func (c *V1ApiClient) SandboxSubmissionUsageGet(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/sandbox/submissionUsage", p)
}

func (c *V1ApiClient) SandboxTasksList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/sandbox/tasks", p)
}

func (c *V1ApiClient) SandboxTaskGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/sandbox/tasks/%s", id), p)
}

func (c *V1ApiClient) SandboxUrlsAnalyze(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/sandbox/urls/analyze", p)
}
