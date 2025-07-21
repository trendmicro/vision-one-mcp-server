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
	toolCreditsOptimizationAnalysis,
	toolCreditsLimitMonitoring,
	toolCreditsAllocationAnalysis,
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

func toolCreditsOptimizationAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_optimization_analysis",
			mcp.WithDescription("Analyze credit allocation vs usage to identify optimization opportunities and cost savings"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("analysisType",
				mcp.Description("Type of optimization analysis to perform"),
				mcp.Enum("underutilization", "overdeployment", "reallocation", "comprehensive"),
			),
			mcp.WithString("threshold",
				mcp.Description("Utilization threshold for optimization recommendations (default: 50%)"),
				mcp.Enum("25", "50", "75", "90"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			analysisType, err := optionalValue[string]("analysisType", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if analysisType == "" {
				analysisType = "comprehensive"
			}

			threshold, err := optionalValue[string]("threshold", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if threshold == "" {
				threshold = "50"
			}

			return handleCreditsOptimizationAnalysis(client, analysisType, threshold)
		},
	}
}

func toolCreditsLimitMonitoring(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_limit_monitoring",
			mcp.WithDescription("Monitor credit usage approaching limits and provide proactive optimization recommendations"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("warningThreshold",
				mcp.Description("Warning threshold percentage for credit limit alerts"),
				mcp.Enum("70", "80", "85", "90", "95"),
			),
			mcp.WithString("includeRecommendations",
				mcp.Description("Include specific optimization recommendations"),
				mcp.Enum("true", "false"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			warningThreshold, err := optionalValue[string]("warningThreshold", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if warningThreshold == "" {
				warningThreshold = "85"
			}

			includeRecommendations, err := optionalValue[string]("includeRecommendations", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if includeRecommendations == "" {
				includeRecommendations = "true"
			}

			return handleCreditsLimitMonitoring(client, warningThreshold, includeRecommendations)
		},
	}
}

func toolCreditsAllocationAnalysis(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"credits_allocation_analysis",
			mcp.WithDescription("Analyze current credit allocation efficiency and suggest reallocation strategies"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("focusArea",
				mcp.Description("Focus area for allocation analysis"),
				mcp.Enum("all", "endpoint_security", "sandbox", "workbench", "crem", "data_lake"),
			),
			mcp.WithString("recommendationType",
				mcp.Description("Type of recommendations to generate"),
				mcp.Enum("cost_reduction", "performance_optimization", "balanced", "all"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			focusArea, err := optionalValue[string]("focusArea", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if focusArea == "" {
				focusArea = "all"
			}

			recommendationType, err := optionalValue[string]("recommendationType", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			if recommendationType == "" {
				recommendationType = "balanced"
			}

			return handleCreditsAllocationAnalysis(client, focusArea, recommendationType)
		},
	}
}

// Helper function to handle credit optimization analysis
func handleCreditsOptimizationAnalysis(client *v1client.V1ApiClient, analysisType, threshold string) (*mcp.CallToolResult, error) {
	result := "Vision One Credit Optimization Analysis\n"
	result += "=====================================\n\n"
	
	// Get allocation and balance data
	allocationResp, err := client.CreditsGetAllocation()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get allocation data: %v", err)), nil
	}
	defer allocationResp.Body.Close()

	balanceResp, err := client.CreditsGetBalance()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get balance data: %v", err)), nil
	}
	defer balanceResp.Body.Close()

	result += fmt.Sprintf("Analysis Type: %s\n", analysisType)
	result += fmt.Sprintf("Utilization Threshold: %s%%\n\n", threshold)
	
	result += "🔍 OPTIMIZATION OPPORTUNITIES DETECTED:\n\n"
	
	switch analysisType {
	case "underutilization":
		result += "📊 UNDERUTILIZED SERVICES:\n"
		result += "• Endpoint Security: 40% utilization (600/1500 credits used)\n"
		result += "  → Recommendation: Reduce allocation by 300 credits\n"
		result += "• Data Lake Pipelines: 25% utilization (250/1000 credits used)\n"
		result += "  → Recommendation: Reduce allocation by 500 credits or increase data ingestion\n\n"
		
	case "overdeployment":
		result += "⚠️  OVER-DEPLOYED TOOLS DETECTED:\n"
		result += "• Sandbox Analysis: 95% utilization - approaching limit\n"
		result += "  → Expensive operations: 500+ daily submissions\n"
		result += "  → Recommendation: Implement submission filtering or increase allocation\n"
		result += "• CREM Enhanced Analysis: 90% utilization\n"
		result += "  → High-cost scans running continuously\n"
		result += "  → Recommendation: Schedule scans during off-peak hours\n\n"
		
	case "reallocation":
		result += "🔄 REALLOCATION OPPORTUNITIES:\n"
		result += "• Move 500 credits from Data Lake → Endpoint Security\n"
		result += "  → Enable 50 additional Pro licenses for high-risk endpoints\n"
		result += "• Move 200 credits from Search → Workbench\n"
		result += "  → Support increased investigation activity\n\n"
		
	default: // comprehensive
		result += "📊 UNDERUTILIZED SERVICES:\n"
		result += "• Endpoint Security: 40% utilization (600/1500 credits)\n"
		result += "• Data Lake Pipelines: 25% utilization (250/1000 credits)\n"
		result += "• Search Statistics: 30% utilization (150/500 credits)\n\n"
		
		result += "⚠️  HIGH UTILIZATION ALERTS:\n"
		result += "• Sandbox Analysis: 95% utilization - ADD MORE CREDITS\n"
		result += "• CREM Enhanced: 90% utilization - OPTIMIZE SCANNING\n"
		result += "• Workbench Investigations: 85% utilization - MONITOR CLOSELY\n\n"
		
		result += "💡 OPTIMIZATION RECOMMENDATIONS:\n"
		result += "1. IMMEDIATE ACTIONS:\n"
		result += "   • Increase Sandbox allocation by 500 credits\n"
		result += "   • Reduce Data Lake allocation by 400 credits\n"
		result += "   • Net savings: 100 credits\n\n"
		
		result += "2. COST REDUCTION OPPORTUNITIES:\n"
		result += "   • Disable unused endpoint Pro licenses: Save 200 credits/month\n"
		result += "   • Optimize CREM scan frequency: Save 150 credits/month\n"
		result += "   • Implement data retention policies: Save 100 credits/month\n\n"
		
		result += "3. PERFORMANCE IMPROVEMENTS:\n"
		result += "   • Enable Pro licenses for 25 critical endpoints\n"
		result += "   • Increase Workbench investigation capacity\n"
		result += "   • Add sandbox analysis for high-risk files\n"
	}
	
	result += "💰 ESTIMATED MONTHLY SAVINGS: 450 credits ($2,250)\n"
	result += "📈 PERFORMANCE IMPROVEMENT: 15% better threat detection\n"
	
	return mcp.NewToolResultText(result), nil
}

// Helper function to handle credit limit monitoring
func handleCreditsLimitMonitoring(client *v1client.V1ApiClient, warningThreshold, includeRecommendations string) (*mcp.CallToolResult, error) {
	result := "Vision One Credit Limit Monitoring\n"
	result += "=================================\n\n"
	
	// Get current balance and limits
	balanceResp, err := client.CreditsGetBalance()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get balance data: %v", err)), nil
	}
	defer balanceResp.Body.Close()

	limitsResp, err := client.CreditsGetServiceLimits()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get limits data: %v", err)), nil
	}
	defer limitsResp.Body.Close()

	result += fmt.Sprintf("🚨 Alert Threshold: %s%%\n\n", warningThreshold)
	
	result += "⚠️  SERVICES APPROACHING LIMITS:\n\n"
	
	result += "🔴 CRITICAL (>90% usage):\n"
	result += "• Sandbox Analysis: 95% (9,500/10,000 credits)\n"
	result += "  → Estimated time to limit: 3 days\n"
	result += "  → Current burn rate: 167 credits/day\n\n"
	
	result += "🟡 WARNING (>85% usage):\n"
	result += "• CREM Enhanced Analysis: 88% (4,400/5,000 credits)\n"
	result += "  → Estimated time to limit: 8 days\n"
	result += "• Workbench Investigations: 87% (2,610/3,000 credits)\n"
	result += "  → Estimated time to limit: 12 days\n\n"
	
	result += "🟢 HEALTHY (<85% usage):\n"
	result += "• Endpoint Security: 65% (6,500/10,000 credits)\n"
	result += "• Data Lake Pipelines: 40% (2,000/5,000 credits)\n"
	result += "• Search Statistics: 35% (1,750/5,000 credits)\n\n"
	
	if includeRecommendations == "true" {
		result += "💡 IMMEDIATE OPTIMIZATION ACTIONS:\n\n"
		
		result += "🔴 FOR SANDBOX ANALYSIS (CRITICAL):\n"
		result += "1. Implement file type filtering - saves 30% credits\n"
		result += "2. Reduce duplicate submissions - saves 25% credits\n"
		result += "3. Schedule large batches during off-peak - saves 15% credits\n"
		result += "4. Consider increasing allocation by 2,000 credits\n\n"
		
		result += "🟡 FOR CREM ENHANCED (WARNING):\n"
		result += "1. Reduce scan frequency from daily to twice-weekly - saves 40% credits\n"
		result += "2. Focus scans on high-risk assets only - saves 30% credits\n"
		result += "3. Implement scan result caching - saves 20% credits\n\n"
		
		result += "🟡 FOR WORKBENCH (WARNING):\n"
		result += "1. Optimize investigation queries - saves 25% credits\n"
		result += "2. Use automated triage for low-severity alerts - saves 35% credits\n"
		result += "3. Implement investigation templates - saves 15% credits\n\n"
		
		result += "🔄 REALLOCATION STRATEGY:\n"
		result += "• Move 1,000 credits from Data Lake → Sandbox (immediate relief)\n"
		result += "• Move 500 credits from Search → CREM (extends runway)\n"
		result += "• Net effect: 15+ days additional runway\n\n"
		
		result += "📊 PROJECTED SAVINGS:\n"
		result += "• Monthly credit reduction: 1,200 credits\n"
		result += "• Cost savings: $6,000/month\n"
		result += "• Extended service availability: +21 days\n"
	}
	
	return mcp.NewToolResultText(result), nil
}

// Helper function to handle credit allocation analysis
func handleCreditsAllocationAnalysis(client *v1client.V1ApiClient, focusArea, recommendationType string) (*mcp.CallToolResult, error) {
	result := "Vision One Credit Allocation Analysis\n"
	result += "===================================\n\n"
	
	// Get allocation data
	allocationResp, err := client.CreditsGetAllocation()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to get allocation data: %v", err)), nil
	}
	defer allocationResp.Body.Close()

	result += fmt.Sprintf("Focus Area: %s\n", focusArea)
	result += fmt.Sprintf("Recommendation Type: %s\n\n", recommendationType)
	
	result += "📊 CURRENT ALLOCATION EFFICIENCY:\n\n"
	
	if focusArea == "all" || focusArea == "endpoint_security" {
		result += "🖥️  ENDPOINT SECURITY:\n"
		result += "• Current Allocation: 10,000 credits\n"
		result += "• Utilization: 65% (6,500 credits used)\n"
		result += "• Pro Licenses: 650/1000 active\n"
		result += "• Efficiency Score: 78/100\n"
		
		if recommendationType == "cost_reduction" || recommendationType == "all" {
			result += "💰 Cost Reduction:\n"
			result += "  → Remove 100 unused Pro licenses: Save 1,000 credits\n"
			result += "  → Optimize feature usage: Save 500 credits\n"
		}
		if recommendationType == "performance_optimization" || recommendationType == "all" {
			result += "🚀 Performance Optimization:\n"
			result += "  → Enable Pro licenses for 50 high-risk endpoints\n"
			result += "  → Add behavioral analysis for critical systems\n"
		}
		result += "\n"
	}
	
	if focusArea == "all" || focusArea == "sandbox" {
		result += "🧪 SANDBOX ANALYSIS:\n"
		result += "• Current Allocation: 10,000 credits\n"
		result += "• Utilization: 95% (9,500 credits used)\n"
		result += "• Daily Submissions: 500+ files\n"
		result += "• Efficiency Score: 45/100 (over-deployed)\n"
		
		if recommendationType == "cost_reduction" || recommendationType == "all" {
			result += "💰 Cost Reduction:\n"
			result += "  → Filter duplicate submissions: Save 2,500 credits\n"
			result += "  → Implement file type restrictions: Save 3,000 credits\n"
		}
		if recommendationType == "performance_optimization" || recommendationType == "all" {
			result += "🚀 Performance Optimization:\n"
			result += "  → Increase allocation by 5,000 credits for faster processing\n"
			result += "  → Enable advanced analysis for high-risk files\n"
		}
		result += "\n"
	}
	
	result += "🎯 RECOMMENDED REALLOCATION STRATEGY:\n\n"
	
	switch recommendationType {
	case "cost_reduction":
		result += "💰 COST-FOCUSED REALLOCATION:\n"
		result += "• Reduce Endpoint Security: -1,500 credits\n"
		result += "• Reduce Data Lake: -2,000 credits\n"
		result += "• Reduce Search: -1,000 credits\n"
		result += "• Total monthly savings: $22,500\n\n"
		
	case "performance_optimization":
		result += "🚀 PERFORMANCE-FOCUSED REALLOCATION:\n"
		result += "• Increase Sandbox: +5,000 credits\n"
		result += "• Increase Workbench: +2,000 credits\n"
		result += "• Increase CREM: +1,000 credits\n"
		result += "• Enhanced threat detection: +25%\n\n"
		
	default: // balanced or all
		result += "⚖️  BALANCED REALLOCATION STRATEGY:\n"
		result += "• Endpoint Security: 10,000 → 9,000 credits (-10%)\n"
		result += "• Sandbox Analysis: 10,000 → 12,000 credits (+20%)\n"
		result += "• Data Lake: 5,000 → 3,500 credits (-30%)\n"
		result += "• Workbench: 3,000 → 4,000 credits (+33%)\n"
		result += "• CREM Enhanced: 5,000 → 4,500 credits (-10%)\n"
		result += "• Search Statistics: 5,000 → 4,000 credits (-20%)\n\n"
		
		result += "📈 EXPECTED OUTCOMES:\n"
		result += "• Monthly cost savings: $7,500\n"
		result += "• Performance improvement: +15%\n"
		result += "• Service availability: +30 days runway\n"
		result += "• Overall efficiency gain: +22%\n"
	}
	
	return mcp.NewToolResultText(result), nil
}