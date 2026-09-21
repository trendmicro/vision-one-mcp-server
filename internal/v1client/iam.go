package v1client

import (
	"fmt"
	"net/http"
)

func (c *V1ApiClient) IAMListAPIKeys(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/iam/apiKeys", filter, queryParams)
}

type DeleteAPIKey struct {
	ID string `json:"id"`
}

func (c *V1ApiClient) IAMDeleteAPIKeys(apiKeyIDs []string) (*http.Response, error) {
	deleteBody := []DeleteAPIKey{}
	for _, id := range apiKeyIDs {
		deleteBody = append(deleteBody, DeleteAPIKey{
			ID: id,
		})
	}
	return c.genericJSONPost("v3.0/iam/apiKeys/delete", deleteBody)
}

type IAMInviteUserInput struct {
	// required
	Email string `json:"email,omitempty"`
	// required
	Role string `json:"role,omitempty"`
	// required
	AuthType    string `json:"authType,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *V1ApiClient) IAMInviteAccount(input IAMInviteUserInput) (*http.Response, error) {
	return c.genericJSONPost("v3.0/iam/accounts", input)
}

func (c *V1ApiClient) IAMListAccounts(filter string, queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/iam/accounts", filter, queryParams)
}

type IAMUpdateAccountInput struct {
	Role        string `json:"role,omitempty"`
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *V1ApiClient) IAMUpdateAccount(accountId string, input IAMUpdateAccountInput) (*http.Response, error) {
	return c.genericJSONPatch(fmt.Sprintf("v3.0/iam/accounts/%s", accountId), input)
}

func (c *V1ApiClient) IAMDeleteAccount(accountId string) (*http.Response, error) {
	return c.genericDelete(fmt.Sprintf("v3.0/iam/accounts/%s", accountId))
}

func (c *V1ApiClient) IamAccountGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/iam/accounts/%s", id), p)
}

func (c *V1ApiClient) IamApiKeysCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/iam/apiKeys", p)
}

func (c *V1ApiClient) IamApiKeyGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/iam/apiKeys/%s", id), p)
}

func (c *V1ApiClient) IamApiKeyUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/iam/apiKeys/%s", id), p)
}

func (c *V1ApiClient) IamIdentityProvidersList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/iam/identityProviders", p)
}

func (c *V1ApiClient) IamIdentityProvidersCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/iam/identityProviders", p)
}

func (c *V1ApiClient) IamIdentityProviderDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/iam/identityProviders/%s", id), p)
}

func (c *V1ApiClient) IamIdentityProviderGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/iam/identityProviders/%s", id), p)
}

func (c *V1ApiClient) IamIdentityProviderUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/iam/identityProviders/%s/update", id), p)
}

func (c *V1ApiClient) IamRolesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/iam/roles", p)
}

func (c *V1ApiClient) IamRolesCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/iam/roles", p)
}

func (c *V1ApiClient) IamRolesPermissionKeysList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/iam/roles/permissionKeys", p)
}

func (c *V1ApiClient) IamRoleDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/iam/roles/%s", id), p)
}

func (c *V1ApiClient) IamRoleGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/iam/roles/%s", id), p)
}

func (c *V1ApiClient) IamRoleUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/iam/roles/%s", id), p)
}

func (c *V1ApiClient) IamRolePermissionsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/iam/roles/%s/permissions", id), p)
}
