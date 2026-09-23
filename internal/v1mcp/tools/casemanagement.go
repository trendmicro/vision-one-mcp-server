package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyCaseManagement = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCaseManagementCasesList,
	toolCaseManagementCaseGet,
	toolCaseManagementCaseAttachmentGet,
	toolCaseManagementCaseContentsList,
	toolCaseManagementCaseContentGet,
	toolCaseManagementCaseHighlightedObjectsList,
	toolCaseManagementCaseHighlightedObjectGet,
	toolCaseManagementCaseRiskIndicatorEventsList,
	toolCaseManagementCaseTasksList,
	toolCaseManagementCaseTaskGet,
	toolCaseManagementCaseTaskAttachmentGet,
	toolCaseManagementCaseTaskContentsList,
	toolCaseManagementCaseTaskContentGet,
}

var ToolsetsWriteCaseManagement = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCaseManagementCasesCreate,
	toolCaseManagementCaseUpdate,
	toolCaseManagementCaseAttachmentsCreate,
	toolCaseManagementCaseAttachmentDelete,
	toolCaseManagementCaseBindOatEvent,
	toolCaseManagementCaseContentsCreate,
	toolCaseManagementCaseContentDelete,
	toolCaseManagementCaseContentUpdate,
	toolCaseManagementCaseHighlightedObjectDelete,
	toolCaseManagementCaseRiskIndicatorEventsClose,
	toolCaseManagementCaseTasksCreate,
	toolCaseManagementCaseTaskUpdate,
	toolCaseManagementCaseTaskAttachmentsCreate,
	toolCaseManagementCaseTaskAttachmentDelete,
	toolCaseManagementCaseTaskContentsCreate,
	toolCaseManagementCaseTaskContentDelete,
	toolCaseManagementCaseTaskContentUpdate,
}

func toolCaseManagementCasesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_cases_list",
			mcp.WithDescription("Get all cases. Allows supported ticketing systems to get case info from Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of cases retrieved. If no value is specified, the request retrieves records for the top 50 cases. One of: 10, 50, 100. Default: 50."), mcp.Enum("10", "50", "100")),
			mcp.WithString("startDateTime", mcp.Description("The date and time of the start of data retrieval in ISO 8601 format. Default: 1970-01-01T00:00:00Z.")),
			mcp.WithString("endDateTime", mcp.Description("The date and time of the end of data retrieval in ISO 8601 format retrieval time range. Default: Current date time.")),
			mcp.WithString("dateTimeTarget", mcp.Description("A parameter that allows you to sort by either created date or updated date. Default: createdDateTime."), mcp.Enum("createdDateTime", "updatedDateTime")),
			mcp.WithString("orderBy", mcp.Description("A parameter to be used for sorting records. Records are returned in descending order by default. To return records in ascending order, add a phrase \"asc\" after the parameter name. Sub-sorts can be specified by a comma-separated list of property names. Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of cases. Supported fields: type - Type of case - Workbench, Forensics, ManagedService, TPIServiceNow, RiskEvent, Compliance, General, Others; priority - Priority of the case - P0, P1, P2, P3; status - Status of the case - Open, In-progress,Closed; Supported operators: 'eq' - Abbreviation of the operator 'equal to'; 'and' - Logical operator 'and'; 'or' - Logical operator 'or'; 'not' - Logical operator 'not'; 'contains' - Abbreviation of the operator 'contains'; '( )' - Symbols for grouping operands with their correct operator. Example: priority eq 'P1'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "dateTimeTarget", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCasesList(params)
			return handleResponse(resp, err, "failed to get all cases")
		},
	}
}

func toolCaseManagementCasesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_cases_create",
			mcp.WithDescription("Create a case. Creates a new case in Case Management. The API key and optional \"type\" field determine the case type."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("type", mcp.Description("The type of case to create."), mcp.Enum("General")),
			mcp.WithString("name", mcp.Required(), mcp.Description("The user-defined name of the case.")),
			mcp.WithString("description", mcp.Description("The user-defined description of the case.")),
			mcp.WithArray("owners", mcp.Description("A required list of user account IDs (UUIDs) assigned as owners of the case. Provide at least one valid account ID. Retrieve available IDs using the ‘GET /v3.0/iam/accounts’ API."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("priority", mcp.Required(), mcp.Description("The priority level of the case."), mcp.Enum("P0", "P1", "P2", "P3")),
			mcp.WithString("status", mcp.Required(), mcp.Description("The status of the case."), mcp.Enum("Open", "In-progress")),
			mcp.WithString("severity", mcp.Description("The severity level of the case."), mcp.Enum("Critical", "High", "Medium", "Low", "Information")),
			mcp.WithArray("mailReceivers", mcp.Description("A list of email recipients (account IDs or email addresses)."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithArray("associatedItemIds", mcp.Description("A list of alert IDs and incident IDs associated with the case."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("externalTicketAlias", mcp.Description("The human-readable ID of the case in the 3rd-party ticketing system.")),
			mcp.WithString("externalTicketId", mcp.Description("The unique ID of the case in the 3rd-party ticketing system.")),
			mcp.WithString("findings", mcp.Description("The investigation findings of the case."), mcp.Enum("NoFindings", "Noteworthy", "TruePositive", "FalsePositive", "BenignTruePositive")),
			mcp.WithArray("relatedCaseIds", mcp.Description("A list of the IDs for related cases."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("holder", mcp.Description("The current role of the owner of the case."), mcp.Enum("Customer", "Analyst")),
			mcp.WithString("externalTicketCreatedDateTime", mcp.Description("The date, in ISO 8601 format, the external ticket was created.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "type", Kind: kindString}, {Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "owners", Kind: kindArray}, {Name: "priority", Kind: kindString, Required: true}, {Name: "status", Kind: kindString, Required: true}, {Name: "severity", Kind: kindString}, {Name: "mailReceivers", Kind: kindArray}, {Name: "associatedItemIds", Kind: kindArray}, {Name: "externalTicketAlias", Kind: kindString}, {Name: "externalTicketId", Kind: kindString}, {Name: "findings", Kind: kindString}, {Name: "relatedCaseIds", Kind: kindArray}, {Name: "holder", Kind: kindString}, {Name: "externalTicketCreatedDateTime", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCasesCreate(params)
			return handleResponse(resp, err, "failed to create a case")
		},
	}
}

func toolCaseManagementCaseGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_get",
			mcp.WithDescription("Get case details. Allows supported ticketing systems to retrieve information about a specific case from Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Case ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get case details")
		},
	}
}

func toolCaseManagementCaseUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_update",
			mcp.WithDescription("Update a case. Allows supported ticketing systems to update the specified case."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("ifMatch", mcp.Description("Provide the ETag of the resource to be updated. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("description", mcp.Description("The user-defined description of the case.")),
			mcp.WithString("name", mcp.Description("The name of the case.")),
			mcp.WithString("priority", mcp.Description("The priority level of the case."), mcp.Enum("P0", "P1", "P2", "P3")),
			mcp.WithString("status", mcp.Description("The status of the case."), mcp.Enum("Open", "In-progress", "Closed")),
			mcp.WithString("severity", mcp.Description("The severity level of the case."), mcp.Enum("Critical", "High", "Medium", "Low", "Information")),
			mcp.WithString("externalTicketUpdatedDateTime", mcp.Description("The date, in ISO 8601 format, the external ticket was last updated.")),
			mcp.WithString("findings", mcp.Description("The investigation findings of the case."), mcp.Enum("NoFindings", "Noteworthy", "TruePositive", "FalsePositive", "BenignTruePositive")),
			mcp.WithArray("associatedItemIds", mcp.Description("A list of alert IDs and incident IDs associated with the case."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("holder", mcp.Description("The current role of the owner of the case."), mcp.Enum("Customer", "Analyst")),
			mcp.WithString("externalTicketId", mcp.Description("The ID of the case in the external ticketing system.")),
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
				BodyFields: []paramDef{{Name: "description", Kind: kindString}, {Name: "name", Kind: kindString}, {Name: "priority", Kind: kindString}, {Name: "status", Kind: kindString}, {Name: "severity", Kind: kindString}, {Name: "externalTicketUpdatedDateTime", Kind: kindString}, {Name: "findings", Kind: kindString}, {Name: "associatedItemIds", Kind: kindArray}, {Name: "holder", Kind: kindString}, {Name: "externalTicketId", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseUpdate(id, params)
			return handleResponse(resp, err, "failed to update a case")
		},
	}
}

func toolCaseManagementCaseAttachmentsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_attachments_create",
			mcp.WithDescription("Add an attachment. Allows supported ticketing systems to add attachments to a case in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("file", mcp.Description("The attachment in binary format Use this field when calling from Jira, or directly via API key. Mutually exclusive with externalAttachmentId. Path to a local file to upload.")),
			mcp.WithString("externalAttachmentId", mcp.Description("The attachment ID from ServiceNow or ServiceDesk Plus (SDP) The service fetches the file from the third-party system on behalf of the caller. Only supported for ServiceNow and SDP integrations.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "externalAttachmentId", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseAttachmentsCreate(id, params)
			return handleResponse(resp, err, "failed to add an attachment")
		},
	}
}

func toolCaseManagementCaseAttachmentDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_attachment_delete",
			mcp.WithDescription("Delete an attachment. Allows supported ticketing systems to delete the specified attachment from a case in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("attachmentId", mcp.Required(), mcp.Description("The attachment ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			attachmentId, err := pathValue("attachmentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseAttachmentDelete(id, attachmentId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete an attachment")
		},
	}
}

func toolCaseManagementCaseAttachmentGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_attachment_get",
			mcp.WithDescription("Download attachment by attachment ID. Allows supported ticketing systems to retrieve the specified attachment from a case in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("attachmentId", mcp.Required(), mcp.Description("The attachment ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			attachmentId, err := pathValue("attachmentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseAttachmentGet(id, attachmentId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to download attachment by attachment ID")
		},
	}
}

func toolCaseManagementCaseBindOatEvent(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_bind_oat_event",
			mcp.WithDescription("Add an OAT event to case. Adds an OAT (Observed Attack Techniques) event to an existing case. When an OAT event is added, the system to the case adds all highlighted objects associated with the event."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("uuid", mcp.Required(), mcp.Description("Unique identifier of the OAT event to add to the case Get this UUID from the GET /oat/detections API response.")),
			mcp.WithString("detectedDateTime", mcp.Description("The date and time, in ISO 8601 format, when the OAT event was detected Get the detectedDateTime using the GET /oat/detections API response. If not provided, the system will search within the last 30 days.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "uuid", Kind: kindString, Required: true}, {Name: "detectedDateTime", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseBindOatEvent(id, params)
			return handleResponse(resp, err, "failed to add an OAT event to case")
		},
	}
}

func toolCaseManagementCaseContentsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_contents_list",
			mcp.WithDescription("Get case contents. Allows supported ticketing systems to retrieve the contents of the specified case in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("top", mcp.Description("The number of case notes retrieved per page. If no value is specified, the request retrieves a maximum of the top 50 notes. One of: 10, 50, 100. Default: 50."), mcp.Enum("10", "50", "100")),
			mcp.WithString("startDateTime", mcp.Description("The date and time of the start of data retrieval in ISO 8601 format. Default: 1970-01-01T00:00:00Z.")),
			mcp.WithString("endDateTime", mcp.Description("The date and time of the end of data retrieval in ISO 8601 format. Default: Current date time.")),
			mcp.WithString("dateTimeTarget", mcp.Description("A parameter that allows you to sort by either created date or updated date. Default: createdDateTime."), mcp.Enum("createdDateTime", "updatedDateTime")),
			mcp.WithString("orderBy", mcp.Description("The fields by which the results are sorted. You can indicate multiple fields separated by commas (,). Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of case contents; Supported fields: creatorId - Creator ID of the case content - Any string value; Supported operators: 'eq' - Abbreviation of the operator 'equal to'; 'and' - Logical operator 'and'; 'or' - Logical operator 'or'; 'not' - Logical operator 'not'; '( )' - Symbols for grouping operands with their correct operator. Example: creatorId eq '320e6df0-a527-4c9e-a7ca-214502c9727c'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "dateTimeTarget", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseContentsList(id, params)
			return handleResponse(resp, err, "failed to get case contents")
		},
	}
}

func toolCaseManagementCaseContentsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_contents_create",
			mcp.WithDescription("Add content to a case. Allows supported ticketing systems to add content to a case in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithArray("attachmentIds", mcp.Description("A list of attachment IDs from case notes. Case Management supports up to 5 attachments per note. Each attachment should only be associated with one note."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("comment", mcp.Required(), mcp.Description("A comment added to the note.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "attachmentIds", Kind: kindArray}, {Name: "comment", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseContentsCreate(id, params)
			return handleResponse(resp, err, "failed to add content to a case")
		},
	}
}

func toolCaseManagementCaseContentDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_content_delete",
			mcp.WithDescription("Delete case content using content ID. Allows supported ticketing systems to delete the specified item from a case in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("contentId", mcp.Required(), mcp.Description("The content ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			contentId, err := pathValue("contentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseContentDelete(id, contentId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete case content using content ID")
		},
	}
}

func toolCaseManagementCaseContentGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_content_get",
			mcp.WithDescription("Get case item. Allows supported ticketing systems to retrieved the specified case item in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("contentId", mcp.Required(), mcp.Description("The content ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			contentId, err := pathValue("contentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseContentGet(id, contentId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get case item")
		},
	}
}

func toolCaseManagementCaseContentUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_content_update",
			mcp.WithDescription("Update the case content. Allows supported ticketing systems to update the specified case item in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("contentId", mcp.Required(), mcp.Description("The content ID.")),
			mcp.WithString("ifMatch", mcp.Description("Provide the ETag of the resource to be updated. Use the ETag returned when the resource was retrieved.")),
			mcp.WithArray("attachmentIds", mcp.Description("A list of updated attachment IDs from case notes. Case Management supports up to five attachments per note. Each attachment should be associated with only one note."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("comment", mcp.Description("The new comment added to the case.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			contentId, err := pathValue("contentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "attachmentIds", Kind: kindArray}, {Name: "comment", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseContentUpdate(id, contentId, params)
			return handleResponse(resp, err, "failed to update the case content")
		},
	}
}

func toolCaseManagementCaseHighlightedObjectsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_highlighted_objects_list",
			mcp.WithDescription("Get highlighted objects. Retrieves highlighted objects associated with a case."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The unique case identifier.")),
			mcp.WithString("top", mcp.Description("The number of highlighted objects retrieved per page. If no value is specified, the request retrieves a maximum of the top 50 highlighted objects. One of: 10, 50, 100. Default: 50."), mcp.Enum("10", "50", "100")),
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

			resp, err := client.CaseManagementCaseHighlightedObjectsList(id, params)
			return handleResponse(resp, err, "failed to get highlighted objects")
		},
	}
}

func toolCaseManagementCaseHighlightedObjectDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_highlighted_object_delete",
			mcp.WithDescription("Remove a highlighted object. Removes a highlighted object from a case."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("highlightedObjectId", mcp.Required(), mcp.Description("The ID of the highlighted object.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			highlightedObjectId, err := pathValue("highlightedObjectId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseHighlightedObjectDelete(id, highlightedObjectId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to remove a highlighted object")
		},
	}
}

func toolCaseManagementCaseHighlightedObjectGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_highlighted_object_get",
			mcp.WithDescription("Get a specific highlighted object. Retrieves a specific highlighted object from a case by its ID."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("highlightedObjectId", mcp.Required(), mcp.Description("The highlighted object ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			highlightedObjectId, err := pathValue("highlightedObjectId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseHighlightedObjectGet(id, highlightedObjectId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get a specific highlighted object")
		},
	}
}

func toolCaseManagementCaseRiskIndicatorEventsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_risk_indicator_events_list",
			mcp.WithDescription("Get risk indicators in specified case. Displays the risk indicators associated to a case ID."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the case on the TrendAI Vision One™ platform.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: detectedDateTime desc.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the cloud asset information list. Supported fields: id - The ID of a risk event - Any value; riskLevel - The risk level of the risk event - high, medium, low riskFactor - The risk factor of the risk event - Threat detection, Security configuration, System configuration, Vulnerability detection, Anomaly detection, Account compromise, Cloud app activity, XDR detection; status - The status of the risk event - new, inProgress, remediated, dismissed, accepted, mitigated; detectedDateTime - The time the event was detected - Any value; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; () - Symbols for grouping operands with their correct operator - -; gt - Operator 'greater than' - Only applicable to detectedDateTime; ge - Operator 'greater than or equal' - Only applicable to detectedDateTime; le - Operator 'less than or equal' - Only applicable to detectedDateTime; lt - Operator 'less than' - Only applicable to detectedDateTime. Example: riskLevel eq 'high'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseRiskIndicatorEventsList(id, params)
			return handleResponse(resp, err, "failed to get risk indicators in specified case")
		},
	}
}

func toolCaseManagementCaseRiskIndicatorEventsClose(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_risk_indicator_events_close",
			mcp.WithDescription("Close risk events related to case. Sets the status of all the risk events associated to a case ID to any of the following statuses:."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the case.")),
			mcp.WithString("status", mcp.Required(), mcp.Description("The new status of the case."), mcp.Enum("remediated", "dismissed", "accepted")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "status", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseRiskIndicatorEventsClose(id, params)
			return handleResponse(resp, err, "failed to close risk events related to case")
		},
	}
}

func toolCaseManagementCaseTasksList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_tasks_list",
			mcp.WithDescription("Get a case's tasks. Retrieves the tasks of the specified case."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("top", mcp.Description("The number of tasks retrieved per page. If no value is specified, the request retrieves a maximum of the top 50 tasks. One of: 10, 50, 100. Default: 50."), mcp.Enum("10", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The field and order for sorting the retrieved tasks. Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of tasks; Supported fields: status - Status of the task - Todo, In-progress, Closed, Rejected; Supported operators: 'eq' - Abbreviation of the operator 'equal to'; 'and' - Logical operator 'and'; 'or' - Logical operator 'or'; 'not' - Logical operator 'not'; '( )' - Symbols for grouping operands with their correct operator. Example: status eq 'Todo'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTasksList(id, params)
			return handleResponse(resp, err, "failed to get a case's tasks")
		},
	}
}

func toolCaseManagementCaseTasksCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_tasks_create",
			mcp.WithDescription("Create a task. Creates a new task on a case in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("description", mcp.Description("A free-text description of the task. If not specified, the task has no description.")),
			mcp.WithString("dueDateTime", mcp.Description("The task's due date and time in ISO 8601 format. If not specified, the task has no due date.")),
			mcp.WithString("name", mcp.Required(), mcp.Description("The display name of the task.")),
			mcp.WithArray("ownerIds", mcp.Description("The IDs of the accounts assigned to the task. If not specified, the task has no assigned owners."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("status", mcp.Required(), mcp.Description("The initial status of the task."), mcp.Enum("Todo", "In-progress", "Closed", "Rejected")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "description", Kind: kindString}, {Name: "dueDateTime", Kind: kindString}, {Name: "name", Kind: kindString, Required: true}, {Name: "ownerIds", Kind: kindArray}, {Name: "status", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTasksCreate(id, params)
			return handleResponse(resp, err, "failed to create a task")
		},
	}
}

func toolCaseManagementCaseTaskGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_get",
			mcp.WithDescription("Get task details. Retrieves the details of the specified task."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskGet(id, taskId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get task details")
		},
	}
}

func toolCaseManagementCaseTaskUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_update",
			mcp.WithDescription("Update a task. Updates the specified task in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithString("ifMatch", mcp.Description("If provided, the value must match the resource’s current ETag. Otherwise, the request returns HTTP 412. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("description", mcp.Description("A free-text description of the task. If not specified, the current description does not change. To clear the description, set to null.")),
			mcp.WithString("dueDateTime", mcp.Description("The due date and time, in ISO 8601 format, of the task. If not specified, the current due date does not change. To clear the due date, set to null.")),
			mcp.WithString("name", mcp.Description("The display name of the task. If not specified, the current name does not change. Setting to null is not accepted.")),
			mcp.WithArray("ownerIds", mcp.Description("The IDs of the accounts responsible for the task. If not specified, the current owners list does not change. To clear all owners, set to an empty array. Setting to null is not accepted."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("status", mcp.Description("The status of the task. If not specified, the current status does not change. Setting to null is not accepted."), mcp.Enum("Todo", "In-progress", "Closed", "Rejected")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "description", Kind: kindString}, {Name: "dueDateTime", Kind: kindString}, {Name: "name", Kind: kindString}, {Name: "ownerIds", Kind: kindArray}, {Name: "status", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskUpdate(id, taskId, params)
			return handleResponse(resp, err, "failed to update a task")
		},
	}
}

func toolCaseManagementCaseTaskAttachmentsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_attachments_create",
			mcp.WithDescription("Upload an attachment to a task. Uploads an attachment to a task in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithString("file", mcp.Required(), mcp.Description("The attachment in binary format. Use this field when calling from Jira, or directly via API key. Mutually exclusive with externalAttachmentId. Path to a local file to upload.")),
			mcp.WithString("externalAttachmentId", mcp.Description("The attachment ID from ServiceNow or ServiceDesk Plus (SDP). The service fetches the file from the third-party system on behalf of the caller. Only supported for ServiceNow and SDP integrations.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "externalAttachmentId", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskAttachmentsCreate(id, taskId, params)
			return handleResponse(resp, err, "failed to upload an attachment to a task")
		},
	}
}

func toolCaseManagementCaseTaskAttachmentDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_attachment_delete",
			mcp.WithDescription("Delete a task attachment. Removes the attachment from a task in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithString("attachmentId", mcp.Required(), mcp.Description("The attachment ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			attachmentId, err := pathValue("attachmentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskAttachmentDelete(id, taskId, attachmentId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete a task attachment")
		},
	}
}

func toolCaseManagementCaseTaskAttachmentGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_attachment_get",
			mcp.WithDescription("Download a task attachment. Downloads the specified attachment from a task in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithString("attachmentId", mcp.Required(), mcp.Description("The attachment ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			attachmentId, err := pathValue("attachmentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskAttachmentGet(id, taskId, attachmentId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to download a task attachment")
		},
	}
}

func toolCaseManagementCaseTaskContentsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_contents_list",
			mcp.WithDescription("Get task contents. Retrieves the contents of the specified task."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithString("top", mcp.Description("The number of content entries retrieved per page. If no value is specified, the request retrieves a maximum of the top 50 entries. One of: 10, 50, 100. Default: 50."), mcp.Enum("10", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The field and order for sorting the retrieved content entries. Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of task content entries; Supported fields: creatorId - Creator ID of the content entry - Any string value; Supported operators: 'eq' - Abbreviation of the operator 'equal to'; 'and' - Logical operator 'and'; 'or' - Logical operator 'or'; 'not' - Logical operator 'not'; '( )' - Symbols for grouping operands with their correct operator. Example: creatorId eq '320e6df0-a527-4c9e-a7ca-214502c9727c'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskContentsList(id, taskId, params)
			return handleResponse(resp, err, "failed to get task contents")
		},
	}
}

func toolCaseManagementCaseTaskContentsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_contents_create",
			mcp.WithDescription("Add content to a task. Adds a content entry (text comment and/or attachments) to a task in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithArray("attachmentIds", mcp.Description("A list of attachment IDs to bind to this content entry. Default: []."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("comment", mcp.Description("The text comment to add to the content entry. Cannot be an empty string.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "attachmentIds", Kind: kindArray}, {Name: "comment", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskContentsCreate(id, taskId, params)
			return handleResponse(resp, err, "failed to add content to a task")
		},
	}
}

func toolCaseManagementCaseTaskContentDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_content_delete",
			mcp.WithDescription("Remove a task's content entry. Removes a content entry from a task in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithString("contentId", mcp.Required(), mcp.Description("The content ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			contentId, err := pathValue("contentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskContentDelete(id, taskId, contentId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to remove a task's content entry")
		},
	}
}

func toolCaseManagementCaseTaskContentGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_content_get",
			mcp.WithDescription("Get a task's content entry. Retrieves the specified content entry of a task."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithString("contentId", mcp.Required(), mcp.Description("The content ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			contentId, err := pathValue("contentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskContentGet(id, taskId, contentId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get a task's content entry")
		},
	}
}

func toolCaseManagementCaseTaskContentUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"case_management_case_task_content_update",
			mcp.WithDescription("Update a task content entry. Updates a content entry of a task in Case Management."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The case ID.")),
			mcp.WithString("taskId", mcp.Required(), mcp.Description("The task ID.")),
			mcp.WithString("contentId", mcp.Required(), mcp.Description("The content ID.")),
			mcp.WithString("ifMatch", mcp.Description("If provided, must match the current ETag of the resource. Returns an HTTP 412 error otherwise. Use the ETag returned when the resource was retrieved.")),
			mcp.WithArray("attachmentIds", mcp.Description("The complete replacement list of attachment IDs for this content entry. The service reconciles the new list against the current state: IDs that are newly uploaded (pending) are bound to this entry; IDs already bound to this entry are retained;."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("comment", mcp.Description("The new text comment of the content entry. If not specified, the current comment does not change. Set to null to clear the comment. Returns an HTTP 400 error if both the attachmentIds and comment are empty.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			taskId, err := pathValue("taskId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			contentId, err := pathValue("contentId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "attachmentIds", Kind: kindArray}, {Name: "comment", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CaseManagementCaseTaskContentUpdate(id, taskId, contentId, params)
			return handleResponse(resp, err, "failed to update a task content entry")
		},
	}
}
