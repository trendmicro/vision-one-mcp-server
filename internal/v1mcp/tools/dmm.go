package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyDMM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolDmmCustomFiltersList,
	toolDmmCustomModelsList,
	toolDmmExceptionsList,
	toolDmmExceptionGet,
	toolDmmModelsList,
}

var ToolsetsWriteDMM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolDmmExceptionsCreate,
	toolDmmExceptionsDelete,
	toolDmmExceptionUpdate,
	toolDmmModelUpdate,
}

func toolDmmCustomFiltersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_custom_filters_list",
			mcp.WithDescription("Get all custom filters. Displays all your custom filters in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.DmmCustomFiltersList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get all custom filters")
		},
	}
}

func toolDmmCustomModelsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_custom_models_list",
			mcp.WithDescription("Get custom detection models. Displays your custom detection models in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.DmmCustomModelsList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get custom detection models")
		},
	}
}

func toolDmmExceptionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_exceptions_list",
			mcp.WithDescription("Get all custom exceptions. Retrieves all your custom exceptions."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.DmmExceptionsList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get all custom exceptions")
		},
	}
}

func toolDmmExceptionsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_exceptions_create",
			mcp.WithDescription("Create an exception. Creates a custom exception."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Description("Creates a custom exception."), mcp.Items(map[string]any{"properties": map[string]any{"criteria": map[string]any{"description": "The objects or events excluded from detections.", "items": map[string]any{"type": "object"}, "type": "array"}, "description": map[string]any{"description": "Description of the custom exception", "type": "string"}, "event": map[string]any{"properties": map[string]any{"category": map[string]any{"description": "Event type", "enum": []string{"ENDPOINT_ACTIVITY"}, "type": "string"}, "id": map[string]any{"description": "The unique identifier of an event type", "enum": []string{"TELEMETRY_PROCESS", "TELEMETRY_FILE", "TELEMETRY_CONNECTION", "TELEMETRY_DNS", "TELEMETRY_REGISTRY", "TELEMETRY_ACCOUNT", "TELEMETRY_INTERNET", "TELEMETRY_MODIFIED_PROCESS", "TELEMETRY_WINDOWS_HOOK", "TELEMETRY_WINDOWS_EVENT", "TELEMETRY_AMSI", "TELEMETRY_WMI", "TELEMETRY_MEMORY", "TELEMETRY_BM", "TELEMETRY_APP", "TELEMETRY_SYSTEM_EVENT"}, "type": "string"}, "subId": map[string]any{"description": "The secondary event identifier of an event type", "type": "array"}}, "type": "object"}, "name": map[string]any{"description": "Name of the custom exception", "type": "string"}, "targetEntities": map[string]any{"description": "The locations of the objects or events excluded from detections.", "properties": map[string]any{"dst": map[string]any{"description": "Destination IP address", "type": "array"}, "endpointGUID": map[string]any{"description": "GUID of the agent which reported the detection", "type": "array"}, "endpointHostName": map[string]any{"description": "Device name of the endpoint", "type": "array"}, "endpointIp": map[string]any{"description": "IP address of the endpoint", "type": "array"}, "mailFromAddress": map[string]any{"description": "Sender address", "type": "array"}, "mailbox": map[string]any{"description": "The mailbox where the email is located", "type": "array"}, "samUser": map[string]any{"description": "Security Account Manager user name", "type": "array"}, "src": map[string]any{"description": "Source IP address", "type": "array"}, "suser": map[string]any{"description": "Email sender", "type": "array"}}, "type": "object"}}, "required": []string{"criteria", "name"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DmmExceptionsCreate(params)
			return handleResponse(resp, err, "failed to create an exception")
		},
	}
}

func toolDmmExceptionsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_exceptions_delete",
			mcp.WithDescription("Delete a custom exception. Deletes the specified custom exception."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The unique identifier of the exception", "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DmmExceptionsDelete(params)
			return handleResponse(resp, err, "failed to delete a custom exception")
		},
	}
}

func toolDmmExceptionGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_exception_get",
			mcp.WithDescription("Get a custom exception. Retrieves the specified custom exception."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of the exception.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DmmExceptionGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get a custom exception")
		},
	}
}

func toolDmmExceptionUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_exception_update",
			mcp.WithDescription("Update a custom exception. Updates the specified custom exception."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique identifier of the exception.")),
			mcp.WithString("name", mcp.Description("Name of the custom exception.")),
			mcp.WithString("description", mcp.Description("Description of the custom exception.")),
			mcp.WithObject("targetEntities", mcp.Description("The locations of the objects or events excluded from detections. Default: {}."), mcp.Properties(map[string]any{"dst": map[string]any{"description": "Destination IP address", "items": map[string]any{"type": "string"}, "type": "array"}, "endpointGUID": map[string]any{"description": "GUID of the agent which reported the detection", "items": map[string]any{"type": "string"}, "type": "array"}, "endpointHostName": map[string]any{"description": "Device name of the endpoint", "items": map[string]any{"type": "string"}, "type": "array"}, "endpointIp": map[string]any{"description": "IP address of the endpoint", "items": map[string]any{"type": "string"}, "type": "array"}, "mailFromAddress": map[string]any{"description": "Sender address", "items": map[string]any{"format": "email", "type": "string"}, "type": "array"}, "mailbox": map[string]any{"description": "The mailbox where the email is located", "items": map[string]any{"format": "email", "type": "string"}, "type": "array"}, "samUser": map[string]any{"description": "Security Account Manager user name", "items": map[string]any{"type": "string"}, "type": "array"}, "src": map[string]any{"description": "Source IP address", "items": map[string]any{"type": "string"}, "type": "array"}, "suser": map[string]any{"description": "Email sender", "items": map[string]any{"format": "email", "type": "string"}, "type": "array"}})),
			mcp.WithObject("event", mcp.Properties(map[string]any{"category": map[string]any{"description": "Event type", "enum": []string{"ENDPOINT_ACTIVITY"}, "type": "string"}, "id": map[string]any{"description": "The unique identifier of an event type", "enum": []string{"TELEMETRY_PROCESS", "TELEMETRY_FILE", "TELEMETRY_CONNECTION", "TELEMETRY_DNS", "TELEMETRY_REGISTRY", "TELEMETRY_ACCOUNT", "TELEMETRY_INTERNET", "TELEMETRY_MODIFIED_PROCESS", "TELEMETRY_WINDOWS_HOOK", "TELEMETRY_WINDOWS_EVENT", "TELEMETRY_AMSI", "TELEMETRY_WMI", "TELEMETRY_MEMORY", "TELEMETRY_BM", "TELEMETRY_APP", "TELEMETRY_SYSTEM_EVENT"}, "type": "string"}, "subId": map[string]any{"description": "The secondary event identifier of an event type", "items": map[string]any{"type": "string"}, "type": "array"}})),
			mcp.WithArray("criteria", mcp.Description("The objects or events excluded from detections."), mcp.Items(map[string]any{"properties": map[string]any{"fieldName": map[string]any{"description": "The field match criteria uses", "enum": []string{"parentFileHashMd5", "processFileHashMd5", "objectFileHashMd5", "srcFileHashMd5", "processPayloadFileHashMd5", "parentPayloadFileHashMd5", "attachmentFileHashMd5"}, "type": "string"}, "fieldType": map[string]any{"description": "The type of match criteria", "enum": []string{"file_md5"}, "type": "string"}, "fieldValues": map[string]any{"items": map[string]any{"description": "The value of the field.", "type": "string"}, "type": "array"}, "matchType": map[string]any{"description": "The match type of the criteria", "enum": []string{"REGEX", "EXACT"}, "type": "string"}}, "required": []string{"fieldName", "fieldType", "fieldValues"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}, {Name: "targetEntities", Kind: kindObject}, {Name: "event", Kind: kindObject}, {Name: "criteria", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DmmExceptionUpdate(id, params)
			return handleResponse(resp, err, "failed to update a custom exception")
		},
	}
}

func toolDmmModelsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_models_list",
			mcp.WithDescription("List detection models. Displays all the detection models in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range, in ISO 8601 format.")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range, in ISO 8601 format.")),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: lastUpdatedDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the detection model list. Supported fields: riskLevel - The severity assigned to the model - Critical, High, Medium, Low; enabled - The status of the model - true,false; modelType - The type of the model - sae, ti; requiredProducts - The products that use the detection model - Any value; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; () - Symbols for grouping operands with their correct operator - -; hassubset - Filter for arrays - hassubset(requiredProducts,['xes']).")),
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

			resp, err := client.DmmModelsList(params)
			return handleResponse(resp, err, "failed to list detection models")
		},
	}
}

func toolDmmModelUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"dmm_model_update",
			mcp.WithDescription("Enable or disable a detection model. Changes the status of the specified detection model."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required()),
			mcp.WithBoolean("enabled", mcp.Required(), mcp.Description("The status of a detection model.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "enabled", Kind: kindBoolean, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.DmmModelUpdate(id, params)
			return handleResponse(resp, err, "failed to enable or disable a detection model")
		},
	}
}
