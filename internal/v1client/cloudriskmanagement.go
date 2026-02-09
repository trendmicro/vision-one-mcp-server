package v1client

import (
	"fmt"
	"net/http"
)

func (c *V1ApiClient) CloudRiskManagementListAccounts(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/cloudRiskManagement/accounts", filter, queryParams)
}

func (c *V1ApiClient) CloudRiskManagementGetAccountScanRules(accountId string, filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(fmt.Sprintf("v3.0/cloudRiskManagement/accounts/%s/scanRules", accountId), filter, queryParams)
}

func (c *V1ApiClient) CloudRiskManagementListServices(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/cloudRiskManagement/services", filter, queryParams)
}
