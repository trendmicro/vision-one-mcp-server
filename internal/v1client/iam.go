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
