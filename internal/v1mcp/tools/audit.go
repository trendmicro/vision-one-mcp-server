package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyAudit = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolAuditLogsList,
}

func toolAuditLogsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"audit_logs_list",
			mcp.WithDescription("Get entries from audit logs. Displays log entries that match the specified search criteria in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the start of the data retrieval range. You can retrieve data for response tasks that were created no later than 180 days ago.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval range.")),
			mcp.WithString("dateTimeTarget", mcp.Description("The timestamp to be used for retrieving the target audit log. Default: loggedDateTime."), mcp.Enum("loggedDateTime", "ingestedDateTime")),
			mcp.WithString("orderBy", mcp.Description("Parameter that allows you to sort the retrieved search results in ascending or descending order. If no order is specified, the results are shown in ascending order. Default: loggedDateTime desc.")),
			mcp.WithString("top", mcp.Description("Number of records displayed on a page. One of: 50, 100, 200. Default: 50."), mcp.Enum("50", "100", "200")),
			mcp.WithString("labels", mcp.Description("Parameter that allows you to retrieve the specified elements from the \"details\" field. Possible values: * 'none' - Retrieves no elements. * 'matched' - Retrieves the specified elements in the \"details\" object. * 'all\" - Retrieves all available elements. Default: matched."), mcp.Enum("none", "matched", "all")),
			mcp.WithString("accept", mcp.Description("Default: application/json."), mcp.Enum("application/json", "text/csv")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the Audit Logs list. Supported fields and operators: 'loggedUser '- The account that was used to perform the activity. 'loggedRole' - Role of the account; 'loggedUserId' - The unique identifier of the console account, or the ID of the API key, that was used to perform the activity. 'loggedUserMailAddress' - Email address of the console account that was used to perform the activity. 'accessType' - Source of the activity. Possible values: 'Console', 'API', 'Service Gateway', 'Mobile App', 'Service Provider API', 'Service Provider Console'. 'category' - Category of the activity. Possible values: _See Response Code 200 description_. 'activity' - The activity that was performed. The 'category' needs to be determined to set this parameter. Possible values: _See Response Code 200 description_. 'result' - Outcome of the activity. Possible values: 'Successful', 'Unsuccessful'. 'eq' - Abbreviation of the operator \"equal to\"; 'and' - Operator \"and\"; 'or' - Operator \"or\"; 'not' - Operator \"not\"; '( )' - Symbols for grouping operands with their correct operator. Example: (category eq 'Account Management') and (activity eq 'Add user account')")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "dateTimeTarget", Kind: kindString}, {Name: "orderBy", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "labels", Kind: kindString}},
				Headers: []headerDef{{Arg: "accept", Header: "Accept"}, {Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.AuditLogsList(params)
			return handleResponse(resp, err, "failed to get entries from audit logs")
		},
	}
}
