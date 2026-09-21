package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyHealthcheck = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolHealthcheckConnectivityGet,
}

func toolHealthcheckConnectivityGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"healthcheck_connectivity_get",
			mcp.WithDescription("Check availability of API service. Checks the connection to the API service and verifies if your authentication token is valid."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.HealthcheckConnectivityGet(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to check availability of API service")
		},
	}
}
