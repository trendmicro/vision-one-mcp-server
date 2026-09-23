package v1client

import "net/http"

func (c *V1ApiClient) TagManagementCustomTagsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/tagManagement/customTags", p)
}

func (c *V1ApiClient) TagManagementCustomTagsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/tagManagement/customTags", p)
}

func (c *V1ApiClient) TagManagementCustomTagsAssign(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/tagManagement/customTags/assign", p)
}

func (c *V1ApiClient) TagManagementCustomTagsDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/tagManagement/customTags/delete", p)
}

func (c *V1ApiClient) TagManagementCustomTagsUnassign(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/tagManagement/customTags/unassign", p)
}

func (c *V1ApiClient) TagManagementCustomTagAssetsList(id string, assetType string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/tagManagement/customTags/%s/%s", id, assetType), p)
}

func (c *V1ApiClient) TagManagementCustomizableTagsList(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, "v3.0/tagManagement/customizableTags", p)
}

func (c *V1ApiClient) TagManagementCustomizableTagsCreate(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/tagManagement/customizableTags", p)
}

func (c *V1ApiClient) TagManagementCustomizableTagsAssign(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/tagManagement/customizableTags/assign", p)
}

func (c *V1ApiClient) TagManagementCustomizableTagsDelete(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/tagManagement/customizableTags/delete", p)
}

func (c *V1ApiClient) TagManagementCustomizableTagsUnassign(p RequestParams) (*http.Response, error) {
	return c.do(http.MethodPost, "v3.0/tagManagement/customizableTags/unassign", p)
}

func (c *V1ApiClient) TagManagementCustomizableTagAssetsList(id string, assetType string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/tagManagement/customizableTags/%s/%s", id, assetType), p)
}

func (c *V1ApiClient) TagManagementTaskGet(id string, p RequestParams) (*http.Response, error) {
	return c.do(http.MethodGet, pathf("v3.0/tagManagement/tasks/%s", id), p)
}
