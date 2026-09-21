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

func (c *V1ApiClient) CamAlibabaAccountsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/alibabaAccounts", p)
}

func (c *V1ApiClient) CamAlibabaAccountsGenerateTerraformPackage(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/alibabaAccounts/generateTerraformPackage", p)
}

func (c *V1ApiClient) CamAlibabaAccountDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/cam/alibabaAccounts/%s", id), p)
}

func (c *V1ApiClient) CamAlibabaAccountUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/cam/alibabaAccounts/%s", id), p)
}

func (c *V1ApiClient) CamAwsAccountsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/awsAccounts", p)
}

func (c *V1ApiClient) CamAwsAccountsFeaturesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/cam/awsAccounts/features", p)
}

func (c *V1ApiClient) CamAwsAccountsGenerateCfnTemplateLinks(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/awsAccounts/generateCfnTemplateLinks", p)
}

func (c *V1ApiClient) CamAwsAccountsGenerateTerraformPackage(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/awsAccounts/generateTerraformPackage", p)
}

func (c *V1ApiClient) CamAwsAccountDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/cam/awsAccounts/%s", id), p)
}

func (c *V1ApiClient) CamAwsAccountUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/cam/awsAccounts/%s", id), p)
}

func (c *V1ApiClient) CamAzureSubscriptionsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/cam/azureSubscriptions", p)
}

func (c *V1ApiClient) CamAzureSubscriptionsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/azureSubscriptions", p)
}

func (c *V1ApiClient) CamAzureSubscriptionsGenerateMgmtGroupTerraformPackage(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/azureSubscriptions/generateManagementGroupTerraformPackage", p)
}

func (c *V1ApiClient) CamAzureSubscriptionsGenerateTerraformPackage(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/azureSubscriptions/generateTerraformPackage", p)
}

func (c *V1ApiClient) CamAzureSubscriptionDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/cam/azureSubscriptions/%s", id), p)
}

func (c *V1ApiClient) CamAzureSubscriptionGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/cam/azureSubscriptions/%s", id), p)
}

func (c *V1ApiClient) CamAzureSubscriptionUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/cam/azureSubscriptions/%s", id), p)
}

func (c *V1ApiClient) CamGcpProjectsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/gcpProjects", p)
}

func (c *V1ApiClient) CamGcpProjectsGenerateTerraformPackage(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/gcpProjects/generateTerraformPackage", p)
}

func (c *V1ApiClient) CamGcpProjectDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/cam/gcpProjects/%s", id), p)
}

func (c *V1ApiClient) CamGcpProjectUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/cam/gcpProjects/%s", id), p)
}

func (c *V1ApiClient) CamOciCompartmentsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/cam/ociCompartments", p)
}

func (c *V1ApiClient) CamOciCompartmentsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/ociCompartments", p)
}

func (c *V1ApiClient) CamOciCompartmentsGenerateTerraformPackage(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/cam/ociCompartments/generateTerraformPackage", p)
}

func (c *V1ApiClient) CamOciCompartmentDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/cam/ociCompartments/%s", id), p)
}

func (c *V1ApiClient) CamOciCompartmentGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/cam/ociCompartments/%s", id), p)
}

func (c *V1ApiClient) CamOciCompartmentUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/cam/ociCompartments/%s", id), p)
}
