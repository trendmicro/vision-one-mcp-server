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

var ToolsetsReadOnlyCloudPostureBeta = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudPostureCustomRulesList,
	toolCloudPostureCustomRuleGet,
	toolCloudPostureCustomRuleTest,
}

var ToolsetsWriteCloudPostureBeta = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCloudPostureCustomRuleCreate,
	toolCloudPostureCustomRuleUpdate,
	toolCloudPostureCustomRuleDelete,
}

// Schema definitions for custom rules (json-rules-engine based)
var customRuleConditionItemSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"fact": map[string]any{
			"type":        "string",
			"description": "The attribute name defined in the attributes array",
		},
		"operator": map[string]any{
			"type":        "string",
			"description": "Comparison operator",
			"enum":        []string{"equal", "notEqual", "lessThan", "lessThanInclusive", "greaterThan", "greaterThanInclusive", "in", "notIn", "contains", "doesNotContain", "dateComparison"},
		},
		"value": map[string]any{
			"description": "Value to compare against. For dateComparison, use object: {days: number, operator: 'within'|'before'|'after'}",
		},
	},
	"required": []string{"fact", "operator", "value"},
}

var customRuleConditionsSchema = map[string]any{
	"type":        "object",
	"description": "Boolean logic container. Use exactly one of: 'any', 'all', or 'not' at root level",
	"properties": map[string]any{
		"any": map[string]any{
			"type":        "array",
			"description": "Array of conditions where at least one must be true",
			"items":       customRuleConditionItemSchema,
		},
		"all": map[string]any{
			"type":        "array",
			"description": "Array of conditions where all must be true",
			"items":       customRuleConditionItemSchema,
		},
		"not": map[string]any{
			"type":        "object",
			"description": "Single condition to negate",
			"properties":  customRuleConditionItemSchema["properties"],
			"required":    customRuleConditionItemSchema["required"],
		},
	},
}

var customRuleAttributeSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"name": map[string]any{
			"type":        "string",
			"description": "The attribute name, referenced as 'fact' in eventRules conditions",
		},
		"path": map[string]any{
			"type":        "string",
			"description": "JSONPath to extract value from resource data. Examples: 'data.AccessKeyMetadata[0].CreateDate' for single value, 'data.sourceRanges[*]' for array elements, 'data.allowed[*].ports[*]' for nested arrays",
		},
		"required": map[string]any{
			"type":        "boolean",
			"description": "If true, rule evaluation fails when attribute path doesn't exist",
		},
	},
	"required": []string{"name", "path"},
}

var customRuleEventRuleSchema = map[string]any{
	"type": "object",
	"properties": map[string]any{
		"conditions": customRuleConditionsSchema,
		"description": map[string]any{
			"type":        "string",
			"description": "Message shown when rule condition triggers (resource is non-compliant)",
		},
	},
	"required": []string{"conditions", "description"},
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
			mcp.WithDescription("(Beta) Creates a custom rule for your organization. Enabled custom rules are immediately available to all your cloud accounts. Requires Master Administrator role."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name",
				mcp.Required(),
				mcp.Description("The name of the custom rule"),
			),
			mcp.WithString("description",
				mcp.Required(),
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
				mcp.Enum("aws", "azure", "gcp", "oci", "alibabaCloud"),
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
				mcp.Description("The attributes to extract from the resource for rule evaluation using JSONPath."),
				mcp.Items(customRuleAttributeSchema),
			),
			mcp.WithArray("eventRules",
				mcp.Required(),
				mcp.Description("Rules using json-rules-engine. Rule passes (resource compliant) when conditions evaluate to true."),
				mcp.Items(customRuleEventRuleSchema),
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

			description, err := requiredValue[string]("description", ctr.GetArguments())
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
				mcp.Enum("aws", "azure", "gcp", "oci", "alibabaCloud"),
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
				mcp.Description("The attributes to extract from the resource for rule evaluation using JSONPath."),
				mcp.Items(customRuleAttributeSchema),
			),
			mcp.WithArray("eventRules",
				mcp.Description("Rules using json-rules-engine. Rule passes (resource compliant) when conditions evaluate to true."),
				mcp.Items(customRuleEventRuleSchema),
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
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId",
				mcp.Description("The Cloud Risk Management account ID to test against. Either accountId or resource must be provided."),
			),
			mcp.WithObject("configuration",
				mcp.Required(),
				mcp.Description("The custom rule configuration to test."),
				mcp.Properties(map[string]any{
					"name": map[string]any{
						"type":        "string",
						"description": "The name of the custom rule",
					},
					"description": map[string]any{
						"type":        "string",
						"description": "The description of the custom rule",
					},
					"categories": map[string]any{
						"type":        "array",
						"items":       map[string]any{"type": "string"},
						"description": "Categories: security, cost-optimisation, reliability, performance-efficiency, operational-excellence, sustainability",
					},
					"riskLevel": map[string]any{
						"type":        "string",
						"enum":        []string{"LOW", "MEDIUM", "HIGH", "VERY_HIGH", "EXTREME"},
						"description": "The risk level of the custom rule",
					},
					"provider": map[string]any{
						"type":        "string",
						"enum":        []string{"aws", "azure", "gcp", "oci", "alibabaCloud"},
						"description": "The cloud provider",
					},
					"enabled": map[string]any{
						"type":        "boolean",
						"description": "Whether the custom rule is enabled",
					},
					"service": map[string]any{
						"type":        "string",
						"description": "The cloud service (e.g., IAM, S3, EC2)",
					},
					"resourceType": map[string]any{
						"type":        "string",
						"description": "The resource type (e.g., iam-user, s3-bucket)",
					},
					"remediationNote": map[string]any{
						"type":        "string",
						"description": "Notes on how to remediate the issue",
					},
					"attributes": map[string]any{
						"type":        "array",
						"description": "Attributes to extract using JSONPath",
						"items":       customRuleAttributeSchema,
					},
					"eventRules": map[string]any{
						"type":        "array",
						"description": "Rules using json-rules-engine",
						"items":       customRuleEventRuleSchema,
					},
				}),
			),
			mcp.WithObject("resource",
				mcp.Description("Mock resource data to test the rule against. Either accountId or resource must be provided."),
				mcp.Properties(map[string]any{
					"metadata": map[string]any{
						"type":        "object",
						"description": "Resource metadata containing the data fields matching the resourceType schema. Place all resource data fields directly inside metadata (e.g., for iam-user: {\"metadata\": {\"AccessKeyMetadata\": [...]}})",
					},
				}),
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
