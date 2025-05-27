package v1client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

func (c *V1ApiClient) IAMListAPIKeys(filter string, queryParams QueryParameters) (*http.Response, error) {
	p, err := query.Values(queryParams)
	if err != nil {
		return nil, err
	}
	r, err := c.newRequest(
		http.MethodGet,
		"v3.0/iam/apiKeys",
		http.NoBody,
		withFilter(filter),
		withUrlParameters(p),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
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
	b, err := json.Marshal(&deleteBody)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"v3.0/iam/apiKeys/delete",
		bytes.NewReader(b),
	)
	if err != nil {
		return nil, err
	}
	contentTypeJSON(r)
	return c.client.Do(r)
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
	b, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"v3.0/iam/accounts",
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
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
	b, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPatch,
		fmt.Sprintf("v3.0/iam/accounts/%s", accountId),
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}

	return c.client.Do(r)
}

func (c *V1ApiClient) IAMDeleteAccount(accountId string) (*http.Response, error) {
	r, err := c.newRequest(
		http.MethodDelete,
		fmt.Sprintf("v3.0/iam/accounts/%s", accountId),
		http.NoBody,
	)
	if err != nil {
		return nil, err
	}

	return c.client.Do(r)
}
