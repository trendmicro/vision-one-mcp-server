package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyWorkench = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolWorkbenchAlertsList,
	toolWorkbenchAlertDetailGet,
	toolObservedAttackTechniquesList,
}

func toolWorkbenchAlertsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_alerts_list",
			mcp.WithDescription("List Trend Vision One Workbench Alerts"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterWorkbenchAlerts)),
			mcp.WithString("orderBy",
				mcp.Description("the field to order by"),
				mcp.Enum(withOrdering(
					asc_desc,
					"id",
					"caseId",
					"name",
					"status",
					"investigationResult",
					"modelId",
					"model",
					"score",
					"severity",
					"createdDateTime",
					"updatedDateTime",
					"firstInvestigatedDateTime",
				)...),
			),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDate, err := optionalTimeValue("startDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDate, err := optionalTimeValue("endDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				OrderBy:       orderBy,
				StartDateTime: startDate,
				EndDateTime:   endDate,
			}

			resp, err := client.WorkbenchAlertsList(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list workbench alerts")
		},
	}
}

func toolWorkbenchAlertDetailGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_alert_detail_get",
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithDescription("Displays information about the specified alert."),
			mcp.WithString("alertId",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			alertId, err := requiredValue[string]("alertId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchGetAlertDetails(alertId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get alerts details")
		},
	}
}

func TookWorkbenchAlertNotesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_alert_notes_list",
			mcp.WithDescription("Displays the notes of the specified Workbench alert."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("alertId",
				mcp.Required(),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterWorkbenchNotes)),
			mcp.WithString(
				"orderBy",
				mcp.Description("the field to order by"),
				mcp.Enum(withOrdering(
					asc_desc,
					"id",
					"creatorMailAddress",
					"creatorName",
					"lastUpdatedBy",
					"createdDateTime",
					"lastUpdatedDateTime",
				)...),
			),
			mcp.WithString(
				"top",
				mcp.Description("The number of records to display per page."),
				mcp.Enum("50", "100", "200"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			alertId, err := requiredValue[string]("alertId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDate, err := optionalTimeValue("startDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDate, err := optionalTimeValue("endDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:           top,
				OrderBy:       orderBy,
				StartDateTime: startDate,
				EndDateTime:   endDate,
			}

			resp, err := client.WorkbenchGetAlertNotes(alertId, filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list alert notes")
		},
	}
}

func toolObservedAttackTechniquesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_observed_attack_techniques_list",
			mcp.WithDescription("List observed attack techniques"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.ObservedAttackFilter)),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum("50", "100", "200"),
			),
			mcp.WithString("detectedStartDateTime",
				mcp.Description("The start of the event detection data retrieval time range in ISO 8601 format."),
			),
			mcp.WithString("detectedEndDateTime",
				mcp.Description("The end of the event detection data retrieval time range in ISO 8601 format."),
			),
			mcp.WithString("ingestedStartDateTime",
				mcp.Description("The beginning of the data ingestion time range in ISO 8601 format."),
			),
			mcp.WithString("ingestedEndDateTime",
				mcp.Description("The end of the data ingestion time range in ISO 8601 format."),
			),
			mcp.WithString("nextBatchToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
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

			nextBatchToken, err := optionalValue[string]("nextBatchToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			detectedStartDate, err := optionalTimeValue("detectedStartDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			detectedEndDate, err := optionalTimeValue("detectedEndDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			ingestedStartDate, err := optionalTimeValue("ingestedStartDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			ingestedEndDate, err := optionalTimeValue("ingestedEndDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:                   top,
				DetectedStartDateTime: detectedStartDate,
				DetectedEndDateTime:   detectedEndDate,
				IngestedStartDateTime: ingestedStartDate,
				IngestedEndDateTime:   ingestedEndDate,
				NextBatchToken:        nextBatchToken,
			}

			resp, err := client.ObservedAttackTechniquesList(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list observed attack techniques")
		},
	}
}
