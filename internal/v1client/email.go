package v1client

import "net/http"

func (c *V1ApiClient) EmailSecurityListAccounts(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/emailAssetInventory/emailAccounts",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) EmailSecurityListDomains(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/emailAssetInventory/emailDomains",
		filter,
		queryParams,
	)
}

func (c *V1ApiClient) EmailSecurityListServers(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/emailAssetInventory/emailServers",
		filter,
		queryParams,
	)
}
