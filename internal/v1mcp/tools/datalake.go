package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyDatalake = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolDatalakeDataPipelinesList,
	toolDatalakeDataPipelineGet,
	toolDatalakeDataPipelinePackagesList,
	toolDatalakeDataPipelinePackageGet,
}

var ToolsetsWriteDatalake = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolDatalakeDataPipelinesCreate,
	toolDatalakeDataPipelinesDelete,
	toolDatalakeDataPipelineUpdate,
}

func toolDatalakeDataPipelinesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"datalake_data_pipelines_list",
			mcp.WithDescription("Get bound data pipelines. Displays all data pipelines that have a data type assigned."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.DatalakeDataPipelinesList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get bound data pipelines")
		},
	}
}

func toolDatalakeDataPipelinesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"datalake_data_pipelines_create",
			mcp.WithDescription("Bind a data type to a pipeline. Binds a data type to a pipeline. After binding, the data pipeline saves the specified data type for seven days."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("type", mcp.Required(), mcp.Description("Data type."), mcp.Enum("telemetry", "detection")),
			mcp.WithString("description", mcp.Description("Notes or comments about the pipeline.")),
			mcp.WithArray("subType", mcp.Description("The type of data to retrieve. Only applicable to the *Telemetry* data source. Default: ['all']."), mcp.Items(map[string]any{"enum": []string{"endpointActivity", "cloudActivity", "emailActivity", "mobileActivity", "networkActivity", "containerActivity", "identityActivity", "all"}, "type": "string"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "type", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "subType", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DatalakeDataPipelinesCreate(params)
			return handleResponse(resp, err, "failed to bind a data type to a pipeline")
		},
	}
}

func toolDatalakeDataPipelinesDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"datalake_data_pipelines_delete",
			mcp.WithDescription("Unbind data type from pipeline. Unbinds data type from pipeline. After unbinding the data pipeline stops saving the specified data type."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of a data pipeline.", "format": "uuid", "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DatalakeDataPipelinesDelete(params)
			return handleResponse(resp, err, "failed to unbind data type from pipeline")
		},
	}
}

func toolDatalakeDataPipelineGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"datalake_data_pipeline_get",
			mcp.WithDescription("Get pipeline information. Displays information about the specified data pipeline."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of a data pipeline.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DatalakeDataPipelineGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get pipeline information")
		},
	}
}

func toolDatalakeDataPipelineUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"datalake_data_pipeline_update",
			mcp.WithDescription("Update pipeline settings. Updates the settings of the specified data pipeline."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of a data pipeline.")),
			mcp.WithString("type", mcp.Description("Data type."), mcp.Enum("telemetry", "detection")),
			mcp.WithArray("subType", mcp.Description("The type of data to retrieve. Only applicable to the *Telemetry* data source. Default: ['all']."), mcp.Items(map[string]any{"enum": []string{"endpointActivity", "cloudActivity", "emailActivity", "mobileActivity", "networkActivity", "containerActivity", "identityActivity", "all"}, "type": "string"})),
			mcp.WithString("description", mcp.Description("Notes or comments about the pipeline.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "type", Kind: kindString}, {Name: "subType", Kind: kindArray}, {Name: "description", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DatalakeDataPipelineUpdate(id, params)
			return handleResponse(resp, err, "failed to update pipeline settings")
		},
	}
}

func toolDatalakeDataPipelinePackagesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"datalake_data_pipeline_packages_list",
			mcp.WithDescription("List available packages. Displays all the available packages from a data pipeline a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of a data pipeline.")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval time range, in ISO 8601 format.")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval time range.")),
			mcp.WithString("top", mcp.Description("Number of records displayed on a page. If no value is specified, 'top' defaults to 500. One of: 50, 100, 200, 500. Default: 500."), mcp.Enum("50", "100", "200", "500")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DatalakeDataPipelinePackagesList(id, params)
			return handleResponse(resp, err, "failed to list available packages")
		},
	}
}

func toolDatalakeDataPipelinePackageGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"datalake_data_pipeline_package_get",
			mcp.WithDescription("Get package. Retrieves the specified data pipeline package."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of a data pipeline.")),
			mcp.WithString("packageId", mcp.Required(), mcp.Description("The ID of a package stored in a data pipeline.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			packageId, err := pathValue("packageId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DatalakeDataPipelinePackageGet(id, packageId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get package")
		},
	}
}
