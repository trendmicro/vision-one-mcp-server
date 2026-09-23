package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyFileSecurity = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolFileSecurityStoragesList,
}

var ToolsetsWriteFileSecurity = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolFileSecurityStoragesUpdate,
}

func toolFileSecurityStoragesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"file_security_storages_list",
			mcp.WithDescription("Get storage container list. Displays a list of the storage containers connected to File Security in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 20, 50, 100. Default: 10."), mcp.Enum("10", "20", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: name desc.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the storage container list. This parameter includes every request that generates paginated output. Supported fields: provider - The cloud storage provider - aws, azure, gcp; accountId - The cloud account that owns the storage container. Its meaning depends on the provider: in AWS it is the account ID, in Azure the subscription ID, and in GCP the project ID. - Any value; name - The name of the storage container - Any value; region - The region of the storage container - Any value; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; ( ) - Symbols for grouping operands - -; Additional functions: Function - Description - Notes; contains() - Searches for a specified string in a field - Only applicable to name.")),
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

			resp, err := client.FileSecurityStoragesList(params)
			return handleResponse(resp, err, "failed to get storage container list")
		},
	}
}

func toolFileSecurityStoragesUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"file_security_storages_update",
			mcp.WithDescription("Update the protection status. Updates the status of the file storage protection for the specified storage containers."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Description("The number of storage containers to update."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the storage container", "type": "string"}, "status": map[string]any{"description": "The protection status to set for the storage container: * protected - Enable File Security scanning for this storage container.", "enum": []string{"protected", "notProtected"}, "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.FileSecurityStoragesUpdate(params)
			return handleResponse(resp, err, "failed to update the protection status")
		},
	}
}
