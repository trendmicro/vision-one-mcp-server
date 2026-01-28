package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyCloudPosture = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudPostureAccountsList,
	toolCloudPostureAccountChecksList,
	toolCloudPostureTemplateScannerRun,
	toolCloudPostureAccountScanSettingsGet,
}

var ToolsetsWriteCloudPosture = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudPostureAccountScan,
	toolCloudPostureAccountScanSettingsUpdate,
}

func toolCloudPostureAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_accounts_list",
			mcp.WithDescription("List CSPM Accounts"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Min(50),
				mcp.Max(200),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				SkipToken: skipToken,
			}

			resp, err := client.CloudPostureListAccounts(queryParams)
			return handleStatusResponse(
				resp,
				err,
				http.StatusOK,
				"failed to list accounts",
			)
		},
	}
}

func toolCloudPostureAccountChecksList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_account_checks_list",
			mcp.WithDescription("List the checks of an account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudPostureChecks)),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Min(50),
				mcp.Max(200),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range")),
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

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDate, err := optionalTimeValue("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDate, err := optionalTimeValue("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:           top,
				StartDateTime: startDate,
				EndDateTime:   endDate,
				SkipToken:     skipToken,
			}

			resp, err := client.CloudPostureListAccountChecks(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list accounts checks")
		},
	}
}

func toolCloudPostureTemplateScannerRun(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_template_scanner_run",
			mcp.WithDescription("Scan an infrastructure as code template using the cloud posture template scanner"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("type",
				mcp.Required(),
				mcp.Enum("cloudformation-template", "terraform-template"),
			),
			mcp.WithString("content",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			templateType, err := requiredValue[string]("type", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			content, err := requiredValue[string]("content", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureScanTemplate(content, templateType)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to scan template")
		},
	}
}

func toolCloudPostureAccountScanSettingsGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_account_scan_settings_get",
			mcp.WithDescription("Get the scan settings for an account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureGetAccountScanSettings(accountId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get account scan settings")
		},
	}
}

func toolCloudPostureAccountScan(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_account_scan",
			mcp.WithDescription("Start scanning cloud posture account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureScanAccount(accountId)
			return handleStatusResponse(resp, err, http.StatusAccepted, "failed to start account scan")
		},
	}
}

func toolCloudPostureAccountScanSettingsUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_account_scan_settings_update",
			mcp.WithDescription("Update an account's scan settings"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
			),
			mcp.WithNumber("interval",
				mcp.Max(12),
				mcp.Min(1),
			),
			mcp.WithBoolean("enabled"),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			interval, err := optionalIntValue("interval", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			enabled, err := optionalPointerValue[bool]("enabled", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureUpdateAccountScanSettings(
				accountId,
				enabled,
				interval,
			)
			return handleStatusResponse(resp, err, http.StatusNoContent, "failed to update account scan settings")
		},
	}
}
