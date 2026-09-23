package v1client

import "net/http"

func (c *V1ApiClient) SecurityAwarenessPhishingSimulationsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/securityAwareness/phishingSimulations", p)
}

func (c *V1ApiClient) SecurityAwarenessPhishingRateTrendsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/securityAwareness/phishingSimulations/phishingRateTrends", p)
}

func (c *V1ApiClient) SecurityAwarenessPhishingSimulationRoundsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/securityAwareness/phishingSimulations/%s/rounds", id), p)
}

func (c *V1ApiClient) SecurityAwarenessPhishingSimulationRoundGet(id string, roundId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/securityAwareness/phishingSimulations/%s/rounds/%s", id, roundId), p)
}

func (c *V1ApiClient) SecurityAwarenessPhishingSimulationRoundRecipientsList(id string, roundId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/securityAwareness/phishingSimulations/%s/rounds/%s/recipients", id, roundId), p)
}

func (c *V1ApiClient) SecurityAwarenessTrainingCampaignsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/securityAwareness/trainingCampaigns", p)
}

func (c *V1ApiClient) SecurityAwarenessTrainingCampaignGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/securityAwareness/trainingCampaigns/%s", id), p)
}

func (c *V1ApiClient) SecurityAwarenessTrainingCampaignRecipientsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/securityAwareness/trainingCampaigns/%s/recipients", id), p)
}
