package v1client

import (
	"fmt"
	"net/http"
)

func (c *V1ApiClient) ContainerSecurityListK8Images(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/kubernetesImages",
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

func (c *V1ApiClient) ContainerSecurityListContainerImageVulnerabilities(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/containerSecurity/vulnerabilities",
		filter,
		qp,
	)
}

func (c *V1ApiClient) ContainerSecurityAmazonEcsClusterGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/containerSecurity/amazonEcsClusters/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityAmazonEcsClusterUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/containerSecurity/amazonEcsClusters/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityAmazonEcsEvaluationEventLogsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/amazonEcsEvaluationEventLogs", p)
}

func (c *V1ApiClient) ContainerSecurityAmazonEcsImageOccurrencesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/amazonEcsImageOccurrences", p)
}

func (c *V1ApiClient) ContainerSecurityAmazonEcsSensorEventLogsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/amazonEcsSensorEventLogs", p)
}

func (c *V1ApiClient) ContainerSecurityAttestorsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/attestors", p)
}

func (c *V1ApiClient) ContainerSecurityAttestorsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/containerSecurity/attestors", p)
}

func (c *V1ApiClient) ContainerSecurityAttestorDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/containerSecurity/attestors/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityAttestorGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/containerSecurity/attestors/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityAttestorUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/containerSecurity/attestors/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityComplianceScanConfigurationGet(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/complianceScanConfiguration", p)
}

func (c *V1ApiClient) ContainerSecurityComplianceScanConfigurationUpdate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, "v3.0/containerSecurity/complianceScanConfiguration", p)
}

func (c *V1ApiClient) ContainerSecurityComplianceScanSummaryGet(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/complianceScanSummary", p)
}

func (c *V1ApiClient) ContainerSecurityComplianceScanVersionsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/complianceScanVersions", p)
}

func (c *V1ApiClient) ContainerSecurityCustomRulesetsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/customRulesets", p)
}

func (c *V1ApiClient) ContainerSecurityCustomRulesetsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/containerSecurity/customRulesets", p)
}

func (c *V1ApiClient) ContainerSecurityCustomRulesetDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/containerSecurity/customRulesets/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityCustomRulesetGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/containerSecurity/customRulesets/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityCustomRulesetRuleFileSourceGet(id string, ruleFileSourceId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/containerSecurity/customRulesets/%s/ruleFileSources/%s", id, ruleFileSourceId), p)
}

func (c *V1ApiClient) ContainerSecurityCustomRulesetUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/containerSecurity/customRulesets/%s/update", id), p)
}

func (c *V1ApiClient) ContainerSecurityFileIntegrityMonitoringRulesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/fileIntegrityMonitoringRules", p)
}

func (c *V1ApiClient) ContainerSecurityFileIntegrityMonitoringRulesCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/containerSecurity/fileIntegrityMonitoringRules", p)
}

func (c *V1ApiClient) ContainerSecurityFileIntegrityMonitoringRuleDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/containerSecurity/fileIntegrityMonitoringRules/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityFileIntegrityMonitoringRuleGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/containerSecurity/fileIntegrityMonitoringRules/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityFileIntegrityMonitoringRuleUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/containerSecurity/fileIntegrityMonitoringRules/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityGenerateServiceGatewayPassword(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/containerSecurity/generateServiceGatewayPassword", p)
}

func (c *V1ApiClient) ContainerSecurityKubernetesAuditEventLogsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/kubernetesAuditEventLogs", p)
}

func (c *V1ApiClient) ContainerSecurityKubernetesClusterGroupsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/kubernetesClusterGroups", p)
}

func (c *V1ApiClient) ContainerSecurityKubernetesClustersCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/containerSecurity/kubernetesClusters", p)
}

func (c *V1ApiClient) ContainerSecurityKubernetesClusterDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/containerSecurity/kubernetesClusters/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityKubernetesClusterUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/containerSecurity/kubernetesClusters/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityKubernetesEvaluationEventLogsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/kubernetesEvaluationEventLogs", p)
}

func (c *V1ApiClient) ContainerSecurityKubernetesImageOccurrencesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/kubernetesImageOccurrences", p)
}

func (c *V1ApiClient) ContainerSecurityKubernetesSensorEventLogsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/kubernetesSensorEventLogs", p)
}

func (c *V1ApiClient) ContainerSecurityManagedRulesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/managedRules", p)
}

func (c *V1ApiClient) ContainerSecurityManagedRuleGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/containerSecurity/managedRules/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityPoliciesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/policies", p)
}

func (c *V1ApiClient) ContainerSecurityPoliciesCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/containerSecurity/policies", p)
}

func (c *V1ApiClient) ContainerSecurityPolicyDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/containerSecurity/policies/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityPolicyGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/containerSecurity/policies/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityPolicyUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/containerSecurity/policies/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityRulesetsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/containerSecurity/rulesets", p)
}

func (c *V1ApiClient) ContainerSecurityRulesetsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/containerSecurity/rulesets", p)
}

func (c *V1ApiClient) ContainerSecurityRulesetDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/containerSecurity/rulesets/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityRulesetGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/containerSecurity/rulesets/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityRulesetUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/containerSecurity/rulesets/%s", id), p)
}

func (c *V1ApiClient) ContainerSecurityStartComplianceScan(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/containerSecurity/startComplianceScan", p)
}
