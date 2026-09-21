package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlySecurityAwareness = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolSecurityAwarenessPhishingSimulationsList,
	toolSecurityAwarenessPhishingRateTrendsList,
	toolSecurityAwarenessPhishingSimulationRoundsList,
	toolSecurityAwarenessPhishingSimulationRoundGet,
	toolSecurityAwarenessPhishingSimulationRoundRecipientsList,
	toolSecurityAwarenessTrainingCampaignsList,
	toolSecurityAwarenessTrainingCampaignGet,
	toolSecurityAwarenessTrainingCampaignRecipientsList,
}

func toolSecurityAwarenessPhishingSimulationsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_awareness_phishing_simulations_list",
			mcp.WithDescription("List phishing simulations. Returns a paginated list of phishing simulations for the company. Each item reflects the state of the simulation's latest round. Results are ordered by creation time descending by default."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description("Maximum number of items to return per page (50–200, default 50). Default: 50.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of phishing simulations. Supported fields: latestRoundStatus - Lifecycle status of the simulation's latest round - initializing, ready, scheduled, preparing, prepared, inProgress, cancelled, completed; createdDateTime - The date and time when the simulation was created in ISO 8601 format - ISO 8601 datetime; latestRoundStartDateTime - Start date and time of the latest round in ISO 8601 format. Scheduled or actual depending on the round's latestRoundStatus. - ISO 8601 datetime; Supported operators: eq - Equal to - All; ge - Greater than or equal - Datetime fields only; le - Less than or equal - Datetime fields only; gt - Greater than - Datetime fields only; lt - Less than - Datetime fields only; and - Logical AND - Combinator; or - Logical OR - Combinator; not - Logical NOT - Combinator; () - Grouping - Combinator; Simulations without a latestRound that have no; rounds scheduled yet do not have a latestRoundStartDateTime; value and are excluded from results when filtering by that; field, regardless of whether not is used. Example: latestRoundStatus eq 'completed' and createdDateTime ge '2024-01-01T00:00:00Z'")),
			mcp.WithString("orderBy", mcp.Description("Sort field and direction. Format: fieldName asc|desc. Supported fields: createdDateTime, latestRoundStartDateTime, name. When omitted, results are sorted by createdDateTime desc. Default: createdDateTime desc.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindNumber}, {Name: "filter", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityAwarenessPhishingSimulationsList(params)
			return handleResponse(resp, err, "failed to list phishing simulations")
		},
	}
}

func toolSecurityAwarenessPhishingRateTrendsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_awareness_phishing_rate_trends_list",
			mcp.WithDescription("Get phishing rate trends. Returns trending data showing phishing rates over time. Used for tracking awareness improvement across simulations."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("period", mcp.Description("Time granularity for the trending data. Defaults to monthly when omitted. The returned data points are aggregated at this granularity level. Default: monthly."), mcp.Enum("monthly", "quarterly", "yearly")),
			mcp.WithString("startDateTime", mcp.Description("Start of the reporting period in ISO 8601 format. When omitted, defaults to 1 year before endDateTime. Maximum date and time range is 3 years.")),
			mcp.WithString("endDateTime", mcp.Description("End of the reporting period in ISO 8601 format. When omitted, defaults to the current time. Maximum date and time range is 3 years.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "period", Kind: kindString}, {Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityAwarenessPhishingRateTrendsList(params)
			return handleResponse(resp, err, "failed to get phishing rate trends")
		},
	}
}

func toolSecurityAwarenessPhishingSimulationRoundsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_awareness_phishing_simulation_rounds_list",
			mcp.WithDescription("List simulation rounds. Returns a paginated list of all rounds for a specific phishing simulation, allowing discovery of historical round IDs. Results are ordered by round number descending (latest first)."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique identifier of the resource specified in the path.")),
			mcp.WithNumber("top", mcp.Description("Maximum number of items to return per page (50–200, default 50). Default: 50.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindNumber}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityAwarenessPhishingSimulationRoundsList(id, params)
			return handleResponse(resp, err, "failed to list simulation rounds")
		},
	}
}

func toolSecurityAwarenessPhishingSimulationRoundGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_awareness_phishing_simulation_round_get",
			mcp.WithDescription("Get simulation round details. Returns detailed information about a specific simulation round, including schedule, template configuration, delivery settings, and a recipient summary with aggregate counts."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique identifier of the resource specified in the path.")),
			mcp.WithString("roundId", mcp.Required(), mcp.Description("Unique identifier of the simulation round.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			roundId, err := pathValue("roundId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityAwarenessPhishingSimulationRoundGet(id, roundId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get simulation round details")
		},
	}
}

func toolSecurityAwarenessPhishingSimulationRoundRecipientsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_awareness_phishing_simulation_round_recipients_list",
			mcp.WithDescription("List simulation round recipients. Returns a paginated list of per-participant details for a simulation round, including delivery status and response actions."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique identifier of the resource specified in the path.")),
			mcp.WithString("roundId", mcp.Required(), mcp.Description("Unique identifier of the simulation round.")),
			mcp.WithNumber("top", mcp.Description("Maximum number of items to return per page (50–200, default 50). Default: 50.")),
			mcp.WithString("orderBy", mcp.Description("Sort field and direction. Format: fieldName asc|desc. Supported fields: name, email, deliveredDateTime. When omitted, results are sorted by deliveredDateTime desc. Default: deliveredDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of recipients. Supported fields: department - Recipient's department name - Any value; location - Recipient's location or office - Any value; To discover valid department and location values for a; specific simulation round, list the recipients using this; endpoint and read the department / location fields from; each item. Server-side enumeration endpoints may be added in; a later release. Supported operators: eq - Equal to; and - Logical AND; or - Logical OR; not - Logical NOT; () - Grouping; Recipients that do not have the optional department or; location property (e.g., when department/location information; is not configured for a user) are excluded from results; when filtering by that field, regardless of whether not; is used. When the header is omitted, all recipients are returned; without filtering. Example: department eq 'Engineering'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			roundId, err := pathValue("roundId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindNumber}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityAwarenessPhishingSimulationRoundRecipientsList(id, roundId, params)
			return handleResponse(resp, err, "failed to list simulation round recipients")
		},
	}
}

func toolSecurityAwarenessTrainingCampaignsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_awareness_training_campaigns_list",
			mcp.WithDescription("List training campaigns. Returns a paginated list of training campaigns. Results are ordered by start time descending by default."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description("Maximum number of items to return per page (50–200, default 50). Default: 50.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of training campaigns. Supported fields: status - Lifecycle status of the campaign - initializing, ready, scheduled, preparing, prepared, inProgress, cancelled, completed; Supported operators: eq - Equal to; and - Logical AND; or - Logical OR; not - Logical NOT; () - Grouping. Example: status eq 'inProgress'")),
			mcp.WithString("orderBy", mcp.Description("Sort field and direction. Format: fieldName asc|desc. Supported fields: startDateTime, name. When omitted, results are sorted by startDateTime desc. Default: startDateTime desc.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindNumber}, {Name: "filter", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityAwarenessTrainingCampaignsList(params)
			return handleResponse(resp, err, "failed to list training campaigns")
		},
	}
}

func toolSecurityAwarenessTrainingCampaignGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_awareness_training_campaign_get",
			mcp.WithDescription("Get training campaign details. Returns detailed information about a specific training campaign, including program, schedule, reminder configuration, and a recipient summary with aggregate completion counts."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique identifier of the resource specified in the path.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityAwarenessTrainingCampaignGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get training campaign details")
		},
	}
}

func toolSecurityAwarenessTrainingCampaignRecipientsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"security_awareness_training_campaign_recipients_list",
			mcp.WithDescription("List training campaign recipients. Returns a paginated list of per-participant training progress for a specific campaign."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique identifier of the resource specified in the path.")),
			mcp.WithNumber("top", mcp.Description("Maximum number of items to return per page (50–200, default 50). Default: 50.")),
			mcp.WithString("orderBy", mcp.Description("Sort field and direction. Format: fieldName asc|desc. Supported fields: name, email, status. When omitted, results are sorted by name asc. Default: name asc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of recipients. Supported fields: status - Training completion status - notStarted, inProgress, completed; department - Recipient's department name - Any value; location - Recipient's location or office - Any value; To discover valid department and location values for a; specific training campaign, list the recipients using this; endpoint and read the department / location fields from; each item. Server-side enumeration endpoints may be added in; a later release. Supported operators: eq - Equal to; and - Logical AND; or - Logical OR; not - Logical NOT; () - Grouping; Recipients that do not have the optional department or; location property are excluded from results when; filtering by that field, regardless of whether not is; used. When the header is omitted, all recipients are returned; without filtering. Example: status eq 'completed' and department eq 'Engineering'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindNumber}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SecurityAwarenessTrainingCampaignRecipientsList(id, params)
			return handleResponse(resp, err, "failed to list training campaign recipients")
		},
	}
}
