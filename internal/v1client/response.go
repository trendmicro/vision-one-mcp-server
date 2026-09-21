package v1client

import "net/http"

func (c *V1ApiClient) ResponseContainersIsolate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/containers/isolate", p)
}

func (c *V1ApiClient) ResponseContainersRestore(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/containers/restore", p)
}

func (c *V1ApiClient) ResponseContainersTerminate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/containers/terminate", p)
}

func (c *V1ApiClient) ResponseCustomScriptsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/response/customScripts", p)
}

func (c *V1ApiClient) ResponseCustomScriptsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/customScripts", p)
}

func (c *V1ApiClient) ResponseCustomScriptDelete(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodDelete, pathf("v3.0/response/customScripts/%s", id), p)
}

func (c *V1ApiClient) ResponseCustomScriptGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/response/customScripts/%s", id), p)
}

func (c *V1ApiClient) ResponseCustomScriptUpdate(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, pathf("v3.0/response/customScripts/%s/update", id), p)
}

func (c *V1ApiClient) ResponseDomainAccountsDisable(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/domainAccounts/disable", p)
}

func (c *V1ApiClient) ResponseDomainAccountsEnable(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/domainAccounts/enable", p)
}

func (c *V1ApiClient) ResponseDomainAccountsResetPassword(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/domainAccounts/resetPassword", p)
}

func (c *V1ApiClient) ResponseDomainAccountsSignOut(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/domainAccounts/signOut", p)
}

func (c *V1ApiClient) ResponseEmailsDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/emails/delete", p)
}

func (c *V1ApiClient) ResponseEmailsQuarantine(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/emails/quarantine", p)
}

func (c *V1ApiClient) ResponseEmailsRestore(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/emails/restore", p)
}

func (c *V1ApiClient) ResponseEndpointActionExceptionsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/response/endpointActionExceptions", p)
}

func (c *V1ApiClient) ResponseEndpointActionExceptionsRegister(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpointActionExceptions/register", p)
}

func (c *V1ApiClient) ResponseEndpointsCollectFile(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpoints/collectFile", p)
}

func (c *V1ApiClient) ResponseEndpointsIsolate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpoints/isolate", p)
}

func (c *V1ApiClient) ResponseEndpointsRestore(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpoints/restore", p)
}

func (c *V1ApiClient) ResponseEndpointsRunOsquery(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpoints/runOsquery", p)
}

func (c *V1ApiClient) ResponseEndpointsRunScript(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpoints/runScript", p)
}

func (c *V1ApiClient) ResponseEndpointsRunYaraRulesCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpoints/runYaraRules", p)
}

func (c *V1ApiClient) ResponseEndpointsStartMalwareScan(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpoints/startMalwareScan", p)
}

func (c *V1ApiClient) ResponseEndpointsTerminateProcess(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/endpoints/terminateProcess", p)
}

func (c *V1ApiClient) ResponseIsolatedTrafficExceptionsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/response/isolatedTrafficExceptions", p)
}

func (c *V1ApiClient) ResponseIsolatedTrafficExceptionsRegister(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/isolatedTrafficExceptions/register", p)
}

func (c *V1ApiClient) ResponseOsqueryStatementsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/response/osqueryStatements", p)
}

func (c *V1ApiClient) ResponseSettingStatusGet(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/response/settingStatus", p)
}

func (c *V1ApiClient) ResponseSettingStatusUpdate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPatch, "v3.0/response/settingStatus", p)
}

func (c *V1ApiClient) ResponseSuspiciousObjectsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/suspiciousObjects", p)
}

func (c *V1ApiClient) ResponseSuspiciousObjectsDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/suspiciousObjects/delete", p)
}

func (c *V1ApiClient) ResponseTasksList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/response/tasks", p)
}

func (c *V1ApiClient) ResponseTasksCancel(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/response/tasks/cancel", p)
}

func (c *V1ApiClient) ResponseTaskGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/response/tasks/%s", id), p)
}

func (c *V1ApiClient) ResponseYaraRuleFilesList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/response/yaraRuleFiles", p)
}
