package v1client

import (
	"fmt"
	"net/http"
)

func (c *V1ApiClient) EndpointSecurityListEndpoints(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/endpointSecurity/endpoints",
		filter,
		qp,
	)
}

func (c *V1ApiClient) EndpointSecurityGetEndpoint(id string) (*http.Response, error) {
	return c.genericGet(
		fmt.Sprintf("v3.0/endpointSecurity/endpoints/%s", id),
	)
}

func (c *V1ApiClient) EndpointSecurityListTasks(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/endpointSecurity/tasks",
		filter,
		qp,
	)
}

func (c *V1ApiClient) EndpointSecurityGetTask(taskID string) (*http.Response, error) {
	return c.genericGet(
		fmt.Sprintf("v3.0/endpointSecurity/tasks/%s", taskID),
	)
}

func (c *V1ApiClient) EndpointSecurityListVersionControlPolicies(qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/endpointSecurity/versionControlPolicies",
		"",
		qp,
	)
}

func (c *V1ApiClient) EndpointSecurityListAgentUpdatePolicies() (*http.Response, error) {
	return c.genericGet("v3.0/endpointSecurity/versionControlPolicies/agentUpdatePolicies")
}

func (c *V1ApiClient) EndpointSecurityEndpointsApplySensorPolicy(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/endpointSecurity/endpoints/applySensorPolicy", p)
}

func (c *V1ApiClient) EndpointSecurityEndpointsDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/endpointSecurity/endpoints/delete", p)
}

func (c *V1ApiClient) EndpointSecurityEndpointsExport(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/endpointSecurity/endpoints/export", p)
}

func (c *V1ApiClient) EndpointSecurityEndpointsRemoveOverriddenSensorPolicy(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/endpointSecurity/endpoints/removeOverriddenSensorPolicy", p)
}

func (c *V1ApiClient) EndpointSecuritySchedulesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/endpointSecurity/schedules", p)
}

func (c *V1ApiClient) EndpointSecuritySchedulesCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/endpointSecurity/schedules", p)
}

func (c *V1ApiClient) EndpointSecuritySchedulesDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/endpointSecurity/schedules/delete", p)
}

func (c *V1ApiClient) EndpointSecuritySchedulesUpdate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/endpointSecurity/schedules/update", p)
}

func (c *V1ApiClient) EndpointSecurityScheduleGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/endpointSecurity/schedules/%s", id), p)
}

func (c *V1ApiClient) EndpointSecurityKernelSupportPackageUpdatePoliciesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/endpointSecurity/versionControlPolicies/kernelSupportPackageUpdatePolicies", p)
}

func (c *V1ApiClient) EndpointSecurityVersionControlPolicyDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/endpointSecurity/versionControlPolicies/%s", id), p)
}

func (c *V1ApiClient) EndpointSecurityVersionControlPolicyPrioritiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/endpointSecurity/versionControlPolicyPriorities", p)
}

func (c *V1ApiClient) EndpointSecurityVersionControlPolicyPriorityUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/endpointSecurity/versionControlPolicyPriorities/%s", id), p)
}
