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
	ToolCREMAttackSurfaceDevicesList,
	ToolCREMAttackSurfaceDomainAccountsList,
	ToolCREMAttackSurfaceServiceAccountsList,
	ToolCREMAttackSurfaceGlobalFQDNsList,
	ToolCREMAttackSurfacePublicIPsList,
	ToolCREMAttackSurfaceCloudAssetsList,
	ToolCREMAttackSurfaceHighRiskUsersList,
	ToolCREMAttackSurfaceCloudAssetProfileGet,
	ToolCREMAttackSurfaceCloudAssetRiskIndicatorsList,
	ToolCREMAttackSurfaceLocalAppsList,
	ToolCREMAttackSurfaceLocalAppProfileGet,
	ToolCREMAttackSurfaceLocalAppRiskIndicatorsList,
	ToolCREMAttackSurfaceLocalAppDevicesList,
	ToolCREMAttackSurfaceLocalAppExecutableFilesList,
	ToolCREMAttackSurfaceCustomTagsList,
}

func ToolCREMAttackSurfaceDevicesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_devices_list",
			mcp.WithDescription("List discovered attack surface devices"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterAttackSurfaceDevices)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
					"deviceName asc",
					"deviceName desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
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
			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceDomainAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_domain_accounts_list",
			mcp.WithDescription("List discovered attack surface domain accounts"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterDomainAccounts)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
					"userAccount asc",
					"userAccount desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceGlobalFQDNsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_global_fqdns_list",
			mcp.WithDescription("List discovered internet facing domains (Fully Qualified Domain Names)"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterFQDNS)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfacePublicIPsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_public_ips_list",
			mcp.WithDescription("List discovered public IP addresses"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterIps)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceCloudAssetsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_cloud_assets_list",
			mcp.WithDescription("List discovered cloud assets"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudAssets)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
				),
				mcp.Description("The field by which the results are sorted"),
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
			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceHighRiskUsersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_high_risk_users_list",
			mcp.WithDescription("List high risk users"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterHighRiskUsers)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
					"userName asc",
					"userName desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceServiceAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_service_accounts_list",
			mcp.WithDescription("List discovered service accounts"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterServiceAccounts)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
					"userAccount asc",
					"userAccount desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceCloudAssetProfileGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
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

func ToolCREMAttackSurfaceCloudAssetRiskIndicatorsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_cloud_asset_risk_indicators_list",
			mcp.WithDescription("List a cloud asset's risk indicators"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("cloudAssetId", mcp.Description("The ID of the cloud asset to retrieve.")),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudAssetRiskIndicators)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"detectedDateTime asc",
					"detectedDateTime desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			cloudAssetId, err := requiredValue[string]("cloudAssetId", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceLocalAppsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_local_apps_list",
			mcp.WithDescription("List discovered local applications"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterLocalApps)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
					"firstSeenDateTime asc",
					"firstSeenDateTime desc",
					"lastDetectedDateTime asc",
					"lastDetectedDateTime desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceLocalAppProfileGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
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

func ToolCREMAttackSurfaceLocalAppRiskIndicatorsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
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
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterLocalAppRiskIndicators)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"detectedDateTime asc",
					"detectedDateTime desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			appID, err := requiredValue[string]("appID", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceLocalAppDevicesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
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
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterLocalAppDevices)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"latestRiskScore asc",
					"latestRiskScore desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			appID, err := requiredValue[string]("appID", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceLocalAppExecutableFilesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
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
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterLocalAppExecutables)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"lastDetectedDateTime asc",
					"lastDetectedDateTime desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			appID, err := requiredValue[string]("appID", request.Params.Arguments)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalIntValue("top", request.Params.Arguments)
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

func ToolCREMAttackSurfaceCustomTagsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_custom_tags_list",
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithDescription("List tag definitions"),
			mcp.WithNumber("top", mcp.Description(tooldescriptions.DefaultTop)),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCustomTags)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"key asc",
					"key desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.Params.Arguments)
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
