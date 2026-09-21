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
	toolCremAccountCompromiseEventDefinitionsList,
	toolCremAccountCompromiseIndicatorsList,
	toolCremAccountCompromiseRiskIndicatorEventsList,
	toolCremAccountCompromiseRiskIndicatorEventGet,
	toolCremAnomalyDetectionRiskIndicatorEventsList,
	toolCremAnomalyDetectionRiskIndicatorEventGet,
	toolCremAssetGroupsList,
	toolCremCloudVmVulnerabilitiesList,
	toolCremContainerVulnerabilitiesList,
	toolCremHighRiskDevicesList,
	toolCremHighRiskDeviceGet,
	toolCremHighRiskUserGet,
	toolCremInternalAssetVulnerabilitiesList,
	toolCremInternetFacingAssetVulnerabilitiesList,
	toolCremSecurityPostureGet,
	toolCremServerlessFunctionVulnerabilitiesList,
	toolCremVulnerabilityGet,
	toolCremVulnerabilityAffectedCloudStorageAssetsList,
	toolCremVulnerabilityAffectedCloudVmsList,
	toolCremVulnerabilityAffectedContainerClustersList,
	toolCremVulnerabilityAffectedContainerImagesList,
	toolCremVulnerabilityAffectedDevicesList,
	toolCremVulnerabilityAffectedGlobalFqdnsList,
	toolCremVulnerabilityAffectedServerlessFunctionLayersList,
	toolCremVulnerabilityAffectedServerlessFunctionsList,
	toolCremVulnerableDevicesList,
}

var ToolsetsWriteCREM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCremAttackSurfaceAssetsUpdateCriticality,
	toolCremAttackSurfaceCloudAssetsUpdate,
	toolCremAttackSurfaceDevicesUpdate,
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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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

			lastDetectedStartDateTime, err := optionalTimeValue("lastDetectedStartDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			lastDetectedEndDateTime, err := optionalTimeValue("lastDetectedEndDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstSeenStartDateTime, err := optionalTimeValue("firstSeenStartDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstSeenEndDateTime, err := optionalTimeValue("firstSeenEndDateTime", request.GetArguments())
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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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

			lastDetectedStartDateTime, err := optionalTimeValue("lastDetectedStartDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			lastDetectedEndDateTime, err := optionalTimeValue("lastDetectedEndDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstSeenStartDateTime, err := optionalTimeValue("firstSeenStartDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstSeenEndDateTime, err := optionalTimeValue("firstSeenEndDateTime", request.GetArguments())
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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
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
			mcp.WithString("cloudAssetId", mcp.Description("The ID of the cloud asset to retrieve."), mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			cloudAssetId, err := requiredValue[string]("cloudAssetId", request.GetArguments())
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
			mcp.WithString("cloudAssetId", mcp.Description("The ID of the cloud asset to retrieve."), mcp.Required()),
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
			cloudAssetId, err := requiredValue[string]("cloudAssetId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			appID, err := requiredValue[string]("appID", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CREMGetAttackSurfaceLocalAppProfile(appID)
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
			appID, err := requiredValue[string]("appID", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			appID, err := requiredValue[string]("appID", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			appID, err := requiredValue[string]("appID", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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
			top, err := optionalStrInt("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

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

func toolCremAccountCompromiseEventDefinitionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_account_compromise_event_definitions_list",
			mcp.WithDescription("Get risk event definitions. Displays detailed information about risk events related to account compromise and the corresponding remediation and suggested actions."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.CremAccountCompromiseEventDefinitionsList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get risk event definitions")
		},
	}
}

func toolCremAccountCompromiseIndicatorsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_account_compromise_indicators_list",
			mcp.WithDescription("Get account compromise indicators. Displays a paginated list of events on user accounts that:."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("Number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
			mcp.WithString("orderBy", mcp.Description("Specifies the field by which the results are sorted. Default: detectedDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of account compromise indicators. Supported fields and operators: 'id' - The ID of a user account on the TrendAI Vision One™ platform. 'account' - The name of the user account. 'name' - The name of the risk event. 'riskLevel' - Risk level of the risk event. Supported values: \"high\", \"medium\", \"low\". 'eq' - Abbreviation of the operator 'equal to'. 'and' - Operator 'and'. 'or' - Operator 'or'. 'not' - Operator 'not'. '( )' - Symbols for grouping operands. Example: account eq 'john'")),
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

			resp, err := client.CremAccountCompromiseIndicatorsList(params)
			return handleResponse(resp, err, "failed to get account compromise indicators")
		},
	}
}

func toolCremAccountCompromiseRiskIndicatorEventsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_account_compromise_risk_indicator_events_list",
			mcp.WithDescription("Get account compromise risk events with summary information. Returns summary information for account compromise risk events detected within the last 30 days as a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The maximum number of records to return per page. The actual number of records returned may be less than the specified value. One of: 20, 50, 100. Default: 100."), mcp.Enum("20", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("Specifies the field and sort direction for ordering results. Default: detectedDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of risk events. Supported fields: 'status' - The status of the risk event. Supported values: new; inProgress; systemRemediated; remediated; dismissed; accepted; 'riskLevel' - The risk level of the risk event. Supported values: low; medium; high; 'assetType' - The asset type impacted by the risk event. Supported values: device; user; ip; domain; service-account; cloud-asset; app; cloud-asset-subscription; All filter fields are optional. If a field is not specified, results are not filtered by that field. Supported operators: eq - Equal to; and - Logical AND; () - Groups filter conditions for precedence or to match multiple values; Example: To retrieve risk events with a specific status, use status eq 'new'; Filter combination limitations: The or and not operators are currently not supported. Support for these operators will be available in a future update. Example: status eq 'new'")),
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

			resp, err := client.CremAccountCompromiseRiskIndicatorEventsList(params)
			return handleResponse(resp, err, "failed to get account compromise risk events with summary information")
		},
	}
}

func toolCremAccountCompromiseRiskIndicatorEventGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_account_compromise_risk_indicator_event_get",
			mcp.WithDescription("Get the detailed information of the specified account compromise risk event. Returns information for the specified risk event in the account compromise risk factor."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the risk event.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremAccountCompromiseRiskIndicatorEventGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get the detailed information of the specified account compromise risk event")
		},
	}
}

func toolCremAnomalyDetectionRiskIndicatorEventsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_anomaly_detection_risk_indicator_events_list",
			mcp.WithDescription("Get activity and behaviors risk events with summary information. Returns summary information for activity and behaviors risk events detected within the last 30 days as a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The maximum number of records to return per page. The actual number of records returned may be less than the specified value. One of: 20, 50, 100. Default: 100."), mcp.Enum("20", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("Specifies the field and sort direction for ordering results. Default: detectedDateTime desc.")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of risk events. Supported fields: 'status' - The status of the risk event. Supported values: new; inProgress; systemRemediated; remediated; dismissed; accepted; 'riskLevel' - The risk level of the risk event. Supported values: low; medium; high; 'assetType' - The asset type impacted by the risk event. Supported values: device; user; ip; domain; service-account; cloud-asset; app; cloud-asset-subscription; All filter fields are optional. If a field is not specified, results are not filtered by that field. Supported operators: eq - Equal to; and - Logical AND; () - Groups filter conditions for precedence or to match multiple values; Example: To retrieve risk events with a specific status, use status eq 'new'; Filter combination limitations: The or and not operators are currently not supported. Support for these operators will be available in a future update. Example: status eq 'new'")),
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

			resp, err := client.CremAnomalyDetectionRiskIndicatorEventsList(params)
			return handleResponse(resp, err, "failed to get activity and behaviors risk events with summary information")
		},
	}
}

func toolCremAnomalyDetectionRiskIndicatorEventGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_anomaly_detection_risk_indicator_event_get",
			mcp.WithDescription("Get the detailed information of the specified activity and behaviors risk event. Returns information for the specified risk event in the activity and behaviors risk factor."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the risk event.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremAnomalyDetectionRiskIndicatorEventGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get the detailed information of the specified activity and behaviors risk event")
		},
	}
}

func toolCremAssetGroupsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_asset_groups_list",
			mcp.WithDescription("Get cyber risk subindexes of asset group data. Displays the cyber risk subindexes of the asset group data and the hierarchical relationships."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200. Default: 100."), mcp.Enum("10", "50", "100", "200")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the asset groups; Supported fields: name - The name of the group - Any value; id - The ID of the group - Any value; riskIndex - The risk score for your organization derived from a comprehensive assessment of a variety of risk categories and factors. - Any value; riskLevel - The risk level of the event - low, medium, high; updatedDateTime - The last date and time when the response was updated should be in ISO 8601 format. - Any value; parentId - The ID of the parent group - Any value; Supported operators: eq - Operator 'equal to'. - -; gt - Operator greater than. - Only applicable to riskIndex and updatedDateTime; ge - Operator greater than or equal. - Only applicable to riskIndex and updatedDateTime; le - Operator less than or equal. - Only applicable to riskIndex and updatedDateTime; lt - Operator less than. - Only applicable to riskIndex and updatedDateTime; and - Operator 'and'. - -; or - Operator 'or'. - -; not - Operator 'not'. - -; ( ) - Symbols for grouping operands. - -. Example: riskIndex ge 50.0")),
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

			resp, err := client.CremAssetGroupsList(params)
			return handleResponse(resp, err, "failed to get cyber risk subindexes of asset group data")
		},
	}
}

func toolCremAttackSurfaceAssetsUpdateCriticality(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_assets_update_criticality",
			mcp.WithDescription("Batch update asset criticality. Updates the criticality of up to 500 assets in a single request. Supported asset types: device, domainAccount, publicIpAddress, globalFqdn, serviceAccount, cloudAsset."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Description("The list of asset criticality updates."), mcp.Items(map[string]any{"description": "Defines a criticality update for a single asset. Either score or systemDefined must be provided;", "properties": map[string]any{"score": map[string]any{"description": "The customer-defined criticality score as a range from 1 to 10. Ranges map to labels as follows: 1 to 3 is low, 4 to 7 is medium, 8 to 10...", "type": "integer"}, "systemDefined": map[string]any{"description": "The flag for resetting the asset criticality to the system-defined default, clearing any customer-defined score.", "enum": []any{true}, "type": "boolean"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremAttackSurfaceAssetsUpdateCriticality(params)
			return handleResponse(resp, err, "failed to batch update asset criticality")
		},
	}
}

func toolCremAttackSurfaceCloudAssetsUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_cloud_assets_update",
			mcp.WithDescription("Update cloud asset tags. Updates the tags of the specified cloud assets."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Description("Example request to update tags in multiple assets."), mcp.Items(map[string]any{"properties": map[string]any{"assetCustomTagIds": map[string]any{"description": "The ID of the tags", "items": map[string]any{"type": "string"}, "type": "array"}, "id": map[string]any{"description": "The ID of the Cloud Asset on the TrendAI Vision One™ platform", "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremAttackSurfaceCloudAssetsUpdate(params)
			return handleResponse(resp, err, "failed to update cloud asset tags")
		},
	}
}

func toolCremAttackSurfaceDevicesUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_attack_surface_devices_update",
			mcp.WithDescription("Update device tags. Updates the tags of the specified devices."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Description("Example request to update tags in multiple assets."), mcp.Items(map[string]any{"properties": map[string]any{"assetCustomTagIds": map[string]any{"description": "The ID of the tags", "items": map[string]any{"type": "string"}, "type": "array"}, "id": map[string]any{"description": "The ID of the device on the TrendAI Vision One™ platform", "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremAttackSurfaceDevicesUpdate(params)
			return handleResponse(resp, err, "failed to update device tags")
		},
	}
}

func toolCremCloudVmVulnerabilitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_cloud_vm_vulnerabilities_list",
			mcp.WithDescription("Get vulnerabilities in cloud VMs. Displays all the highly-exploitable CVEs detected in your cloud VMs in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: cveRiskLevel desc."), mcp.Enum("cveRiskLevel desc", "cveRiskLevel asc")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200. Default: 100."), mcp.Enum("10", "50", "100", "200")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the highly-exploitable CVEs detected in your cloud VMs; Supported fields: cveId - The Common Vulnerabilities and Exposures identifier (CVE ID) - Any value; cveRiskLevel - The risk level of the risk event triggered by the CVE - high, medium, low; globalExploitActivityLevel - Indicates how often attackers are exploiting a vulnerability - high, medium, low; firstSeenDateTime - The date the CVE was first detected in your environment - Any value; publishedDateTime - The date MITRE published the CVE - Any value; Supported operators: eq - Operator equal to. - -; gt - Operator greater than. - Only applicable to firstSeenDateTime and publishedDateTime; ge - Operator greater than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; le - Operator less than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; lt - Operator less than. - Only applicable to firstSeenDateTime and publishedDateTime; and - Operator and. - -; or - Operator or. - -; not - Operator not. - -; ( ) - Symbols for grouping operands. - -. Example: cveRiskLevel eq 'high'")),
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

			resp, err := client.CremCloudVmVulnerabilitiesList(params)
			return handleResponse(resp, err, "failed to get vulnerabilities in cloud VMs")
		},
	}
}

func toolCremContainerVulnerabilitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_container_vulnerabilities_list",
			mcp.WithDescription("Get CVEs in containers. Displays all the highly-exploitable CVEs detected in your containers assets in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: cveRiskLevel desc."), mcp.Enum("cveRiskLevel desc", "cveRiskLevel asc")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200. Default: 100."), mcp.Enum("10", "50", "100", "200")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the highly-exploitable CVEs detected in your containers. Supported fields: cveId - The Common Vulnerabilities and Exposures identifier (CVE ID) - Any value; cveRiskLevel - The risk level of the risk event triggered by the CVE - high, medium, low; globalExploitActivityLevel - Indicates how often attackers are exploiting a vulnerability - high, medium, low; firstSeenDateTime - The date the CVE was first detected in your environment - Any value; publishedDateTime - The date MITRE published the CVE - Any value; Supported operators: eq - Operator equal to. - -; gt - Operator greater than. - Only applicable to firstSeenDateTime and publishedDateTime; ge - Operator greater than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; le - Operator less than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; lt - Operator less than. - Only applicable to firstSeenDateTime and publishedDateTime; and - Operator and. - -; or - Operator or. - -; not - Operator not. - -; ( ) - Symbols for grouping operands. - -. Example: cveRiskLevel eq 'high'")),
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

			resp, err := client.CremContainerVulnerabilitiesList(params)
			return handleResponse(resp, err, "failed to get CVEs in containers")
		},
	}
}

func toolCremHighRiskDevicesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_high_risk_devices_list",
			mcp.WithDescription("Get at-risk devices list. Displays information about at-risk devices that match specified criteria in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: riskScore desc.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "1000")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the at-risk devices list. Supported fields: id - The ID of the device on the TrendAI Vision One™ platform. deviceName - The name of the device; ip - The IP addresses of the device. os - The operating system of the device. riskScore - The risk score of a user. lastLogonUser - The user who last signed into the device. Supported operators: eq - Operator 'equal to'. and - Operator 'and'. or - Operator 'or'. not - Operator 'not'. ( ) - Symbols for grouping operands. gt - Operator 'greater than'. ge - Operator 'greater than or equal'. le - Operator 'less than or equal'. lt - Operator 'less than'. Example: deviceName eq 'SASE-PC1'")),
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

			resp, err := client.CremHighRiskDevicesList(params)
			return handleResponse(resp, err, "failed to get at-risk devices list")
		},
	}
}

func toolCremHighRiskDeviceGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_high_risk_device_get",
			mcp.WithDescription("Get device risk profile. Displays the risk profile of the specified device."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the device on the TrendAI Vision One™ platform.")),
			mcp.WithNumber("riskyEventScore", mcp.Description("The minimum risk score of the risk events this API endpoint returns. Default: 70.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "riskyEventScore", Kind: kindNumber}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremHighRiskDeviceGet(id, params)
			return handleResponse(resp, err, "failed to get device risk profile")
		},
	}
}

func toolCremHighRiskUserGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_high_risk_user_get",
			mcp.WithDescription("Get user risk profile. Displays the risk profile of the specified user."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of a user on the TrendAI Vision One™ platform.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremHighRiskUserGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get user risk profile")
		},
	}
}

func toolCremInternalAssetVulnerabilitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_internal_asset_vulnerabilities_list",
			mcp.WithDescription("Get CVEs in internal assets. Displays all the highly-exploitable CVEs detected in your internal assets in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: cveRiskLevel desc."), mcp.Enum("cveRiskLevel desc", "cveRiskLevel asc")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200. Default: 100."), mcp.Enum("10", "50", "100", "200")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the highly-exploitable CVEs detected in your internal assets. Supported fields: cveId - The Common Vulnerabilities and Exposures identifier (CVE ID) - Any value; cveRiskLevel - The risk level of the risk event triggered by the CVE - high, medium, low; globalExploitActivityLevel - Indicates how often attackers are exploiting a vulnerability - high, medium, low; firstSeenDateTime - The date the CVE was first detected in your environment - Any value; publishedDateTime - The date MITRE published the CVE - Any value; Supported operators: eq - Operator equal to. - -; gt - Operator greater than. - Only applicable to firstSeenDateTime and publishedDateTime; ge - Operator greater than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; le - Operator less than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; lt - Operator less than. - Only applicable to firstSeenDateTime and publishedDateTime; and - Operator and. - -; or - Operator or. - -; not - Operator not. - -; ( ) - Symbols for grouping operands. - -. Example: cveRiskLevel eq 'high'")),
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

			resp, err := client.CremInternalAssetVulnerabilitiesList(params)
			return handleResponse(resp, err, "failed to get CVEs in internal assets")
		},
	}
}

func toolCremInternetFacingAssetVulnerabilitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_internet_facing_asset_vulnerabilities_list",
			mcp.WithDescription("Get CVEs in internet-facing assets. Displays all the highly-exploitable CVEs detected in your internet-facing assets in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: cveRiskLevel desc."), mcp.Enum("cveRiskLevel desc", "cveRiskLevel asc")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200. Default: 100."), mcp.Enum("10", "50", "100", "200")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the highly-exploitable CVEs detected in your internet-facing assets. Supported fields: cveId - The Common Vulnerabilities and Exposures identifier (CVE ID) - Any value; cveRiskLevel - The risk level of the risk event triggered by the CVE - high, medium, low; globalExploitActivityLevel - Indicates how often attackers are exploiting a vulnerability - high, medium, low; firstSeenDateTime - The date the CVE was first detected in your environment - Any value; publishedDateTime - The date MITRE published the CVE - Any value; Supported operators: eq - Operator equal to. - -; gt - Operator greater than. - Only applicable to firstSeenDateTime and publishedDateTime; ge - Operator greater than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; le - Operator less than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; lt - Operator less than. - Only applicable to firstSeenDateTime and publishedDateTime; and - Operator and. - -; or - Operator or. - -; not - Operator not. - -; ( ) - Symbols for grouping operands. - -. Example: cveRiskLevel eq 'high'")),
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

			resp, err := client.CremInternetFacingAssetVulnerabilitiesList(params)
			return handleResponse(resp, err, "failed to get CVEs in internet-facing assets")
		},
	}
}

func toolCremSecurityPostureGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_security_posture_get",
			mcp.WithDescription("Get security posture data. Displays information about your organization's security posture."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.CremSecurityPostureGet(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get security posture data")
		},
	}
}

func toolCremServerlessFunctionVulnerabilitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_serverless_function_vulnerabilities_list",
			mcp.WithDescription("Get vulnerabilities in serverless functions. Displays all the highly-exploitable CVEs detected in your serverless functions in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted. Default: cveRiskLevel desc."), mcp.Enum("cveRiskLevel desc", "cveRiskLevel asc")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200. Default: 100."), mcp.Enum("10", "50", "100", "200")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the prioritized CVEs. Supported fields: cveId - The Common Vulnerabilities and Exposures identifier (CVE ID) - Any value; cveRiskLevel - The risk level of the risk event triggered by the CVE - high, medium, low; globalExploitActivityLevel - Indicates how often attackers are exploiting a vulnerability - high, medium, low; firstSeenDateTime - The date the CVE was first detected in your environment - Any value; publishedDateTime - The date MITRE published the CVE - Any value; Supported operators: eq - Operator equal to. - -; gt - Operator greater than. - Only applicable to firstSeenDateTime and publishedDateTime; ge - Operator greater than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; le - Operator less than or equal. - Only applicable to firstSeenDateTime and publishedDateTime; lt - Operator less than. - Only applicable to firstSeenDateTime and publishedDateTime; and - Operator and. - -; or - Operator or. - -; not - Operator not. - -; ( ) - Symbols for grouping operands. - -. Example: cveRiskLevel eq 'high'")),
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

			resp, err := client.CremServerlessFunctionVulnerabilitiesList(params)
			return handleResponse(resp, err, "failed to get vulnerabilities in serverless functions")
		},
	}
}

func toolCremVulnerabilityGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_get",
			mcp.WithDescription("Get basic CVE information. Displays basic information of the specified CVE, including mitigation options."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremVulnerabilityGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get basic CVE information")
		},
	}
}

func toolCremVulnerabilityAffectedCloudStorageAssetsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_affected_cloud_storage_assets_list",
			mcp.WithDescription("Get cloud storage affected by CVEs. Displays the details of the cloud storage assets affected by the specified CVE."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
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

			resp, err := client.CremVulnerabilityAffectedCloudStorageAssetsList(id, params)
			return handleResponse(resp, err, "failed to get cloud storage affected by CVEs")
		},
	}
}

func toolCremVulnerabilityAffectedCloudVmsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_affected_cloud_vms_list",
			mcp.WithDescription("Get cloud VMs affected by CVEs. Displays the details of the cloud VMs affected by the specified CVE."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
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

			resp, err := client.CremVulnerabilityAffectedCloudVmsList(id, params)
			return handleResponse(resp, err, "failed to get cloud VMs affected by CVEs")
		},
	}
}

func toolCremVulnerabilityAffectedContainerClustersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_affected_container_clusters_list",
			mcp.WithDescription("Get container clusters affected by CVEs. Displays the details of container clusters affected by the specified CVE."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
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

			resp, err := client.CremVulnerabilityAffectedContainerClustersList(id, params)
			return handleResponse(resp, err, "failed to get container clusters affected by CVEs")
		},
	}
}

func toolCremVulnerabilityAffectedContainerImagesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_affected_container_images_list",
			mcp.WithDescription("Get container images affected by CVEs. Displays the details of the container images affected by the specified CVE."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
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

			resp, err := client.CremVulnerabilityAffectedContainerImagesList(id, params)
			return handleResponse(resp, err, "failed to get container images affected by CVEs")
		},
	}
}

func toolCremVulnerabilityAffectedDevicesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_affected_devices_list",
			mcp.WithDescription("Get devices affected by CVEs. Displays the details of the devices affected by the specified CVE in internal assets."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
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

			resp, err := client.CremVulnerabilityAffectedDevicesList(id, params)
			return handleResponse(resp, err, "failed to get devices affected by CVEs")
		},
	}
}

func toolCremVulnerabilityAffectedGlobalFqdnsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_affected_global_fqdns_list",
			mcp.WithDescription("Get FQDNs affected by CVEs. Displays the details of the FQDNs affected by the specified CVE in internet-facing assets."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
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

			resp, err := client.CremVulnerabilityAffectedGlobalFqdnsList(id, params)
			return handleResponse(resp, err, "failed to get FQDNs affected by CVEs")
		},
	}
}

func toolCremVulnerabilityAffectedServerlessFunctionLayersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_affected_serverless_function_layers_list",
			mcp.WithDescription("Get serverless function layers affected by CVEs. Retrieves details of the serverless function layers affected by the specified CVE."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
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

			resp, err := client.CremVulnerabilityAffectedServerlessFunctionLayersList(id, params)
			return handleResponse(resp, err, "failed to get serverless function layers affected by CVEs")
		},
	}
}

func toolCremVulnerabilityAffectedServerlessFunctionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerability_affected_serverless_functions_list",
			mcp.WithDescription("Get serverless functions affected by CVEs. Displays the details of serverless functions affected by the specified CVEs."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The CVE ID.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200, 500, 1000. Default: 100."), mcp.Enum("10", "50", "100", "200", "500", "1000")),
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

			resp, err := client.CremVulnerabilityAffectedServerlessFunctionsList(id, params)
			return handleResponse(resp, err, "failed to get serverless functions affected by CVEs")
		},
	}
}

func toolCremVulnerableDevicesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"crem_vulnerable_devices_list",
			mcp.WithDescription("Get CVEs detected in a device. Displays all the highly-exploitable CVEs found in a device in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The field by which the results are sorted.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 10, 50, 100, 200. Default: 100."), mcp.Enum("10", "50", "100", "200")),
			mcp.WithString("cveDetectionStatus", mcp.Description("The CVE detection status of the devices included on the list. Default: affected."), mcp.Enum("affected", "any")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the prioritized CVEs. Supported fields: id - The ID of the device on the TrendAI Vision One™ platform. - Any value; deviceName - The name of the device. - Any value; ip - The IP addresses of a device. - Any value; criticality - The criticality of the device. - high, medium, low; cvssScore - The CVSS score. - Any value; globalExploitActivityLevel - Indicates how often attackers are exploiting a vulnerability. - high, medium, low; cveId - The Common Vulnerabilities and Exposures identifier (CVE ID). - Any value; cveMitigationStatus - The mitigation status of the CVE detected on the device. - new, inProgress, closed, dismissed, accepted, mitigated; cveEventRiskLevel - The risk level of the risk event triggered by a CVE - high, medium, low; Supported operators: eq - Operator 'equal to'. gt - Operator 'greater than'. ge - Operator 'greater than or equal'. le - Operator 'less than or equal'. lt - Operator 'less than'. and - Operator 'and'. or - Operator 'or'. not - Operator 'not'. ( ) - Symbols for grouping operands. Example: deviceName eq 'SASE-PC1'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "orderBy", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "cveDetectionStatus", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CremVulnerableDevicesList(params)
			return handleResponse(resp, err, "failed to get CVEs detected in a device")
		},
	}
}
