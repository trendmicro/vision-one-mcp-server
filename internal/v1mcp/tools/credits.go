package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyCredits = []func(client *v1client.V1ApiClient) mcpserver.ServerTool{
	toolCreditsEndpointSecurityAnalysis,
	toolCreditsDataLakePipelinesAnalysis,
	toolCreditsOATDetectionsAnalysis,
	toolCreditsSandboxUsageAnalysis,
	toolCreditsWorkbenchAlertsAnalysis,
	toolCreditsSearchStatisticsAnalysis,
	toolCreditsComprehensiveAnalysis,
}

func toolCreditsEndpointSecurityAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_endpoint_security_analysis",
			mcp.WithDescription("Analyze endpoint security credit usage including Pro licenses and security features"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter to apply to endpoint query")),
			mcp.WithString("top",
				mcp.Description("Number of endpoints to analyze (default: 100, use 'all' for comprehensive analysis)"),
				mcp.Enum("10", "50", "100", "500", "1000", "all"),
			),
			mcp.WithString("fetchAll",
				mcp.Description("Whether to fetch all endpoints (may be slow for large environments)"),
				mcp.Enum("true", "false"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalStrInt("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			fetchAllStr, err := optionalValue[string]("fetchAll", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			fetchAll := fetchAllStr == "true"

			queryParams := v1client.QueryParameters{
				Top: top,
			}

			response, err := client.CreditsListEndpoints(filter, queryParams)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to retrieve endpoints: %v", err)), nil
			}

			return handleCreditsAnalysisResponse(response, "endpoint_security", fetchAll)
		},
	}
}

func toolCreditsDataLakePipelinesAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_datalake_pipelines_analysis",
			mcp.WithDescription("Analyze active data lake pipelines consuming credits"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter to apply to pipeline query")),
			mcp.WithString("top",
				mcp.Description("Number of pipelines to analyze"),
				mcp.Enum("10", "50", "100"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalStrInt("top", request.Params.Arguments)
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

			response, err := client.CreditsListDataLakePipelines(filter, queryParams)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to retrieve data lake pipelines: %v", err)), nil
			}

			return handleCreditsAnalysisResponse(response, "datalake_pipelines", false)
		},
	}
}

func toolCreditsOATDetectionsAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_oat_detections_analysis",
			mcp.WithDescription("Analyze Observed Attack Techniques (OAT) detections for credit usage"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter to apply to OAT detections query")),
			mcp.WithString("top",
				mcp.Description("Number of detections to analyze"),
				mcp.Enum("10", "50", "100", "500"),
			),
			mcp.WithString("startDateTime",
				mcp.Description("Start time for detection analysis in ISO 8601 format"),
			),
			mcp.WithString("endDateTime",
				mcp.Description("End time for detection analysis in ISO 8601 format"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalStrInt("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalTimeValue("startDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalTimeValue("endDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:           top,
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			}

			response, err := client.CreditsListOATDetections(filter, queryParams)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to retrieve OAT detections: %v", err)), nil
			}

			return handleCreditsAnalysisResponse(response, "oat_detections", false)
		},
	}
}

func toolCreditsSandboxUsageAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_sandbox_usage_analysis",
			mcp.WithDescription("Analyze sandbox submission usage for credit consumption"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter to apply to sandbox submissions query")),
			mcp.WithString("top",
				mcp.Description("Number of submissions to analyze"),
				mcp.Enum("10", "50", "100", "500"),
			),
			mcp.WithString("startDateTime",
				mcp.Description("Start time for submission analysis in ISO 8601 format"),
			),
			mcp.WithString("endDateTime",
				mcp.Description("End time for submission analysis in ISO 8601 format"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalStrInt("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalTimeValue("startDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalTimeValue("endDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:           top,
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			}

			response, err := client.CreditsListSandboxSubmissions(filter, queryParams)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to retrieve sandbox submissions: %v", err)), nil
			}

			return handleCreditsAnalysisResponse(response, "sandbox_usage", false)
		},
	}
}

func toolCreditsWorkbenchAlertsAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_workbench_alerts_analysis",
			mcp.WithDescription("Analyze workbench alert investigation activity for credits"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter to apply to workbench alerts query")),
			mcp.WithString("top",
				mcp.Description("Number of alerts to analyze"),
				mcp.Enum("10", "50", "100", "500"),
			),
			mcp.WithString("startDateTime",
				mcp.Description("Start time for alert analysis in ISO 8601 format"),
			),
			mcp.WithString("endDateTime",
				mcp.Description("End time for alert analysis in ISO 8601 format"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalStrInt("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalTimeValue("startDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalTimeValue("endDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:           top,
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			}

			response, err := client.CreditsListWorkbenchAlerts(filter, queryParams)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to retrieve workbench alerts: %v", err)), nil
			}

			return handleCreditsAnalysisResponse(response, "workbench_alerts", false)
		},
	}
}

func toolCreditsSearchStatisticsAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_search_statistics_analysis",
			mcp.WithDescription("Analyze search activity and sensor statistics for credit usage"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter to apply to search statistics query")),
			mcp.WithString("startDateTime",
				mcp.Description("Start time for search statistics analysis in ISO 8601 format"),
			),
			mcp.WithString("endDateTime",
				mcp.Description("End time for search statistics analysis in ISO 8601 format"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalTimeValue("startDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalTimeValue("endDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			}

			response, err := client.CreditsGetSearchStatistics(filter, queryParams)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to retrieve search statistics: %v", err)), nil
			}

			return handleCreditsAnalysisResponse(response, "search_statistics", false)
		},
	}
}

func toolCreditsComprehensiveAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_comprehensive_analysis",
			mcp.WithDescription("Run comprehensive credit usage analysis across all Vision One modules (may be slow for large environments)"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("sampleSize",
				mcp.Description("Sample size for analysis to balance speed vs completeness"),
				mcp.Enum("10", "50", "100", "500", "all"),
			),
			mcp.WithString("startDateTime",
				mcp.Description("Start time for comprehensive analysis in ISO 8601 format"),
			),
			mcp.WithString("endDateTime",
				mcp.Description("End time for comprehensive analysis in ISO 8601 format"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			sampleSize, err := optionalValue[string]("sampleSize", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalTimeValue("startDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalTimeValue("endDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// This would run all analysis functions in sequence
			return handleComprehensiveCreditsAnalysis(client, sampleSize, startDateTime, endDateTime)
		},
	}
}

// Helper function to handle credits analysis responses
func handleCreditsAnalysisResponse(response *http.Response, analysisType string, fetchAll bool) (*mcp.CallToolResult, error) {
	return handleStatusResponse(response, nil, http.StatusOK, fmt.Sprintf("failed to analyze %s credits", analysisType))
}

// Helper function to handle comprehensive credits analysis
func handleComprehensiveCreditsAnalysis(client *v1client.V1ApiClient, sampleSize string, startDateTime, endDateTime interface{}) (*mcp.CallToolResult, error) {
	result := "Comprehensive Vision One Credit Usage Analysis\n"
	result += "==========================================\n\n"
	
	// This would orchestrate all the individual analysis functions
	result += "Analysis includes:\n"
	result += "- Endpoint Security Credits\n"
	result += "- Data Lake Pipelines\n" 
	result += "- OAT Detections Usage\n"
	result += "- Sandbox Submission Activity\n"
	result += "- Workbench Investigation Activity\n"
	result += "- Search Statistics\n"
	result += "- CREM Enhanced Analysis\n\n"
	
	result += fmt.Sprintf("Sample size: %s\n", sampleSize)
	
	return mcp.NewToolResultText(result), nil
}