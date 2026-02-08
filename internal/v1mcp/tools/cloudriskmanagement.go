package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyCloudRiskManagement = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudRiskManagementAccountsList,
	toolCloudRiskManagementAccountScanRulesGet,
	toolCloudRiskManagementServicesList,
}

func toolCloudRiskManagementAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_risk_management_accounts_list",
			mcp.WithDescription("Displays the cloud accounts you can access in a paginated list"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudRiskManagementAccounts)),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Min(50),
				mcp.Max(200),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.GetArguments())
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

			resp, err := client.CloudRiskManagementListAccounts(filter, queryParams)
			return handleStatusResponse(
				resp,
				err,
				http.StatusOK,
				"failed to list cloud accounts",
			)
		},
	}
}

func toolCloudRiskManagementAccountScanRulesGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_risk_management_account_scan_rules_get",
			mcp.WithDescription("Displays the settings for all rules of the specified account in a paginated list"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
				mcp.Description("The Cloud Risk Management ID of the account"),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudRiskManagementScanRules)),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Min(50),
				mcp.Max(200),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalIntValue("top", request.GetArguments())
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

			resp, err := client.CloudRiskManagementGetAccountScanRules(accountId, filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get account scan rules")
		},
	}
}

func toolCloudRiskManagementServicesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_risk_management_services_list",
			mcp.WithDescription("Retrieves a list of cloud services and their associated rules supported by Cloud Risk Management"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudRiskManagementServices)),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Min(50),
				mcp.Max(200),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.GetArguments())
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

			resp, err := client.CloudRiskManagementListServices(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list cloud services")
		},
	}
}
