package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlySandbox = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolSandboxAnalysisResultsList,
	toolSandboxAnalysisResultGet,
	toolSandboxAnalysisResultInvestigationPackageGet,
	toolSandboxAnalysisResultReportGet,
	toolSandboxAnalysisResultSuspiciousObjectsList,
	toolSandboxSubmissionUsageGet,
	toolSandboxTasksList,
	toolSandboxTaskGet,
}

var ToolsetsWriteSandbox = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolSandboxFilesAnalyze,
	toolSandboxUrlsAnalyze,
}

func toolSandboxAnalysisResultsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_analysis_results_list",
			mcp.WithDescription("Get list of analysis results. Displays the analysis results of all submitted objects in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the start of the data retrieval time range. If no value is specified, 'startDateTime' defaults to 180 days before the time the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("orderBy", mcp.Description("Parameter that allows you to sort the retrieved elements in ascending or descending order. If no order is specified, the elements are shown in ascending order. Supported fields and operators: * analysisCompletionDateTime * riskLevel. Default: analysisCompletionDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the analysis results list. Supported fields and operators: 'sha1' - SHA-1 hash value of an object; 'sha256' - SHA-256 hash value of an object; 'md5' - MD5 hash value of an object; 'riskLevel' - The risk level assigned to the object by the sandbox. Possible values: 'high', 'medium', 'low', 'noRisk'; 'type' - Object type. Possible values: 'file', 'url'; 'id' - Unique alphanumeric string that identifies the analysis results of a submission; 'eq' - Abbreviation of the operator 'equal to'; 'and' - Operator 'and'; 'or' - Operator 'or'; 'not' - Operator 'not'; '( )' - Symbols for grouping operands with their correct operator. Example: (riskLevel eq 'high') or (riskLevel eq 'medium')")),
			mcp.WithString("top", mcp.Description("Number of records displayed on a page. One of: 50, 100, 200. Default: 50."), mcp.Enum("50", "100", "200")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "orderBy", Kind: kindString}, {Name: "filter", Kind: kindString}, {Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxAnalysisResultsList(params)
			return handleResponse(resp, err, "failed to get list of analysis results")
		},
	}
}

func toolSandboxAnalysisResultGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_analysis_result_get",
			mcp.WithDescription("Get analysis results. Displays the analysis results of the specified object."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies the analysis results of a submission.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxAnalysisResultGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get analysis results")
		},
	}
}

func toolSandboxAnalysisResultInvestigationPackageGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_analysis_result_investigation_package_get",
			mcp.WithDescription("Download Investigation Package. Downloads the Investigation Package of the specified object."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies the analysis results of a submission.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxAnalysisResultInvestigationPackageGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to download Investigation Package")
		},
	}
}

func toolSandboxAnalysisResultReportGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_analysis_result_report_get",
			mcp.WithDescription("Download analysis results. Downloads the analysis results of the specified object as PDF."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies the analysis results of a submission.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxAnalysisResultReportGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to download analysis results")
		},
	}
}

func toolSandboxAnalysisResultSuspiciousObjectsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_analysis_result_suspicious_objects_list",
			mcp.WithDescription("Download suspicious object list. Downloads the suspicious object list associated to the specified object."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies the analysis results of a submission.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxAnalysisResultSuspiciousObjectsList(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to download suspicious object list")
		},
	}
}

func toolSandboxFilesAnalyze(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_files_analyze",
			mcp.WithDescription("Submit file to sandbox. Submits a file to the sandbox for analysis."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("file", mcp.Required(), mcp.Description("File submitted to the sandbox. Path to a local file to upload.")),
			mcp.WithString("documentPassword", mcp.Description("Password encoded in Base64 used to decrypt the submitted file sample. The maximum password length (without encoding) is 128 bytes.")),
			mcp.WithString("archivePassword", mcp.Description("Password encoded in Base64 used to decrypt the submitted archive. The maximum password length (without encoding) is 128 bytes.")),
			mcp.WithString("arguments", mcp.Description("Parameter that allows you to specify Base64-encoded command line arguments to run the submitted file. The maximum argument length before encoding is 1024 bytes. Arguments are only available for Portable Executable (PE) files and script files.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "documentPassword", Kind: kindString}, {Name: "archivePassword", Kind: kindString}, {Name: "arguments", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxFilesAnalyze(params)
			return handleResponse(resp, err, "failed to submit file to sandbox")
		},
	}
}

func toolSandboxSubmissionUsageGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_submission_usage_get",
			mcp.WithDescription("Get daily reserve. Retrieves the maximum number of samples your company can submit to the sandbox per day. Note: Samples marked as \"Not analyzed\" do not count towards the daily reserve."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.SandboxSubmissionUsageGet(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get daily reserve")
		},
	}
}

func toolSandboxTasksList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_tasks_list",
			mcp.WithDescription("List submissions. Displays all submitted objects in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the start of the data retrieval time range. If no value is specified, 'startDateTime' defaults to 180 days before the time the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("dateTimeTarget", mcp.Description("Parameter that indicates which field is used to sort the submission list. Default: createdDateTime."), mcp.Enum("createdDateTime", "lastActionDateTime")),
			mcp.WithString("orderBy", mcp.Description("Parameter that allows you to sort the retrieved elements in ascending or descending order. If no order is specified, the elements are shown in ascending order. Supported fields and operators: * createdDateTime * lastActionDateTime. Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the submissions list. Supported fields and operators: 'status' - Analysis status of an object. Possible values: 'succeeded', 'running', 'failed'; 'action' - Action applied to an object. Possible values: 'analyzeFile', 'analyzeUrl'; 'sha1' - SHA-1 hash value of an object; 'sha256' - SHA-256 hash value of an object; 'md5' - MD5 hash value of an object; 'id' - Unique alphanumeric string that identifies a submission; 'eq' - Abbreviation of the operator 'equal to'; 'and' - Operator 'and'; 'or' - Operator 'or'; 'not' - Operator 'not'; '( )' - Symbols for grouping operands with their correct operator. Example: status eq 'succeeded'")),
			mcp.WithString("top", mcp.Description("Number of records displayed on a page. One of: 50, 100, 200. Default: 50."), mcp.Enum("50", "100", "200")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "dateTimeTarget", Kind: kindString}, {Name: "orderBy", Kind: kindString}, {Name: "filter", Kind: kindString}, {Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxTasksList(params)
			return handleResponse(resp, err, "failed to list submissions")
		},
	}
}

func toolSandboxTaskGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_task_get",
			mcp.WithDescription("Get submission status. Displays information about the specified submission."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Displays information about the specified submission.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxTaskGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get submission status")
		},
	}
}

func toolSandboxUrlsAnalyze(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"sandbox_urls_analyze",
			mcp.WithDescription("Submit URLs to sandbox. Submits URLs to the sandbox for analysis."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Description("URLs to be submitted."), mcp.Items(map[string]any{"properties": map[string]any{"url": map[string]any{"description": "URL to be submitted.", "type": "string"}}, "required": []string{"url"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SandboxUrlsAnalyze(params)
			return handleResponse(resp, err, "failed to submit URLs to sandbox")
		},
	}
}
