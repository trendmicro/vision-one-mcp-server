package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyTagManagement = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolTagManagementCustomTagsList,
	toolTagManagementCustomTagAssetsList,
	toolTagManagementCustomizableTagsList,
	toolTagManagementCustomizableTagAssetsList,
	toolTagManagementTaskGet,
}

var ToolsetsWriteTagManagement = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolTagManagementCustomTagsCreate,
	toolTagManagementCustomTagsAssign,
	toolTagManagementCustomTagsDelete,
	toolTagManagementCustomTagsUnassign,
	toolTagManagementCustomizableTagsCreate,
	toolTagManagementCustomizableTagsAssign,
	toolTagManagementCustomizableTagsDelete,
	toolTagManagementCustomizableTagsUnassign,
}

func toolTagManagementCustomTagsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_custom_tags_list",
			mcp.WithDescription("List custom tag property-value pairs. Returns a flat list of (property, value) rows for all customer-created custom tags, sorted alphabetically by property. Results are paginated; follow nextLink to retrieve subsequent pages."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description("The maximum number of rows to return in a single page. Accepts values from 50 to 1000. If omitted, 100 rows are returned. Default: 100.")),
			mcp.WithString("orderBy", mcp.Description("The sort expression for the result set. The only supported field is property. Append asc (ascending) or desc (descending) to choose direction. If the direction is omitted, the service defaults to ascending. Default: property asc.")),
			mcp.WithString("filter", mcp.Description("The filter for narrowing results, mirroring the Tag Management console search. Supported fields: property - The display name of the tag property - Any value in field; value - The tag value name - Any value in field; id - The unique identifier of the tag value - Any value in field; Supported operators: contains - Substring match — supported only on property and value; eq - Operator 'equal to' — supported only on id; or - Operator 'or' — combines multiple id eq clauses only; Match is case-sensitive. Maximum length: 1024 characters. Resend this header unchanged on every request while following nextLink; a changed value restarts pagination (400).")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindNumber}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomTagsList(params)
			return handleResponse(resp, err, "failed to list custom tag property-value pairs")
		},
	}
}

func toolTagManagementCustomTagsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_custom_tags_create",
			mcp.WithDescription("Create custom tags. Creates custom tag properties and values. If the record already exists, it is counted as successful and no changes are made."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("tags", mcp.Required(), mcp.Description("The tag property and value pairs to create. Maximum 1,000 records per request."), mcp.Items(map[string]any{"properties": map[string]any{"property": map[string]any{"description": "The name of the custom tag property. Created if it does not exist. Allowed characters: Unicode letters and digits, space, and .", "type": "string"}, "value": map[string]any{"description": "The value to create under the tag property. Created if it does not exist. Allowed characters: Unicode letters and digits, space, and .", "type": "string"}}, "required": []string{"property", "value"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "tags", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomTagsCreate(params)
			return handleResponse(resp, err, "failed to create custom tags")
		},
	}
}

func toolTagManagementCustomTagsAssign(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_custom_tags_assign",
			mcp.WithDescription("Assign custom tags to assets. Assigns existing custom tags to the specified assets."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("assignments", mcp.Required(), mcp.Description("Asset type groups containing tag assignment mappings. Currently limited to 1 group per request."), mcp.Items(map[string]any{"properties": map[string]any{"assetType": map[string]any{"description": "The type of asset for this group", "enum": []string{"device", "domainAccount", "publicIpAddress", "globalFqdn", "serviceAccount", "cloudAsset", "localApp"}, "type": "string"}, "mappings": map[string]any{"description": "Tag assignment mappings. Maximum 1,000 mappings total across all groups.", "items": map[string]any{"type": "object"}, "type": "array"}}, "required": []string{"assetType", "mappings"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "assignments", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomTagsAssign(params)
			return handleResponse(resp, err, "failed to assign custom tags to assets")
		},
	}
}

func toolTagManagementCustomTagsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_custom_tags_delete",
			mcp.WithDescription("Delete custom tags. Deletes individual custom tag values in bulk."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("tags", mcp.Required(), mcp.Description("The custom tag values to delete. Maximum 1,000 records per request."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The unique identifier of the tag value to delete, obtained from the List Tags API.", "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "tags", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomTagsDelete(params)
			return handleResponse(resp, err, "failed to delete custom tags")
		},
	}
}

func toolTagManagementCustomTagsUnassign(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_custom_tags_unassign",
			mcp.WithDescription("Unassign custom tags from assets. Removes existing custom tag values from the specified assets. This is the inverse of assign; it removes the asset-to-value relationship only and never deletes the tag definition itself."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("assignments", mcp.Required(), mcp.Description("Asset type groups containing tag unassignment mappings. Currently limited to 1 group per request."), mcp.Items(map[string]any{"properties": map[string]any{"assetType": map[string]any{"description": "The type of asset for this group", "enum": []string{"device", "domainAccount", "publicIpAddress", "globalFqdn", "serviceAccount", "cloudAsset", "localApp"}, "type": "string"}, "mappings": map[string]any{"description": "Tag unassignment mappings. Maximum 1,000 mappings total across all groups.", "items": map[string]any{"type": "object"}, "type": "array"}}, "required": []string{"assetType", "mappings"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "assignments", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomTagsUnassign(params)
			return handleResponse(resp, err, "failed to unassign custom tags from assets")
		},
	}
}

func toolTagManagementCustomTagAssetsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_custom_tag_assets_list",
			mcp.WithDescription("List assets that have a given custom tag. Returns the assets of the given assetType that currently have the custom tag identified by id assigned. The tag id is obtained from the tag List API and maps directly to the internal tag — no name resolution is performed. Results are paginated; follow nextLink to retrieve subsequent pages."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The opaque tag identifier, obtained from the tag List API. Maps directly to the internal tag; no name resolution is performed.")),
			mcp.WithString("assetType", mcp.Required(), mcp.Description("The asset type that scopes the relationship lookup. Relationships are keyed by asset type."), mcp.Enum("devices", "domainAccounts", "publicIpAddresses", "globalFqdns", "serviceAccounts", "cloudAssets", "localApps")),
			mcp.WithNumber("top", mcp.Description("The maximum number of assets to return in a single page. Accepts values from 50 to 1000. If omitted, 1000 assets are returned. Default: 1000.")),
			mcp.WithString("orderBy", mcp.Description("The sort expression for the result set. The only supported field is name. Append asc (ascending) or desc (descending) to choose direction. If the direction is omitted, the service defaults to ascending. Default: name asc.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			assetType, err := pathValue("assetType", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindNumber}, {Name: "orderBy", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomTagAssetsList(id, assetType, params)
			return handleResponse(resp, err, "failed to list assets that have a given custom tag")
		},
	}
}

func toolTagManagementCustomizableTagsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_customizable_tags_list",
			mcp.WithDescription("List customizable platform tag property-value pairs. Returns a flat list of (property, value) rows for the three customizable platform tag properties — Asset group, Cloud project, and Custom compliance frameworks. A property with no values contributes no rows."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description("The maximum number of rows to return in a single page. Accepts values from 50 to 1000. If omitted, 100 rows are returned. Default: 100.")),
			mcp.WithString("orderBy", mcp.Description("The sort expression for the result set. The only supported field is property. Append asc (ascending) or desc (descending) to choose direction. If the direction is omitted, the service defaults to ascending. Default: property asc.")),
			mcp.WithString("filter", mcp.Description("The filter for narrowing results, mirroring the Tag Management console search. Supported fields: property - The display name of the tag property - Asset group, Cloud project, Custom compliance frameworks; value - The tag value name - Any value in field; id - The unique identifier of the tag value - Any value in field; Supported operators: contains - Substring match — supported only on property and value; eq - Operator 'equal to' — supported only on id; or - Operator 'or' — combines multiple id eq clauses only; Match is case-sensitive. Maximum length: 1024 characters. Resend this header unchanged on every request while following nextLink; a changed value restarts pagination (400).")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindNumber}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomizableTagsList(params)
			return handleResponse(resp, err, "failed to list customizable platform tag property-value pairs")
		},
	}
}

func toolTagManagementCustomizableTagsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_customizable_tags_create",
			mcp.WithDescription("Create customizable tags. Creates values under customizable platform tag properties. If the value already exists, the record is treated as successful and no duplicate is created."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("tags", mcp.Required(), mcp.Description("The customizable tag property and value pairs to create. Maximum 1,000 records per request."), mcp.Items(map[string]any{"properties": map[string]any{"property": map[string]any{"description": "The name of the customizable platform tag property.", "enum": []string{"Asset group", "Cloud project"}, "type": "string"}, "value": map[string]any{"description": "The value to create under the customizable tag property. Created if it does not exist.", "type": "string"}}, "required": []string{"property", "value"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "tags", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomizableTagsCreate(params)
			return handleResponse(resp, err, "failed to create customizable tags")
		},
	}
}

func toolTagManagementCustomizableTagsAssign(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_customizable_tags_assign",
			mcp.WithDescription("Assign customizable tags to assets. Assigns existing customizable platform tags to the specified assets."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("assignments", mcp.Required(), mcp.Description("Asset type groups containing customizable tag assignment mappings. Currently limited to 1 group per request."), mcp.Items(map[string]any{"properties": map[string]any{"assetType": map[string]any{"description": "The type of asset for this group", "enum": []string{"device", "domainAccount", "publicIpAddress", "globalFqdn", "serviceAccount", "cloudAsset", "localApp"}, "type": "string"}, "mappings": map[string]any{"description": "Customizable tag assignment mappings. Maximum 1,000 mappings total across all groups.", "items": map[string]any{"type": "object"}, "type": "array"}}, "required": []string{"assetType", "mappings"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "assignments", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomizableTagsAssign(params)
			return handleResponse(resp, err, "failed to assign customizable tags to assets")
		},
	}
}

func toolTagManagementCustomizableTagsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_customizable_tags_delete",
			mcp.WithDescription("Delete customizable tags. Deletes values under customizable platform tag properties in bulk."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("tags", mcp.Required(), mcp.Description("The customizable tag values to delete. Maximum 1,000 records per request."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The unique identifier of the tag value to delete, obtained from the List Tags API.", "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "tags", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomizableTagsDelete(params)
			return handleResponse(resp, err, "failed to delete customizable tags")
		},
	}
}

func toolTagManagementCustomizableTagsUnassign(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_customizable_tags_unassign",
			mcp.WithDescription("Unassign customizable tags from assets. Removes existing customizable platform tag values from the specified assets. This is the inverse of assign; it removes the asset-to-value relationship only and never deletes the tag definition itself."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("assignments", mcp.Required(), mcp.Description("Asset type groups containing customizable tag unassignment mappings. Currently limited to 1 group per request."), mcp.Items(map[string]any{"properties": map[string]any{"assetType": map[string]any{"description": "The type of asset for this group", "enum": []string{"device", "domainAccount", "publicIpAddress", "globalFqdn", "serviceAccount", "cloudAsset", "localApp"}, "type": "string"}, "mappings": map[string]any{"description": "Customizable tag unassignment mappings. Maximum 1,000 mappings total across all groups.", "items": map[string]any{"type": "object"}, "type": "array"}}, "required": []string{"assetType", "mappings"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "assignments", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomizableTagsUnassign(params)
			return handleResponse(resp, err, "failed to unassign customizable tags from assets")
		},
	}
}

func toolTagManagementCustomizableTagAssetsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_customizable_tag_assets_list",
			mcp.WithDescription("List assets that have a given customizable platform tag. Returns the assets of the given assetType that currently have the customizable platform tag identified by id assigned. The id identifies a value of one of the customizable platform properties — Asset group, Cloud project, or Custom compliance frameworks — and is obtained from the tag List API;."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The opaque tag identifier, obtained from the tag List API. Maps directly to the internal tag; no name resolution is performed.")),
			mcp.WithString("assetType", mcp.Required(), mcp.Description("The asset type that scopes the relationship lookup. Relationships are keyed by asset type."), mcp.Enum("devices", "domainAccounts", "publicIpAddresses", "globalFqdns", "serviceAccounts", "cloudAssets", "localApps")),
			mcp.WithNumber("top", mcp.Description("The maximum number of assets to return in a single page. Accepts values from 50 to 1000. If omitted, 1000 assets are returned. Default: 1000.")),
			mcp.WithString("orderBy", mcp.Description("The sort expression for the result set. The only supported field is name. Append asc (ascending) or desc (descending) to choose direction. If the direction is omitted, the service defaults to ascending. Default: name asc.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			assetType, err := pathValue("assetType", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindNumber}, {Name: "orderBy", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementCustomizableTagAssetsList(id, assetType, params)
			return handleResponse(resp, err, "failed to list assets that have a given customizable platform tag")
		},
	}
}

func toolTagManagementTaskGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"tag_management_task_get",
			mcp.WithDescription("Get tag operation task status. Returns the current status and results of an asynchronous tag operation (create, assign, unassign, or delete) by its task ID."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of the task, from the Operation-Location header.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.TagManagementTaskGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get tag operation task status")
		},
	}
}
