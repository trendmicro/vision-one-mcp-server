package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyThreatIntel = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolThreatIntelSuspiciousObjectsList,
	toolThreatIntelExceptionsList,
	toolThreatIntelIntelligenceReportsList,
	toolThreatIntelIntelligenceReportGet,
	toolThreatIntelTasksList,
	toolThreatIntelTaskResultsGet,
	toolThreatIntelFeedIndicatorsList,
	toolThreatIntelFeedsList,
	toolThreatIntelFeedFilterDefinitionGet,
}

var ToolsetsWriteThreatIntel = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolThreatIntelSuspiciousObjectsAdd,
	toolThreatIntelSuspiciousObjectsDelete,
	toolThreatIntelExceptionsAdd,
	toolThreatIntelExceptionsDelete,
	toolThreatIntelIntelligenceReportsDelete,
	toolThreatIntelSweepTrigger,
}

func toolThreatIntelSuspiciousObjectsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_suspicious_objects_list",
			mcp.WithDescription("Retrieves information about domains, file SHA-1, file SHA-256, IP addresses, email addresses, or URLs in the Suspicious Object List"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterSuspiciousObjects)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "riskLevel", "lastModifiedDateTime", "expiredDateTime")...),
			),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum("50", "100", "200"),
			),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range in ISO 8601 format")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalValue[string]("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalValue[string]("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.ThreatIntelQueryParameters{
				OrderBy:       orderBy,
				Top:           top,
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			}

			resp, err := client.ThreatIntelListSuspiciousObjects(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list suspicious objects")
		},
	}
}

func toolThreatIntelSuspiciousObjectsAdd(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_suspicious_objects_add",
			mcp.WithDescription("Adds information about domains, file SHA-1, file SHA-256, IP addresses, email addresses, or URLs to the Suspicious Object List"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("type",
				mcp.Required(),
				mcp.Description("The type of suspicious object"),
				mcp.Enum("url", "domain", "ip", "senderMailAddress", "fileSha1", "fileSha256"),
			),
			mcp.WithString("value",
				mcp.Required(),
				mcp.Description("The value of the suspicious object (URL, domain, IP, email address, or file hash)"),
			),
			mcp.WithString("description",
				mcp.Description("Brief description of the suspicious object"),
			),
			mcp.WithString("scanAction",
				mcp.Description("Action that connected products apply after detecting a suspicious object"),
				mcp.Enum("block", "log"),
			),
			mcp.WithString("riskLevel",
				mcp.Description("Risk level of the suspicious object"),
				mcp.Enum("high", "medium", "low"),
			),
			mcp.WithNumber("daysToExpiration",
				mcp.Description("Number of days before the object expires from the list"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			objType, err := requiredValue[string]("type", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			value, err := requiredValue[string]("value", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			scanAction, err := optionalValue[string]("scanAction", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			riskLevel, err := optionalValue[string]("riskLevel", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			daysToExpiration, err := optionalIntValue("daysToExpiration", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			obj := v1client.SuspiciousObject{
				Description:      description,
				ScanAction:       scanAction,
				RiskLevel:        riskLevel,
				DaysToExpiration: daysToExpiration,
			}

			switch objType {
			case "url":
				obj.URL = value
			case "domain":
				obj.Domain = value
			case "ip":
				obj.IP = value
			case "senderMailAddress":
				obj.SenderMailAddress = value
			case "fileSha1":
				obj.FileSha1 = value
			case "fileSha256":
				obj.FileSha256 = value
			}

			resp, err := client.ThreatIntelAddSuspiciousObjects([]v1client.SuspiciousObject{obj})
			return handleStatusResponse(resp, err, http.StatusMultiStatus, "failed to add suspicious object")
		},
	}
}

func toolThreatIntelSuspiciousObjectsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_suspicious_objects_delete",
			mcp.WithDescription("Deletes information about domains, file SHA-1, file SHA-256, IP addresses, email addresses, or URLs from the Suspicious Object List"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("type",
				mcp.Required(),
				mcp.Description("The type of suspicious object"),
				mcp.Enum("url", "domain", "ip", "senderMailAddress", "fileSha1", "fileSha256"),
			),
			mcp.WithString("value",
				mcp.Required(),
				mcp.Description("The value of the suspicious object to delete"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			objType, err := requiredValue[string]("type", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			value, err := requiredValue[string]("value", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			obj := v1client.SuspiciousObjectDelete{}

			switch objType {
			case "url":
				obj.URL = value
			case "domain":
				obj.Domain = value
			case "ip":
				obj.IP = value
			case "senderMailAddress":
				obj.SenderMailAddress = value
			case "fileSha1":
				obj.FileSha1 = value
			case "fileSha256":
				obj.FileSha256 = value
			}

			resp, err := client.ThreatIntelDeleteSuspiciousObjects([]v1client.SuspiciousObjectDelete{obj})
			return handleStatusResponse(resp, err, http.StatusMultiStatus, "failed to delete suspicious object")
		},
	}
}

func toolThreatIntelExceptionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_exceptions_list",
			mcp.WithDescription("Retrieves information about domains, file SHA-1, file SHA-256, IP addresses, sender addresses, or URLs in the Exception List"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterSuspiciousObjectExceptions)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "lastModifiedDateTime")...),
			),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum("50", "100", "200"),
			),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range in ISO 8601 format")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalValue[string]("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalValue[string]("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.ThreatIntelQueryParameters{
				OrderBy:       orderBy,
				Top:           top,
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			}

			resp, err := client.ThreatIntelListExceptions(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list exception objects")
		},
	}
}

func toolThreatIntelExceptionsAdd(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_exceptions_add",
			mcp.WithDescription("Adds domains, file SHA-1, file SHA-256, IP addresses, sender addresses, or URLs to the Exception List"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("type",
				mcp.Required(),
				mcp.Description("The type of exception object"),
				mcp.Enum("url", "domain", "ip", "senderMailAddress", "fileSha1", "fileSha256"),
			),
			mcp.WithString("value",
				mcp.Required(),
				mcp.Description("The value of the exception object (URL, domain, IP, email address, or file hash)"),
			),
			mcp.WithString("description",
				mcp.Description("Brief description of the exception object"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			objType, err := requiredValue[string]("type", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			value, err := requiredValue[string]("value", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			obj := v1client.SuspiciousObjectException{
				Description: description,
			}

			switch objType {
			case "url":
				obj.URL = value
			case "domain":
				obj.Domain = value
			case "ip":
				obj.IP = value
			case "senderMailAddress":
				obj.SenderMailAddress = value
			case "fileSha1":
				obj.FileSha1 = value
			case "fileSha256":
				obj.FileSha256 = value
			}

			resp, err := client.ThreatIntelAddExceptions([]v1client.SuspiciousObjectException{obj})
			return handleStatusResponse(resp, err, http.StatusMultiStatus, "failed to add exception object")
		},
	}
}

func toolThreatIntelExceptionsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_exceptions_delete",
			mcp.WithDescription("Deletes the specified objects from the Exception List"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("type",
				mcp.Required(),
				mcp.Description("The type of exception object"),
				mcp.Enum("url", "domain", "ip", "senderMailAddress", "fileSha1", "fileSha256"),
			),
			mcp.WithString("value",
				mcp.Required(),
				mcp.Description("The value of the exception object to delete"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			objType, err := requiredValue[string]("type", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			value, err := requiredValue[string]("value", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			obj := v1client.SuspiciousObjectDelete{}

			switch objType {
			case "url":
				obj.URL = value
			case "domain":
				obj.Domain = value
			case "ip":
				obj.IP = value
			case "senderMailAddress":
				obj.SenderMailAddress = value
			case "fileSha1":
				obj.FileSha1 = value
			case "fileSha256":
				obj.FileSha256 = value
			}

			resp, err := client.ThreatIntelDeleteExceptions([]v1client.SuspiciousObjectDelete{obj})
			return handleStatusResponse(resp, err, http.StatusMultiStatus, "failed to delete exception object")
		},
	}
}

func toolThreatIntelIntelligenceReportsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_intelligence_reports_list",
			mcp.WithDescription("Retrieves a list of custom intelligence reports created from imported or retrieved data"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterIntelligenceReports)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "updatedDateTime", "createdDateTime")...),
			),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum("50", "100", "200"),
			),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range in ISO 8601 format")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalValue[string]("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalValue[string]("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.ThreatIntelQueryParameters{
				Filter:        filter,
				OrderBy:       orderBy,
				Top:           top,
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			}

			resp, err := client.ThreatIntelListIntelligenceReports(qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list intelligence reports")
		},
	}
}

func toolThreatIntelIntelligenceReportGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_intelligence_report_get",
			mcp.WithDescription("Downloads a custom intelligence report as a STIX Bundle"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("reportId",
				mcp.Required(),
				mcp.Description("The unique identifier of the intelligence report"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			reportId, err := requiredValue[string]("reportId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ThreatIntelGetIntelligenceReport(reportId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get intelligence report")
		},
	}
}

func toolThreatIntelIntelligenceReportsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_intelligence_reports_delete",
			mcp.WithDescription("Deletes the specified custom intelligence reports"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("reportIds",
				mcp.Required(),
				mcp.Description("Array of intelligence report IDs to delete"),
				mcp.Items(map[string]any{"type": "string"}),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			reportIds := []string{}
			if ids, ok := request.GetArguments()["reportIds"].([]any); ok && len(ids) > 0 {
				for _, id := range ids {
					reportId, ok := id.(string)
					if !ok {
						return mcp.NewToolResultError("each report ID must be a string"), nil
					}
					reportIds = append(reportIds, reportId)
				}
			}

			resp, err := client.ThreatIntelDeleteIntelligenceReports(reportIds)
			return handleStatusResponse(resp, err, http.StatusMultiStatus, "failed to delete intelligence reports")
		},
	}
}

func toolThreatIntelSweepTrigger(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_sweep_trigger",
			mcp.WithDescription("Searches your environment for threat indicators specified in a custom intelligence report"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("reportId",
				mcp.Required(),
				mcp.Description("The unique identifier of the intelligence report to sweep"),
			),
			mcp.WithString("sweepType",
				mcp.Required(),
				mcp.Description("The type of sweeping task"),
				mcp.Enum("manual", "stixShifter"),
			),
			mcp.WithString("description",
				mcp.Description("Brief description of the sweep task"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			reportId, err := requiredValue[string]("reportId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			sweepType, err := requiredValue[string]("sweepType", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			sweep := v1client.IntelligenceReportSweep{
				ID:          reportId,
				SweepType:   sweepType,
				Description: description,
			}

			resp, err := client.ThreatIntelTriggerSweep([]v1client.IntelligenceReportSweep{sweep})
			return handleStatusResponse(resp, err, http.StatusMultiStatus, "failed to trigger sweep")
		},
	}
}

func toolThreatIntelTasksList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_tasks_list",
			mcp.WithDescription("Displays information about threat intelligence tasks and asynchronous jobs"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterThreatIntelTasks)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "lastActionDateTime")...),
			),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum("50", "100", "200"),
			),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range in ISO 8601 format")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalValue[string]("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalValue[string]("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.ThreatIntelQueryParameters{
				Filter:        filter,
				OrderBy:       orderBy,
				Top:           top,
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
			}

			resp, err := client.ThreatIntelListTasks(qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list tasks")
		},
	}
}

func toolThreatIntelTaskResultsGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_task_results_get",
			mcp.WithDescription("Retrieves the results of a threat intelligence task"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("taskId",
				mcp.Required(),
				mcp.Description("The unique identifier of the task"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			taskId, err := requiredValue[string]("taskId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ThreatIntelGetTaskResults(taskId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get task results")
		},
	}
}

func toolThreatIntelFeedIndicatorsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_feed_indicators_list",
			mcp.WithDescription("Retrieves a list of IoCs from Trend Threat Intelligence Feed"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range in ISO 8601 format")),
			mcp.WithString("top",
				mcp.Description("The number of IoCs returned by a query (maximum 10,000)"),
				mcp.Enum("1000", "5000", "10000"),
			),
			mcp.WithString("indicatorObjectFormat",
				mcp.Description("The desired format for the query response"),
				mcp.Enum("stixBundle", "taxiiEnvelope"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			startDateTime, err := optionalValue[string]("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalValue[string]("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			indicatorObjectFormat, err := optionalValue[string]("indicatorObjectFormat", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.ThreatIntelFeedParameters{
				StartDateTime:         startDateTime,
				EndDateTime:           endDateTime,
				Top:                   top,
				IndicatorObjectFormat: indicatorObjectFormat,
			}

			resp, err := client.ThreatIntelListFeedIndicators(qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list feed indicators")
		},
	}
}

func toolThreatIntelFeedsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_feeds_list",
			mcp.WithDescription("Retrieves a list of intelligence reports from the Trend Threat Intelligence Feed with associated objects and relationships"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("contextualFilter", mcp.Description(tooldescriptions.FilterThreatIntelFeeds)),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range in ISO 8601 format")),
			mcp.WithString("topReport",
				mcp.Description("The number of reports returned by a query (maximum 20)"),
				mcp.Enum("5", "10", "20"),
			),
			mcp.WithString("responseObjectFormat",
				mcp.Description("The preferred format for the query response"),
				mcp.Enum("stixBundle", "taxiiEnvelope"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			contextualFilter, err := optionalValue[string]("contextualFilter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalValue[string]("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalValue[string]("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			topReport, err := optionalStrInt("topReport", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			responseObjectFormat, err := optionalValue[string]("responseObjectFormat", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.ThreatIntelFeedParameters{
				StartDateTime:        startDateTime,
				EndDateTime:          endDateTime,
				TopReport:            topReport,
				ResponseObjectFormat: responseObjectFormat,
			}

			resp, err := client.ThreatIntelListFeeds(contextualFilter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list feeds")
		},
	}
}

func toolThreatIntelFeedFilterDefinitionGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"threatintel_feed_filter_definition_get",
			mcp.WithDescription("Retrieves supported filter keys and values for Trend Threat Intelligence Feed queries"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.ThreatIntelGetFeedFilterDefinition()
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get feed filter definition")
		},
	}
}
