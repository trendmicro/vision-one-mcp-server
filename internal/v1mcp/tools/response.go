package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlyResponse = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolResponseCustomScriptsList,
	toolResponseCustomScriptGet,
	toolResponseEndpointActionExceptionsList,
	toolResponseIsolatedTrafficExceptionsList,
	toolResponseOsqueryStatementsList,
	toolResponseSettingStatusGet,
	toolResponseTasksList,
	toolResponseTaskGet,
	toolResponseYaraRuleFilesList,
}

var ToolsetsWriteResponse = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolResponseContainersIsolate,
	toolResponseContainersRestore,
	toolResponseContainersTerminate,
	toolResponseCustomScriptsCreate,
	toolResponseCustomScriptDelete,
	toolResponseCustomScriptUpdate,
	toolResponseDomainAccountsDisable,
	toolResponseDomainAccountsEnable,
	toolResponseDomainAccountsResetPassword,
	toolResponseDomainAccountsSignOut,
	toolResponseEmailsDelete,
	toolResponseEmailsQuarantine,
	toolResponseEmailsRestore,
	toolResponseEndpointActionExceptionsRegister,
	toolResponseEndpointsCollectFile,
	toolResponseEndpointsIsolate,
	toolResponseEndpointsRestore,
	toolResponseEndpointsRunOsquery,
	toolResponseEndpointsRunScript,
	toolResponseEndpointsRunYaraRulesCreate,
	toolResponseEndpointsStartMalwareScan,
	toolResponseEndpointsTerminateProcess,
	toolResponseIsolatedTrafficExceptionsRegister,
	toolResponseSettingStatusUpdate,
	toolResponseSuspiciousObjectsCreate,
	toolResponseSuspiciousObjectsDelete,
	toolResponseTasksCancel,
}

func toolResponseContainersIsolate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_containers_isolate",
			mcp.WithDescription("Isolate container. Stops and isolates the Kubernetes pod containing the container from the environment."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "The description of a response task.", "type": "string"}, "kubernetesClusterId": map[string]any{"description": "The Kubernetes cluster Id, This value can be gotten from the Container Security API.", "type": "string"}, "kubernetesPodId": map[string]any{"description": "The Kubernetes pod Id , This value can be gotten from the Container Security API.", "format": "uuid", "type": "string"}}, "required": []string{"kubernetesClusterId", "kubernetesPodId"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseContainersIsolate(params)
			return handleResponse(resp, err, "failed to isolate container")
		},
	}
}

func toolResponseContainersRestore(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_containers_restore",
			mcp.WithDescription("Restore Container connection. Restarts and restores network connectivity to one or more containers that applied the \"Isolate Container\" action."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "The description of a response task.", "type": "string"}, "kubernetesClusterId": map[string]any{"description": "The Kubernetes cluster Id, This value can be gotten from the Container Security API.", "type": "string"}, "kubernetesPodId": map[string]any{"description": "The Kubernetes pod Id , This value can be gotten from the Container Security API.", "format": "uuid", "type": "string"}}, "required": []string{"kubernetesClusterId", "kubernetesPodId"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseContainersRestore(params)
			return handleResponse(resp, err, "failed to restore Container connection")
		},
	}
}

func toolResponseContainersTerminate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_containers_terminate",
			mcp.WithDescription("Terminate container. Terminates the specified Kubernetes pod or Amazon ECS containing the container. Terminating a container does not prevent further instances of suspicious activity."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"amazonEcsClusterId": map[string]any{"description": "The amazon ecs cluster id ,This value can be obtained from the Container Security API.", "type": "string"}, "amazonEcsTaskId": map[string]any{"description": "The ECS task ID. This value can be obtained from the Container Security API.", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "kubernetesClusterId": map[string]any{"description": "The Kubernetes cluster Id, This value can be gotten from the Container Security API.", "type": "string"}, "kubernetesPodId": map[string]any{"description": "The Kubernetes pod Id , This value can be gotten from the Container Security API.", "format": "uuid", "type": "string"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseContainersTerminate(params)
			return handleResponse(resp, err, "failed to terminate container")
		},
	}
}

func toolResponseCustomScriptsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_custom_scripts_list",
			mcp.WithDescription("List custom scripts. Retrieves information about the available custom scripts and displays the information in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the custom scripts list. Supported fields and operators: 'fileName' - File name of a custom script; 'fileType' - File extension of a custom script; 'eq' - Abbreviation of the operator \"equal to\"; 'and' - Operator \"and\"; 'or' - Operator \"or\"; 'not' - Operator \"not\"; '( )' - Symbols for grouping operands with their correct operator.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "filter", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseCustomScriptsList(params)
			return handleResponse(resp, err, "failed to list custom scripts")
		},
	}
}

func toolResponseCustomScriptsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_custom_scripts_create",
			mcp.WithDescription("Add custom script. Uploads a custom script. Supported file extensions: * PowerShell script (.ps1) * Bash script (.sh)."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("fileType", mcp.Required(), mcp.Description("File type of custom script."), mcp.Enum("powershell", "bash")),
			mcp.WithString("description", mcp.Description("Description of a custom script.")),
			mcp.WithString("file", mcp.Required(), mcp.Description("Custom script. Path to a local file to upload.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "fileType", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseCustomScriptsCreate(params)
			return handleResponse(resp, err, "failed to add custom script")
		},
	}
}

func toolResponseCustomScriptDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_custom_script_delete",
			mcp.WithDescription("Delete custom script. Deletes custom script."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a script file.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseCustomScriptDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete custom script")
		},
	}
}

func toolResponseCustomScriptGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_custom_script_get",
			mcp.WithDescription("Download custom script. Downloads custom script."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a script file.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseCustomScriptGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to download custom script")
		},
	}
}

func toolResponseCustomScriptUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_custom_script_update",
			mcp.WithDescription("Update custom script. Updates custom script. Supported file extensions: * PowerShell script (.ps1) * Bash script (.sh)."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a script file.")),
			mcp.WithString("fileType", mcp.Required(), mcp.Description("File type of custom script."), mcp.Enum("powershell", "bash")),
			mcp.WithString("description", mcp.Description("Description of a custom script.")),
			mcp.WithString("file", mcp.Required(), mcp.Description("Custom script. Path to a local file to upload.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "fileType", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseCustomScriptUpdate(id, params)
			return handleResponse(resp, err, "failed to update custom script")
		},
	}
}

func toolResponseDomainAccountsDisable(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_domain_accounts_disable",
			mcp.WithDescription("Disable user account. Signs the user out of all active application and browser sessions, and prevents the user from signing in any new session."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"accountName": map[string]any{"description": "User account", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}}, "required": []string{"accountName"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseDomainAccountsDisable(params)
			return handleResponse(resp, err, "failed to disable user account")
		},
	}
}

func toolResponseDomainAccountsEnable(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_domain_accounts_enable",
			mcp.WithDescription("Enable user account. Allows the user to sign in to new application and browser sessions."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"accountName": map[string]any{"description": "User account", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}}, "required": []string{"accountName"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseDomainAccountsEnable(params)
			return handleResponse(resp, err, "failed to enable user account")
		},
	}
}

func toolResponseDomainAccountsResetPassword(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_domain_accounts_reset_password",
			mcp.WithDescription("Force password reset. Signs the user out of all active application and browser sessions, and forces the user to create a new password during the next sign-in attempt."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"accountName": map[string]any{"description": "User account", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}}, "required": []string{"accountName"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseDomainAccountsResetPassword(params)
			return handleResponse(resp, err, "failed to force password reset")
		},
	}
}

func toolResponseDomainAccountsSignOut(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_domain_accounts_sign_out",
			mcp.WithDescription("Force sign out. Signs the user out of all active application and browser sessions."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"accountName": map[string]any{"description": "User account", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}}, "required": []string{"accountName"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseDomainAccountsSignOut(params)
			return handleResponse(resp, err, "failed to force sign out")
		},
	}
}

func toolResponseEmailsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_emails_delete",
			mcp.WithDescription("Delete email message. Deletes a message from one or more mailboxes."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "The description of a response task.", "type": "string"}, "mailBox": map[string]any{"description": "Email address.", "type": "object"}, "messageId": map[string]any{"description": "The ID of an email message.", "type": "string"}, "uniqueId": map[string]any{"description": "Unique alphanumeric string that identifies an email message within one mailbox.", "type": "string"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEmailsDelete(params)
			return handleResponse(resp, err, "failed to delete email message")
		},
	}
}

func toolResponseEmailsQuarantine(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_emails_quarantine",
			mcp.WithDescription("Quarantine email message. Moves a message from one or more mailboxes to their corresponding quarantine folders."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "The description of a response task.", "type": "string"}, "mailBox": map[string]any{"description": "Email address.", "type": "object"}, "messageId": map[string]any{"description": "The ID of an email message.", "type": "string"}, "uniqueId": map[string]any{"description": "Unique alphanumeric string that identifies an email message within one mailbox.", "type": "string"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEmailsQuarantine(params)
			return handleResponse(resp, err, "failed to quarantine email message")
		},
	}
}

func toolResponseEmailsRestore(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_emails_restore",
			mcp.WithDescription("Restore email message. Restore quarantined email messages."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "The description of a response task.", "type": "string"}, "mailBox": map[string]any{"description": "Email address.", "type": "object"}, "messageId": map[string]any{"description": "The ID of an email message.", "type": "string"}, "uniqueId": map[string]any{"description": "Unique alphanumeric string that identifies an email message within one mailbox.", "type": "string"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEmailsRestore(params)
			return handleResponse(resp, err, "failed to restore email message")
		},
	}
}

func toolResponseEndpointActionExceptionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoint_action_exceptions_list",
			mcp.WithDescription("Get endpoint response action exclusion list. Retrieves all exclusions including response actions and excluded endpoints in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the endpoint response action exclusion list. Supported fields: action - The response action name. - isolate, runCustomScript, remoteShell, collectFile, dumpProcessMemory; endpointName - The endpoint host name. - Any valid host name; endpointAgentGuid - The endpoint GUID. - Any valid GUID; Supported operators: eq - Operator 'equal to'; and - Operator 'and'; or - Operator 'or'; not - Operator 'not'; () - Symbols for grouping operands with their correct operator.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointActionExceptionsList(params)
			return handleResponse(resp, err, "failed to get endpoint response action exclusion list")
		},
	}
}

func toolResponseEndpointActionExceptionsRegister(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoint_action_exceptions_register",
			mcp.WithDescription("Exclude endpoints from response actions. Configures exclusions to prevent specific response actions from affecting selected endpoints. Note: Each exclusion supports a maximum of 2000 endpoints."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Description("The list of endpoint response action exclusions. Note: Each action can only exclude up to 2000 endpoints."), mcp.Items(map[string]any{"description": "The object that contains an endpoint response action exclusion.", "properties": map[string]any{"action": map[string]any{"description": "The response action name. Note: The isolate action is used for isolating and restoring.", "enum": []string{"isolate", "remoteShell", "collectFile", "dumpProcessMemory", "runCustomScript", "terminateProcess"}, "type": "string"}, "endpoints": map[string]any{"description": "The list of endpoint identification information.", "items": map[string]any{"type": "object"}, "type": "array"}}, "required": []string{"action", "endpoints"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointActionExceptionsRegister(params)
			return handleResponse(resp, err, "failed to exclude endpoints from response actions")
		},
	}
}

func toolResponseEndpointsCollectFile(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoints_collect_file",
			mcp.WithDescription("Collect file. Collects a file from one or more endpoints and then sends the files to TrendAI Vision One™ in a password-protected archive Note: You can specify either the computer name (\"endpointName\") or the GUID of the installed agent program (\"agentGuid\")."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of an installed agent.", "format": "uuid", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "endpointName": map[string]any{"description": "The endpoint name of the target endpoint", "type": "string"}, "filePath": map[string]any{"description": "The file path of the file to be collected from the object.", "type": "object"}}, "required": []string{"filePath"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointsCollectFile(params)
			return handleResponse(resp, err, "failed to collect file")
		},
	}
}

func toolResponseEndpointsIsolate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoints_isolate",
			mcp.WithDescription("Isolate endpoints. Disconnects one or more endpoints from the network (but allows communication with the managing TrendAI™ server product) Note: You can specify either the computer name (\"endpointName\") or the GUID of the installed agent program (\"agentGuid\")."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of an installed agent.", "format": "uuid", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "endpointName": map[string]any{"description": "The endpoint name of the target endpoint", "type": "string"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointsIsolate(params)
			return handleResponse(resp, err, "failed to isolate endpoints")
		},
	}
}

func toolResponseEndpointsRestore(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoints_restore",
			mcp.WithDescription("Restore endpoint connection. Restores network connectivity to one or more endpoints that applied the \"Isolate endpoint\" action Note: You can specify either the computer name (\"endpointName\") or the GUID of the installed agent program (\"agentGuid\")."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of an installed agent.", "format": "uuid", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "endpointName": map[string]any{"description": "The endpoint name of the target endpoint", "type": "string"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointsRestore(params)
			return handleResponse(resp, err, "failed to restore endpoint connection")
		},
	}
}

func toolResponseEndpointsRunOsquery(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoints_run_osquery",
			mcp.WithDescription("Run osquery. Run SQL-based queries on the specified endpoints."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of an installed agent.", "format": "uuid", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "endpointName": map[string]any{"description": "The endpoint name of the target endpoint", "type": "string"}, "statement": map[string]any{"description": "The SQL statement of an osquery. Please refer to the osquery specification at https://osquery.readthedocs.io/en/stable/introduction/sql/.", "type": "string"}, "statementId": map[string]any{"description": "The ID of an osquery statement in hash format", "type": "string"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointsRunOsquery(params)
			return handleResponse(resp, err, "failed to run osquery")
		},
	}
}

func toolResponseEndpointsRunScript(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoints_run_script",
			mcp.WithDescription("Run custom script. Runs custom script."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of an installed agent.", "format": "uuid", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "endpointName": map[string]any{"description": "The endpoint name of the target endpoint", "type": "string"}, "fileName": map[string]any{"description": "File name of a custom script", "type": "string"}, "parameter": map[string]any{"description": "Options passed to the script during execution", "type": "string"}}, "required": []string{"fileName"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointsRunScript(params)
			return handleResponse(resp, err, "failed to run custom script")
		},
	}
}

func toolResponseEndpointsRunYaraRulesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoints_run_yara_rules_create",
			mcp.WithDescription("Run YARA rules. Run custom YARA rules on the specified endpoints."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of an installed agent.", "format": "uuid", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "endpointName": map[string]any{"description": "The endpoint name of the target endpoint", "type": "string"}, "target": map[string]any{"description": "The target type of a YARA scan", "enum": []string{"File"}, "type": "string"}, "targetFileLocation": map[string]any{"description": "The location of the file to be scanned", "type": "string"}, "targetFileOption": map[string]any{"description": "The scanning options of a YARA scan. SCAN_ALL scans all files and subfolders, so might cause performance issues.", "enum": []string{"SCAN_ALL", "SCAN_TOP"}, "type": "string"}, "targetFileSize": map[string]any{"description": "The maximum size of the file to be scanned", "enum": []string{"1M", "2M", "3M", "4M"}, "type": "string"}, "targetProcessName": map[string]any{"description": "The target process name of a YARA scan", "type": "string"}, "yaraRuleFileContent": map[string]any{"description": "The YARA rules to be used in a YARA scan", "type": "string"}, "yaraRuleFileId": map[string]any{"description": "The ID of a YARA rule file", "format": "uuid", "type": "string"}, "yaraRuleFileName": map[string]any{"description": "File name of a Yara rule", "type": "string"}}, "required": []string{"target"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointsRunYaraRulesCreate(params)
			return handleResponse(resp, err, "failed to run YARA rules")
		},
	}
}

func toolResponseEndpointsStartMalwareScan(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoints_start_malware_scan",
			mcp.WithDescription("Scan for malware. Runs a one-time malware scan on one or more endpoints. Note: When choosing endpoints, you may specify either the computer name (\"endpointName\") or the GUID of the installed agent program (\"agentGuid\")."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("endpoints", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of an installed agent.", "format": "uuid", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "endpointName": map[string]any{"description": "The endpoint name of the target endpoint", "type": "string"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "endpoints", Kind: kindArray, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointsStartMalwareScan(params)
			return handleResponse(resp, err, "failed to scan for malware")
		},
	}
}

func toolResponseEndpointsTerminateProcess(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_endpoints_terminate_process",
			mcp.WithDescription("Terminate process. Terminates a process that is running on one or more endpoints. This operation is currently only supported on the Windows platform. Note: You can specify either the computer name (\"endpointName\") or the GUID of the installed agent program (\"agentGuid\")."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"agentGuid": map[string]any{"description": "The ID of an installed agent.", "format": "uuid", "type": "string"}, "description": map[string]any{"description": "The description of a response task.", "type": "string"}, "endpointName": map[string]any{"description": "The endpoint name of the target endpoint", "type": "string"}, "fileName": map[string]any{"description": "File name of a response task target", "type": "object"}, "fileSha1": map[string]any{"description": "SHA1 hash of the terminated process's executable file. Either 'fileName' or 'fileSha1' must be provided.", "type": "object"}, "processId": map[string]any{"description": "Unique numeric string that identifies an active process", "type": "object"}}, "required": []string{"processId"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseEndpointsTerminateProcess(params)
			return handleResponse(resp, err, "failed to terminate process")
		},
	}
}

func toolResponseIsolatedTrafficExceptionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_isolated_traffic_exceptions_list",
			mcp.WithDescription("Get network traffic exceptions for isolated endpoints. Retrieves the list of exceptions that allow network traffic to and from isolated endpoints. Note: This method can only retrieve the entire exception list at once."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the network traffic exceptions list. Supported fields: direction - The direction of network traffic of the traffic exception. - inbound, outbound; ip - The IP address of the traffic exception. - Any IPv4 address; protocol - The network protocol of the traffic exception. - \"Any\", \"TCP/UDP\", \"TCP\", \"UDP\", \"ICMP\"; port - The port numbers associated with the traffic exception. - 0-65535; Supported operators: eq - Operator 'equal to'; and - Operator 'and'; or - Operator 'or'; not - Operator 'not'; () - Symbols for grouping operands with their correct operator.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseIsolatedTrafficExceptionsList(params)
			return handleResponse(resp, err, "failed to get network traffic exceptions for isolated endpoints")
		},
	}
}

func toolResponseIsolatedTrafficExceptionsRegister(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_isolated_traffic_exceptions_register",
			mcp.WithDescription("Configure network traffic exceptions. Configures the exceptions that allow network traffic to and from isolated endpoints. Note: You can specify up to 50 inbound and 50 outbound exceptions."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Description("The list of both inbound and outbound network traffic exceptions for isolated endpoints."), mcp.Items(map[string]any{"description": "The object that contains a network traffic exception.", "properties": map[string]any{"direction": map[string]any{"description": "The direction of network traffic of the traffic exception.", "enum": []string{"inbound", "outbound"}, "type": "string"}, "ip": map[string]any{"description": "The IP address of the traffic exception.", "type": "string"}, "port": map[string]any{"description": "The port numbers associated with the traffic exception. Multiple ports can be specified by separating them with commas.", "type": "string"}, "protocol": map[string]any{"description": "The network protocol of the traffic exception.", "enum": []string{"Any", "TCP/UDP", "TCP", "UDP", "ICMP"}, "type": "string"}}, "required": []string{"ip", "direction"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseIsolatedTrafficExceptionsRegister(params)
			return handleResponse(resp, err, "failed to configure network traffic exceptions")
		},
	}
}

func toolResponseOsqueryStatementsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_osquery_statements_list",
			mcp.WithDescription("List osquery statements. Retrieves information about the available osquery and displays the information in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The maximum number of items to return in the response. One of: 50, 100, 200. Default: 50."), mcp.Enum("50", "100", "200")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseOsqueryStatementsList(params)
			return handleResponse(resp, err, "failed to list osquery statements")
		},
	}
}

func toolResponseSettingStatusGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_setting_status_get",
			mcp.WithDescription("Get endpoint response settings. Retrieves the status of the endpoint response settings."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.ResponseSettingStatusGet(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get endpoint response settings")
		},
	}
}

func toolResponseSettingStatusUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_setting_status_update",
			mcp.WithDescription("Update endpoint response settings. Enables or disables isolated endpoint network traffic or endpoint response action exclusions."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("endpointActionException", mcp.Description("The status of the endpoint response action exclusion setting."), mcp.Enum("enabled", "disabled")),
			mcp.WithString("isolatedTrafficException", mcp.Description("The status of the isolated endpoint network traffic setting."), mcp.Enum("enabled", "disabled")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "endpointActionException", Kind: kindString}, {Name: "isolatedTrafficException", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseSettingStatusUpdate(params)
			return handleResponse(resp, err, "failed to update endpoint response settings")
		},
	}
}

func toolResponseSuspiciousObjectsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_suspicious_objects_create",
			mcp.WithDescription("Add to block list. Adds an email address, file SHA-1, domain, IP address, or URL to the Suspicious Object List, which blocks the objects on subsequent detections."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "The description of a response task.", "type": "string"}, "domain": map[string]any{"description": "Domain", "type": "object"}, "fileSha1": map[string]any{"description": "The SHA1 hash of file", "type": "object"}, "fileSha256": map[string]any{"description": "The SHA256 hash of a file", "type": "object"}, "ip": map[string]any{"description": "IP address. This fields supports both IPv4 and IPv6 addresses", "type": "object"}, "senderMailAddress": map[string]any{"description": "Email address", "type": "object"}, "url": map[string]any{"description": "Universal Resource Locator", "type": "object"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseSuspiciousObjectsCreate(params)
			return handleResponse(resp, err, "failed to add to block list")
		},
	}
}

func toolResponseSuspiciousObjectsDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_suspicious_objects_delete",
			mcp.WithDescription("Remove from block list. Removes an email address, file SHA-1, domain, IP address, or URL that was added to the Suspicious Object List using the \"Add to block list\" action."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"description": map[string]any{"description": "The description of a response task.", "type": "string"}, "domain": map[string]any{"description": "Domain", "type": "object"}, "fileSha1": map[string]any{"description": "The SHA1 hash of file", "type": "object"}, "fileSha256": map[string]any{"description": "The SHA256 hash of a file", "type": "object"}, "ip": map[string]any{"description": "IP address. This fields supports both IPv4 and IPv6 addresses", "type": "object"}, "senderMailAddress": map[string]any{"description": "Email address", "type": "object"}, "url": map[string]any{"description": "Universal Resource Locator", "type": "object"}}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseSuspiciousObjectsDelete(params)
			return handleResponse(resp, err, "failed to remove from block list")
		},
	}
}

func toolResponseTasksList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_tasks_list",
			mcp.WithDescription("Get response tasks. Displays a list of response tasks in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the start of the data retrieval range. You can retrieve data for response tasks that were created no later than 180 days ago.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval range.")),
			mcp.WithString("dateTimeTarget", mcp.Description("Parameter that allows you to filter the results. Default: createdDateTime."), mcp.Enum("createdDateTime", "lastActionDateTime")),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the response task list. Supported fields: id - The identifier of a response task - Any value; status - The status of a command sent to the management server - queued, running, succeeded, failed, canceled, pendingApproval, rejected action - The command sent to a target - isolate, isolateForMultiple, restoreIsolate, restoreIsolateForMultiple, collectFile, terminateProcess, block, restoreBlock, quarantineMessage, restoreMessage, deleteMessage, remoteShell, investigationKit, runCustomScript, runCustomScriptForMultiple, submitSandbox, dumpProcessMemory, disableAccount, enableAccount, forceSignOut, resetPassword, collectNetworkAnalysisPackage, collectEvidence, runOsquery, runOsqueryForMultiple, runYaraRules, runYaraForMultiple, isolateContainer, restoreContainer, terminateContainer, malwareScanForMultiple; account - The user that created the task - Any value; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; () - Symbols for grouping operands with their correct operator - -.")),
			mcp.WithString("top", mcp.Description("Number of records displayed on a page. If no value is specified, 'top' defaults to 50. One of: 50, 100, 200. Default: 50."), mcp.Enum("50", "100", "200")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "dateTimeTarget", Kind: kindString}, {Name: "filter", Kind: kindString}, {Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseTasksList(params)
			return handleResponse(resp, err, "failed to get response tasks")
		},
	}
}

func toolResponseTasksCancel(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_tasks_cancel",
			mcp.WithDescription("Cancel response task. Stops a currently running response task. Note: Only the Scan for Malware task is currently supported."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithArray("items", mcp.Required(), mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The ID of the response action.", "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body: bodyArray,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseTasksCancel(params)
			return handleResponse(resp, err, "failed to cancel response task")
		},
	}
}

func toolResponseTaskGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_task_get",
			mcp.WithDescription("Download response task results. Retrieves an object containing the results of a response task in JSON format."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("Unique alphanumeric string that identifies a response task.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseTaskGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to download response task results")
		},
	}
}

func toolResponseYaraRuleFilesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"response_yara_rule_files_list",
			mcp.WithDescription("List YARA rule files. Retrieves information about the available YARA rules and displays the information in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("filter", mcp.Description("Filter for retrieving a subset of the YARA rules file list. Supported fields and operators: 'name' - File name of a YARA rule file; 'eq' - Abbreviation of the operator \"equal to\"; 'and' - Operator \"and\"; 'or' - Operator \"or\"; 'not' - Operator \"not\"; '( )' - Symbols for grouping operands with their correct operator.")),
			mcp.WithString("top", mcp.Description("The maximum number of items to return in the response. One of: 50, 100, 200. Default: 50."), mcp.Enum("50", "100", "200")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "filter", Kind: kindString}, {Name: "top", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ResponseYaraRuleFilesList(params)
			return handleResponse(resp, err, "failed to list YARA rule files")
		},
	}
}
