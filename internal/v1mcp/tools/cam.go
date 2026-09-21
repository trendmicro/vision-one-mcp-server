package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyCAM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCAMAwsAccountsList,
	toolCAMAwsAccountGet,
	toolCAMGcpAccountsList,
	toolCAMGcpAccountGet,
	toolCAMAlibabaAccountsList,
	toolCAMAlibabaAccountGet,
	toolCamAlibabaAccountsGenerateTerraformPackage,
	toolCamAwsAccountsFeaturesList,
	toolCamAwsAccountsGenerateCfnTemplateLinks,
	toolCamAwsAccountsGenerateTerraformPackage,
	toolCamAzureSubscriptionsList,
	toolCamAzureSubscriptionsGenerateMgmtGroupTerraformPackage,
	toolCamAzureSubscriptionsGenerateTerraformPackage,
	toolCamAzureSubscriptionGet,
	toolCamGcpProjectsGenerateTerraformPackage,
	toolCamOciCompartmentsList,
	toolCamOciCompartmentsGenerateTerraformPackage,
	toolCamOciCompartmentGet,
}

var ToolsetsWriteCAM = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolCamAlibabaAccountsCreate,
	toolCamAlibabaAccountDelete,
	toolCamAlibabaAccountUpdate,
	toolCamAwsAccountsCreate,
	toolCamAwsAccountDelete,
	toolCamAwsAccountUpdate,
	toolCamAzureSubscriptionsCreate,
	toolCamAzureSubscriptionDelete,
	toolCamAzureSubscriptionUpdate,
	toolCamGcpProjectsCreate,
	toolCamGcpProjectDelete,
	toolCamGcpProjectUpdate,
	toolCamOciCompartmentsCreate,
	toolCamOciCompartmentDelete,
	toolCamOciCompartmentUpdate,
}

func toolCAMAwsAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_accounts_list",
			mcp.WithDescription("List AWS accounts managed by Cloud Accounts Management"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(camTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterAWSAccounts)),
			mcp.WithString("nextBatchToken", mcp.Description("Token used to retrieve the next page of results")),
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

			nextBatchToken, err := optionalValue[string]("nextBatchToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:            top,
				NextBatchToken: nextBatchToken,
			}

			resp, err := client.CAMListAWSAccounts(filter, qp)
			if err != nil {
				return nil, err
			}

			return handleStatusResponse(resp, err, http.StatusOK, "failed to list cam aws accounts")
		},
	}
}

func toolCAMAwsAccountGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_account_get",
			mcp.WithDescription("Get the details of an AWS account managed by Cloud Account Management"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			resp, err := client.CAMGetAWSAccount(accountId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get aws account details")
		},
	}
}

func toolCAMGcpAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_gcp_accounts_list",
			mcp.WithDescription("List Google Cloud Projects managed by Cloud Account Management"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(camTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.CAMListGCPProjectsFilterDescription)),
			mcp.WithString("nextBatchToken", mcp.Description("Token used to retrieve the next page of results")),
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

			nextBatchToken, err := optionalValue[string]("nextBatchToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:            top,
				NextBatchToken: nextBatchToken,
			}

			resp, err := client.CAMListGCPAccounts(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list gcp accounts")

		},
	}
}

func toolCAMGcpAccountGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_gcp_account_get",
			mcp.WithDescription("Get the details of a GCP project managed by Cloud Account Manangement"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CAMGetGCPAccountDetails(accountId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get gcp project details")
		},
	}
}

func toolCAMAlibabaAccountsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_alibaba_accounts_list",
			mcp.WithDescription("Displays all Alibaba Cloud accounts connected to Trend Vision One in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top",
				mcp.Description(tooldescriptions.DefaultTop),
				mcp.Enum(camTop()...),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterAlibabaAccounts)),
			mcp.WithString("nextBatchToken", mcp.Description("Token used to retrieve the next page of results")),
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

			nextBatchToken, err := optionalValue[string]("nextBatchToken", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				Top:            top,
				NextBatchToken: nextBatchToken,
			}

			resp, err := client.CAMListAlibabaAccounts(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list alibaba accounts")
		},
	}
}

func toolCAMAlibabaAccountGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_alibaba_account_get",
			mcp.WithDescription("Get the details of an Alibaba account managed by Cloud Account Manangement"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("accountId", mcp.Required()),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			accountId, err := requiredValue[string]("accountId", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CAMGetAlibabaAccountDetails(accountId)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get gcp project details")
		},
	}
}

func camTop() []string {
	return []string{
		"25",
		"50",
		"100",
		"500",
		"1000",
		"5000",
	}
}

func toolCamAlibabaAccountsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_alibaba_accounts_create",
			mcp.WithDescription("Add Alibaba Cloud account. Adds and connects the specified Alibaba Cloud account to TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("alibabaRegion", mcp.Required(), mcp.Description("The region of the Terraform backend where the state files are stored. > Important > > - The default region is based on your TrendAI Vision One™ region.")),
			mcp.WithString("roleArn", mcp.Required(), mcp.Description("The Alibaba Cloud resource name (ARN) of the user role for TrendAI Vision One™.")),
			mcp.WithString("oidcProviderId", mcp.Required(), mcp.Description("The ID of the Alibaba Cloud OpenID Connect (OIDC) provider.")),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the Alibaba Cloud account to be used in Cloud Account Management.")),
			mcp.WithString("description", mcp.Description("The description of the Alibaba Cloud account.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security service associated with the connected Alibaba Cloud account."), mcp.Items(map[string]any{"properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the security services connected to the system.", "enum": []string{"workload"}, "type": "string"}}, "required": []string{"name", "instanceIds"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "alibabaRegion", Kind: kindString, Required: true}, {Name: "roleArn", Kind: kindString, Required: true}, {Name: "oidcProviderId", Kind: kindString, Required: true}, {Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "connectedSecurityServices", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAlibabaAccountsCreate(params)
			return handleResponse(resp, err, "failed to add Alibaba Cloud account")
		},
	}
}

func toolCamAlibabaAccountsGenerateTerraformPackage(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_alibaba_accounts_generate_terraform_package",
			mcp.WithDescription("Get Alibaba Cloud Terraform template. Generates and downloads the Alibaba Cloud Terraform template."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("cloudAccountName", mcp.Required(), mcp.Description("The name of the Alibaba Cloud account.")),
			mcp.WithString("cloudAccountDescription", mcp.Description("The description of the Alibaba Cloud account.")),
			mcp.WithString("alibabaRegion", mcp.Required(), mcp.Description("The region of Terraform backend where the state files are stored.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security service associated with the connected Alibaba Cloud account."), mcp.Items(map[string]any{"properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the security services connected to the system.", "enum": []string{"workload"}, "type": "string"}}, "required": []string{"name", "instanceIds"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "cloudAccountName", Kind: kindString, Required: true}, {Name: "cloudAccountDescription", Kind: kindString}, {Name: "alibabaRegion", Kind: kindString, Required: true}, {Name: "connectedSecurityServices", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAlibabaAccountsGenerateTerraformPackage(params)
			return handleResponse(resp, err, "failed to get Alibaba Cloud Terraform template")
		},
	}
}

func toolCamAlibabaAccountDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_alibaba_account_delete",
			mcp.WithDescription("Remove Alibaba Cloud account. Removes and disconnects the specified Alibaba Cloud account from TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Alibaba Cloud account.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAlibabaAccountDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to remove Alibaba Cloud account")
		},
	}
}

func toolCamAlibabaAccountUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_alibaba_account_update",
			mcp.WithDescription("Modify Alibaba Cloud account details. Changes the settings for the specified Alibaba Cloud account in TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Alibaba Cloud account.")),
			mcp.WithString("ifMatch", mcp.Description("Entity tag in MD5 format to perform version control on the resource to be updated. This is returned in response header ETag of the v3 GET method. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("name", mcp.Description("The name of the Alibaba Cloud account to be used in Cloud Account Management.")),
			mcp.WithString("description", mcp.Description("The description of the Alibaba Cloud account. The default value is an empty string if the field is omitted.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAlibabaAccountUpdate(id, params)
			return handleResponse(resp, err, "failed to modify Alibaba Cloud account details")
		},
	}
}

func toolCamAwsAccountsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_accounts_create",
			mcp.WithDescription("Add an account. Adds and connects the specified AWS account to TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("roleArn", mcp.Required(), mcp.Description("The Amazon Resource Name (ARN) of the user role for TrendAI Vision One™.")),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the AWS account.")),
			mcp.WithString("description", mcp.Description("The description of the AWS account.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security service associated with the connected AWS account."), mcp.Items(map[string]any{"description": "Server & Workload Protection service configuration", "properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the associated security service.", "enum": []string{"workload"}, "type": "string"}, "regions": map[string]any{"description": "The AWS regions where Server & Workload Protection will be deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"name", "instanceIds"}, "type": "object"})),
			mcp.WithArray("customTags", mcp.Description("The custom tags attached to the resources created by TrendAI Vision One™."), mcp.Items(map[string]any{"properties": map[string]any{"key": map[string]any{"description": "The key of the custom tag", "type": "string"}, "value": map[string]any{"description": "The value of the custom tag", "type": "string"}}, "required": []string{"key", "value"}, "type": "object"})),
			mcp.WithArray("features", mcp.Description("The list of features you selected and the corresponding regions for deployment."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature. For the up-to-date list of available feature IDs, call GET /beta/cam/awsAccounts/features or GET...", "type": "string"}, "regions": map[string]any{"description": "The regions where features are deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "roleArn", Kind: kindString, Required: true}, {Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "connectedSecurityServices", Kind: kindArray}, {Name: "customTags", Kind: kindArray}, {Name: "features", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAwsAccountsCreate(params)
			return handleResponse(resp, err, "failed to add an account")
		},
	}
}

func toolCamAwsAccountsFeaturesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_accounts_features_list",
			mcp.WithDescription("Get feature list. Displays the available AWS account features TrendAI Vision One™ offers."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.CamAwsAccountsFeaturesList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get feature list")
		},
	}
}

func toolCamAwsAccountsGenerateCfnTemplateLinks(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_accounts_generate_cfn_template_links",
			mcp.WithDescription("Generate AWS CloudFormation template. Generates and downloads the AWS CloudFormation template."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("awsRegion", mcp.Description("The AWS region for template deployment.")),
			mcp.WithArray("features", mcp.Description("The list of features you selected and the corresponding regions for deployment."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature. For the up-to-date list of available feature IDs, call GET /beta/cam/awsAccounts/features or GET...", "type": "string"}, "regions": map[string]any{"description": "The regions where features are deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"id"}, "type": "object"})),
			mcp.WithString("templateAccountType", mcp.Description("The type of template generation for a single account scope or organization scope."), mcp.Enum("single", "organization")),
			mcp.WithString("awsAccountName", mcp.Description("The cloud account name assigned to the created account.")),
			mcp.WithString("awsAccountDescription", mcp.Description("The cloud account description assigned to the created account.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security service associated with the connected AWS account."), mcp.Items(map[string]any{"description": "Server & Workload Protection service configuration", "properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the associated security service.", "enum": []string{"workload"}, "type": "string"}, "regions": map[string]any{"description": "The AWS regions where Server & Workload Protection will be deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"name", "instanceIds"}, "type": "object"})),
			mcp.WithArray("customTags", mcp.Description("The custom tags attached to the resources created by TrendAI Vision One™."), mcp.Items(map[string]any{"properties": map[string]any{"key": map[string]any{"description": "The key of the custom tag", "type": "string"}, "value": map[string]any{"description": "The value of the custom tag", "type": "string"}}, "required": []string{"key", "value"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "awsRegion", Kind: kindString}, {Name: "features", Kind: kindArray}, {Name: "templateAccountType", Kind: kindString}, {Name: "awsAccountName", Kind: kindString}, {Name: "awsAccountDescription", Kind: kindString}, {Name: "connectedSecurityServices", Kind: kindArray}, {Name: "customTags", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAwsAccountsGenerateCfnTemplateLinks(params)
			return handleResponse(resp, err, "failed to generate AWS CloudFormation template")
		},
	}
}

func toolCamAwsAccountsGenerateTerraformPackage(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_accounts_generate_terraform_package",
			mcp.WithDescription("Generate AWS Terraform template. Generates and downloads the package as a ZIP file for Terraform deployment."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("awsRegion", mcp.Description("The AWS region for template deployment.")),
			mcp.WithArray("features", mcp.Description("The list of features you selected and the corresponding regions for deployment."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature. For the up-to-date list of available feature IDs, call GET /beta/cam/awsAccounts/features or GET...", "type": "string"}, "regions": map[string]any{"description": "The regions where features are deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"id"}, "type": "object"})),
			mcp.WithString("awsAccountName", mcp.Description("The cloud account name that will be assigned to the account created.")),
			mcp.WithString("awsAccountDescription", mcp.Description("The cloud account description that will be assigned to the account created.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security service associated with the connected AWS account."), mcp.Items(map[string]any{"description": "Server & Workload Protection service configuration", "properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the associated security service.", "enum": []string{"workload"}, "type": "string"}, "regions": map[string]any{"description": "The AWS regions where Server & Workload Protection will be deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"name", "instanceIds"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "awsRegion", Kind: kindString}, {Name: "features", Kind: kindArray}, {Name: "awsAccountName", Kind: kindString}, {Name: "awsAccountDescription", Kind: kindString}, {Name: "connectedSecurityServices", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAwsAccountsGenerateTerraformPackage(params)
			return handleResponse(resp, err, "failed to generate AWS Terraform template")
		},
	}
}

func toolCamAwsAccountDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_account_delete",
			mcp.WithDescription("Remove account. Removes and disconnects the specified AWS account from TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the AWS account.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAwsAccountDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to remove account")
		},
	}
}

func toolCamAwsAccountUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_aws_account_update",
			mcp.WithDescription("Modify account details. Modifies the settings of the specified AWS account in TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the AWS account.")),
			mcp.WithString("ifMatch", mcp.Description("Entity tag in MD5 format to perform version control on the resource to be updated. This is returned in response header ETag of the v3 GET method. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("roleArn", mcp.Description("The Amazon Resource Name (ARN) of the user role for TrendAI Vision One™.")),
			mcp.WithString("name", mcp.Description("The name of the AWS account.")),
			mcp.WithString("description", mcp.Description("The description of the AWS account.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security services associated with the connected AWS account."), mcp.Items(map[string]any{"properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the associated security service.", "enum": []string{"workload"}, "type": "string"}, "regions": map[string]any{"description": "The AWS regions where Server & Workload Protection will be deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"instanceIds", "name"}, "type": "object"})),
			mcp.WithArray("customTags", mcp.Description("The custom tags attached to the resources created by TrendAI Vision One™."), mcp.Items(map[string]any{"properties": map[string]any{"key": map[string]any{"description": "The key of the custom tag", "type": "string"}, "value": map[string]any{"description": "The value of the custom tag", "type": "string"}}, "required": []string{"key", "value"}, "type": "object"})),
			mcp.WithArray("features", mcp.Description("The list of features you selected and the corresponding regions for deployment."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature. For the up-to-date list of available feature IDs, call GET /beta/cam/awsAccounts/features or GET...", "type": "string"}, "regions": map[string]any{"description": "The regions where features are deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "roleArn", Kind: kindString}, {Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}, {Name: "connectedSecurityServices", Kind: kindArray}, {Name: "customTags", Kind: kindArray}, {Name: "features", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAwsAccountUpdate(id, params)
			return handleResponse(resp, err, "failed to modify account details")
		},
	}
}

func toolCamAzureSubscriptionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_azure_subscriptions_list",
			mcp.WithDescription("Get connected Azure subscriptions. Displays all connected Azure subscriptions in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100, 500, 1000, 5000. Default: 25."), mcp.Enum("25", "50", "100", "500", "1000", "5000")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the connected Azure accounts. Supported fields: name - The name of the Azure subscription - Any value; state - The state of the Azure subscription - managed, outdated, failed; id - The ID of the Azure subscription. - Any value; tenantId - The tenant ID of the Azure subscription. - Any value; Supported operators: eq - Operator 'equal to'; and - Operator 'and'; or - Operator 'or'; not - Operator 'not'; ( ) - Symbols for grouping operands with their correct operator. contains - Operator that allows you to search for a specified string in a field; Note: Include this parameter in every request that generates paginated output.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAzureSubscriptionsList(params)
			return handleResponse(resp, err, "failed to get connected Azure subscriptions")
		},
	}
}

func toolCamAzureSubscriptionsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_azure_subscriptions_create",
			mcp.WithDescription("Add an Azure subscription. Adds and connects the specified Azure subscription to TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("tenantId", mcp.Required(), mcp.Description("The tenant ID of the Azure subscription.")),
			mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("The target subscription ID for TrendAI Vision One™.")),
			mcp.WithString("applicationId", mcp.Required(), mcp.Description("The ID of an Azure App Registration resource created by the deployed template.")),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the Azure subscription used in Cloud Account Management.")),
			mcp.WithString("description", mcp.Description("The description of the Azure subscription.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security service associated with the connected Azure subscription."), mcp.Items(map[string]any{"properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the associated security service.", "enum": []string{"workload"}, "type": "string"}}, "required": []string{"name", "instanceIds"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "tenantId", Kind: kindString, Required: true}, {Name: "subscriptionId", Kind: kindString, Required: true}, {Name: "applicationId", Kind: kindString, Required: true}, {Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "connectedSecurityServices", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAzureSubscriptionsCreate(params)
			return handleResponse(resp, err, "failed to add an Azure subscription")
		},
	}
}

func toolCamAzureSubscriptionsGenerateMgmtGroupTerraformPackage(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_azure_subscriptions_generate_mgmt_group_terraform_package",
			mcp.WithDescription("Generate Azure management group Terraform deployment package. Generates a pre-configured Azure Terraform deployment package (.zip) for an Azure management group deployment across multiple subscriptions and returns a pre-signed S3 URL for download."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("managementGroupId", mcp.Required(), mcp.Description("The ID of the Azure management group.")),
			mcp.WithString("azureRegion", mcp.Required(), mcp.Description("The Azure region where Cloud Accounts resources are deployed.")),
			mcp.WithString("azureSubscriptionDescription", mcp.Description("The description applied to all Azure subscriptions in the management group deployment.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security service associated with the connected cloud account."), mcp.Items(map[string]any{"properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the associated security service.", "enum": []string{"workload"}, "type": "string"}}, "required": []string{"name", "instanceIds"}, "type": "object"})),
			mcp.WithArray("excludedSubscriptionIds", mcp.Description("A list of Azure subscription IDs to exclude from the management group deployment."), mcp.Items(map[string]any{"type": "string"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "managementGroupId", Kind: kindString, Required: true}, {Name: "azureRegion", Kind: kindString, Required: true}, {Name: "azureSubscriptionDescription", Kind: kindString}, {Name: "connectedSecurityServices", Kind: kindArray}, {Name: "excludedSubscriptionIds", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAzureSubscriptionsGenerateMgmtGroupTerraformPackage(params)
			return handleResponse(resp, err, "failed to generate Azure management group Terraform deployment package")
		},
	}
}

func toolCamAzureSubscriptionsGenerateTerraformPackage(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_azure_subscriptions_generate_terraform_package",
			mcp.WithDescription("Generate Azure Terraform deployment package. Generates a pre-configured Azure Terraform deployment package (.zip) for a single Azure subscription and returns a pre-signed S3 URL for download."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("azureSubscriptionName", mcp.Required(), mcp.Description("The name of the cloud account.")),
			mcp.WithString("azureSubscriptionDescription", mcp.Description("The description of the cloud account.")),
			mcp.WithString("subscriptionId", mcp.Required(), mcp.Description("The subscription ID of the Azure account.")),
			mcp.WithArray("connectedSecurityServices", mcp.Description("The security service associated with the connected cloud account."), mcp.Items(map[string]any{"properties": map[string]any{"instanceIds": map[string]any{"description": "The instance ID of the associated security service.", "items": map[string]any{"type": "string"}, "type": "array"}, "name": map[string]any{"description": "The name of the associated security service.", "enum": []string{"workload"}, "type": "string"}}, "required": []string{"name", "instanceIds"}, "type": "object"})),
			mcp.WithString("azureRegion", mcp.Required(), mcp.Description("The Azure region where Cloud Accounts resources are deployed.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "azureSubscriptionName", Kind: kindString, Required: true}, {Name: "azureSubscriptionDescription", Kind: kindString}, {Name: "subscriptionId", Kind: kindString, Required: true}, {Name: "connectedSecurityServices", Kind: kindArray}, {Name: "azureRegion", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAzureSubscriptionsGenerateTerraformPackage(params)
			return handleResponse(resp, err, "failed to generate Azure Terraform deployment package")
		},
	}
}

func toolCamAzureSubscriptionDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_azure_subscription_delete",
			mcp.WithDescription("Remove Azure subscription. Removes and disconnects the specified Azure subscription from TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Azure subscription.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAzureSubscriptionDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to remove Azure subscription")
		},
	}
}

func toolCamAzureSubscriptionGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_azure_subscription_get",
			mcp.WithDescription("Get Azure subscription details. Displays information about the specified Azure subscription."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Azure subscription.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAzureSubscriptionGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get Azure subscription details")
		},
	}
}

func toolCamAzureSubscriptionUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_azure_subscription_update",
			mcp.WithDescription("Modify Azure subscription details. Modifies the details of the specified Azure subscription in TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Azure subscription.")),
			mcp.WithString("ifMatch", mcp.Description("Entity tag in MD5 format to perform version control on the resource to be updated. This is returned in response header ETag of the v3 GET method. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("name", mcp.Description("The name of the Azure subscription used in Cloud Account Management.")),
			mcp.WithString("description", mcp.Description("The description of the Azure subscription.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamAzureSubscriptionUpdate(id, params)
			return handleResponse(resp, err, "failed to modify Azure subscription details")
		},
	}
}

func toolCamGcpProjectsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_gcp_projects_create",
			mcp.WithDescription("Add Google Cloud project. Adds and connects the specified Google Cloud project to TrendAI Vision One™. Used in cases where template deployment fails when using the template from GET /v3.0/gcpProjects/generateTerraformPackage."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("workloadIdentityPoolId", mcp.Required(), mcp.Description("The ID of the Google Cloud workload identity pool.")),
			mcp.WithString("oidcProviderId", mcp.Required(), mcp.Description("The ID of the Google Cloud OpenID Connect (OIDC) provider.")),
			mcp.WithString("serviceAccountId", mcp.Required(), mcp.Description("The ID of the Google Cloud service account.")),
			mcp.WithString("projectNumber", mcp.Required(), mcp.Description("The Google Cloud project number.")),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the Google Cloud project to be used in Cloud Accounts.")),
			mcp.WithString("description", mcp.Description("The description of the Google Cloud project. If there is no description, the field is empty.")),
			mcp.WithString("gcpRegion", mcp.Required(), mcp.Description("The region where Cloud Accounts is deployed."), mcp.Enum("us-west1", "us-west2", "us-west3", "us-west4", "us-central1", "us-east1", "us-east4", "northamerica-northeast1", "southamerica-east1", "europe-west1", "europe-west2", "europe-west3", "europe-west6", "europe-central2", "asia-south1", "asia-southeast1", "asia-southeast2", "asia-east1", "asia-east", "asia-northeast1", "asia-northeast2", "asia-northeast3", "australia-southeast1")),
			mcp.WithArray("features", mcp.Description("The list of selected features and the corresponding regions for deployment."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature.", "enum": []string{"cloud-sentry"}, "type": "string"}, "regions": map[string]any{"description": "The regions where the features are deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "workloadIdentityPoolId", Kind: kindString, Required: true}, {Name: "oidcProviderId", Kind: kindString, Required: true}, {Name: "serviceAccountId", Kind: kindString, Required: true}, {Name: "projectNumber", Kind: kindString, Required: true}, {Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "gcpRegion", Kind: kindString, Required: true}, {Name: "features", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamGcpProjectsCreate(params)
			return handleResponse(resp, err, "failed to add Google Cloud project")
		},
	}
}

func toolCamGcpProjectsGenerateTerraformPackage(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_gcp_projects_generate_terraform_package",
			mcp.WithDescription("Get Google Cloud project Terraform template package. Generates and downloads a .zip file containing the Google Cloud platform Terraform template package. The Terraform template contains the role and policies needed for TrendAI Vision One™ to access the resources stored in your Google Cloud projects."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("gcpProjectName", mcp.Required(), mcp.Description("The name of the Google Cloud project. Use the Google Cloud project name in the name field for POST /v3.0/cam/gcpProjects.")),
			mcp.WithString("gcpProjectDescription", mcp.Description("The description of the Google Cloud project. If there is no description, the field value is empty. Use the Google Cloud project description in the description field for 'POST /v3.0/cam/gcpProjects.")),
			mcp.WithString("gcpRegion", mcp.Required(), mcp.Description("The region where Cloud Accounts is deployed."), mcp.Enum("us-west1", "us-west2", "us-west3", "us-west4", "us-central1", "us-east1", "us-east4", "northamerica-northeast1", "southamerica-east1", "europe-west1", "europe-west2", "europe-west3", "europe-west6", "europe-central2", "asia-south1", "asia-southeast1", "asia-southeast2", "asia-east1", "asia-east", "asia-northeast1", "asia-northeast2", "asia-northeast3", "australia-southeast1")),
			mcp.WithArray("features", mcp.Description("The list of selected features and the corresponding regions for deployment."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature.", "enum": []string{"cloud-sentry"}, "type": "string"}, "regions": map[string]any{"description": "The regions where the features are deployed.", "items": map[string]any{"type": "string"}, "type": "array"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "gcpProjectName", Kind: kindString, Required: true}, {Name: "gcpProjectDescription", Kind: kindString}, {Name: "gcpRegion", Kind: kindString, Required: true}, {Name: "features", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamGcpProjectsGenerateTerraformPackage(params)
			return handleResponse(resp, err, "failed to get Google Cloud project Terraform template package")
		},
	}
}

func toolCamGcpProjectDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_gcp_project_delete",
			mcp.WithDescription("Remove Google Cloud project. Removes and disconnects the specified Google Cloud project from TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The Google Cloud project number Cloud Account Management uses for managing the connected Google Cloud project.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamGcpProjectDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to remove Google Cloud project")
		},
	}
}

func toolCamGcpProjectUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_gcp_project_update",
			mcp.WithDescription("Modify Google Cloud project details. Changes the TrendAI Vision One™ settings for the specified Google Cloud project."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The Google Cloud project number Cloud Account Management uses for managing the connected Google Cloud project.")),
			mcp.WithString("ifMatch", mcp.Description("Entity tag in MD5 format to perform version control on the resource to be updated. This is returned in response header ETag of the v3 GET method. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("name", mcp.Description("The name of the Google Cloud project to be used in Cloud Accounts.")),
			mcp.WithString("description", mcp.Description("The description of the Google Cloud project. If there is no description, the field is empty.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamGcpProjectUpdate(id, params)
			return handleResponse(resp, err, "failed to modify Google Cloud project details")
		},
	}
}

func toolCamOciCompartmentsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_oci_compartments_list",
			mcp.WithDescription("Show connected OCI compartments. Displays all OCI compartments connected to TrendAI Vision One™ in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100, 500, 1000, 5000. Default: 25."), mcp.Enum("25", "50", "100", "500", "1000", "5000")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of connected OCI compartments. Supported fields: id - The id of the OCI compartment. - Any value; name - The name of the cloud account - Any value; state - The state of the cloud account - managed, outdated, failed; Supported operators: eq - Operator 'equal to'; and - Operator 'and'; or - Operator 'or'; not - Operator 'not'; ( ) - Symbols for grouping operands with their correct operator. contains - Operator that allows you to search for a specified string in a field; Note: Include this parameter in every request that generates paginated output.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamOciCompartmentsList(params)
			return handleResponse(resp, err, "failed to show connected OCI compartments")
		},
	}
}

func toolCamOciCompartmentsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_oci_compartments_create",
			mcp.WithDescription("Add OCI compartment. Adds and connects the specified Oracle Cloud Infrastructure (OCI) compartment to TrendAI Vision One™."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("serviceUserId", mcp.Required(), mcp.Description("The service user ID for the OCI account.")),
			mcp.WithString("compartmentId", mcp.Required(), mcp.Description("The OCI compartment ID.")),
			mcp.WithString("tenancyId", mcp.Required(), mcp.Description("The OCI tenancy ID.")),
			mcp.WithString("ociHomeRegion", mcp.Required(), mcp.Description("The region of the OCI tenancy.")),
			mcp.WithString("name", mcp.Required(), mcp.Description("The display name for the cloud account.")),
			mcp.WithString("description", mcp.Description("Optional description for the cloud account.")),
			mcp.WithString("identityDomainId", mcp.Required(), mcp.Description("The OCI identity domain ID.")),
			mcp.WithString("compartmentPolicyId", mcp.Required(), mcp.Description("The OCI policy OCID for the compartment.")),
			mcp.WithArray("features", mcp.Required(), mcp.Description("The list of selected features and the corresponding regions for deployment. The \"cloud-account-assessment\" feature is a primary feature that includes several child features. If it is disabled, all associated child features will also be disabled."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature.", "enum": []string{"cloud-account-assessment"}, "type": "string"}}, "required": []string{"id"}, "type": "object"})),
			mcp.WithObject("freeformTags", mcp.Description("The customer-defined OCI freeform tags to apply to every taggable resource created by the template. If omitted, no custom freeform tags are applied. Freeform tags are arbitrary key-value pairs that do not need to be pre-created in the OCI tenancy.")),
			mcp.WithObject("definedTags", mcp.Description("The customer-defined OCI defined tags to apply to every taggable resource created by the template. If omitted, no custom defined tags are applied. The tag namespace and key definitions must exist in the OCI tenancy at the time of template deployment.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "serviceUserId", Kind: kindString, Required: true}, {Name: "compartmentId", Kind: kindString, Required: true}, {Name: "tenancyId", Kind: kindString, Required: true}, {Name: "ociHomeRegion", Kind: kindString, Required: true}, {Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "identityDomainId", Kind: kindString, Required: true}, {Name: "compartmentPolicyId", Kind: kindString, Required: true}, {Name: "features", Kind: kindArray, Required: true}, {Name: "freeformTags", Kind: kindObject}, {Name: "definedTags", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamOciCompartmentsCreate(params)
			return handleResponse(resp, err, "failed to add OCI compartment")
		},
	}
}

func toolCamOciCompartmentsGenerateTerraformPackage(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_oci_compartments_generate_terraform_package",
			mcp.WithDescription("Get OCI Terraform template. Generates and downloads the Oracle Cloud Infrastructure (OCI) Terraform template."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("ociCompartmentName", mcp.Required(), mcp.Description("The user-defined name of the OCI compartment used in Cloud Accounts. Use this value as the 'name' field in the 'POST /v3.0/cam/ociCompartments' API.")),
			mcp.WithString("ociCompartmentDescription", mcp.Description("The description of the OCI compartment. Use this value as the 'description' field in the 'POST /v3.0/cam/ociCompartments' API.")),
			mcp.WithString("compartmentId", mcp.Required(), mcp.Description("The OCID of the OCI compartment. Use this value as the 'compartmentId' field in the 'POST /v3.0/cam/ociCompartments' API.")),
			mcp.WithString("identityDomainId", mcp.Required(), mcp.Description("The OCID of the OCI identity domain to be deployed. Use this value as the 'identityDomainId' field in the 'POST /v3.0/cam/ociCompartments' API.")),
			mcp.WithArray("features", mcp.Required(), mcp.Description("The list of selected features and the corresponding regions for deployment. The \"cloud-account-assessment\" feature is a primary feature that includes several child features. If it is disabled, all associated child features will also be disabled."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature.", "enum": []string{"cloud-account-assessment"}, "type": "string"}}, "required": []string{"id"}, "type": "object"})),
			mcp.WithObject("freeformTags", mcp.Description("The customer-defined OCI freeform tags to apply to every taggable resource created by the template. If omitted, no custom freeform tags are applied. Freeform tags are arbitrary key-value pairs that do not need to be pre-created in the OCI tenancy.")),
			mcp.WithObject("definedTags", mcp.Description("The customer-defined OCI defined tags to apply to every taggable resource created by the template. If omitted, no custom defined tags are applied. The tag namespace and key definitions must exist in the OCI tenancy at the time of template deployment.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "ociCompartmentName", Kind: kindString, Required: true}, {Name: "ociCompartmentDescription", Kind: kindString}, {Name: "compartmentId", Kind: kindString, Required: true}, {Name: "identityDomainId", Kind: kindString, Required: true}, {Name: "features", Kind: kindArray, Required: true}, {Name: "freeformTags", Kind: kindObject}, {Name: "definedTags", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamOciCompartmentsGenerateTerraformPackage(params)
			return handleResponse(resp, err, "failed to get OCI Terraform template")
		},
	}
}

func toolCamOciCompartmentDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_oci_compartment_delete",
			mcp.WithDescription("Remove OCI compartment. Disconnects and removes the specified OCI compartment from Trend Vision One."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Cloud Accounts uses the compartment OCID as the ID of the record for managing the OCI.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamOciCompartmentDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to remove OCI compartment")
		},
	}
}

func toolCamOciCompartmentGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_oci_compartment_get",
			mcp.WithDescription("Show OCI compartment details. Displays information about the specified OCI compartment."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Cloud Accounts uses the compartment OCID as the ID of the record for managing the OCI.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamOciCompartmentGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to show OCI compartment details")
		},
	}
}

func toolCamOciCompartmentUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"cam_oci_compartment_update",
			mcp.WithDescription("Modify OCI compartment details. Changes the settings for the specified OCI compartment in Trend Vision One."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Cloud Accounts uses the compartment OCID as the ID of the record for managing the OCI.")),
			mcp.WithString("ifMatch", mcp.Description("Entity tag in MD5 format to perform version control on the resource to be updated. This is returned in response header ETag of the v3 GET method. Use the ETag returned when the resource was retrieved.")),
			mcp.WithString("name", mcp.Description("The user-defined name of the OCI compartment to be used in Cloud Accounts.")),
			mcp.WithString("description", mcp.Description("The description of the OCI compartment. The default value is an empty string if the field is omitted.")),
			mcp.WithArray("features", mcp.Description("The list of selected features and the corresponding regions for deployment. The \"cloud-account-assessment\" feature is a primary feature that includes several child features. If it is disabled, all associated child features will also be disabled."), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the feature.", "enum": []string{"cloud-account-assessment"}, "type": "string"}}, "required": []string{"id"}, "type": "object"})),
			mcp.WithObject("freeformTags", mcp.Description("The customer-defined OCI freeform tags: omit the field to leave the tags unchanged, or pass an empty array to remove all of them. Freeform tags are arbitrary key-value pairs that do not need to be pre-created in the OCI tenancy.")),
			mcp.WithObject("definedTags", mcp.Description("The customer-defined OCI defined tags: omit the field to leave the tags unchanged, or pass an empty array to remove all of them. The tag namespace and key definitions must exist in the OCI tenancy at the time of template deployment.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Headers:    []headerDef{{Arg: "ifMatch", Header: "If-Match"}},
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}, {Name: "features", Kind: kindArray}, {Name: "freeformTags", Kind: kindObject}, {Name: "definedTags", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.CamOciCompartmentUpdate(id, params)
			return handleResponse(resp, err, "failed to modify OCI compartment details")
		},
	}
}
