package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyCloudPosture = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudPostureAccountsList,
	toolCloudPostureAccountChecksList,
	toolCloudPostureTemplateScannerRun,
	toolCloudPostureAccountScanSettingsGet,
}

var ToolsetsWriteCloudPosture = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudPostureAccountScan,
	toolCloudPostureAccountScanSettingsUpdate,
}

// Cloud Posture Custom Rules (Beta) toolsets
var ToolsetsReadOnlyCloudPostureBeta = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudPostureCustomRulesList,
	toolCloudPostureCustomRuleGet,
}

var ToolsetsWriteCloudPostureBeta = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudPostureCustomRuleCreate,
	toolCloudPostureCustomRuleUpdate,
	toolCloudPostureCustomRuleDelete,
	toolCloudPostureCustomRuleTest,
}

func toolCloudPostureAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_accounts_list",
			mcp.WithDescription("List CSPM Accounts"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Min(50),
				mcp.Max(200),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				SkipToken: skipToken,
			}

			resp, err := client.CloudPostureListAccounts(queryParams)
			return handleStatusResponse(
				resp,
				err,
				http.StatusOK,
				"failed to list accounts",
			)
		},
	}
}

func toolCloudPostureAccountChecksList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_account_checks_list",
			mcp.WithDescription("List the checks of an account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterCloudPostureChecks)),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Min(50),
				mcp.Max(200),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval range")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			startDate, err := optionalTimeValue("startDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			endDate, err := optionalTimeValue("endDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:           top,
				StartDateTime: startDate,
				EndDateTime:   endDate,
				SkipToken:     skipToken,
			}

			resp, err := client.CloudPostureListAccountChecks(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list accounts checks")
		},
	}
}

func toolCloudPostureTemplateScannerRun(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_template_scanner_run",
			mcp.WithDescription("Scan an infrastructure as code template using the cloud posture template scanner"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("type",
				mcp.Required(),
				mcp.Enum("cloudformation-template", "terraform-template"),
			),
			mcp.WithString("content",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			templateType, err := requiredValue[string]("type", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			content, err := requiredValue[string]("content", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureScanTemplate(content, templateType)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to scan template")
		},
	}
}

func toolCloudPostureAccountScanSettingsGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_account_scan_settings_get",
			mcp.WithDescription("Get the scan settings for an account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureGetAccountScanSettings(accountId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get account scan settings")
		},
	}
}

func toolCloudPostureAccountScan(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_account_scan",
			mcp.WithDescription("Start scanning cloud posture account"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureScanAccount(accountId)
			return handleStatusResponse(resp, err, http.StatusAccepted, "failed to start account scan")
		},
	}
}

func toolCloudPostureAccountScanSettingsUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_account_scan_settings_update",
			mcp.WithDescription("Update an account's scan settings"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("accountId",
				mcp.Required(),
			),
			mcp.WithNumber("interval",
				mcp.Max(12),
				mcp.Min(1),
			),
			mcp.WithBoolean("enabled"),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			interval, err := optionalIntValue("interval", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			enabled, err := optionalPointerValue[bool]("enabled", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureUpdateAccountScanSettings(
				accountId,
				enabled,
				interval,
			)
			return handleStatusResponse(resp, err, http.StatusNoContent, "failed to update account scan settings")
		},
	}
}

// Cloud Posture Custom Rules (Beta) tools

func toolCloudPostureCustomRulesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_custom_rules_list",
			mcp.WithDescription("(Beta) Displays the custom rules of your company in a paginated list"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithNumber("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Min(50),
				mcp.Max(200),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token used to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			top, err := optionalIntValue("top", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			skipToken, err := optionalValue[string]("skipToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			queryParams := v1client.QueryParameters{
				Top:       top,
				SkipToken: skipToken,
			}

			resp, err := client.CloudPostureListCustomRules(queryParams)
			return handleStatusResponse(
				resp,
				err,
				http.StatusOK,
				"failed to list custom rules",
			)
		},
	}
}

func toolCloudPostureCustomRuleGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_custom_rule_get",
			mcp.WithDescription("(Beta) Returns the configuration of the specified custom rule"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("ruleId",
				mcp.Required(),
				mcp.Description("The Cloud Risk Management ID of the custom rule"),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			ruleId, err := requiredValue[string]("ruleId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureGetCustomRule(ruleId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get custom rule")
		},
	}
}

func toolCloudPostureCustomRuleCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_custom_rule_create",
			mcp.WithDescription("(Beta) Creates a custom rule for your company. Enabled custom rules are immediately available to all your cloud accounts. Requires Master Administrator role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("The name of the custom rule"),
			),
			mcp.WithString("description",
				mcp.Description("The description of the custom rule"),
			),
			mcp.WithArray("categories",
				mcp.Required(),
				mcp.Description("A list of categories for the custom rule (e.g., security, cost-optimisation, reliability, performance-efficiency, operational-excellence, sustainability)"),
				mcp.Items(map[string]any{"type": "string"}),
			),
			mcp.WithString("riskLevel",
				mcp.Required(),
				mcp.Description("The risk level of the custom rule"),
				mcp.Enum("LOW", "MEDIUM", "HIGH", "VERY_HIGH", "EXTREME"),
			),
			mcp.WithString("provider",
				mcp.Required(),
				mcp.Description("The cloud provider for the custom rule"),
				mcp.Enum("aws", "azure", "gcp", "alibaba"),
			),
			mcp.WithString("resolutionReferenceLink",
				mcp.Description("A URL to the resolution reference documentation"),
			),
			mcp.WithString("remediationNote",
				mcp.Description("Notes on how to remediate the issue"),
			),
			mcp.WithBoolean("enabled",
				mcp.Required(),
				mcp.Description("Whether the custom rule is enabled"),
			),
			mcp.WithString("service",
				mcp.Required(),
				mcp.Description("The cloud service for the custom rule (e.g., S3, EC2)"),
			),
			mcp.WithString("resourceType",
				mcp.Required(),
				mcp.Description("The resource type for the custom rule"),
			),
			mcp.WithArray("attributes",
				mcp.Required(),
				mcp.Description("The attributes to extract from the resource for rule evaluation. Each attribute should have name, path, and optionally required fields."),
				mcp.Items(map[string]any{"type": "object"}),
			),
			mcp.WithArray("eventRules",
				mcp.Required(),
				mcp.Description("The event rules defining conditions for the custom rule. Each rule should have conditions and description fields."),
				mcp.Items(map[string]any{"type": "object"}),
			),
			mcp.WithString("slug",
				mcp.Description("The slug of the custom rule used to form the rule ID. If not provided, a random string will be used."),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			name, err := requiredValue[string]("name", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			categoriesRaw, ok := ctr.GetArguments()["categories"].([]any)
			if !ok || len(categoriesRaw) == 0 {
				return mcp.NewToolResultError("categories is required and must be a non-empty array"), nil
			}
			categories := make([]string, 0, len(categoriesRaw))
			for _, c := range categoriesRaw {
				if s, ok := c.(string); ok {
					categories = append(categories, s)
				}
			}

			riskLevel, err := requiredValue[string]("riskLevel", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			provider, err := requiredValue[string]("provider", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resolutionReferenceLink, err := optionalValue[string]("resolutionReferenceLink", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			remediationNote, err := optionalValue[string]("remediationNote", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			enabled, err := requiredValue[bool]("enabled", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			service, err := requiredValue[string]("service", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resourceType, err := requiredValue[string]("resourceType", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			attributes, ok := ctr.GetArguments()["attributes"].([]any)
			if !ok {
				return mcp.NewToolResultError("attributes is required and must be an array"), nil
			}

			eventRules, ok := ctr.GetArguments()["eventRules"].([]any)
			if !ok {
				return mcp.NewToolResultError("eventRules is required and must be an array"), nil
			}

			slug, err := optionalValue[string]("slug", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			input := v1client.CustomRuleInput{
				Name:                    name,
				Description:             description,
				Categories:              categories,
				RiskLevel:               riskLevel,
				Provider:                provider,
				ResolutionReferenceLink: resolutionReferenceLink,
				RemediationNote:         remediationNote,
				Enabled:                 enabled,
				Service:                 service,
				ResourceType:            resourceType,
				Attributes:              attributes,
				EventRules:              eventRules,
				Slug:                    slug,
			}

			resp, err := client.CloudPostureCreateCustomRule(input)
			return handleStatusResponse(resp, err, http.StatusCreated, "failed to create custom rule")
		},
	}
}

func toolCloudPostureCustomRuleUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_custom_rule_update",
			mcp.WithDescription("(Beta) Updates the specified custom rule. Requires Master Administrator role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("ruleId",
				mcp.Required(),
				mcp.Description("The Cloud Risk Management ID of the custom rule to update"),
			),
			mcp.WithString("name",
				mcp.Description("The name of the custom rule"),
			),
			mcp.WithString("description",
				mcp.Description("The description of the custom rule"),
			),
			mcp.WithArray("categories",
				mcp.Description("A list of categories for the custom rule"),
				mcp.Items(map[string]any{"type": "string"}),
			),
			mcp.WithString("riskLevel",
				mcp.Description("The risk level of the custom rule"),
				mcp.Enum("LOW", "MEDIUM", "HIGH", "VERY_HIGH", "EXTREME"),
			),
			mcp.WithString("provider",
				mcp.Description("The cloud provider for the custom rule"),
				mcp.Enum("aws", "azure", "gcp", "alibaba"),
			),
			mcp.WithString("resolutionReferenceLink",
				mcp.Description("A URL to the resolution reference documentation"),
			),
			mcp.WithString("remediationNote",
				mcp.Description("Notes on how to remediate the issue"),
			),
			mcp.WithBoolean("enabled",
				mcp.Description("Whether the custom rule is enabled"),
			),
			mcp.WithString("service",
				mcp.Description("The cloud service for the custom rule"),
			),
			mcp.WithString("resourceType",
				mcp.Description("The resource type for the custom rule"),
			),
			mcp.WithArray("attributes",
				mcp.Description("The attributes to extract from the resource for rule evaluation"),
				mcp.Items(map[string]any{"type": "object"}),
			),
			mcp.WithArray("eventRules",
				mcp.Description("The event rules defining conditions for the custom rule"),
				mcp.Items(map[string]any{"type": "object"}),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			ruleId, err := requiredValue[string]("ruleId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			name, err := optionalValue[string]("name", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			description, err := optionalValue[string]("description", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var categories []string
			if categoriesRaw, ok := ctr.GetArguments()["categories"].([]any); ok && len(categoriesRaw) > 0 {
				categories = make([]string, 0, len(categoriesRaw))
				for _, c := range categoriesRaw {
					if s, ok := c.(string); ok {
						categories = append(categories, s)
					}
				}
			}

			riskLevel, err := optionalValue[string]("riskLevel", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			provider, err := optionalValue[string]("provider", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resolutionReferenceLink, err := optionalValue[string]("resolutionReferenceLink", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			remediationNote, err := optionalValue[string]("remediationNote", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			enabled, err := optionalPointerValue[bool]("enabled", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			service, err := optionalValue[string]("service", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resourceType, err := optionalValue[string]("resourceType", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			var attributes []any
			if attrsRaw, ok := ctr.GetArguments()["attributes"].([]any); ok {
				attributes = attrsRaw
			}

			var eventRules []any
			if rulesRaw, ok := ctr.GetArguments()["eventRules"].([]any); ok {
				eventRules = rulesRaw
			}

			input := v1client.CustomRuleUpdateInput{
				Name:                    name,
				Description:             description,
				Categories:              categories,
				RiskLevel:               riskLevel,
				Provider:                provider,
				ResolutionReferenceLink: resolutionReferenceLink,
				RemediationNote:         remediationNote,
				Enabled:                 enabled,
				Service:                 service,
				ResourceType:            resourceType,
				Attributes:              attributes,
				EventRules:              eventRules,
			}

			resp, err := client.CloudPostureUpdateCustomRule(ruleId, input)
			return handleStatusResponse(resp, err, http.StatusNoContent, "failed to update custom rule")
		},
	}
}

func toolCloudPostureCustomRuleDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_custom_rule_delete",
			mcp.WithDescription("(Beta) Deletes the specified custom rule permanently. Requires Master Administrator role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("ruleId",
				mcp.Required(),
				mcp.Description("The Cloud Risk Management ID of the custom rule to delete"),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			ruleId, err := requiredValue[string]("ruleId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CloudPostureDeleteCustomRule(ruleId)
			return handleStatusResponse(resp, err, http.StatusNoContent, "failed to delete custom rule")
		},
	}
}

func toolCloudPostureCustomRuleTest(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cloud_posture_custom_rule_test",
			mcp.WithDescription("(Beta) Tests the provided custom rule configuration against the specified Cloud Risk Management account or mock resource data. Requires Master Administrator role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("accountId",
				mcp.Description("The Cloud Risk Management account ID to test against. Either accountId or resource must be provided."),
			),
			mcp.WithObject("configuration",
				mcp.Required(),
				mcp.Description("The custom rule configuration to test. Should include name, description, categories, riskLevel, provider, service, resourceType, attributes, and eventRules."),
			),
			mcp.WithObject("resource",
				mcp.Description("Mock resource data to test the rule against. Either accountId or resource must be provided."),
			),
		),
		Handler: func(ctx context.Context, ctr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := optionalValue[string]("accountId", ctr.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			configuration, ok := ctr.GetArguments()["configuration"]
			if !ok {
				return mcp.NewToolResultError("configuration is required"), nil
			}

			resource := ctr.GetArguments()["resource"]

			if accountId == "" && resource == nil {
				return mcp.NewToolResultError("either accountId or resource must be provided"), nil
			}

			input := v1client.CustomRuleTestInput{
				AccountId:     accountId,
				Configuration: configuration,
				Resource:      resource,
			}

			resp, err := client.CloudPostureTestCustomRule(input)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to test custom rule")
		},
	}
}
