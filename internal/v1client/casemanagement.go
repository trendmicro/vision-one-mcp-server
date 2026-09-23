package v1client

import "net/http"

func (c *V1ApiClient) CaseManagementCasesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/caseManagement/cases", p)
}

func (c *V1ApiClient) CaseManagementCasesCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/caseManagement/cases", p)
}

func (c *V1ApiClient) CaseManagementCaseGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s", id), p)
}

func (c *V1ApiClient) CaseManagementCaseUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/caseManagement/cases/%s", id), p)
}

func (c *V1ApiClient) CaseManagementCaseAttachmentsCreate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/caseManagement/cases/%s/attachments", id), p)
}

func (c *V1ApiClient) CaseManagementCaseAttachmentDelete(id string, attachmentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/caseManagement/cases/%s/attachments/%s", id, attachmentId), p)
}

func (c *V1ApiClient) CaseManagementCaseAttachmentGet(id string, attachmentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/attachments/%s", id, attachmentId), p)
}

func (c *V1ApiClient) CaseManagementCaseBindOatEvent(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/caseManagement/cases/%s/bindOatEvent", id), p)
}

func (c *V1ApiClient) CaseManagementCaseContentsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/contents", id), p)
}

func (c *V1ApiClient) CaseManagementCaseContentsCreate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/caseManagement/cases/%s/contents", id), p)
}

func (c *V1ApiClient) CaseManagementCaseContentDelete(id string, contentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/caseManagement/cases/%s/contents/%s", id, contentId), p)
}

func (c *V1ApiClient) CaseManagementCaseContentGet(id string, contentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/contents/%s", id, contentId), p)
}

func (c *V1ApiClient) CaseManagementCaseContentUpdate(id string, contentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/caseManagement/cases/%s/contents/%s", id, contentId), p)
}

func (c *V1ApiClient) CaseManagementCaseHighlightedObjectsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/highlightedObjects", id), p)
}

func (c *V1ApiClient) CaseManagementCaseHighlightedObjectDelete(id string, highlightedObjectId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/caseManagement/cases/%s/highlightedObjects/%s", id, highlightedObjectId), p)
}

func (c *V1ApiClient) CaseManagementCaseHighlightedObjectGet(id string, highlightedObjectId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/highlightedObjects/%s", id, highlightedObjectId), p)
}

func (c *V1ApiClient) CaseManagementCaseRiskIndicatorEventsList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/riskIndicatorEvents", id), p)
}

func (c *V1ApiClient) CaseManagementCaseRiskIndicatorEventsClose(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/caseManagement/cases/%s/riskIndicatorEvents/close", id), p)
}

func (c *V1ApiClient) CaseManagementCaseTasksList(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/tasks", id), p)
}

func (c *V1ApiClient) CaseManagementCaseTasksCreate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/caseManagement/cases/%s/tasks", id), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskGet(id string, taskId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/tasks/%s", id, taskId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskUpdate(id string, taskId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/caseManagement/cases/%s/tasks/%s", id, taskId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskAttachmentsCreate(id string, taskId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/caseManagement/cases/%s/tasks/%s/attachments", id, taskId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskAttachmentDelete(id string, taskId string, attachmentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/caseManagement/cases/%s/tasks/%s/attachments/%s", id, taskId, attachmentId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskAttachmentGet(id string, taskId string, attachmentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/tasks/%s/attachments/%s", id, taskId, attachmentId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskContentsList(id string, taskId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/tasks/%s/contents", id, taskId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskContentsCreate(id string, taskId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/caseManagement/cases/%s/tasks/%s/contents", id, taskId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskContentDelete(id string, taskId string, contentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/caseManagement/cases/%s/tasks/%s/contents/%s", id, taskId, contentId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskContentGet(id string, taskId string, contentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/caseManagement/cases/%s/tasks/%s/contents/%s", id, taskId, contentId), p)
}

func (c *V1ApiClient) CaseManagementCaseTaskContentUpdate(id string, taskId string, contentId string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, pathf("v3.0/caseManagement/cases/%s/tasks/%s/contents/%s", id, taskId, contentId), p)
}
