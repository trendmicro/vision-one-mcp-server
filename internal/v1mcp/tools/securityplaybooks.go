package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlySecurityPlaybooks = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolSecurityPlaybooksPlaybooksList,
	toolSecurityPlaybooksTasksList,
	toolSecurityPlaybooksTaskGet,
	toolSecurityPlaybooksTaskActionsList,
}

var ToolsetsWriteSecurityPlaybooks = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolSecurityPlaybooksPlaybooksRun,
}

func toolSecurityPlaybooksPlaybooksList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_playbooks_playbooks_list",
			mcp.WithDescription("Get playbook list. Retrieves a list of playbooks in the Security Playbooks app."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of items to return in the result set. Default is 10. One of: 10, 20, 50, 100. Default: 10."), mcp.Enum("10", "20", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("Specifies the field by which the results are sorted. Default: updatedDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of playbooks. Supported fields: type - The playbook type - activityAndBehavior, general, systemConfiguration, securityPolicy, threatDetection, vulnerability, xdrDetection; name - The name of playbook - Any value; creator - User that created a playbook - Any value; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; ( ) - Symbols for grouping operands with their correct operator. - -; Additional functions: Function - Description - Notes; contains() - Allows you to search for a specified string in a field - Only applicable to 'name'; Note: Include this parameter in every request that generates paginated output.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityPlaybooksPlaybooksList(params)
			return handleResponse(resp, err, "failed to get playbook list")
		},
	}
}

func toolSecurityPlaybooksPlaybooksRun(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_playbooks_playbooks_run",
			mcp.WithDescription("Run playbooks."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "Playbook ID", "type": "string"}, "parameter": map[string]any{"description": "Additional parameter to trigger a task", "properties": map[string]any{"id": map[string]any{"description": "Alert ID", "type": "string"}, "type": map[string]any{"description": "A string that indicates the type of the target.", "enum": []string{"workbenchAlert"}, "type": "string"}}, "required": []string{"id", "type"}, "type": "object"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityPlaybooksPlaybooksRun(params)
			return handleResponse(resp, err, "failed to run playbooks")
		},
	}
}

func toolSecurityPlaybooksTasksList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_playbooks_tasks_list",
			mcp.WithDescription("Get task list. Get a list of tasks."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the start of the data retrieval range. You can retrieve data for tasks that were created no later than 180 days ago.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval range.")),
			mcp.WithString("dateTimeTarget", mcp.Description("Parameter that allows you to filter the results. Default: createdDateTime."), mcp.Enum("createdDateTime")),
			mcp.WithString("top", mcp.Description("The number of items to return in the result set. Default is 10. One of: 10, 20, 50, 100. Default: 10."), mcp.Enum("10", "20", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("Specifies the field by which the results are sorted. Default: lastActionDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of tasks. Supported fields: id - The ID of task - Any value; playbookName - The name of playbook - Any value; status - The status of task - succeeded, partiallySucceeded, failed, queued, running, pendingApproval; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; ( ) - Symbols for grouping operands with their correct operator. - -; Additional functions: Function - Description - Notes; contains() - Allows you to search for a specified string in a field - Only applicable to 'playbookName' and 'id'; Note: Include this parameter in every request that generates paginated output.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "dateTimeTarget", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityPlaybooksTasksList(params)
			return handleResponse(resp, err, "failed to get task list")
		},
	}
}

func toolSecurityPlaybooksTaskGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_playbooks_task_get",
			mcp.WithDescription("Get task details. Retrieves the execution result of a task."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Task ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityPlaybooksTaskGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get task details")
		},
	}
}

func toolSecurityPlaybooksTaskActionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_playbooks_task_actions_list",
			mcp.WithDescription("Get action list. Get a list of actions for a task."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Task ID.")),
			mcp.WithString("top", mcp.Description("The number of items to return in the result set. Default is 10. One of: 10, 20, 50, 100. Default: 10."), mcp.Enum("10", "20", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("Specifies the field by which the results are sorted. Default: lastActionDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of results. Supported fields: id - The ID of action - Any value; workflowNodeId - The ID of workflow node - Any value; target - The name of target - Any value; status - The status of action - succeeded, partiallySucceeded, failed, running, pendingApproval, approved, rejected, timeout, noAction; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; ( ) - Symbols for grouping operands with their correct operator - -; Additional functions: Function - Description - Notes; contains() - Allows you to search for a specified string in a field - Only applicable to 'target'; Note: Only supported for a task that runs a customizable playbook. Otherwise it will return an 400 error. Include this parameter in every request that generates paginated output.")),
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

			resp, err := client.SecurityPlaybooksTaskActionsList(id, params)
			return handleResponse(resp, err, "failed to get action list")
		},
	}
}
