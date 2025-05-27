package v1client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

func (c *V1ApiClient) CloudPostureListAccounts(queryParams QueryParameters) (*http.Response, error) {
	p, err := query.Values(queryParams)
	if err != nil {
		return nil, err
	}
	r, err := c.newRequest(
		http.MethodGet,
		"beta/cloudPosture/accounts",
		http.NoBody,
		withUrlParameters(p),
	)
	if err != nil {
		return nil, err
	}

	return c.client.Do(r)
}

func (c *V1ApiClient) CloudPostureListAccountChecks(filter string, qp QueryParameters) (*http.Response, error) {
	p, err := query.Values(qp)
	if err != nil {
		return nil, err
	}
	r, err := c.newRequest(
		http.MethodGet,
		"beta/cloudPosture/checks",
		http.NoBody,
		withFilter(filter),
		withUrlParameters(p),
	)
	if err != nil {
		return nil, err
	}

	return c.client.Do(r)
}

func (c *V1ApiClient) CloudPostureScanTemplate(content string, templateType string) (*http.Response, error) {
	jsonInput := map[string]any{
		"content": content,
		"type":    templateType,
	}
	b, err := json.Marshal(jsonInput)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"beta/cloudPosture/scanTemplate",
		bytes.NewReader(b),
	)
	if err != nil {
		return nil, err
	}

	contentTypeJSON(r)

	return c.client.Do(r)
}

func (c *V1ApiClient) CloudPostureScanAccount(accountId string) (*http.Response, error) {
	r, err := c.newRequest(
		http.MethodPost,
		fmt.Sprintf("beta/cloudPosture/accounts/%s/scan", accountId),
		http.NoBody,
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) CloudPostureGetAccountScanSettings(accountId string) (*http.Response, error) {
	r, err := c.newRequest(
		http.MethodGet,
		fmt.Sprintf("beta/cloudPosture/accounts/%s/scanSetting", accountId),
		http.NoBody,
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

type UpdateAccountScanSettings struct {
	Enabled  *bool `json:"enabled,omitempty"`
	Interval int   `json:"interval,omitempty"`
}

func (c *V1ApiClient) CloudPostureUpdateAccountScanSettings(
	accountId string,
	enabled *bool,
	interval int,
) (*http.Response, error) {
	apiInput := UpdateAccountScanSettings{
		Enabled:  enabled,
		Interval: interval,
	}
	b, err := json.Marshal(apiInput)
	if err != nil {
		return nil, err
	}
	r, err := c.newRequest(
		http.MethodPatch,
		fmt.Sprintf("beta/cloudPosture/accounts/%s/scanSetting", accountId),
		bytes.NewReader(b),
	)
	if err != nil {
		return nil, err
	}

	contentTypeJSON(r)

	return c.client.Do(r)
}
