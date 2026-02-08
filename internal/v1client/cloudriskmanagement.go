package v1client

import (
	"fmt"
	"net/http"
)

// CloudRiskManagementListAccounts lists all cloud accounts.
// API: GET /v3.0/cloudRiskManagement/accounts
func (c *V1ApiClient) CloudRiskManagementListAccounts(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/cloudRiskManagement/accounts", filter, queryParams)
}

// CloudRiskManagementGetAccountScanRules retrieves the scan rule settings for a specific account.
// API: GET /v3.0/cloudRiskManagement/accounts/{id}/scanRules
func (c *V1ApiClient) CloudRiskManagementGetAccountScanRules(accountId string, filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(fmt.Sprintf("v3.0/cloudRiskManagement/accounts/%s/scanRules", accountId), filter, queryParams)
}

// CloudRiskManagementListServices lists cloud services and their associated rules supported by Cloud Risk Management.
// API: GET /v3.0/cloudRiskManagement/services
func (c *V1ApiClient) CloudRiskManagementListServices(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/cloudRiskManagement/services", filter, queryParams)
}
