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
		"v3.0/workbench/alerts/%s/notes",
		filter,
		qp,
	)
}
