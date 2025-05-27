package v1client

import "net/http"

func (c *V1ApiClient) ObservedAttackTechniquesList(filter string, qp QueryParameters) (*http.Response, error) {
	return c.searchAndFilter(
		"v3.0/oat/detections",
		filter,
		qp,
	)
}
