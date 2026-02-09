package v1client

import (
	"fmt"
	"net/http"
)

func (c *V1ApiClient) CloudPostureListAccounts(queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("beta/cloudPosture/accounts", "", queryParams)
}

func (c *V1ApiClient) CloudPostureListAccountChecks(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("beta/cloudPosture/checks", filter, qp)
}

func (c *V1ApiClient) CloudPostureScanTemplate(content string, templateType string) (*http.Response, error) {
	body := map[string]any{
		"content": content,
		"type":    templateType,
	}
	return c.genericJSONPost("beta/cloudPosture/scanTemplate", body)
}

func (c *V1ApiClient) CloudPostureScanAccount(accountId string) (*http.Response, error) {
	return c.genericPost(fmt.Sprintf("beta/cloudPosture/accounts/%s/scan", accountId))
}

func (c *V1ApiClient) CloudPostureGetAccountScanSettings(accountId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("beta/cloudPosture/accounts/%s/scanSetting", accountId))
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
	body := UpdateAccountScanSettings{
		Enabled:  enabled,
		Interval: interval,
	}
	return c.genericJSONPatch(fmt.Sprintf("beta/cloudPosture/accounts/%s/scanSetting", accountId), body)
}

// Custom Rules (Beta) API methods

func (c *V1ApiClient) CloudPostureListCustomRules(queryParams QueryParameters) (*http.Response, error) {
	return c.searchAndFilter("beta/cloudPosture/customRules", "", queryParams)
}

func (c *V1ApiClient) CloudPostureGetCustomRule(ruleId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("beta/cloudPosture/customRules/%s", ruleId))
}

type CustomRuleInput struct {
	Name                    string   `json:"name"`
	Description             string   `json:"description,omitempty"`
	Categories              []string `json:"categories"`
	RiskLevel               string   `json:"riskLevel"`
	Provider                string   `json:"provider"`
	ResolutionReferenceLink string   `json:"resolutionReferenceLink,omitempty"`
	RemediationNote         string   `json:"remediationNote,omitempty"`
	Enabled                 bool     `json:"enabled"`
	Service                 string   `json:"service"`
	ResourceType            string   `json:"resourceType"`
	Attributes              []any    `json:"attributes"`
	EventRules              []any    `json:"eventRules"`
	Slug                    string   `json:"slug,omitempty"`
}

func (c *V1ApiClient) CloudPostureCreateCustomRule(input CustomRuleInput) (*http.Response, error) {
	return c.genericJSONPost("beta/cloudPosture/customRules", input)
}

type CustomRuleUpdateInput struct {
	Name                    string   `json:"name,omitempty"`
	Description             string   `json:"description,omitempty"`
	Categories              []string `json:"categories,omitempty"`
	RiskLevel               string   `json:"riskLevel,omitempty"`
	Provider                string   `json:"provider,omitempty"`
	ResolutionReferenceLink string   `json:"resolutionReferenceLink,omitempty"`
	RemediationNote         string   `json:"remediationNote,omitempty"`
	Enabled                 *bool    `json:"enabled,omitempty"`
	Service                 string   `json:"service,omitempty"`
	ResourceType            string   `json:"resourceType,omitempty"`
	Attributes              []any    `json:"attributes,omitempty"`
	EventRules              []any    `json:"eventRules,omitempty"`
}

func (c *V1ApiClient) CloudPostureUpdateCustomRule(ruleId string, input CustomRuleUpdateInput) (*http.Response, error) {
	return c.genericJSONPatch(fmt.Sprintf("beta/cloudPosture/customRules/%s", ruleId), input)
}

func (c *V1ApiClient) CloudPostureDeleteCustomRule(ruleId string) (*http.Response, error) {
	return c.genericDelete(fmt.Sprintf("beta/cloudPosture/customRules/%s", ruleId))
}

type CustomRuleTestInput struct {
	AccountId     string `json:"accountId,omitempty"`
	Configuration any    `json:"configuration"`
	Resource      any    `json:"resource,omitempty"`
}

func (c *V1ApiClient) CloudPostureTestCustomRule(input CustomRuleTestInput) (*http.Response, error) {
	return c.genericJSONPost("beta/cloudPosture/customRules/test", input)
}
