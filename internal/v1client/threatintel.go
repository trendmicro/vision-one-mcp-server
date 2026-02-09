package v1client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ThreatIntelQueryParameters struct {
	OrderBy       string `url:"orderBy,omitempty"`
	Top           int    `url:"top,omitempty"`
	SkipToken     string `url:"skipToken,omitempty"`
	StartDateTime string `url:"startDateTime,omitempty"`
	EndDateTime   string `url:"endDateTime,omitempty"`
	Filter        string `url:"filter,omitempty"`
}

type ThreatIntelFeedParameters struct {
	StartDateTime         string `url:"startDateTime,omitempty"`
	EndDateTime           string `url:"endDateTime,omitempty"`
	Top                   int    `url:"top,omitempty"`
	TopReport             int    `url:"topReport,omitempty"`
	IndicatorObjectFormat string `url:"indicatorObjectFormat,omitempty"`
	ResponseObjectFormat  string `url:"responseObjectFormat,omitempty"`
}

type SuspiciousObject struct {
	URL               string `json:"url,omitempty"`
	Domain            string `json:"domain,omitempty"`
	IP                string `json:"ip,omitempty"`
	SenderMailAddress string `json:"senderMailAddress,omitempty"`
	FileSha1          string `json:"fileSha1,omitempty"`
	FileSha256        string `json:"fileSha256,omitempty"`
	Description       string `json:"description,omitempty"`
	ScanAction        string `json:"scanAction,omitempty"`
	RiskLevel         string `json:"riskLevel,omitempty"`
	DaysToExpiration  int    `json:"daysToExpiration,omitempty"`
}

type SuspiciousObjectException struct {
	URL               string `json:"url,omitempty"`
	Domain            string `json:"domain,omitempty"`
	IP                string `json:"ip,omitempty"`
	SenderMailAddress string `json:"senderMailAddress,omitempty"`
	FileSha1          string `json:"fileSha1,omitempty"`
	FileSha256        string `json:"fileSha256,omitempty"`
	Description       string `json:"description,omitempty"`
}

type SuspiciousObjectDelete struct {
	URL               string `json:"url,omitempty"`
	Domain            string `json:"domain,omitempty"`
	IP                string `json:"ip,omitempty"`
	SenderMailAddress string `json:"senderMailAddress,omitempty"`
	FileSha1          string `json:"fileSha1,omitempty"`
	FileSha256        string `json:"fileSha256,omitempty"`
}

type IntelligenceReportDelete struct {
	ID string `json:"id"`
}

type IntelligenceReportSweep struct {
	ID          string `json:"id"`
	SweepType   string `json:"sweepType"`
	Description string `json:"description,omitempty"`
}

func (c *V1ApiClient) ThreatIntelListSuspiciousObjects(filter string, queryParams ThreatIntelQueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/threatintel/suspiciousObjects", filter, queryParams)
}

func (c *V1ApiClient) ThreatIntelAddSuspiciousObjects(objects []SuspiciousObject) (*http.Response, error) {
	b, err := json.Marshal(&objects)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"v3.0/threatintel/suspiciousObjects",
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) ThreatIntelDeleteSuspiciousObjects(objects []SuspiciousObjectDelete) (*http.Response, error) {
	b, err := json.Marshal(&objects)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"v3.0/threatintel/suspiciousObjects/delete",
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) ThreatIntelListExceptions(filter string, queryParams ThreatIntelQueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/threatintel/suspiciousObjectExceptions", filter, queryParams)
}

func (c *V1ApiClient) ThreatIntelAddExceptions(objects []SuspiciousObjectException) (*http.Response, error) {
	b, err := json.Marshal(&objects)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"v3.0/threatintel/suspiciousObjectExceptions",
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) ThreatIntelDeleteExceptions(objects []SuspiciousObjectDelete) (*http.Response, error) {
	b, err := json.Marshal(&objects)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"v3.0/threatintel/suspiciousObjectExceptions/delete",
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) ThreatIntelListIntelligenceReports(queryParams ThreatIntelQueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/threatintel/intelligenceReports", "", queryParams)
}

func (c *V1ApiClient) ThreatIntelGetIntelligenceReport(reportId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("v3.0/threatintel/intelligenceReports/%s", reportId))
}

func (c *V1ApiClient) ThreatIntelDeleteIntelligenceReports(reportIds []string) (*http.Response, error) {
	deleteBody := []IntelligenceReportDelete{}
	for _, id := range reportIds {
		deleteBody = append(deleteBody, IntelligenceReportDelete{ID: id})
	}
	b, err := json.Marshal(&deleteBody)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"v3.0/threatintel/intelligenceReports/delete",
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) ThreatIntelTriggerSweep(sweeps []IntelligenceReportSweep) (*http.Response, error) {
	b, err := json.Marshal(&sweeps)
	if err != nil {
		return nil, err
	}

	r, err := c.newRequest(
		http.MethodPost,
		"v3.0/threatintel/intelligenceReports/sweep",
		bytes.NewReader(b),
		withContentTypeJSON(),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) ThreatIntelListTasks(queryParams ThreatIntelQueryParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/threatintel/tasks", "", queryParams)
}

func (c *V1ApiClient) ThreatIntelGetTaskResults(taskId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("v3.0/threatintel/tasks/%s", taskId))
}

func (c *V1ApiClient) ThreatIntelListFeedIndicators(queryParams ThreatIntelFeedParameters) (*http.Response, error) {
	return c.searchAndFilter("v3.0/threatintel/feedIndicators", "", queryParams)
}

func (c *V1ApiClient) ThreatIntelListFeeds(contextualFilter string, queryParams ThreatIntelFeedParameters) (*http.Response, error) {
	p, err := query.Values(queryParams)
	if err != nil {
		return nil, err
	}
	r, err := c.newRequest(
		http.MethodGet,
		"v3.0/threatintel/feeds",
		http.NoBody,
		withHeader("TMV1-Contextual-Filter", contextualFilter),
		withUrlParameters(p),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(r)
}

func (c *V1ApiClient) ThreatIntelGetFeedFilterDefinition() (*http.Response, error) {
	return c.genericGet("v3.0/threatintel/feeds/filterDefinition")
}
