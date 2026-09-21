package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyEIQS = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolEiqsEndpointsList,
}

func toolEiqsEndpointsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"eiqs_endpoints_list",
			mcp.WithDescription("Get detailed endpoint list. Displays detailed information about your endpoints in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 50, 100, 200. Default: 50."), mcp.Enum("50", "100", "200")),
			mcp.WithString("query", mcp.Required(), mcp.Description("The filter for retrieving a subset of the collected endpoint information data. Supported fields: agentGuid - The ID of the endpoint on the TrendAI Vision One™ platform. - Any value; loginAccount - - - Any value; endpointName - The name of the endpoint. - Any value; macAddress - The MAC address of the endpoint - Any value; ip - The IP address of the endpoint - Any value; protectionManager - The name of your protection manager. - Any value; policyName - The name of a policy from your protection manager. - Any value; componentUpdatePolicy - The update policy for the module/pattern of the agent installed on the endpoint. - n represents the latest version. n - x represents x snapshots prior (Example: n - 2). componentUpdateStatus - The status of the module/pattern updates of the agent installed on the endpoint. - pause, onSchedule, notSupported; componentVersion - outdatedVersion, latestVersion, controlledLatestVersion, unknownVersions; osName - The operating system of the endpoint - Linux, Windows, macOS, macOSX; osVersion - The version of the operating system of the endpoint. - Any value; productCode - The 3-character code that identifies TrendAI™ products. - sao, sds, xes; installedProductCodes - The installed TrendAI™ products on an endpoint - sao, sds, xes; Supported operators: eq - Operator 'equal to'; and - Operator 'and'; or - Operator 'or'; not - Operator 'not'.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.EiqsEndpointsList(params)
			return handleResponse(resp, err, "failed to get detailed endpoint list")
		},
	}
}
