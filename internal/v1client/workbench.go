package v1client

import (
	"fmt"
	"net/http"
)

func (c *V1ApiClient) WorkbenchAlertsList(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/workbench/alerts",
		filter,
		qp,
	)
}

func (c *V1ApiClient) WorkbenchGetAlertDetails(alertId string) (*http.Response, error) {
	return c.genericGet(fmt.Sprintf("v3.0/workbench/alerts/%s", alertId))
}

func (c *V1ApiClient) WorkbenchGetAlertNotes(alertId string, filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		fmt.Sprintf("v3.0/workbench/alerts/%s/notes", alertId),
		filter,
		qp,
	)
}

func (c *V1ApiClient) WorkbenchAlertNotesCreate(alertId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/workbench/alerts/%s/notes", alertId), p)
}

func (c *V1ApiClient) WorkbenchAlertNotesDelete(alertId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/workbench/alerts/%s/notes/delete", alertId), p)
}

func (c *V1ApiClient) WorkbenchAlertNoteGet(alertId string, id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/workbench/alerts/%s/notes/%s", alertId, id), p)
}

func (c *V1ApiClient) WorkbenchAlertNoteUpdate(alertId string, id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/workbench/alerts/%s/notes/%s", alertId, id), p)
}

func (c *V1ApiClient) WorkbenchAlertUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/workbench/alerts/%s", id), p)
}

func (c *V1ApiClient) WorkbenchInsightsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/workbench/insights", p)
}

func (c *V1ApiClient) WorkbenchInsightGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/workbench/insights/%s", id), p)
}

func (c *V1ApiClient) WorkbenchInsightImpactScopeEntitiesList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/workbench/insights/%s/impactScopeEntities", id), p)
}

func (c *V1ApiClient) WorkbenchInsightIndicatorsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/workbench/insights/%s/indicators", id), p)
}

func (c *V1ApiClient) WorkbenchInsightMatchedHighlightsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/workbench/insights/%s/matchedHighlights", id), p)
}
