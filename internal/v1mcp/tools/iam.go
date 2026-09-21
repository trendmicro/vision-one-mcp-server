package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"

	mcpserver "github.com/mark3labs/mcp-go/server"
)

var ToolsetsReadOnlyIAM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolIamApiKeysList,
	toolIamAccountsList,
	toolIamAccountGet,
	toolIamApiKeyGet,
	toolIamIdentityProvidersList,
	toolIamIdentityProviderGet,
	toolIamRolesList,
	toolIamRolesPermissionKeysList,
	toolIamRoleGet,
	toolIamRolePermissionsList,
}

var ToolsetsWriteIAM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolIamApiKeysDelete,
	toolIamAccountInvite,
	toolIamAccountUpdate,
	toolIamAccountDelete,
	toolIamApiKeysCreate,
	toolIamApiKeyUpdate,
	toolIamIdentityProvidersCreate,
	toolIamIdentityProviderDelete,
	toolIamIdentityProviderUpdate,
	toolIamRolesCreate,
	toolIamRoleDelete,
	toolIamRoleUpdate,
}

func toolIamApiKeysList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_api_keys_list",
			mcp.WithDescription("List Vision One API Keys"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterApiKeys)),
			mcp.WithString("orderBy",
				mcp.Enum(
					withOrdering(
						asc_desc,
						"lastUsedDateTime",
						"lastModifiedDateTime",
						"expiredDateTime",
						"createdDateTim",
					)...,
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString(
				"top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum("50", "100", "200"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.IAMListAPIKeys(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list api keys")
		},
	}
}

func toolIamApiKeysDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_api_keys_delete",
			mcp.WithDescription("Delete Vision One API Keys"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("apiKeyIds",
				mcp.Description("Array of API Key Ids to delete"),
				mcp.Required(),
				mcp.Items(
					map[string]any{
						"type": "string",
					},
				),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			keysToDelete := []string{}
			if keyIds, ok := request.GetArguments()["apiKeyIds"].([]any); ok && len(keyIds) > 0 {
				for _, id := range keyIds {
					keyId, ok := id.(string)
					if !ok {
						return mcp.NewToolResultError("each key ID must be a string"), nil
					}
					keysToDelete = append(keysToDelete, keyId)
				}
			}
			resp, err := client.IAMDeleteAPIKeys(keysToDelete)
			return handleStatusResponse(resp, err, http.StatusMultiStatus, "failed to delete api keys")
		},
	}
}

func toolIamAccountInvite(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_account_invite",
			mcp.WithDescription("Sends an invitation to the specified email address to be added as an account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("email",
				mcp.Required(),
				mcp.Description("Email address of the user"),
			),
			mcp.WithString("role",
				mcp.Required(),
				mcp.Description("The role to assign to the user"),
			),
			mcp.WithString("authType",
				mcp.Required(),
				mcp.Description("The type of the user account"),
				mcp.Enum("local", "saml", "samlGroup"),
			),
			mcp.WithString("description",
				mcp.Description("Brief note for the user account"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			email, err := requiredValue[string]("email", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			role, err := requiredValue[string]("role", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			authType, err := requiredValue[string]("authType", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			input := v1client.IAMInviteUserInput{
				Email:       email,
				Role:        role,
				AuthType:    authType,
				Description: description,
			}

			resp, err := client.IAMInviteAccount(input)
			return handleStatusResponse(resp, err, http.StatusCreated, "failed to invite user")
		},
	}
}

func toolIamAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_accounts_list",
			mcp.WithDescription("Displays users, groups, and invitations in the account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterUserAccounts)),
			mcp.WithString(
				"top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum("50", "100", "200"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top: top,
			}

			resp, err := client.IAMListAccounts(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list accounts")
		},
	}
}

func toolIamAccountUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_account_update",
			mcp.WithDescription("Updates the specified account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
				mcp.Description("The ID of the account to update"),
			),
			mcp.WithString("role",
				mcp.Description("The role to assign to the user"),
			),
			mcp.WithString("status",
				mcp.Description("The status of the user account"),
				mcp.Enum("enabled", "disabled"),
			),
			mcp.WithString("description",
				mcp.Description("Brief note for the user account"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			role, err := optionalValue[string]("role", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			status, err := optionalValue[string]("status", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			input := v1client.IAMUpdateAccountInput{
				Role:        role,
				Status:      status,
				Description: description,
			}

			resp, err := client.IAMUpdateAccount(accountId, input)
			return handleStatusResponse(resp, err, http.StatusNoContent, "failed to update account")
		},
	}
}

func toolIamAccountDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_account_delete",
			mcp.WithDescription("Deletes the specified account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
				mcp.Description("The ID of the account to delete"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IAMDeleteAccount(accountId)
			return handleStatusResponse(resp, err, http.StatusNoContent, "failed to delete account")
		},
	}
}

func toolIamAccountGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_account_get",
			mcp.WithDescription("Retrieve details of a user account, SAML group, or invitation. Returns the details of the specified user account, SAML group, or invitation."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a user account, SAML group, or invitation.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamAccountGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to retrieve details of a user account, SAML group, or invitation")
		},
	}
}

func toolIamApiKeysCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_api_keys_create",
			mcp.WithDescription("Create API Keys. Generates API keys that allow third-party applications to access the TrendAI Vision One™ APIs."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "A brief note about the API key", "type": "string"}, "monthsToExpiration": map[string]any{"description": "The duration of validity for the API key.", "enum": []any{1, 3, 6, 12, 0}, "type": "integer"}, "name": map[string]any{"description": "The unique name of an API key", "type": "string"}, "role": map[string]any{"description": "The user role assigned to the API key. Please do not use Master Administrator role for API key creation.", "type": "string"}, "status": map[string]any{"description": "The status of an API key.", "enum": []string{"enabled", "disabled"}, "type": "string"}}, "required": []string{"role", "name"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamApiKeysCreate(params)
			return handleResponse(resp, err, "failed to create API Keys")
		},
	}
}

func toolIamApiKeyGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_api_key_get",
			mcp.WithDescription("Get API key. Displays information of the specified API key."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of the API key.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamApiKeyGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get API key")
		},
	}
}

func toolIamApiKeyUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_api_key_update",
			mcp.WithDescription("Update API key. Updates the specified API key."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of the API key.")),
			mcp.WithString("ifMatch", mcp.Description("The ETag of the resource you want to update. Note: The resource is updated only if the provided value matches the ETag of the resource. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("role", mcp.Description("The user role assigned to the API key. Please do not use Master Administrator role for API key creation Note: Not every role can be assigned to API keys.")),
			mcp.WithString("name", mcp.Description("The unique name of the API key.")),
			mcp.WithString("status", mcp.Description("The status of an API key. Available values: * enabled * disabled."), mcp.Enum("enabled", "disabled")),
			mcp.WithString("description"),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "role", Kind: kindString}, {Name: "name", Kind: kindString}, {Name: "status", Kind: kindString}, {Name: "description", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamApiKeyUpdate(id, params)
			return handleResponse(resp, err, "failed to update API key")
		},
	}
}

func toolIamIdentityProvidersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_identity_providers_list",
			mcp.WithDescription("List SAML identity providers. Returns a paginated list of SAML identity providers."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records returned per page. One of: 50, 100, 200. Default: 50."), mcp.Enum("50", "100", "200")),
			mcp.WithString("orderBy", mcp.Description("The field that specifies the order in which the results are sorted. To display records in ascending or descending order, add the phrase asc or desc after the parameter name. If no order is specified, the results are shown in ascending order. Default: expiredDateTime.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the identity provider list. Supported fields: entityId - The entity ID extracted from the identity provider's metadata XML - Any value; name - The name to display for this SAML identity provider - Any value; expiredDateTime - The latest expiry time from the certificates in the identity provider's metadata XML - Date and time in ISO 8601 format (yyyy-MM-ddThh:mm:ssZ in UTC); Supported operators: eq - Operator 'equal to' - Not applicable to expiredDateTime; gt - Operator 'greater than' - Only applicable to expiredDateTime; ge - Operator 'greater than or equal' - Only applicable to expiredDateTime; lt - Operator 'less than' - Only applicable to expiredDateTime; le - Operator 'less than or equal' - Only applicable to expiredDateTime; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; ( ) - Symbols for grouping operands with their correct operator - -.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamIdentityProvidersList(params)
			return handleResponse(resp, err, "failed to list SAML identity providers")
		},
	}
}

func toolIamIdentityProvidersCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_identity_providers_create",
			mcp.WithDescription("Add a SAML identity provider. Adds a SAML identity provider to TrendAI Vision One."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name to display for this SAML identity provider.")),
			mcp.WithString("description", mcp.Description("An optional description for this SAML identity provider.")),
			mcp.WithString("file", mcp.Required(), mcp.Description("The metadata XML file describing this SAML identity provider. Must be a valid XML file under 1 MB. Path to a local file to upload.")),
			mcp.WithBoolean("enabled", mcp.Required(), mcp.Description("Specifies whether the identity provider is active. Default: True.")),
			mcp.WithString("samlAssertionNameAttribute", mcp.Description("The attribute name in the SAML assertion that the identity provider uses to identify users. Defaults to the NameID from the SAML assertion.")),
			mcp.WithString("samlAssertionDisplayNameAttribute", mcp.Description("The attribute name in the SAML assertion determining the display name of the user. Defaults to the NameID from the SAML assertion.")),
			mcp.WithString("samlAssertionGroupAttribute", mcp.Description("The attribute name in the SAML assertion used to determine if the user belongs to an IdP-Only SAML group.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "enabled", Kind: kindBoolean, Required: true}, {Name: "samlAssertionNameAttribute", Kind: kindString}, {Name: "samlAssertionDisplayNameAttribute", Kind: kindString}, {Name: "samlAssertionGroupAttribute", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamIdentityProvidersCreate(params)
			return handleResponse(resp, err, "failed to add a SAML identity provider")
		},
	}
}

func toolIamIdentityProviderDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_identity_provider_delete",
			mcp.WithDescription("Delete a SAML identity provider. Deletes the specified SAML identity provider."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a SAML identity provider.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamIdentityProviderDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete a SAML identity provider")
		},
	}
}

func toolIamIdentityProviderGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_identity_provider_get",
			mcp.WithDescription("Retrieve SAML identity provider details. Returns the details of the specified SAML identity provider."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a SAML identity provider.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamIdentityProviderGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to retrieve SAML identity provider details")
		},
	}
}

func toolIamIdentityProviderUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_identity_provider_update",
			mcp.WithDescription("Update a SAML identity provider. Updates a SAML identity provider. Fields omitted from the request are not updated."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a SAML identity provider.")),
			mcp.WithString("ifMatch", mcp.Description("The ETag value of the SAML identity provider to update. Retrieve the ETag by sending a GET request to the same endpoint. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("name", mcp.Description("The name to display for this SAML identity provider.")),
			mcp.WithString("description", mcp.Description("An optional description for this SAML identity provider.")),
			mcp.WithString("file", mcp.Description("The metadata XML file describing this SAML identity provider. Must be a valid XML file under 1 MB. Path to a local file to upload.")),
			mcp.WithBoolean("enabled", mcp.Description("Specifies whether the identity provider is active.")),
			mcp.WithString("samlAssertionNameAttribute", mcp.Description("The attribute name in the SAML assertion that the identity provider uses to identify users. Defaults to the NameID from the SAML assertion.")),
			mcp.WithString("samlAssertionDisplayNameAttribute", mcp.Description("The attribute name in the SAML assertion determining the display name of the user. Defaults to the NameID from the SAML assertion.")),
			mcp.WithString("samlAssertionGroupAttribute", mcp.Description("The attribute name in the SAML assertion used to determine if the user belongs to an IdP-Only SAML group.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}, {Name: "enabled", Kind: kindBoolean}, {Name: "samlAssertionNameAttribute", Kind: kindString}, {Name: "samlAssertionDisplayNameAttribute", Kind: kindString}, {Name: "samlAssertionGroupAttribute", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamIdentityProviderUpdate(id, params)
			return handleResponse(resp, err, "failed to update a SAML identity provider")
		},
	}
}

func toolIamRolesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_roles_list",
			mcp.WithDescription("List user roles. Displays the available user roles in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 50, 100, 200."), mcp.Enum("50", "100", "200")),
			mcp.WithString("orderBy", mcp.Description("Specifies the field by which the results are sorted. Default: lastUpdatedDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the role list. Supported fields: name - The name of a user role; id - The unique identifier of a user role; Supported operators: eq - Operator 'equal to'; and - Operator 'and'; or - Operator 'or'; not - Operator 'not'; ( ) - Symbols for grouping operands with their correct operator. Example: name eq 'full-access'\n")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamRolesList(params)
			return handleResponse(resp, err, "failed to list user roles")
		},
	}
}

func toolIamRolesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_roles_create",
			mcp.WithDescription("Create user role. Creates a new user role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the role.")),
			mcp.WithString("description", mcp.Required(), mcp.Description("The description of the role.")),
			mcp.WithArray("permissionKeys", mcp.Required(), mcp.Description("The permission keys assigned to the role."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithBoolean("allowApiKeyAssignment", mcp.Required(), mcp.Description("Whether the role can be assigned to an API key.")),
			mcp.WithBoolean("allowUserAccountAssignment", mcp.Required(), mcp.Description("Whether the role can be assigned to a user account.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString, Required: true}, {Name: "permissionKeys", Kind: kindArray, Required: true}, {Name: "allowApiKeyAssignment", Kind: kindBoolean, Required: true}, {Name: "allowUserAccountAssignment", Kind: kindBoolean, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamRolesCreate(params)
			return handleResponse(resp, err, "failed to create user role")
		},
	}
}

func toolIamRolesPermissionKeysList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_roles_permission_keys_list",
			mcp.WithDescription("Get all available permission keys. Displays all available permission keys."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.IamRolesPermissionKeysList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get all available permission keys")
		},
	}
}

func toolIamRoleDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_role_delete",
			mcp.WithDescription("Delete user role. Deletes the specified user role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a role.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamRoleDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete user role")
		},
	}
}

func toolIamRoleGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_role_get",
			mcp.WithDescription("Get user role. Displays information about the specified user role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a role.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamRoleGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get user role")
		},
	}
}

func toolIamRoleUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_role_update",
			mcp.WithDescription("Update user role. Updates the specified user role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a role.")),
			mcp.WithString("ifMatch", mcp.Description("Allows you to update an existing role. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("name", mcp.Description("The name of the role.")),
			mcp.WithString("description", mcp.Description("The description of the role.")),
			mcp.WithArray("permissionKeys", mcp.Description("The permission keys assigned to the role."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithBoolean("allowApiKeyAssignment", mcp.Description("Whether the role can be assigned to an API key.")),
			mcp.WithBoolean("allowUserAccountAssignment", mcp.Description("Whether the role can be assigned to a user account.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}, {Name: "permissionKeys", Kind: kindArray}, {Name: "allowApiKeyAssignment", Kind: kindBoolean}, {Name: "allowUserAccountAssignment", Kind: kindBoolean}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamRoleUpdate(id, params)
			return handleResponse(resp, err, "failed to update user role")
		},
	}
}

func toolIamRolePermissionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"iam_role_permissions_list",
			mcp.WithDescription("Get permissions of a specific role. Displays the permissions of the specified user role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a role.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IamRolePermissionsList(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get permissions of a specific role")
		},
	}
}
