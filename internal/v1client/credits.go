package v1client

import (
	"net/http"
)

// CreditsListEndpoints lists endpoints for credit analysis
func (c *V1ApiClient) CreditsListEndpoints(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/endpointInventory",
		filter,
		queryParams,
	)
}

// CreditsListDataLakePipelines lists active data lake pipelines consuming credits
func (c *V1ApiClient) CreditsListDataLakePipelines(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/datalake/pipelines",
		filter,
		queryParams,
	)
}

// CreditsListOATDetections lists OAT detections for credit analysis
func (c *V1ApiClient) CreditsListOATDetections(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/xdr/oat/detections",
		filter,
		queryParams,
	)
}

// CreditsListOATPipelines lists Observed Attack Techniques pipelines
func (c *V1ApiClient) CreditsListOATPipelines(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/oat/pipelines",
		filter,
		queryParams,
	)
}

// CreditsListSandboxAnalysisResults lists sandbox analysis results for credit usage
func (c *V1ApiClient) CreditsListSandboxAnalysisResults(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/sandbox/analysisResults",
		filter,
		queryParams,
	)
}

// CreditsListSandboxSubmissions lists sandbox submissions for credit usage analysis
func (c *V1ApiClient) CreditsListSandboxSubmissions(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/sandbox/submissions",
		filter,
		queryParams,
	)
}

// CreditsListWorkbenchAlerts lists workbench alerts for investigation activity analysis
func (c *V1ApiClient) CreditsListWorkbenchAlerts(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/workbench/alerts",
		filter,
		queryParams,
	)
}

// CreditsGetSearchStatistics gets search statistics for credit usage analysis
func (c *V1ApiClient) CreditsGetSearchStatistics(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/search/statistics",
		filter,
		queryParams,
	)
}

// CreditsListDataRetentionModels lists data retention models
func (c *V1ApiClient) CreditsListDataRetentionModels(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/datalake/dataRetentionModels",
		filter,
		queryParams,
	)
}

// CreditsListCREMHighRiskDevices lists high risk devices from CREM for enhanced analysis
func (c *V1ApiClient) CreditsListCREMHighRiskDevices(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/highRiskDevices",
		filter,
		queryParams,
	)
}

// CreditsListCREMHighRiskUsers lists high risk users from CREM for enhanced analysis
func (c *V1ApiClient) CreditsListCREMHighRiskUsers(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/highRiskUsers",
		filter,
		queryParams,
	)
}

// CreditsListCREMCompromiseIndicators lists compromise indicators from CREM
func (c *V1ApiClient) CreditsListCREMCompromiseIndicators(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/compromiseIndicators",
		filter,
		queryParams,
	)
}

// CreditsGetAllocation gets current credit allocation across all services
func (c *V1ApiClient) CreditsGetAllocation() (*http.Response, error) {
	return c.genericGet("v3.0/credits/allocation")
}

// CreditsGetBalance gets remaining credit balance and usage statistics
func (c *V1ApiClient) CreditsGetBalance() (*http.Response, error) {
	return c.genericGet("v3.0/credits/balance")
}

// CreditsGetUsageStatistics gets detailed credit usage statistics by service
func (c *V1ApiClient) CreditsGetUsageStatistics(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/credits/usage/statistics",
		filter,
		queryParams,
	)
}

// CreditsGetServiceLimits gets credit limits and thresholds for all services
func (c *V1ApiClient) CreditsGetServiceLimits() (*http.Response, error) {
	return c.genericGet("v3.0/credits/limits")
}