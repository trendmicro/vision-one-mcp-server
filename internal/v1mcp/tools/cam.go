package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyCAM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCAMAwsAccountsList,
	toolCAMAwsAccountGet,
	toolCAMGcpAccountsList,
	toolCAMGcpAccountGet,
	toolCAMAlibabaAccountsList,
	toolCAMAlibabaAccountGet,
}

func toolCAMAwsAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_accounts_list",
			mcp.WithDescription("List AWS accounts managed by Cloud Accounts Management"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description("The number of records to display per page")),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterAWSAccounts)),
			mcp.WithString("nextBatchToken", mcp.Description("Token used to retrieve the next page of results")),
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

			nextBatchToken, err := optionalValue[string]("nextBatchToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:            top,
				NextBatchToken: nextBatchToken,
			}

			resp, err := client.CAMListAWSAccounts(filter, qp)
			if err != nil {
				return nil, err
			}

			return handleStatusResponse(resp, err, http.StatusOK, "failed to list cam aws accounts")
		},
	}
}

func toolCAMAwsAccountGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_account_get",
			mcp.WithDescription("Get the details of an AWS account managed by Cloud Account Management"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			resp, err := client.CAMGetAWSAccount(accountId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get aws account details")
		},
	}
}

func toolCAMGcpAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_gcp_accounts_list",
			mcp.WithDescription("List Google Cloud Projects managed by Cloud Account Management"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description("The number of records to display per page")),
			mcp.WithString("filter", mcp.Description(tooldescriptions.CAMListGCPProjectsFilterDescription)),
			mcp.WithString("nextBatchToken", mcp.Description("Token used to retrieve the next page of results")),
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

			nextBatchToken, err := optionalValue[string]("nextBatchToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:            top,
				NextBatchToken: nextBatchToken,
			}

			resp, err := client.CAMListGCPAccounts(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list gcp accounts")

		},
	}
}

func toolCAMGcpAccountGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_gcp_account_get",
			mcp.WithDescription("Get the details of a GCP project managed by Cloud Account Manangement"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CAMGetGCPAccountDetails(accountId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get gcp project details")
		},
	}
}

func toolCAMAlibabaAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_alibaba_accounts_list",
			mcp.WithDescription("Displays all Alibaba Cloud accounts connected to Trend Vision One in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description("The number of records to display per page")),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterAlibabaAccounts)),
			mcp.WithString("nextBatchToken", mcp.Description("Token used to retrieve the next page of results")),
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

			nextBatchToken, err := optionalValue[string]("nextBatchToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:            top,
				NextBatchToken: nextBatchToken,
			}

			resp, err := client.CAMListAlibabaAccounts(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list alibaba accounts")
		},
	}
}

func toolCAMAlibabaAccountGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_alibaba_account_get",
			mcp.WithDescription("Get the details of an Alibaba account managed by Cloud Account Manangement"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CAMGetAlibabaAccountDetails(accountId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get gcp project details")
		},
	}
}
