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
	toolWorkbenchAlertNotesList,
	toolObservedAttackTechniquesList,
	toolOatDataPipelinesList,
	toolOatDataPipelineGet,
	toolOatDataPipelinePackagesList,
	toolOatDataPipelinePackageGet,
	toolWorkbenchAlertNoteGet,
	toolWorkbenchInsightsList,
	toolWorkbenchInsightGet,
	toolWorkbenchInsightImpactScopeEntitiesList,
	toolWorkbenchInsightIndicatorsList,
	toolWorkbenchInsightMatchedHighlightsList,
}

var ToolsetsWriteWorkench = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolOatDataPipelinesCreate,
	toolOatDataPipelinesDelete,
	toolOatDataPipelineUpdate,
	toolWorkbenchAlertNotesCreate,
	toolWorkbenchAlertNotesDelete,
	toolWorkbenchAlertNoteUpdate,
	toolWorkbenchAlertUpdate,
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
			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDate, err := optionalTimeValue("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDate, err := optionalTimeValue("endDateTime", request.GetArguments())
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
			alertId, err := requiredValue[string]("alertId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchGetAlertDetails(alertId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get alerts details")
		},
	}
}

func toolWorkbenchAlertNotesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
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
			alertId, err := requiredValue[string]("alertId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDate, err := optionalTimeValue("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDate, err := optionalTimeValue("endDateTime", request.GetArguments())
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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			nextBatchToken, err := optionalValue[string]("nextBatchToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			detectedStartDate, err := optionalTimeValue("detectedStartDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			detectedEndDate, err := optionalTimeValue("detectedEndDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			ingestedStartDate, err := optionalTimeValue("ingestedStartDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			ingestedEndDate, err := optionalTimeValue("ingestedEndDateTime", request.GetArguments())
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

func toolOatDataPipelinesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"oat_data_pipelines_list",
			mcp.WithDescription("Get active data pipelines. Displays all data pipelines that have registered users."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.OatDataPipelinesList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get active data pipelines")
		},
	}
}

func toolOatDataPipelinesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"oat_data_pipelines_create",
			mcp.WithDescription("Registers a customer to the Observed Attack Techniques data pipeline."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("riskLevels", mcp.Required(), mcp.Description("The list of severity levels to include in the results."), mcp.Items(map[string]any{"enum": []string{"info", "low", "medium", "high", "critical"}, "type": "string"})),
			mcp.WithBoolean("hasDetail", mcp.Required(), mcp.Description("Parameter that allows you to retrieve detailed logs from the Observed Attack Techniques data pipeline.")),
			mcp.WithString("description", mcp.Description("(Optional) Notes or comments about the pipeline.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "riskLevels", Kind: kindArray, Required: true}, {Name: "hasDetail", Kind: kindBoolean, Required: true}, {Name: "description", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.OatDataPipelinesCreate(params)
			return handleResponse(resp, err, "failed to registers a customer to the Observed Attack Techniques data pipeline")
		},
	}
}

func toolOatDataPipelinesDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"oat_data_pipelines_delete",
			mcp.WithDescription("Unregister from data pipeline. Unregisters a customer from the Observed Attack Techniques data pipeline."),
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

			resp, err := client.OatDataPipelinesDelete(params)
			return handleResponse(resp, err, "failed to unregister from data pipeline")
		},
	}
}

func toolOatDataPipelineGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"oat_data_pipeline_get",
			mcp.WithDescription("Get pipeline settings. Displays the settings of the specified data pipeline."),
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

			resp, err := client.OatDataPipelineGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get pipeline settings")
		},
	}
}

func toolOatDataPipelineUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"oat_data_pipeline_update",
			mcp.WithDescription("Modify data pipeline settings. Modifies the settings of the specified data pipeline."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of a data pipeline.")),
			mcp.WithString("ifMatch", mcp.Description("Parameter that allows you to update an existing resource. Use the ETag returned when the resource was retrieved.")),
			mcp.WithArray("riskLevels", mcp.Description("The severity levels to include in the results."), mcp.Items(map[string]any{"enum": []string{"info", "low", "medium", "high", "critical"}, "type": "string"})),
			mcp.WithBoolean("hasDetail", mcp.Description("Parameter that allows you to retrieve detailed logs from the Observed Attack Techniques data pipeline.")),
			mcp.WithString("description", mcp.Description("Notes or comments about the pipeline.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "riskLevels", Kind: kindArray}, {Name: "hasDetail", Kind: kindBoolean}, {Name: "description", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.OatDataPipelineUpdate(id, params)
			return handleResponse(resp, err, "failed to modify data pipeline settings")
		},
	}
}

func toolOatDataPipelinePackagesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"oat_data_pipeline_packages_list",
			mcp.WithDescription("Get Observed Attack Techniques event packages. Displays all the available packages from a data pipeline a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of a data pipeline.")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval time range, in ISO 8601 format.")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval time range, in ISO 8601 format.")),
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

			resp, err := client.OatDataPipelinePackagesList(id, params)
			return handleResponse(resp, err, "failed to get Observed Attack Techniques event packages")
		},
	}
}

func toolOatDataPipelinePackageGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"oat_data_pipeline_package_get",
			mcp.WithDescription("Get Observed Attack Techniques package. Retrieves the specified Observed Attack Techniques package."),
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

			resp, err := client.OatDataPipelinePackageGet(id, packageId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get Observed Attack Techniques package")
		},
	}
}

func toolWorkbenchAlertNotesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_alert_notes_create",
			mcp.WithDescription("Add alert note. Adds a note to the specified Workbench alert."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("alertId", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a Workbench alert.")),
			mcp.WithString("content", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a Workbench alert.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			alertId, err := pathValue("alertId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "content", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchAlertNotesCreate(alertId, params)
			return handleResponse(resp, err, "failed to add alert note")
		},
	}
}

func toolWorkbenchAlertNotesDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_alert_notes_delete",
			mcp.WithDescription("Delete alert notes. Deletes the specified notes from a Workbench alert."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("alertId", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a Workbench alert.")),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "Numeric string that identifies a Workbench alert note", "type": "integer"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			alertId, err := pathValue("alertId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchAlertNotesDelete(alertId, params)
			return handleResponse(resp, err, "failed to delete alert notes")
		},
	}
}

func toolWorkbenchAlertNoteGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_alert_note_get",
			mcp.WithDescription("Get alert note. Displays the specified Workbench alert note."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("alertId", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a Workbench alert.")),
			mcp.WithNumber("id", mcp.Required(), mcp.Description("Numeric string that identifies a Workbench alert note.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			alertId, err := pathValue("alertId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchAlertNoteGet(alertId, id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get alert note")
		},
	}
}

func toolWorkbenchAlertNoteUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_alert_note_update",
			mcp.WithDescription("Edit alert note. Modifies the content of the specified Workbench alert note."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("alertId", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a Workbench alert.")),
			mcp.WithNumber("id", mcp.Required(), mcp.Description("Numeric string that identifies a Workbench alert note.")),
			mcp.WithString("ifMatch", mcp.Description("Parameter that allows you to specify the version of the resource to be updated. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("content", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			alertId, err := pathValue("alertId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "content", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchAlertNoteUpdate(alertId, id, params)
			return handleResponse(resp, err, "failed to edit alert note")
		},
	}
}

func toolWorkbenchAlertUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_alert_update",
			mcp.WithDescription("Modify alert status. Modifies the status of an alert or investigation triggered in Workbench."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of a Workbench alert.")),
			mcp.WithString("ifMatch", mcp.Description("The ETag of the resource you want to update. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("investigationStatus", mcp.Description("The status of an investigation."), mcp.Enum("New", "In Progress", "True Positive", "False Positive", "Benign True Positive", "Closed")),
			mcp.WithString("status", mcp.Description("The status of a case or investigation."), mcp.Enum("Open", "In Progress", "Closed")),
			mcp.WithString("investigationResult", mcp.Description("The findings of a case or investigation."), mcp.Enum("No Findings", "Noteworthy", "True Positive", "False Positive", "Benign True Positive", "Other Findings")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "investigationStatus", Kind: kindString}, {Name: "status", Kind: kindString}, {Name: "investigationResult", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchAlertUpdate(id, params)
			return handleResponse(resp, err, "failed to modify alert status")
		},
	}
}

func toolWorkbenchInsightsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_insights_list",
			mcp.WithDescription("Get insights list. Displays information about Workbench insights that match the specified criteria in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field that specifies the order in which the results are sorted. Default: updatedDateTime desc.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 50. Default: 50."), mcp.Enum("50")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the insight list. Supported fields: id - The unique identifier of an insight - Any value; caseId - The unique identifier of a case - Any value; inProcess - Whether the insight graph is in process - true, false; createdDateTime - The date and time when the insight was created - Datetime in ISO 8601 format (yyyy-MM-ddThh:mm:ssZ in UTC); updatedDateTime - The date and time when the insight was last updated - Datetime in ISO 8601 format (yyyy-MM-ddThh:mm:ssZ in UTC); Supported operators: eq - Operator 'equal to'. - -; gt - Operator 'greater than'. - Only applicable to createdDateTime and updatedDateTime; ge - Operator 'greater than or equal'. - Only applicable to createdDateTime and updatedDateTime; le - Operator 'less than or equal'. - Only applicable to createdDateTime and updatedDateTime; lt - Operator 'less than'. - Only applicable to createdDateTime and updatedDateTime; and - Operator 'and'. - -; or - Operator 'or'. - -; not - Operator 'not'. - -; ( ) - Symbols for grouping operands with their correct operator. - -; Note: Include this parameter in every request that generates paginated output.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "orderBy", Kind: kindString}, {Name: "top", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchInsightsList(params)
			return handleResponse(resp, err, "failed to get insights list")
		},
	}
}

func toolWorkbenchInsightGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_insight_get",
			mcp.WithDescription("Get insight details. Displays details about the specified Workbench insight."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The Workbench insight ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchInsightGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get insight details")
		},
	}
}

func toolWorkbenchInsightImpactScopeEntitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_insight_impact_scope_entities_list",
			mcp.WithDescription("Get insight details (impact scope). Displays impact scope of the specified Workbench insight."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The Workbench insight ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 50. Default: 50."), mcp.Enum("50")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchInsightImpactScopeEntitiesList(id, params)
			return handleResponse(resp, err, "failed to get insight details (impact scope)")
		},
	}
}

func toolWorkbenchInsightIndicatorsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_insight_indicators_list",
			mcp.WithDescription("Get insight details (indicators). Displays indicators of the specified Workbench insight."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The Workbench insight ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 50. Default: 50."), mcp.Enum("50")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchInsightIndicatorsList(id, params)
			return handleResponse(resp, err, "failed to get insight details (indicators)")
		},
	}
}

func toolWorkbenchInsightMatchedHighlightsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"workbench_insight_matched_highlights_list",
			mcp.WithDescription("Get insight details (matched highlights). Displays the matched SAE filters and threat intelligence indicator patterns of the specified Workbench insight."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The Workbench insight ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 50. Default: 50."), mcp.Enum("50")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.WorkbenchInsightMatchedHighlightsList(id, params)
			return handleResponse(resp, err, "failed to get insight details (matched highlights)")
		},
	}
}
