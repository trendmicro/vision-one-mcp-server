package v1client

import (
	"fmt"
	"net/http"
)

func (c *V1ApiClient) ContainerSecurityListPolicies(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/policies",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityGetPolicy(policyID string) (*http.Response, error) {
	return c.genericGet(
		fmt.Sprintf("v3.0/containerSecurity/policies/%s", policyID),
	)
}

func (c *V1ApiClient) ContainerSecurityListRuntimeRules(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/managedRules",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityGetRuntimeRule(ruleID string) (*http.Response, error) {
	return c.genericGet(
		fmt.Sprintf("v3.0/containerSecurity/managedRules/%s", ruleID),
	)
}

func (c *V1ApiClient) ContainerSecurityListRulesets(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/rulesets",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityGetRuleset(rulesetID string) (*http.Response, error) {
	return c.genericGet(
		fmt.Sprintf("v3.0/containerSecurity/rulesets/%s", rulesetID),
	)
}

func (c *V1ApiClient) ContainerSecurityListK8Images(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/kubernetesImages",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityListK8ImageOccurences(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/kubernetesImageOccurrences",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityListECSImageOccurences(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/amazonEcsImageOccurrences",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityListK8Clusters(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/kubernetesClusters",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityGetK8ClusterDetails(clusterID string) (*http.Response, error) {
	return c.genericGet(
		fmt.Sprintf("v3.0/containerSecurity/kubernetesClusters/%s", clusterID),
	)
}

func (c *V1ApiClient) ContainerSecurityListECSClusters(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/amazonEcsClusters",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityGetECSClusterDetails(clusterID string) (*http.Response, error) {
	return c.genericGet(
		fmt.Sprintf("v3.0/containerSecurity/amazonEcsClusters/%s", clusterID),
	)
}

func (c *V1ApiClient) ContainerSecurityListContainerImageVulnerabilities(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/vulnerabilities",
		filter,
		qp,
	)
}
