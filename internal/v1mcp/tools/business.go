package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyBusiness = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolBusinessProfileGet,
}

var ToolsetsWriteBusiness = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolBusinessProfileUpdate,
}

func toolBusinessProfileGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"business_profile_get",
			mcp.WithDescription("Get business information. Retrieves the current business profile including the business identifier and business name."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.BusinessProfileGet(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get business information")
		},
	}
}

func toolBusinessProfileUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"business_profile_update",
			mcp.WithDescription("Change business name. Updates the business name for the organization."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("businessName", mcp.Description("Update the name for the business.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "businessName", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.BusinessProfileUpdate(params)
			return handleResponse(resp, err, "failed to change business name")
		},
	}
}
