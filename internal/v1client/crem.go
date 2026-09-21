package v1client

import (
	"fmt"
	"net/http"
	"time"
)

type QueryParameters struct {
	OrderBy                string    `url:"orderBy,omitempty"`
	Top                    int       `url:"top,omitempty"`
	FirstSeenStartDateTime time.Time `url:"firstSeenStartDateTime,omitempty"`
	FirstSeenEndDateTime   time.Time `url:"firstSeenEndDateTime,omitempty"`
	SkipToken              string    `url:"skipToken,omitempty"`
	NextBatchToken         string    `url:"nextBatchToken,omitempty"`
	NextLink               string    `url:"-"`

	StartDateTime time.Time `url:"startDateTime,omitempty"`
	EndDateTime   time.Time `url:"endDateTime,omitempty"`

	LastDetectedStartDateTime time.Time `url:"lastDetectedStartDateTime,omitempty"`
	LastDetectedEndDateTime   time.Time `url:"lastDetectedEndDateTime,omitempty"`

	FirstDetectedStartDateTime time.Time `url:"firstDetectedStartDateTime,omitempty"`
	FirstDetectedEndDateTime   time.Time `url:"firstDetectedEndDateTime,omitempty"`

	DetectedStartDateTime time.Time `url:"detectedStartDateTime,omitempty"`
	DetectedEndDateTime   time.Time `url:"detectedEndDateTime,omitempty"`

	IngestedStartDateTime time.Time `url:"ingestedStartDateTime,omitempty"`
	IngestedEndDateTime   time.Time `url:"ingestedEndDateTime,omitempty"`
}

func (c *V1ApiClient) CREMListAttackSurfaceDevices(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/attackSurfaceDevices",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListAttackSurfaceDomainAccounts(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/attackSurfaceDomainAccounts",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListAttackSurfaceServiceAccounts(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/attackSurfaceServiceAccounts",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListAttackSurfaceGlobalFQDNs(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/attackSurfaceGlobalFqdns",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListAttackSurfacePublicIPs(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/attackSurfacePublicIpAddresses",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListAttackSurfaceCloudAssets(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/attackSurfaceCloudAssets",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListHighRiskUsers(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/highRiskUsers",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMGetAttackSurfaceCloudAssetProfile(resourceId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("v3.0/asrm/attackSurfaceCloudAssets/%s", resourceId))
}

func (c *V1ApiClient) CREMListAttackSurfaceCloudAssetRiskIndicators(
	resourceId string,
	filter string,
	queryParams QueryParameters,
) (*http.Response, error) {
	return c.searchAndFilter(
		fmt.Sprintf("v3.0/asrm/attackSurfaceCloudAssets/%s/riskIndicatorEvents", resourceId),
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListAttackSurfaceLocalApps(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/attackSurfaceLocalApps",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMGetAttackSurfaceLocalAppProfile(resourceId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("v3.0/asrm/attackSurfaceLocalApps/%s", resourceId))
}

func (c *V1ApiClient) CREMGetAttackSurfaceLocalAppRiskIndicators(
	resourceId,
	filter string,
	queryParams QueryParameters,
) (*http.Response, error) {
	return c.searchAndFilter(
		fmt.Sprintf("v3.0/asrm/attackSurfaceLocalApps/%s/riskIndicatorEvents", resourceId),
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListAttackSurfaceLocalAppDevices(
	resourceId,
	filter string,
	queryParams QueryParameters,
) (*http.Response, error) {
	return c.searchAndFilter(
		fmt.Sprintf("v3.0/asrm/attackSurfaceLocalApps/%s/devices", resourceId),
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListAttackSurfaceLocalAppExecutableFiles(
	resourceId,
	filter string,
	queryParams QueryParameters,
) (*http.Response, error) {
	return c.searchAndFilter(
		fmt.Sprintf("v3.0/asrm/attackSurfaceLocalApps/%s/executableFiles", resourceId),
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CREMListCustomTags(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/asrm/attackSurfaceCustomTags",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) CremAccountCompromiseEventDefinitionsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/accountCompromiseEventDefinitions", p)
}

func (c *V1ApiClient) CremAccountCompromiseIndicatorsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/accountCompromiseIndicators", p)
}

func (c *V1ApiClient) CremAccountCompromiseRiskIndicatorEventsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/accountCompromiseRiskIndicatorEvents", p)
}

func (c *V1ApiClient) CremAccountCompromiseRiskIndicatorEventGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/accountCompromiseRiskIndicatorEvents/%s", id), p)
}

func (c *V1ApiClient) CremAnomalyDetectionRiskIndicatorEventsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/anomalyDetectionRiskIndicatorEvents", p)
}

func (c *V1ApiClient) CremAnomalyDetectionRiskIndicatorEventGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/anomalyDetectionRiskIndicatorEvents/%s", id), p)
}

func (c *V1ApiClient) CremAssetGroupsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/assetGroups", p)
}

func (c *V1ApiClient) CremAttackSurfaceAssetsUpdateCriticality(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/asrm/attackSurfaceAssets/updateCriticality", p)
}

func (c *V1ApiClient) CremAttackSurfaceCloudAssetsUpdate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/asrm/attackSurfaceCloudAssets/update", p)
}

func (c *V1ApiClient) CremAttackSurfaceDevicesUpdate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/asrm/attackSurfaceDevices/update", p)
}

func (c *V1ApiClient) CremCloudVmVulnerabilitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/cloudVmVulnerabilities", p)
}

func (c *V1ApiClient) CremContainerVulnerabilitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/containerVulnerabilities", p)
}

func (c *V1ApiClient) CremHighRiskDevicesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/highRiskDevices", p)
}

func (c *V1ApiClient) CremHighRiskDeviceGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/highRiskDevices/%s", id), p)
}

func (c *V1ApiClient) CremHighRiskUserGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/highRiskUsers/%s", id), p)
}

func (c *V1ApiClient) CremInternalAssetVulnerabilitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/internalAssetVulnerabilities", p)
}

func (c *V1ApiClient) CremInternetFacingAssetVulnerabilitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/internetFacingAssetVulnerabilities", p)
}

func (c *V1ApiClient) CremSecurityPostureGet(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/securityPosture", p)
}

func (c *V1ApiClient) CremServerlessFunctionVulnerabilitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/serverlessFunctionVulnerabilities", p)
}

func (c *V1ApiClient) CremVulnerabilityGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s", id), p)
}

func (c *V1ApiClient) CremVulnerabilityAffectedCloudStorageAssetsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s/affectedCloudStorageAssets", id), p)
}

func (c *V1ApiClient) CremVulnerabilityAffectedCloudVmsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s/affectedCloudVms", id), p)
}

func (c *V1ApiClient) CremVulnerabilityAffectedContainerClustersList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s/affectedContainerClusters", id), p)
}

func (c *V1ApiClient) CremVulnerabilityAffectedContainerImagesList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s/affectedContainerImages", id), p)
}

func (c *V1ApiClient) CremVulnerabilityAffectedDevicesList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s/affectedDevices", id), p)
}

func (c *V1ApiClient) CremVulnerabilityAffectedGlobalFqdnsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s/affectedGlobalFqdns", id), p)
}

func (c *V1ApiClient) CremVulnerabilityAffectedServerlessFunctionLayersList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s/affectedServerlessFunctionLayers", id), p)
}

func (c *V1ApiClient) CremVulnerabilityAffectedServerlessFunctionsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/asrm/vulnerabilities/%s/affectedServerlessFunctions", id), p)
}

func (c *V1ApiClient) CremVulnerableDevicesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/asrm/vulnerableDevices", p)
}
