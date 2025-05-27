package v1client

import (
	"fmt"
	"net/http"
	"time"
)

type QueryParameters struct {
	OrderBy                   string    `url:"orderBy,omitempty"`
	Top                       int       `url:"top,omitempty"`
	LastDetectedStartDateTime time.Time `url:"lastDetectedStartDateTime,omitempty"`
	LastDetectedEndDateTime   time.Time `url:"lastDetectedEndDateTime,omitempty"`
	FirstSeenStartDateTime    time.Time `url:"firstSeenStartDateTime,omitempty"`
	FirstSeenEndDateTime      time.Time `url:"firstSeenEndDateTime,omitempty"`
	SkipToken                 string    `url:"skipToken,omitempty"`
	NextBatchToken            string    `url:"nextBatchToken,omitempty"`
	NextLink                  string    `url:"-"`

	StartDateTime time.Time `url:"startDateTime,omitempty"`
	EndDateTime   time.Time `url:"endDateTime,omitempty"`

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
		"v3.0/asrm/attackSurfaceServiceAccoutns",
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

func (c *V1ApiClient) CREMGetSecurityPosture() (*http.Response, error) {
	return c.genericGet("v3.0/asrm/securityPosture")
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
		"v3.0/asrm/attackSurfaceLocalApps/%s/riskIndicatorEvents",
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
		"v3.0/asrm/attackSurfaceLocalApps/%s/devices",
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
		"v3.0/asrm/attackSurfaceLocalApps/%s/executableFiles",
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
