package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyEmail = []func(client *v1client.V1ApiClient) mcpserver.ServerTool{
	toolEmailSecurityAccountsList,
	toolEmailSecurityDomainsList,
	toolEmailSecurityServersList,
}

func toolEmailSecurityAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"email_security_accounts_list",
			mcp.WithDescription(
				"Returns all email accounts managed by an email protection solution or with email sensor detection enabled.",
			),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top",
				mcp.Description("The number of email accounts returned on each page. Supported range: 10-1000"),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterEmailAccounts)),
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

			qp := v1client.QueryParameters{
				Top: top,
			}

			resp, err := client.EmailSecurityListAccounts(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list email accounts")
		},
	}
}

func toolEmailSecurityDomainsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"email_security_domains_list",
			mcp.WithDescription("Returns all email domains managed by an email protection solution."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top",
				mcp.Description("The number of email domains returned on each page. Supported range: 10-1000"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			qp := v1client.QueryParameters{
				Top: top,
			}

			resp, err := client.EmailSecurityListAccounts("", qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list email accounts")
		},
	}
}

func toolEmailSecurityServersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"email_security_servers_list",
			mcp.WithDescription("Returns all email servers managed by an on-premises email protection solution."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top",
				mcp.Description("The number of email servers returned on each page. Supported range: 10-1000"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top: top,
			}

			resp, err := client.EmailSecurityListServers("", qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list email accounts")
		},
	}
}
