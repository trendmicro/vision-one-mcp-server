package tools

import (
	"context"
	"net/http"

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
}

func toolEndpointSecurityEndpointsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"endpoint_security_endpoints_list",
			mcp.WithDescription("Displays a detailed list of your endpoints"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterEndpoints)),
			mcp.WithString("orderBy",
				mcp.Enum(
					withOrdering(
						asc_desc,
						"agentGuid",
						"eppAgentLastConnectedDateTime",
						"eppAgentLastScannedDateTime",
						"edrSensorLastConnectedDateTime",
					)...,
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
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

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
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
			endpointID, err := requiredValue[string]("endpointID", request.Params.Arguments)
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
				mcp.Enum(
					withOrdering(
						asc_desc,
						"createdDateTime",
						"lastActionDateTime",
					)...,
				),
				mcp.Description("The field by which the results are sorted"),
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
			filter, err := optionalValue[string]("filter", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
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
			taskID, err := requiredValue[string]("taskID", request.Params.Arguments)
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
			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
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
