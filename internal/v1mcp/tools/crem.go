package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyCREM = []func(client *v1client.V1ApiClient) mcpserver.ServerTool{
	toolCREMAttackSurfaceDevicesList,
	toolCREMAttackSurfaceDomainAccountsList,
	toolCREMAttackSurfaceServiceAccountsList,
	toolCREMAttackSurfaceGlobalFQDNsList,
	toolCREMAttackSurfacePublicIPsList,
	toolCREMAttackSurfaceCloudAssetsList,
	toolCREMAttackSurfaceHighRiskUsersList,
	toolCREMAttackSurfaceCloudAssetProfileGet,
	toolCREMAttackSurfaceCloudAssetRiskIndicatorsList,
	toolCREMAttackSurfaceLocalAppsList,
	toolCREMAttackSurfaceLocalAppProfileGet,
	toolCREMAttackSurfaceLocalAppRiskIndicatorsList,
	toolCREMAttackSurfaceLocalAppDevicesList,
	toolCREMAttackSurfaceLocalAppExecutableFilesList,
	toolCREMAttackSurfaceCustomTagsList,
}

func toolCREMAttackSurfaceDevicesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_devices_list",
			mcp.WithDescription("List discovered attack surface devices"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterAttackSurfaceDevices)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(
					withOrdering(
						asc_desc,
						"deviceName",
						"latestRiskScore",
					)...,
				),
			),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("lastDetectedStartDateTime",
				mcp.Description("The start time of the data retrieval range, in ISO 8601 format."),
			),
			mcp.WithString("lastDetectedEndDateTime",
				mcp.Description("The end time of the data retrieval range, represented in ISO 8601 format."),
			),
			mcp.WithString("firstSeenStartDateTime",
				mcp.Description("The start time of the data retrieval range, represented in ISO 8601 format."),
			),
			mcp.WithString("firstSeenStartDateTime",
				mcp.Description("The end time of the data retrieval range, represented in ISO 8601 format."),
			),
			mcp.WithString("skipToken",
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			lastDetectedStartDateTime, err := optionalTimeValue("lastDetectedStartDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			lastDetectedEndDateTime, err := optionalTimeValue("lastDetectedEndDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstSeenStartDateTime, err := optionalTimeValue("firstSeenStartDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstSeenEndDateTime, err := optionalTimeValue("firstSeenEndDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				OrderBy:                   orderBy,
				Top:                       top,
				LastDetectedStartDateTime: lastDetectedStartDateTime,
				LastDetectedEndDateTime:   lastDetectedEndDateTime,
				FirstSeenStartDateTime:    firstSeenStartDateTime,
				FirstSeenEndDateTime:      firstSeenEndDateTime,
				SkipToken:                 skipToken,
			}

			resp, err := client.CREMListAttackSurfaceDevices(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface devices")
		},
	}
}

func toolCREMAttackSurfaceDomainAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_domain_accounts_list",
			mcp.WithDescription("List discovered attack surface domain accounts"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterDomainAccounts)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(
					withOrdering(
						asc_desc,
						"latestRiskScore",
						"userAccount",
					)...,
				),
			),
			mcp.WithString("skipToken",
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMListAttackSurfaceDomainAccounts(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface domain accounts")
		},
	}
}

func toolCREMAttackSurfaceGlobalFQDNsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_global_fqdns_list",
			mcp.WithDescription("List discovered internet facing domains (Fully Qualified Domain Names)"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterFQDNS)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "latestRiskScore")...),
			),
			mcp.WithString("skipToken",
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMListAttackSurfaceGlobalFQDNs(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface public domains")
		},
	}
}

func toolCREMAttackSurfacePublicIPsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_public_ips_list",
			mcp.WithDescription("List discovered public IP addresses"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterIps)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "latestRiskScore")...),
			),
			mcp.WithString("skipToken",
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMListAttackSurfacePublicIPs(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface public ips")
		},
	}
}

func toolCREMAttackSurfaceCloudAssetsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_cloud_assets_list",
			mcp.WithDescription("List discovered cloud assets"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudAssets)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "latestRiskScore")...),
			),
			mcp.WithString("lastDetectedStartDateTime",
				mcp.Description("The start time of the data retrieval range, in ISO 8601 format."),
			),
			mcp.WithString("lastDetectedEndDateTime",
				mcp.Description("The end time of the data retrieval range, represented in ISO 8601 format."),
			),
			mcp.WithString("firstSeenStartDateTime",
				mcp.Description("The start time of the data retrieval range, represented in ISO 8601 format."),
			),
			mcp.WithString("firstSeenStartDateTime",
				mcp.Description("The end time of the data retrieval range, represented in ISO 8601 format."),
			),
			mcp.WithString("skipToken",
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			lastDetectedStartDateTime, err := optionalTimeValue("lastDetectedStartDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			lastDetectedEndDateTime, err := optionalTimeValue("lastDetectedEndDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstSeenStartDateTime, err := optionalTimeValue("firstSeenStartDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstSeenEndDateTime, err := optionalTimeValue("firstSeenEndDateTime", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				OrderBy:                   orderBy,
				Top:                       top,
				LastDetectedStartDateTime: lastDetectedStartDateTime,
				LastDetectedEndDateTime:   lastDetectedEndDateTime,
				FirstSeenStartDateTime:    firstSeenStartDateTime,
				FirstSeenEndDateTime:      firstSeenEndDateTime,
				SkipToken:                 skipToken,
			}

			resp, err := client.CREMListAttackSurfaceCloudAssets(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface cloud assets")
		},
	}
}

func toolCREMAttackSurfaceHighRiskUsersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_high_risk_users_list",
			mcp.WithDescription("List high risk users"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterHighRiskUsers)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(
					asc_desc,
					"riskScore",
					"userName",
				)...),
			),
			mcp.WithString("skipToken",
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}
			resp, err := client.CREMListHighRiskUsers(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list high risk users")
		},
	}
}

func toolCREMAttackSurfaceServiceAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_service_accounts_list",
			mcp.WithDescription("List discovered service accounts"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterServiceAccounts)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(
					asc_desc,
					"latestRiskScore",
					"userAccount",
				)...,
				),
			),
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:     top,
				OrderBy: orderBy,
			}

			resp, err := client.CREMListAttackSurfaceServiceAccounts(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface service accounts")
		},
	}
}

func toolCREMAttackSurfaceCloudAssetProfileGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_cloud_asset_profile_get",
			mcp.WithDescription("Get a cloud asset's profile"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("cloudAssetId", mcp.Description("The ID of the cloud asset to retrieve.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			cloudAssetId, err := requiredValue[string]("cloudAssetId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CREMGetAttackSurfaceCloudAssetProfile(cloudAssetId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get attack surface cloud asset profile")
		},
	}
}

func toolCREMAttackSurfaceCloudAssetRiskIndicatorsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_cloud_asset_risk_indicators_list",
			mcp.WithDescription("List a cloud asset's risk indicators"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("cloudAssetId", mcp.Description("The ID of the cloud asset to retrieve.")),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudAssetRiskIndicators)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(
					asc_desc,
					"detectedDateTime",
				)...),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			cloudAssetId, err := requiredValue[string]("cloudAssetId", request.Params.Arguments)
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

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParames := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMListAttackSurfaceCloudAssetRiskIndicators(cloudAssetId, filter, queryParames)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get attack surface cloud asset risk indicators")
		},
	}
}

func toolCREMAttackSurfaceLocalAppsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_local_apps_list",
			mcp.WithDescription("List discovered local applications"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterLocalApps)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(
					withOrdering(
						asc_desc,
						"latestRiskScore",
						"firstSeenDateTime",
						"lastDetectedDateTime",
					)...,
				),
			),
			mcp.WithString("skipToken",
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMListAttackSurfaceLocalApps(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface local apps")
		},
	}
}

func toolCREMAttackSurfaceLocalAppProfileGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_local_app_profile_get",
			mcp.WithDescription("Get a local app's profile"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("appID",
				mcp.Description("The ID of the local app to retrieve."),
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			appID, err := requiredValue[string]("appID", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CREMGetAttackSurfaceCloudAssetProfile(appID)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get attack surface local app profile")
		},
	}
}

func toolCREMAttackSurfaceLocalAppRiskIndicatorsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_local_app_risk_indicators_list",
			mcp.WithDescription("List a local app's risk indicators"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("appID",
				mcp.Description("The ID of the local app to retrieve."),
				mcp.Required(),
			),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterLocalAppRiskIndicators)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "detectedDateTime")...),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			appID, err := requiredValue[string]("appID", request.Params.Arguments)
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

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMGetAttackSurfaceLocalAppRiskIndicators(appID, filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface local app risk indicators")
		},
	}
}

func toolCREMAttackSurfaceLocalAppDevicesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_local_app_devices_list",
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithDescription("Displays the devices with the specified local application installed"),
			mcp.WithString("appID",
				mcp.Description("The ID of the local app to retrieve devices for."),
				mcp.Required(),
			),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterLocalAppDevices)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "latestRiskScore")...),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			appID, err := requiredValue[string]("appID", request.Params.Arguments)
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

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMListAttackSurfaceLocalAppDevices(appID, filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface local app devices")
		},
	}
}

func toolCREMAttackSurfaceLocalAppExecutableFilesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_local_app_executable_files_list",
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithDescription("Displays the local applications installed executable files"),
			mcp.WithString("appID",
				mcp.Description("The ID of the local app to retrieve executable files for."),
				mcp.Required(),
			),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterLocalAppExecutables)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "lastDetectedDateTime")...),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			appID, err := requiredValue[string]("appID", request.Params.Arguments)
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

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMListAttackSurfaceLocalAppExecutableFiles(appID, filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list attack surface local app executable files")
		},
	}
}

func toolCREMAttackSurfaceCustomTagsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_custom_tags_list",
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithDescription("List tag definitions"),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(cremTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCustomTags)),
			mcp.WithString("orderBy",
				mcp.Description("The field by which the results are sorted"),
				mcp.Enum(withOrdering(asc_desc, "key")...),
			),
			mcp.WithString("skipToken",
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

			orderBy, err := optionalValue[string]("orderBy", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				OrderBy:   orderBy,
				SkipToken: skipToken,
			}

			resp, err := client.CREMListCustomTags(filter, queryParams)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list custom tags")
		},
	}
}

func cremTop() []string {
	return []string{
		"10",
		"50",
		"100",
		"200",
		"500",
		"1000",
	}
}
