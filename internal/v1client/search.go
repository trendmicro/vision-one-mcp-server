package v1client

import "net/http"

func (c *V1ApiClient) SearchActivityStatisticsGet(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/activityStatistics", p)
}

func (c *V1ApiClient) SearchCloudActivitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/cloudActivities", p)
}

func (c *V1ApiClient) SearchContainerActivitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/containerActivities", p)
}

func (c *V1ApiClient) SearchDetectionsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/detections", p)
}

func (c *V1ApiClient) SearchEmailActivitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/emailActivities", p)
}

func (c *V1ApiClient) SearchEndpointActivitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/endpointActivities", p)
}

func (c *V1ApiClient) SearchIdentityActivitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/identityActivities", p)
}

func (c *V1ApiClient) SearchMobileActivitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/mobileActivities", p)
}

func (c *V1ApiClient) SearchNetworkActivitiesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/networkActivities", p)
}

func (c *V1ApiClient) SearchSensorStatisticsGet(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/search/sensorStatistics", p)
}
