package v1client

import "net/http"

func (c *V1ApiClient) AuditLogsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/audit/logs", p)
}
