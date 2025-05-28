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
