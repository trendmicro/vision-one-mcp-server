package tools

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyEndpoint = []func(client *v1client.V1ApiClient) mcpserver.ServerTool{
	toolEndpointSecurityEndpointsList,
	toolEndpointSecurityEndpointGet,
	toolEndpointSecurityTaskList,
	toolEndpointSecurityTaskGet,
	toolEndpointSecurityVersionControlPoliciesList,
	toolEndpointSecurityAgentUpdatePoliciesList,
	toolEndpointSecuritySchedulesList,
	toolEndpointSecurityScheduleGet,
	toolEndpointSecurityKernelSupportPackageUpdatePoliciesList,
	toolEndpointSecurityVersionControlPolicyPrioritiesList,
}

var ToolsetsWriteEndpoint = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolEndpointSecurityEndpointsApplySensorPolicy,
	toolEndpointSecurityEndpointsDelete,
	toolEndpointSecurityEndpointsExport,
	toolEndpointSecurityEndpointsRemoveOverriddenSensorPolicy,
	toolEndpointSecuritySchedulesCreate,
	toolEndpointSecuritySchedulesDelete,
	toolEndpointSecuritySchedulesUpdate,
	toolEndpointSecurityVersionControlPolicyDelete,
	toolEndpointSecurityVersionControlPolicyPriorityUpdate,
}

func toolEndpointSecurityEndpointsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_endpoints_list",
			mcp.WithDescription("Displays a detailed list of your endpoints"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterEndpoints)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(
					withOrdering(
						asc_desc,
						"agentGuid",
						"eppAgentLastConnectedDateTime",
						"eppAgentLastScannedDateTime",
						"edrSensorLastConnectedDateTime",
					)...,
				),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
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

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.EndpointSecurityListEndpoints(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list endpoints")
		},
	}
}

func toolEndpointSecurityEndpointGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_endpoint_get",
			mcp.WithDescription("Displays the detailed profile of the specified endpoint"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("endpointID", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			endpointID, err := requiredValue[string]("endpointID", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityGetEndpoint(endpointID)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get endpoint details")
		},
	}
}

func toolEndpointSecurityTaskList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_tasks_list",
			mcp.WithDescription("Displays the tasks of your endpoints in a paginated list"),
			mcp.WithReadOnlyHintAnnotation(true),

			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterEndpointTasks)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(
					withOrdering(
						asc_desc,
						"createdDateTime",
						"lastActionDateTime",
					)...,
				),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),

			mcp.WithString("startDateTime",
				mcp.Description("The start time of the data retrieval range, in ISO 8601 format."),
			),

			mcp.WithString("endDateTime",
				mcp.Description("The end time of the data retrieval range, in ISO 8601 format."),
			),
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

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDateTime, err := optionalTimeValue("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDateTime, err := optionalTimeValue("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				StartDateTime: startDateTime,
				EndDateTime:   endDateTime,
				OrderBy:       orderBy,
				SkipToken:     skipToken,
			}

			resp, err := client.EndpointSecurityListTasks(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list tasks")
		},
	}
}

func toolEndpointSecurityTaskGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_task_get",
			mcp.WithDescription("Displays the status of the specified task"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("taskID", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			taskID, err := requiredValue[string]("taskID", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityGetTask(taskID)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get task details")
		},
	}
}

func toolEndpointSecurityVersionControlPoliciesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_version_control_policies_list",
			mcp.WithDescription("Displays your Endpoint Version Control policies"),
			mcp.WithReadOnlyHintAnnotation(true),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.EndpointSecurityListVersionControlPolicies(qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list version control policies")
		},
	}
}

func toolEndpointSecurityAgentUpdatePoliciesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_agent_update_policies_list",
			mcp.WithDescription("Displays the available agent update policies"),
			mcp.WithReadOnlyHintAnnotation(true),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.EndpointSecurityListAgentUpdatePolicies()
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get agent update policies")
		},
	}
}

func toolEndpointSecurityEndpointsApplySensorPolicy(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_endpoints_apply_sensor_policy",
			mcp.WithDescription("Override endpoint sensor policy. Overrides endpoint sensor policy settings for multiple endpoints."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("agentGuids", mcp.Required(), mcp.Description("A list of agent GUIDs."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("sensorDetectionAndResponse", mcp.Description("The configuration for sensor detection and response on the endpoint."), mcp.Enum("enabled", "disabled")),
			mcp.WithString("advancedRiskTelemetry", mcp.Description("The configuration for advanced risk telemetry on the endpoint."), mcp.Enum("enabled", "disabled")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "agentGuids", Kind: kindArray, Required: true}, {Name: "sensorDetectionAndResponse", Kind: kindString}, {Name: "advancedRiskTelemetry", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityEndpointsApplySensorPolicy(params)
			return handleResponse(resp, err, "failed to override endpoint sensor policy")
		},
	}
}

func toolEndpointSecurityEndpointsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_endpoints_delete",
			mcp.WithDescription("Remove endpoints. Removes the specified endpoints from Endpoint Inventory."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of the endpoint on Endpoint Inventory.", "format": "uuid", "type": "string"}}, "required": []string{"agentGuid"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityEndpointsDelete(params)
			return handleResponse(resp, err, "failed to remove endpoints")
		},
	}
}

func toolEndpointSecurityEndpointsExport(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_endpoints_export",
			mcp.WithDescription("Export information about endpoints. Generates a CSV file with the information about your endpoints in Endpoint Inventory. The CSV file is compressed into a ZIP file for download."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the endpoint information list.")),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: agentGuid.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "filter", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityEndpointsExport(params)
			return handleResponse(resp, err, "failed to export information about endpoints")
		},
	}
}

func toolEndpointSecurityEndpointsRemoveOverriddenSensorPolicy(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_endpoints_remove_overridden_sensor_policy",
			mcp.WithDescription("Remove overriden endpoint security policy settings. Removes endpoint security policy overrides from the specified endpoints, including: - Overrides applied through the /v3.0/endpointSecurity/endpoints/applySensorPolicy API endpoint - Overrides applied through the TrendAI Vision One™ console - The hypersensitive detection mode' setting."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The agent GUID", "type": "string"}}, "required": []string{"agentGuid"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityEndpointsRemoveOverriddenSensorPolicy(params)
			return handleResponse(resp, err, "failed to remove overriden endpoint security policy settings")
		},
	}
}

func toolEndpointSecuritySchedulesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_schedules_list",
			mcp.WithDescription("Search Schedules."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("Sort order for schedules. Default: updatedDateTime asc.")),
			mcp.WithNumber("top", mcp.Description("Number of schedules per page. Default: 200.")),
			mcp.WithString("filter", mcp.Description("Filter for schedules by name keyword or resource IDs. Supported fields: id - The array of schedule IDs. - Any value; name - The keyword in schedule name. - Any value; isPredefined - Whether the schedule is predefined. - true, false; Supported operators: eq - Operator 'equal to'. and - Operator 'and'. in - Operator 'in'. Supported functions: contains - The string partially matches.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "orderBy", Kind: kindString}, {Name: "top", Kind: kindNumber}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecuritySchedulesList(params)
			return handleResponse(resp, err, "failed to search Schedules")
		},
	}
}

// scheduleVEventSchema is the iCalendar VEvent shape accepted by the endpoint security
// schedules API. dtStart, duration and rRule each only accept a restricted subset of
// iCalendar syntax; the patterns below mirror the API's own validation so invalid values
// are rejected before a request is ever sent. See
// https://icalendar.org/iCalendar-RFC-5545/3-6-1-event-component.html
var scheduleVEventSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"dtStart": map[string]any{
			"type":        "string",
			"description": "iCalendar DTSTART in local time format. The schedule is triggered at the specified local time in each endpoint's time zone. (Z-suffix is not accepted)",
			"pattern":     "^[0-9]{8}T[0-9]{6}$",
		},
		"duration": map[string]any{
			"type":        "string",
			"description": "iCalendar Duration (only days and hours are supported), e.g. P1D, PT2H, P1DT2H.",
			"pattern":     "^P(\\d+D)?(T(\\d+H)?)?$",
		},
		"rRule": map[string]any{
			"type":        "string",
			"description": "iCalendar RRule (only DAILY, WEEKLY and MONTHLY frequencies are supported), e.g. FREQ=DAILY or FREQ=WEEKLY;BYDAY=MO,WE.",
			"pattern":     "^FREQ=(DAILY|WEEKLY|MONTHLY)(;[A-Z]+=[^;]+)*$",
		},
	},
	"required": []string{"dtStart", "duration", "rRule"},
}

var (
	scheduleDTStartPattern  = regexp.MustCompile(`^[0-9]{8}T[0-9]{6}$`)
	scheduleDurationPattern = regexp.MustCompile(`^P(\d+D)?(T(\d+H)?)?$`)
	scheduleRRulePattern    = regexp.MustCompile(`^FREQ=(DAILY|WEEKLY|MONTHLY)(;[A-Z]+=[^;]+)*$`)
)

// validateScheduleItems checks every vEvent in a schedule create/update request against the
// same constraints advertised in scheduleVEventSchema. The Trend Vision One API enforces these
// server-side, but its error response for a violation is an opaque error code with no
// explanation (e.g. {"error":{"code":"Error_001001","message":"TraceId: ..."}}), so callers
// need this check to get a specific, human-readable error before the request is ever sent.
func validateScheduleItems(args map[string]any) error {
	items, ok := args["items"].([]any)
	if !ok {
		return nil
	}

	for _, item := range items {
		itemMap, ok := item.(map[string]any)
		if !ok {
			continue
		}

		vEvents, ok := itemMap["vEvents"].([]any)
		if !ok {
			continue
		}

		for _, ve := range vEvents {
			veMap, ok := ve.(map[string]any)
			if !ok {
				continue
			}

			if dtStart, ok := veMap["dtStart"].(string); ok && !scheduleDTStartPattern.MatchString(dtStart) {
				return fmt.Errorf("invalid dtStart %q: expected iCalendar local time format YYYYMMDDTHHMMSS (e.g. 20261001T020000)", dtStart)
			}
			if duration, ok := veMap["duration"].(string); ok && !scheduleDurationPattern.MatchString(duration) {
				return fmt.Errorf("invalid duration %q: only days and/or hours are supported (e.g. P1D, PT2H, P1DT2H)", duration)
			}
			if rRule, ok := veMap["rRule"].(string); ok && !scheduleRRulePattern.MatchString(rRule) {
				return fmt.Errorf("invalid rRule %q: only DAILY, WEEKLY, and MONTHLY frequencies are supported (e.g. FREQ=DAILY, FREQ=WEEKLY;BYDAY=MO,WE)", rRule)
			}
		}
	}

	return nil
}

func toolEndpointSecuritySchedulesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_schedules_create",
			mcp.WithDescription("Create Schedule. Create a new schedule."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "Schedule description", "type": "string"}, "name": map[string]any{"description": "Schedule name", "type": "string"}, "vEvents": map[string]any{"description": "iCalendar VEvent list. Currently only accepts exactly one VEvent per schedule.", "items": scheduleVEventSchema, "minItems": 1, "maxItems": 1, "type": "array"}}, "required": []string{"name", "vEvents"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			if err := validateScheduleItems(args); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecuritySchedulesCreate(params)
			return handleResponse(resp, err, "failed to create Schedule")
		},
	}
}

func toolEndpointSecuritySchedulesDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_schedules_delete",
			mcp.WithDescription("Delete Schedule. Delete multiple schedules by ID."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "Schedule ID for deletion", "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecuritySchedulesDelete(params)
			return handleResponse(resp, err, "failed to delete Schedule")
		},
	}
}

func toolEndpointSecuritySchedulesUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_schedules_update",
			mcp.WithDescription("Update Schedule. Update multiple existing schedules."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "Schedule description", "type": "string"}, "id": map[string]any{"description": "Schedule ID", "type": "string"}, "name": map[string]any{"description": "Schedule name", "type": "string"}, "vEvents": map[string]any{"description": "iCalendar VEvent list. Currently only accepts exactly one VEvent per schedule.", "items": scheduleVEventSchema, "minItems": 1, "maxItems": 1, "type": "array"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			if err := validateScheduleItems(args); err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecuritySchedulesUpdate(params)
			return handleResponse(resp, err, "failed to update Schedule")
		},
	}
}

func toolEndpointSecurityScheduleGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_schedule_get",
			mcp.WithDescription("Get Schedule by ID. Retrieve schedule by ID."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Schedule ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityScheduleGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get Schedule by ID")
		},
	}
}

func toolEndpointSecurityKernelSupportPackageUpdatePoliciesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_kernel_support_package_update_policies_list",
			mcp.WithDescription("Get kernel support package versions. Displays a list of available kernel support package versions that can be applied in endpoint version control policies."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200. Default: 10."), mcp.Enum("10", "50", "100", "200")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the kernel support package version list. Supported fields: version - The version of the kernel support package list - Any value; Supported operators: eq - Operator 'equal to' - -; and - Operator and. - -; or - Operator or. - -; not - Operator not. - -; ( ) - Symbols for grouping operands. - -.")),
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

			resp, err := client.EndpointSecurityKernelSupportPackageUpdatePoliciesList(params)
			return handleResponse(resp, err, "failed to get kernel support package versions")
		},
	}
}

func toolEndpointSecurityVersionControlPolicyDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_version_control_policy_delete",
			mcp.WithDescription("Delete version control policy. Deletes the specified endpoint version control policy."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the policy.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityVersionControlPolicyDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete version control policy")
		},
	}
}

func toolEndpointSecurityVersionControlPolicyPrioritiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_version_control_policy_priorities_list",
			mcp.WithDescription("Get version control policy priorities. Displays your endpoint version control policy priorities in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 50, 100, 200. Default: 100."), mcp.Enum("50", "100", "200")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the endpoint version control policies. Supported fields: agentUpdateStatus - The status of agent updates - pause, onSchedule, disable; agentUpdatePolicy - The version the agent or sensor will update to - Any value; componentUpdateStatus - The status of component updates - pause, onSchedule; componentUpdatePolicy - The version the component will update to - n, n-1, n-2, n-3, n-4, n-5, n-6, n-7, n-8; kernelSupportPackageUpdateStatus - The status of kernel support package updates - pause, always, kernelCompatibility, never; kernelSupportPackageUpdatePolicy - The version the kernel support package will update to - Any value; Supported operators: eq - Operator 'equal to'. and - Operator 'and'. or - Operator 'or'. not - Operator 'not'. ( ) - Symbols for grouping operands.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityVersionControlPolicyPrioritiesList(params)
			return handleResponse(resp, err, "failed to get version control policy priorities")
		},
	}
}

func toolEndpointSecurityVersionControlPolicyPriorityUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_version_control_policy_priority_update",
			mcp.WithDescription("Modify version control policy priority. Modifies the settings of the specified endpoint version control policy priority."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the policy priority.")),
			mcp.WithString("agentUpdateStatus", mcp.Description("The status of agent updates."), mcp.Enum("onSchedule", "disable")),
			mcp.WithString("agentUpdatePolicy", mcp.Description("The version the agent or sensor will update to.")),
			mcp.WithString("componentUpdatePolicy", mcp.Description("The version that component will update to'."), mcp.Enum("n", "n-1", "n-2", "n-3", "n-4", "n-5", "n-6", "n-7", "n-8")),
			mcp.WithString("kernelSupportPackageUpdatePolicy", mcp.Description("The version the kernel support package will update to.")),
			mcp.WithString("kernelSupportPackageUpdateStatus", mcp.Description("The status of kernel support package updates."), mcp.Enum("always", "kernelCompatibility", "never")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "agentUpdateStatus", Kind: kindString}, {Name: "agentUpdatePolicy", Kind: kindString}, {Name: "componentUpdatePolicy", Kind: kindString}, {Name: "kernelSupportPackageUpdatePolicy", Kind: kindString}, {Name: "kernelSupportPackageUpdateStatus", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EndpointSecurityVersionControlPolicyPriorityUpdate(id, params)
			return handleResponse(resp, err, "failed to modify version control policy priority")
		},
	}
}
