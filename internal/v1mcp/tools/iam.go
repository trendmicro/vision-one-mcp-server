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
}

var ToolsetsWriteIAM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolIamApiKeysDelete,
	toolIamAccountInvite,
	toolIamAccountUpdate,
	toolIamAccountDelete,
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
					"lastUsedDateTime asc",
					"lastUsedDateTime desc",
					"lastModifiedDateTime asc",
					"lastModifiedDateTime desc",
					"expiredDateTime asc",
					"expiredDateTime desc",
					"createdDateTime asc",
					"createdDateTime desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithNumber("top", mcp.Description("The number of records to display per page.")),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
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
				mcp.Required(),
				mcp.Items(
					map[string]any{
						"type": "string",
					},
				),
				mcp.Description("Array of API Key Ids to delete"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			keysToDelete := []string{}
			if keyIds, ok := request.Params.Arguments["apiKeyIds"].([]any); ok && len(keyIds) > 0 {
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
			email, err := requiredValue[string]("email", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			role, err := requiredValue[string]("role", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			authType, err := requiredValue[string]("authType", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", request.Params.Arguments)
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
			mcp.WithNumber("top", mcp.Description("The number of records to display per page.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.Params.Arguments)
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
			accountId, err := requiredValue[string]("accountId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			role, err := optionalValue[string]("role", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			status, err := optionalValue[string]("status", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", request.Params.Arguments)
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
			accountId, err := requiredValue[string]("accountId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.IAMDeleteAccount(accountId)
			return handleStatusResponse(resp, err, http.StatusNoContent, "failed to delete account")
		},
	}
}
