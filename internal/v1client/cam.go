package v1client

import (
	"fmt"
	"net/http"
)

func (c *V1ApiClient) CAMListAWSAccounts(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/cam/awsAccounts",
		filter,
		qp,
	)
}

func (c *V1ApiClient) CAMGetAWSAccount(accountId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("v3.0/cam/awsAccounts/%s", accountId))
}

func (c *V1ApiClient) CAMListAlibabaAccounts(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/cam/alibabaAccounts",
		filter,
		qp,
	)
}

func (c *V1ApiClient) CAMGetAlibabaAccountDetails(accountId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("v3.0/cam/alibabaAccounts/%s", accountId))
}

func (c *V1ApiClient) CAMListGCPAccounts(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/cam/gcpProjects",
		filter,
		qp,
	)
}

func (c *V1ApiClient) CAMGetGCPAccountDetails(accountId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("v3.0/cam/gcpProjects/%s", accountId))
}
