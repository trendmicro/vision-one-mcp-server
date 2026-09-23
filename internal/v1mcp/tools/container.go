package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tooldescriptions"
)

var ToolsetsReadOnlyContainer = []func(client *v1client.V1ApiClient) mcpserver.ServerTool{
	toolContainerSecurityImageVulnerabilitiesList,
	toolContainerSecurityK8ClustersList,
	toolContainerSecurityK8ClusterGet,
	toolContainerSecurityECSClustersList,
	toolContainerSecurityK8ImagesList,
	toolContainerSecurityAmazonEcsClusterGet,
	toolContainerSecurityAmazonEcsEvaluationEventLogsList,
	toolContainerSecurityAmazonEcsImageOccurrencesList,
	toolContainerSecurityAmazonEcsSensorEventLogsList,
	toolContainerSecurityAttestorsList,
	toolContainerSecurityAttestorGet,
	toolContainerSecurityComplianceScanConfigurationGet,
	toolContainerSecurityComplianceScanSummaryGet,
	toolContainerSecurityComplianceScanVersionsList,
	toolContainerSecurityCustomRulesetsList,
	toolContainerSecurityCustomRulesetGet,
	toolContainerSecurityCustomRulesetRuleFileSourceGet,
	toolContainerSecurityFileIntegrityMonitoringRulesList,
	toolContainerSecurityFileIntegrityMonitoringRuleGet,
	toolContainerSecurityGenerateServiceGatewayPassword,
	toolContainerSecurityKubernetesAuditEventLogsList,
	toolContainerSecurityKubernetesClusterGroupsList,
	toolContainerSecurityKubernetesEvaluationEventLogsList,
	toolContainerSecurityKubernetesImageOccurrencesList,
	toolContainerSecurityKubernetesSensorEventLogsList,
	toolContainerSecurityManagedRulesList,
	toolContainerSecurityManagedRuleGet,
	toolContainerSecurityPoliciesList,
	toolContainerSecurityPolicyGet,
	toolContainerSecurityRulesetsList,
	toolContainerSecurityRulesetGet,
}

var ToolsetsWriteContainer = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolContainerSecurityAmazonEcsClusterUpdate,
	toolContainerSecurityAttestorsCreate,
	toolContainerSecurityAttestorDelete,
	toolContainerSecurityAttestorUpdate,
	toolContainerSecurityComplianceScanConfigurationUpdate,
	toolContainerSecurityCustomRulesetsCreate,
	toolContainerSecurityCustomRulesetDelete,
	toolContainerSecurityCustomRulesetUpdate,
	toolContainerSecurityFileIntegrityMonitoringRulesCreate,
	toolContainerSecurityFileIntegrityMonitoringRuleDelete,
	toolContainerSecurityFileIntegrityMonitoringRuleUpdate,
	toolContainerSecurityKubernetesClustersCreate,
	toolContainerSecurityKubernetesClusterDelete,
	toolContainerSecurityKubernetesClusterUpdate,
	toolContainerSecurityPoliciesCreate,
	toolContainerSecurityPolicyDelete,
	toolContainerSecurityPolicyUpdate,
	toolContainerSecurityRulesetsCreate,
	toolContainerSecurityRulesetDelete,
	toolContainerSecurityRulesetUpdate,
	toolContainerSecurityStartComplianceScan,
}

func toolContainerSecurityImageVulnerabilitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_image_vulnerabilities_list",
			mcp.WithDescription(
				"Displays the container image vulnerabilities detected in Kubernetes and Amazon ECS clusters for your account",
			),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterContainerVuln)),
			mcp.WithString("orderBy",
				mcp.Enum(
					"riskLevel desc",
					"firstDetectedDateTime desc",
					"lastDetectedDateTime desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("lastDetectedStartDateTime",
				mcp.Description("The start time of the data retrieval range, in ISO 8601 format."),
			),
			mcp.WithString("lastDetectedEndDateTime",
				mcp.Description("The end time of the data retrieval range, represented in ISO 8601 format."),
			),
			mcp.WithString("firstDetectedStartDateTime",
				mcp.Description("The start time of the data retrieval range, in ISO 8601 format."),
			),
			mcp.WithString("firstDetectedEndDateTime",
				mcp.Description("The end time of the data retrieval range, represented in ISO 8601 format."),
			),
			mcp.WithString("skipToken",
				mcp.Description("The token use to paginate. Used to retrieve the next page of information.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

			firstDetectedStartDateTime, err := optionalTimeValue("firstDetectedStartDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			firstDetectedEndDateTime, err := optionalTimeValue("firstDetectedEndDateTime", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				OrderBy:                    orderBy,
				LastDetectedStartDateTime:  lastDetectedStartDateTime,
				LastDetectedEndDateTime:    lastDetectedEndDateTime,
				FirstDetectedStartDateTime: firstDetectedStartDateTime,
				FirstDetectedEndDateTime:   firstDetectedEndDateTime,
				SkipToken:                  skipToken,
			}

			resp, err := client.ContainerSecurityListContainerImageVulnerabilities(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list container image vulnerabilities")
		},
	}
}

func toolContainerSecurityK8ClustersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_k8_clusters_list",
			mcp.WithDescription("Displays all registered Kubernetes clusters"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("orderBy",
				mcp.Enum(
					"createdDateTime desc",
					"updatedDateTime desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterK8s)),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				OrderBy: orderBy,
			}

			resp, err := client.ContainerSecurityListK8Clusters(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list k8 clusters")
		},
	}
}

func toolContainerSecurityK8ClusterGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_k8_cluster_get",
			mcp.WithDescription("Displays the details of the specified Kubernetes cluster"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("clusterID",
				mcp.Required(),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			clusterID, err := requiredValue[string]("clusterID", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityGetK8ClusterDetails(clusterID)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to get cluster details")
		},
	}
}

func toolContainerSecurityECSClustersList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_ecs_clusters_list",
			mcp.WithDescription("Displays all registered Amazon Elastic Container Service (ECS) clusters in a paginated list"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("orderBy",
				mcp.Enum(
					"createdDateTime desc",
					"updatedDateTime desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterECS)),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				OrderBy: orderBy,
			}

			resp, err := client.ContainerSecurityListECSClusters(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list ecs clusters")
		},
	}
}

func toolContainerSecurityK8ImagesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_k8_images_list",
			mcp.WithDescription("Displays the Kubernetes images that are running in all clusters for your account"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("orderBy",
				mcp.Enum(
					"id desc",
				),
				mcp.Description("The field by which the results are sorted"),
			),
			mcp.WithString("filter", mcp.Description(tooldescriptions.FilterK8Images)),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			filter, err := optionalValue[string]("filter", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			orderBy, err := optionalValue[string]("orderBy", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			qp := v1client.QueryParameters{
				OrderBy: orderBy,
			}

			resp, err := client.ContainerSecurityListK8Images(filter, qp)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list kubernetes images")
		},
	}
}

func toolContainerSecurityAmazonEcsClusterGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_amazon_ecs_cluster_get",
			mcp.WithDescription("Get cluster details (Amazon ECS). Displays the details of the specified Amazon Elastic Container Service (ECS) cluster."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Amazon ECS cluster.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityAmazonEcsClusterGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get cluster details (Amazon ECS)")
		},
	}
}

func toolContainerSecurityAmazonEcsClusterUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_amazon_ecs_cluster_update",
			mcp.WithDescription("Modify cluster settings (Amazon ECS). Modifies the settings of a registered Amazon Elastic Container Service (ECS) cluster."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Amazon ECS cluster.")),
			mcp.WithString("description", mcp.Description("The description of the cluster.")),
			mcp.WithString("policyId", mcp.Description("The ID of the policy associated with the cluster.")),
			mcp.WithBoolean("vulnerabilityScanEnabled", mcp.Description("If true, enables vulnerability scan for the cluster.")),
			mcp.WithBoolean("runtimeSecurityEnabled", mcp.Description("If true, enables runtime security for the cluster.")),
			mcp.WithArray("customizableTagIds", mcp.Description("The custom tags and platform tags associated with the cluster."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithObject("connection", mcp.Description("Connection setting configuration for the cluster."), mcp.Properties(map[string]any{"customProxyCredential": map[string]any{"properties": map[string]any{"password": map[string]any{"description": "Password of custom proxy authentication.", "type": "string"}, "username": map[string]any{"description": "Username of custom proxy authentication.", "type": "string"}}, "required": []string{"username", "password"}, "type": "object"}, "customProxyUrl": map[string]any{"description": "The proxy url connect to the proxy server.", "type": "string"}, "serviceGatewayApplianceIds": map[string]any{"description": "Array of Service Gateway appliance IDs for Service Gateway connection.", "items": map[string]any{"description": "Service Gateway Appliance ID.", "type": "string"}, "type": "array"}, "type": map[string]any{"description": "Direct connection.", "enum": []string{"direct"}, "type": "string"}, "useAllAppliances": map[string]any{"description": "Uses all available Service Gateway appliances. When specified, serviceGatewayApplianceIds is not required.", "enum": []any{true}, "type": "boolean"}})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "description", Kind: kindString}, {Name: "policyId", Kind: kindString}, {Name: "vulnerabilityScanEnabled", Kind: kindBoolean}, {Name: "runtimeSecurityEnabled", Kind: kindBoolean}, {Name: "customizableTagIds", Kind: kindArray}, {Name: "connection", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityAmazonEcsClusterUpdate(id, params)
			return handleResponse(resp, err, "failed to modify cluster settings (Amazon ECS)")
		},
	}
}

func toolContainerSecurityAmazonEcsEvaluationEventLogsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_amazon_ecs_evaluation_event_logs_list",
			mcp.WithDescription("Get evaluation event logs (Amazon ECS). Displays all the evaluation events of your Amazon ECS clusters in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval time range in ISO 8601 format. Default: now() - 30 days.")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval time range in ISO 8601 format. Default: Current date time.")),
			mcp.WithString("orderBy", mcp.Description("The field used to sort results. Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("The filter that retrieves a subset of the Evaluation events of your Amazon ECS clusters. Include this parameter in every request that generates paginated output. Supported fields: action - The policy action for the event - allow, block, log; clusterId - The ID of the cluster in Container Security - Any value; clusterName - The name of the cluster - Any value; decision - The evaluation decision for the event - allow, deny; policyId - The ID of the policy for an event - Any value; policyName - The name of a policy for an event - Any value; kind - The type of ECS workload of the event - Task, Service; workloadArn - The ARN of the ECS workload of the event - Any value; taskDefinitionArn - The ARN of the ECS task definition for the workload evaluated - Any value; operation - The operation of the evaluation event - Any value; Supported operators: eq - Abbreviation of the operator \"equal to\"; and - Operator \"and\"; or - Operator \"or\"; contains - Operator that allows you to search for a specified string in a field; not - Operator \"not\"; () - Symbols for grouping operands. Example: contains(clusterName,'example_cluster') and action eq 'allow'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityAmazonEcsEvaluationEventLogsList(params)
			return handleResponse(resp, err, "failed to get evaluation event logs (Amazon ECS)")
		},
	}
}

func toolContainerSecurityAmazonEcsImageOccurrencesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_amazon_ecs_image_occurrences_list",
			mcp.WithDescription("Get image occurrences (Amazon ECS). Displays the occurrences of Amazon ECS images that are running in all clusters for your account."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the ECS Images Occurrences List. Include this parameter in every request that generates paginated output. Supported fields: imageId - The The ID of the container image; Supported operators: eq - Operator \"equal to\"; not - Operator \"not\"; or - Operator \"or\". Example: imageId eq 'imageId_1'")),
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

			resp, err := client.ContainerSecurityAmazonEcsImageOccurrencesList(params)
			return handleResponse(resp, err, "failed to get image occurrences (Amazon ECS)")
		},
	}
}

func toolContainerSecurityAmazonEcsSensorEventLogsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_amazon_ecs_sensor_event_logs_list",
			mcp.WithDescription("Get runtime sensor events (Amazon ECS). Displays a list of all sensor events on your Amazon ECS clusters."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval time range, in ISO 8601 format. Default: now() - 30 days.")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval time range, in ISO 8601 format. Default: Current date time.")),
			mcp.WithString("filter", mcp.Description("The Filter for retrieving a subset of the sensor events On Amazon ECS clusters. Include this parameter in every request that generates paginated output. Supported fields: clusterId - The ID of the cluster in Container Security - Any value; clusterName - The name of the cluster - Any value; policyId - The ID of the policy for an event - Any value; policyName - The name of a policy for an event - Any value; mitigation - The mitigation action to take when a rule fails during runtime. - Any value; ruleName - The name of the rule that the event triggers - Any; ruleId - The ID of the rule that the event triggers - Any; eventType - The event type of the sensor event - syscall, fileIntegrity; Supported operators: eq - Operator \"equal to\"; and - Operator \"and\"; or - Operator \"or\"; contains - Operator that allows you to search for a specified string in a field; not - Operator \"not\"; () - Symbols for grouping operands. Example: contains(clusterName, 'example_cluster') and policyName eq 'mitre'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityAmazonEcsSensorEventLogsList(params)
			return handleResponse(resp, err, "failed to get runtime sensor events (Amazon ECS)")
		},
	}
}

func toolContainerSecurityAttestorsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_attestors_list",
			mcp.WithDescription("Get Container Security attestors. Displays your Container Security attestors in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The fields by which the results are sorted. You can indicate multiple fields separated by commas (,). Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the attestors list. Include this parameter in every request that generates paginated output. Supported fields: name - The name of the attestor - Any value; type - The type of the attestor - cosignPublicKey; managementType - The management type of the attestor - userManaged, clusterManaged; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; Additional functions: Function - Description - Notes; contains - Operator that allows you to search for a specified string in a field - Only applicable to name.")),
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

			resp, err := client.ContainerSecurityAttestorsList(params)
			return handleResponse(resp, err, "failed to get Container Security attestors")
		},
	}
}

func toolContainerSecurityAttestorsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_attestors_create",
			mcp.WithDescription("Create Container Security attestor. Creates a new Container Security attestor."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the attestor.")),
			mcp.WithString("description", mcp.Description("An optional description of the attestor.")),
			mcp.WithString("type", mcp.Required(), mcp.Description("The type of attestor."), mcp.Enum("cosignPublicKey")),
			mcp.WithString("publicKey", mcp.Required(), mcp.Description("The public key of an attestor of type cosignPublicKey.")),
			mcp.WithBoolean("transparencyLogEnabled", mcp.Description("Whether to verify the signature using a transparency log server. Default: False.")),
			mcp.WithObject("transparencyLogOptions", mcp.Description("The options for the transparency log server."), mcp.Properties(map[string]any{"publicKey": map[string]any{"description": "The public key used to verify the transparency log entry.", "type": "string"}, "server": map[string]any{"description": "The server used to verify the cosign signature using a transparency log.", "type": "string"}})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "type", Kind: kindString, Required: true}, {Name: "publicKey", Kind: kindString, Required: true}, {Name: "transparencyLogEnabled", Kind: kindBoolean}, {Name: "transparencyLogOptions", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityAttestorsCreate(params)
			return handleResponse(resp, err, "failed to create Container Security attestor")
		},
	}
}

func toolContainerSecurityAttestorDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_attestor_delete",
			mcp.WithDescription("Delete a Container Security attestor. Deletes the specified Container Security attestor."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Container Security attestor.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityAttestorDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete a Container Security attestor")
		},
	}
}

func toolContainerSecurityAttestorGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_attestor_get",
			mcp.WithDescription("Get Container Security attestor. Displays the details of the specified Container Security attestor."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Container Security attestor.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityAttestorGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get Container Security attestor")
		},
	}
}

func toolContainerSecurityAttestorUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_attestor_update",
			mcp.WithDescription("Modify a Container Security attestor. Modifies the settings of the specified Container Security attestor."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the Container Security attestor.")),
			mcp.WithString("description", mcp.Description("A description of the attestor.")),
			mcp.WithString("publicKey", mcp.Description("The public key of an attestor of type cosignPublicKey.")),
			mcp.WithBoolean("transparencyLogEnabled", mcp.Description("Whether to verify the signature using a transparency log server.")),
			mcp.WithObject("transparencyLogOptions", mcp.Description("The options for the transparency log server."), mcp.Properties(map[string]any{"publicKey": map[string]any{"description": "The public key used to verify the transparency log entry.", "type": "string"}, "server": map[string]any{"description": "The server used to verify the cosign signature using a transparency log.", "type": "string"}})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "description", Kind: kindString}, {Name: "publicKey", Kind: kindString}, {Name: "transparencyLogEnabled", Kind: kindBoolean}, {Name: "transparencyLogOptions", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityAttestorUpdate(id, params)
			return handleResponse(resp, err, "failed to modify a Container Security attestor")
		},
	}
}

func toolContainerSecurityComplianceScanConfigurationGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_compliance_scan_configuration_get",
			mcp.WithDescription("Get compliance scans configuration. Displays the configuration of compliance scans, including any exceptions."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.ContainerSecurityComplianceScanConfigurationGet(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get compliance scans configuration")
		},
	}
}

func toolContainerSecurityComplianceScanConfigurationUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_compliance_scan_configuration_update",
			mcp.WithDescription("Update compliance scans configuration. Updates the configuration of the specified compliance scan."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithBoolean("scheduleScanEnabled", mcp.Description("Whether the scheduled compliance scan is enabled or disabled.")),
			mcp.WithArray("exceptions", mcp.Items(map[string]any{"properties": map[string]any{"benchmarkVersion": map[string]any{"description": "The benchmark version of the exception.", "type": "string"}, "parameters": map[string]any{"description": "The categories of the exception.", "items": map[string]any{"type": "object"}, "type": "array"}, "type": map[string]any{"description": "The benchmark type of the exception.", "enum": []string{"Kubernetes", "Amazon EKS", "Google GKE", "Microsoft AKS", "NSA"}, "type": "string"}}, "required": []string{"type", "benchmarkVersion", "parameters"}, "type": "object"})),
			mcp.WithArray("benchmarkVersions", mcp.Items(map[string]any{"description": "The settings of the request.", "properties": map[string]any{"activated": map[string]any{"description": "The active benchmark version of the compliance scan.", "type": "string"}, "type": map[string]any{"description": "The type of the benchmark.", "enum": []string{"Kubernetes", "Amazon EKS", "OpenShift", "Google GKE", "Microsoft AKS", "NSA"}, "type": "string"}}, "required": []string{"type", "activated"}, "type": "object"})),
			mcp.WithArray("frameworks", mcp.Items(map[string]any{"properties": map[string]any{"activationStatus": map[string]any{"description": "The activation status of compliance scan frameworks used to run scans.", "enum": []string{"activated", "deactivated"}, "type": "string"}, "name": map[string]any{"description": "The framework type of the compliance scans.", "enum": []string{"CIS", "NSA"}, "type": "string"}}, "required": []string{"name", "activationStatus"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "scheduleScanEnabled", Kind: kindBoolean}, {Name: "exceptions", Kind: kindArray}, {Name: "benchmarkVersions", Kind: kindArray}, {Name: "frameworks", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityComplianceScanConfigurationUpdate(params)
			return handleResponse(resp, err, "failed to update compliance scans configuration")
		},
	}
}

func toolContainerSecurityComplianceScanSummaryGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_compliance_scan_summary_get",
			mcp.WithDescription("Get compliance scan summary. Displays the most recent status summary of your compliance scans."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.ContainerSecurityComplianceScanSummaryGet(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get compliance scan summary")
		},
	}
}

func toolContainerSecurityComplianceScanVersionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_compliance_scan_versions_list",
			mcp.WithDescription("Get status of compliance scan benchmarks. Display the status of the compliance scan bechmarks in the account."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.ContainerSecurityComplianceScanVersionsList(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get status of compliance scan benchmarks")
		},
	}
}

func toolContainerSecurityCustomRulesetsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_custom_rulesets_list",
			mcp.WithDescription("Get Container Security custom rulesets. Displays your Container Security custom rulesets in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The fields by which the results are sorted. You can indicate multiple fields separated by commas (,). Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the custom rulesets list. Include this parameter in every request that generates paginated output. Supported fields: name - The name of the custom ruleset - Any value; type - The type of the custom ruleset - userManaged, clusterManaged; Supported operators: eq - Operator 'equal to' - -; and - Operator 'and' - -; or - Operator 'or' - -; not - Operator 'not' - -; Additional functions: Function - Description - Notes; contains - Operator that allows you to search for a specified string in a field - Only applicable to name.")),
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

			resp, err := client.ContainerSecurityCustomRulesetsList(params)
			return handleResponse(resp, err, "failed to get Container Security custom rulesets")
		},
	}
}

func toolContainerSecurityCustomRulesetsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_custom_rulesets_create",
			mcp.WithDescription("Create Container Security custom ruleset. Creates your Container Security user-managed custom ruleset, including rule file upload. Note that only one rules file is supported per user-managed custom ruleset that is created through this endpoint."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the custom ruleset.")),
			mcp.WithString("description", mcp.Description("An optional description of the custom ruleset.")),
			mcp.WithString("file", mcp.Required(), mcp.Description("The rules file to upload. Must be a valid YAML file under 1 MB. The name of the file will be used as the ruleFileName property of the custom ruleset. The length of the file name must be under 128 characters and follow the pattern ^[a-zA-Z0-9_.-]+.ya?ml$. Path to a local file to upload.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityCustomRulesetsCreate(params)
			return handleResponse(resp, err, "failed to create Container Security custom ruleset")
		},
	}
}

func toolContainerSecurityCustomRulesetDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_custom_ruleset_delete",
			mcp.WithDescription("Delete Container Security custom ruleset. Deletes the specified Container Security custom ruleset."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the custom ruleset.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityCustomRulesetDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete Container Security custom ruleset")
		},
	}
}

func toolContainerSecurityCustomRulesetGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_custom_ruleset_get",
			mcp.WithDescription("Get ruleset details. Displays the details of the specified Container Security custom ruleset."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the custom ruleset.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityCustomRulesetGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get ruleset details")
		},
	}
}

func toolContainerSecurityCustomRulesetRuleFileSourceGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_custom_ruleset_rule_file_source_get",
			mcp.WithDescription("Fetch custom ruleset file content. Downloads the specified Container Security custom ruleset YAML file."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the custom ruleset.")),
			mcp.WithString("ruleFileSourceId", mcp.Required(), mcp.Description("The ID of the custom ruleset file source.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			ruleFileSourceId, err := pathValue("ruleFileSourceId", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityCustomRulesetRuleFileSourceGet(id, ruleFileSourceId, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to fetch custom ruleset file content")
		},
	}
}

func toolContainerSecurityCustomRulesetUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_custom_ruleset_update",
			mcp.WithDescription("Updates Container Security custom ruleset. Updates the custom ruleset with a new description or rule file. Omitted fields will not be updated. Only userUploaded type custom rulesets can be updated with this API."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the custom ruleset.")),
			mcp.WithString("description", mcp.Description("An optional description of the custom ruleset.")),
			mcp.WithString("file", mcp.Description("The rules file to upload. Must be a valid YAML file under 1 MB. The name of the file will be used as the ruleFileName property of the custom ruleset. The length of the file name must be under 128 characters and follow the pattern ^[a-zA-Z0-9_.-]+.ya?ml$. Path to a local file to upload.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyMultipart,
				BodyFields: []paramDef{{Name: "description", Kind: kindString}},
				FileFields: []paramDef{{Name: "file", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityCustomRulesetUpdate(id, params)
			return handleResponse(resp, err, "failed to updates Container Security custom ruleset")
		},
	}
}

func toolContainerSecurityFileIntegrityMonitoringRulesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_file_integrity_monitoring_rules_list",
			mcp.WithDescription("Get FIM rules. Displays your FIM rules in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The fields by which the results are sorted. You can indicate multiple fields separated by commas (,). Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("The filter that retrieves a subset of the FIM rules list, which is included in every request that generates paginated output. Supported fields: name - The name of the FIM rule; type - The management type of the FIM rule. Supported values: [ custom, managed ]; Supported operators: eq - Operator \"equal to\"; not - Operator \"not\"; and - Operator \"and\"; or - Operator \"or\"; () - Symbols for grouping operands; Supported functions: contains - The string partially matches; Example: contains(name,'my-rule') and type eq 'custom'.")),
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

			resp, err := client.ContainerSecurityFileIntegrityMonitoringRulesList(params)
			return handleResponse(resp, err, "failed to get FIM rules")
		},
	}
}

func toolContainerSecurityFileIntegrityMonitoringRulesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_file_integrity_monitoring_rules_create",
			mcp.WithDescription("Create a FIM rule. Creates a new File Integrity Monitoring (FIM) rule."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name", mcp.Required(), mcp.Description("A descriptive name for the FIM rule.")),
			mcp.WithString("description", mcp.Description("A description for the FIM rule.")),
			mcp.WithObject("content", mcp.Required(), mcp.Description("The content configuration for a FIM rule."), mcp.Properties(map[string]any{"excludedFileNames": map[string]any{"description": "File name exclusion patterns. Supports wildcards (* and ?). Validation rules: - Max length: 255 characters per pattern - Max wildcards: 10...", "items": map[string]any{"type": "string"}, "type": "array"}, "scanDirectoryPaths": map[string]any{"description": "The directory paths to scan for FIM. Each path supports wildcards (* and ?).", "items": map[string]any{"type": "string"}, "type": "array"}, "scanFileNames": map[string]any{"description": "File name inclusion patterns. Supports wildcards (* and ?). Validation rules: - Max length: 255 characters per pattern - Max wildcards: 10...", "items": map[string]any{"type": "string"}, "type": "array"}, "scanTarget": map[string]any{"description": "Specifies whether to scan the container's file system or the host's file system.", "enum": []string{"container", "host"}, "type": "string"}})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "content", Kind: kindObject, Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityFileIntegrityMonitoringRulesCreate(params)
			return handleResponse(resp, err, "failed to create a FIM rule")
		},
	}
}

func toolContainerSecurityFileIntegrityMonitoringRuleDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_file_integrity_monitoring_rule_delete",
			mcp.WithDescription("Delete FIM rule. Deletes the specified FIM rule."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the FIM rule.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityFileIntegrityMonitoringRuleDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete FIM rule")
		},
	}
}

func toolContainerSecurityFileIntegrityMonitoringRuleGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_file_integrity_monitoring_rule_get",
			mcp.WithDescription("Get FIM rule. Displays the details of the specified FIM rule."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the FIM rule.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityFileIntegrityMonitoringRuleGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get FIM rule")
		},
	}
}

func toolContainerSecurityFileIntegrityMonitoringRuleUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_file_integrity_monitoring_rule_update",
			mcp.WithDescription("Update FIM rule. Updates the specified FIM rule."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the FIM rule.")),
			mcp.WithString("name", mcp.Description("A descriptive name for the FIM rule.")),
			mcp.WithString("description", mcp.Description("A description for the FIM rule.")),
			mcp.WithObject("content", mcp.Description("The content configuration for a FIM rule."), mcp.Properties(map[string]any{"excludedFileNames": map[string]any{"description": "File name exclusion patterns. Supports wildcards (* and ?). Validation rules: - Max length: 255 characters per pattern - Max wildcards: 10...", "items": map[string]any{"type": "string"}, "type": "array"}, "scanDirectoryPaths": map[string]any{"description": "The directory paths to scan for FIM. Each path supports wildcards (* and ?).", "items": map[string]any{"type": "string"}, "type": "array"}, "scanFileNames": map[string]any{"description": "File name inclusion patterns. Supports wildcards (* and ?). Validation rules: - Max length: 255 characters per pattern - Max wildcards: 10...", "items": map[string]any{"type": "string"}, "type": "array"}, "scanTarget": map[string]any{"description": "Specifies whether to scan the container's file system or the host's file system.", "enum": []string{"container", "host"}, "type": "string"}})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString}, {Name: "description", Kind: kindString}, {Name: "content", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityFileIntegrityMonitoringRuleUpdate(id, params)
			return handleResponse(resp, err, "failed to update FIM rule")
		},
	}
}

func toolContainerSecurityGenerateServiceGatewayPassword(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_generate_service_gateway_password",
			mcp.WithDescription("Generate service gateway authentication password. Generates a service gateway authentication password for Kubernetes cluster auto-registration using helm commands."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.ContainerSecurityGenerateServiceGatewayPassword(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to generate service gateway authentication password")
		},
	}
}

func toolContainerSecurityKubernetesAuditEventLogsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_kubernetes_audit_event_logs_list",
			mcp.WithDescription("Get runtime audit events (Kubernetes). Displays all the Kubernetes audit events in a paginated list. These events are collected from agent audit logs."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval time range, in ISO 8601 format. Default: now - 30 days.")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval time range, in ISO 8601 format. Default: Current date time.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the audit events of your Kubernetes clusters. Include this parameter in every request that generates paginated output. Supported fields: clusterId - The ID of the cluster in Container Security; priority - The priority level for the event. source - The service that the request was made to. Supported operators: eq - Operator \"equal to\"; and - Operator \"and\"; or - Operator \"or\"; not - Operator \"not\"; () - Symbols for grouping operands. Example: priority eq 'Highest'  \n")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityKubernetesAuditEventLogsList(params)
			return handleResponse(resp, err, "failed to get runtime audit events (Kubernetes)")
		},
	}
}

func toolContainerSecurityKubernetesClusterGroupsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_kubernetes_cluster_groups_list",
			mcp.WithDescription("Get Kubernetes cluster group list. Displays all groups of registered Kubernetes clusters that can be managed."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The fields by which the results are sorted. Default: orchestrator desc.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the Kubernetes cluster list. This parameter should be included in every request that generates paginated output. Supported fields: name - The name of the cluster group. Predefined parent groups are named Amazon EKS, Microsoft AKS, Google GKE, Alibaba Cloud ACK, Oracle Cloud OKE, Self-managed. orchestrator - The orchestrator of the Cluster. Support values: [ Self-managed, Amazon EKS, Microsoft AKS, Google GKE, Alibaba Cloud ACK, Oracle Cloud OKE] parentId - The UUID of the parent cluster group. It is not possible to find default groups by searching using the eq '' expression. The default groups are: Self-managed - '00000000-0000-0000-0000-000000000000', Amazon EKS - '00000000-0000-0000-0000-000000000001', Microsoft AKS - '00000000-0000-0000-0000-000000000002', Google GKE - '00000000-0000-0000-0000-000000000003', Alibaba Cloud ACK - '00000000-0000-0000-0000-000000000004', Oracle Cloud OKE - '00000000-0000-0000-0000-000000000005'. Supported operators: eq - Operator \"equal to\"; and - Operator \"and\"; or - Operator \"or\"; not - Operator \"not\"; contains - Operator for searching for a specified string within a field; () - Symbols for grouping operands. Example: name eq 'example_cluster'\n")),
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

			resp, err := client.ContainerSecurityKubernetesClusterGroupsList(params)
			return handleResponse(resp, err, "failed to get Kubernetes cluster group list")
		},
	}
}

func toolContainerSecurityKubernetesClustersCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_kubernetes_clusters_create",
			mcp.WithDescription("Register a cluster (Kubernetes). Registers a Kubernetes cluster resource in Container Inventory."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the cluster.")),
			mcp.WithString("groupId", mcp.Required(), mcp.Description("The ID of the group associated with the cluster To get IDs of the groups within the user's management scope, use the Kubernetes cluster groups API to list these IDs.")),
			mcp.WithString("description", mcp.Description("The description of the cluster.")),
			mcp.WithString("policyId", mcp.Description("The ID of the policy associated with the cluster.")),
			mcp.WithString("resourceId", mcp.Description("The ID of the cluster of a different cloud provider.")),
			mcp.WithArray("customizableTagIds", mcp.Description("The custom tags and platform tags associated with the cluster."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithString("connectionType", mcp.Description("Connection type for the cluster. Default: direct."), mcp.Enum("direct", "customProxy", "serviceGateway")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString, Required: true}, {Name: "groupId", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "policyId", Kind: kindString}, {Name: "resourceId", Kind: kindString}, {Name: "customizableTagIds", Kind: kindArray}, {Name: "connectionType", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityKubernetesClustersCreate(params)
			return handleResponse(resp, err, "failed to register a cluster (Kubernetes)")
		},
	}
}

func toolContainerSecurityKubernetesClusterDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_kubernetes_cluster_delete",
			mcp.WithDescription("Remove cluster from Container Security (Kubernetes). Unregisters and removes a Kubernetes cluster from Container Security."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the cluster.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityKubernetesClusterDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to remove cluster from Container Security (Kubernetes)")
		},
	}
}

func toolContainerSecurityKubernetesClusterUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_kubernetes_cluster_update",
			mcp.WithDescription("Modify cluster settings (Kubernetes). Modifies the settings of the specified Kubernetes cluster."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the cluster.")),
			mcp.WithString("description", mcp.Description("The description of the cluster.")),
			mcp.WithString("policyId", mcp.Description("The ID of the policy associated with the cluster.")),
			mcp.WithString("resourceId", mcp.Description("The ID of the cluster of a different cloud provider.")),
			mcp.WithString("groupId", mcp.Description("The ID of the group associated with the cluster.")),
			mcp.WithArray("customizableTagIds", mcp.Description("The custom tags and platform tags associated with the cluster."), mcp.Items(map[string]any{"type": "string"})),
			mcp.WithObject("connection", mcp.Properties(map[string]any{"type": map[string]any{"description": "The type of connection setting.", "enum": []string{"direct", "customProxy", "serviceGateway"}, "type": "string"}})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "description", Kind: kindString}, {Name: "policyId", Kind: kindString}, {Name: "resourceId", Kind: kindString}, {Name: "groupId", Kind: kindString}, {Name: "customizableTagIds", Kind: kindArray}, {Name: "connection", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityKubernetesClusterUpdate(id, params)
			return handleResponse(resp, err, "failed to modify cluster settings (Kubernetes)")
		},
	}
}

func toolContainerSecurityKubernetesEvaluationEventLogsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_kubernetes_evaluation_event_logs_list",
			mcp.WithDescription("Get evaluation event logs (Kubernetes). Displays all the evaluation events of your Kubernetes clusters in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval time range, in ISO 8601 format. Default: now() - 30 days.")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval time range, in ISO 8601 format. Default: Current date time.")),
			mcp.WithString("orderBy", mcp.Description("The field used to sort results. Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the evaluation events of your Kubernetes clusters. Include this parameter in every request that generates paginated output. Supported fields: action - The policy action for the event - allow, block, log; clusterId - The ID of the cluster in Container Security - Any value; clusterName - The name of the cluster - Any value; decision - The evaluation decision for the event - allow, deny; mitigation - The policy mitigation action for the event - log, isolate, terminate; policyId - The ID of the policy for an event - Any value; policyName - The name of a policy for an event - Any value; kind - The type of Kubernetes workload for the event - Pod, EphemeralContainers, ReplicaSet, ReplicationController, Deployment, StatefulSet, DaemonSet, Job, CronJob; workloadName - The name of the Kubernetes workload for the event - Any value; namespace - The cluster namespace of the event - Any value; operation - The operation of the evaluation event - Any value; Supported operators: eq - Operator \"equal to\"; and - Operator \"and\"; or - Operator \"or\"; contains - Operator that allows you to search for a specified string in a field; not - Operator \"not\"; () - Symbols for grouping operands. Example: contains(clusterName,'example_cluster') and action eq 'allow'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "orderBy", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityKubernetesEvaluationEventLogsList(params)
			return handleResponse(resp, err, "failed to get evaluation event logs (Kubernetes)")
		},
	}
}

func toolContainerSecurityKubernetesImageOccurrencesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_kubernetes_image_occurrences_list",
			mcp.WithDescription("Get image occurrences (Kubernetes). Displays the occurrences of Kubernetes images that are running in all clusters for your account."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the Kubernetes image occurences list. Include this parameter in every request that generates paginated output. Supported fields: clusterid - The ID of the Kubernetes cluster; imageid - The The ID of the container image; Supported operators: eq - Operator \"equal to\"; and - Operator \"and\"; or - Operator \"or\"; not - Operator \"not\"; () - Symbols for grouping operands. Example: imageId eq 'imageId_1' and clusterId eq 'clusterId_1'")),
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

			resp, err := client.ContainerSecurityKubernetesImageOccurrencesList(params)
			return handleResponse(resp, err, "failed to get image occurrences (Kubernetes)")
		},
	}
}

func toolContainerSecurityKubernetesSensorEventLogsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_kubernetes_sensor_event_logs_list",
			mcp.WithDescription("Get runtime sensor events (Kubernetes). Displays a list of sensor events on your Kubernetes clusters."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval time range, in ISO 8601 format. Default: now() - 30 days.")),
			mcp.WithString("endDateTime", mcp.Description("The end of the data retrieval time range, in ISO 8601 format. Default: Current date time.")),
			mcp.WithString("filter", mcp.Description("The Filter for retrieving a subset of the sensor events of your Kubernetes clusters. Include this parameter in every request that generates paginated output. Supported fields: clusterId - The ID of the cluster in Container Security - Any value; clusterName - The name of the cluster - Any value; policyId - The ID of the policy for an event - Any value; policyName - The name of a policy for an event - Any value; mitigation - The mitigation action to take when a rule fails during runtime. - Any value; ruleName - The name of the rule that the event triggers - Any; ruleId - The ID of the rule that the event triggers - Any; eventType - The event type of the sensor event - malware, syscall, secretScan, auditLog, fileIntegrity; ruleType - The rule type of the sensor event - managed, custom; Supported operators: eq - Operator \"equal to\"; and - Operator \"and\"; or - Operator \"or\"; contains - Operator that allows you to search for a specified string in a field; not - Operator \"not\"; () - Symbols for grouping operands. Example: contains(clusterName,'example_cluster') and policyName eq 'mitre'")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}},
				Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityKubernetesSensorEventLogsList(params)
			return handleResponse(resp, err, "failed to get runtime sensor events (Kubernetes)")
		},
	}
}

func toolContainerSecurityManagedRulesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_managed_rules_list",
			mcp.WithDescription("Get managed runtime rules. Displays the list of runtime rules managed by Container Security in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("orderBy", mcp.Description("The fields by which the results are sorted. You can indicate multiple fields separated by commas (,). Default: createdDateTime desc.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the runtime rules list. Include this parameter in every request that generates paginated output. Supported fields and operators: name - The name of the runtime rule; Supported operators: and - Operator \"and\"; or - Operator \"or\"; eq - Operator \"equal to\"; contains - Operator that allows you to search for a specified string in a field; not - Operator \"not\"; () - Symbols for grouping operands. Example: contains(name,'rules_name')\n")),
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

			resp, err := client.ContainerSecurityManagedRulesList(params)
			return handleResponse(resp, err, "failed to get managed runtime rules")
		},
	}
}

func toolContainerSecurityManagedRuleGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_managed_rule_get",
			mcp.WithDescription("Get runtime rule details. Displays the details of the specified runtime rule managed by Container Security."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the runtime rule managed by Container Security.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityManagedRuleGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get runtime rule details")
		},
	}
}

func toolContainerSecurityPoliciesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_policies_list",
			mcp.WithDescription("Get Container Security policies. Displays your Container Security policies in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The fields by which the results are sorted. You can indicate multiple fields separated by commas (,). Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the policies list. Include this parameter in every request that generates paginated output. Supported fields: name - The name of the policy; xdrEnabled - Whether XDR telemetry is enabled; type - The type of the policy. Supported values: userManaged, clusterManaged; Supported operators: eq - Operator \"equal to\"; and - Operator \"and\"; or - Operator \"or\"; not - Operator \"not\"; contains - Operator that allows you to search for a specified string in a field; () - Symbols for grouping operands. Example: contains(name,'policy_name') and xdrEnabled eq true and not xdrEnabled eq false")),
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

			resp, err := client.ContainerSecurityPoliciesList(params)
			return handleResponse(resp, err, "failed to get Container Security policies")
		},
	}
}

func toolContainerSecurityPoliciesCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_policies_create",
			mcp.WithDescription("Create a Container Security policy. Creates a new Container Security policy."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name", mcp.Required(), mcp.Description("A descriptive name for the policy.")),
			mcp.WithString("description", mcp.Description("(optional) A description of the policy.")),
			mcp.WithBoolean("xdrEnabled", mcp.Description("If true, enables XDR telemetry. Default: True.")),
			mcp.WithObject("malwareScan", mcp.Description("The Runtime Malware Scanning configuration for this policy."), mcp.Properties(map[string]any{"fileScan": map[string]any{"description": "File scan filter configuration. If omitted, not included in the response.", "properties": map[string]any{"filter": map[string]any{"description": "The file scan filter settings. Contains enabled flag, mode, and filter rules. Optional — if omitted, no file scan filter is applied.", "type": "object"}}, "type": "object"}, "imageScan": map[string]any{"description": "Image scan filter configuration. If omitted, not included in the response.", "properties": map[string]any{"filter": map[string]any{"description": "The image scan filter settings. Contains enabled flag, mode, and filter rules. Optional — if omitted, no image scan filter is applied.", "type": "object"}}, "type": "object"}, "mitigation": map[string]any{"description": "The mitigation action for malware.", "enum": []string{"log", "isolate", "terminate"}, "type": "string"}, "schedule": map[string]any{"description": "The schedule for the malware scan. If the schedule is not configured, the scheduled scan will be disabled, and the cron configuration will...", "properties": map[string]any{"cron": map[string]any{"description": "The cron expression for the schedule currently supports only daily and weekly schedules.", "type": "string"}}, "required": []string{"cron"}, "type": "object"}})),
			mcp.WithObject("secretScan", mcp.Description("The Runtime Secret Scanning configuration for the policy."), mcp.Properties(map[string]any{"excludedPaths": map[string]any{"description": "The list of file paths excluded from Runtime Secret Scans. You can use wildcards to match multiple files or directories. For example:", "items": map[string]any{"type": "string"}, "type": "array"}, "mitigation": map[string]any{"description": "The mitigation or action to apply to the secret event type.", "enum": []string{"log"}, "type": "string"}, "schedule": map[string]any{"description": "The schedule for the Runtime Secret Scan. If the schedule is not configured, the scheduled scan is disabled, and the cron configuration is...", "properties": map[string]any{"cron": map[string]any{"description": "The cron expression for the scheduled scan based on UTC. Only daily and weekly schedules are currently supported.", "type": "string"}, "skipIfRuleNotChanged": map[string]any{"description": "Whether scheduled Runtime Secret Scans will be skipped. If true, scheduled Runtime Secret Scans of the container image are skipped when the...", "type": "boolean"}}, "required": []string{"cron"}, "type": "object"}})),
			mcp.WithString("platform", mcp.Description("The platform the policy will be applied to. Cannot be updated after the policy is created. Default: kubernetes."), mcp.Enum("kubernetes")),
			mcp.WithObject("default", mcp.Required(), mcp.Properties(map[string]any{"exceptions": map[string]any{"description": "The set of policy rules. The rules are OR together.", "items": map[string]any{"type": "object"}, "type": "array"}, "rules": map[string]any{"description": "The set of policy rules. The rules are OR together.", "items": map[string]any{"type": "object"}, "type": "array"}})),
			mcp.WithArray("namespaced", mcp.Description("The definition of all the namespaced policies. Properties specified in a namespaced item override the corresponding parent configuration; properties not specified inherit from the parent. - rules, exceptions — override the corresponding fields under default."), mcp.Items(map[string]any{"properties": map[string]any{"exceptions": map[string]any{"description": "The set of policy rules. The rules are OR together.", "items": map[string]any{"type": "object"}, "type": "array"}, "malwareScan": map[string]any{"description": "Malware scan filter configuration. Only imageScan and fileScan filter fields are used.", "properties": map[string]any{"fileScan": map[string]any{"description": "File scan filter configuration.", "type": "object"}, "imageScan": map[string]any{"description": "Image scan filter configuration.", "type": "object"}}, "type": "object"}, "name": map[string]any{"description": "Descriptive name for the namespaced policy definition.", "type": "string"}, "namespaces": map[string]any{"description": "The namespaces that are associated with this policy definition.", "items": map[string]any{"type": "string"}, "type": "array"}, "rules": map[string]any{"description": "The set of policy rules. The rules are OR together.", "items": map[string]any{"type": "object"}, "type": "array"}}, "required": []string{"name", "namespaces", "rules"}, "type": "object"})),
			mcp.WithObject("runtime", mcp.Description("The runtime properties for this policy."), mcp.Properties(map[string]any{"customRuleset": map[string]any{"description": "The custom ruleset associated with this policy.", "properties": map[string]any{"id": map[string]any{"description": "The ID of the custom ruleset.", "type": "string"}}, "required": []string{"id"}, "type": "object"}, "rulesets": map[string]any{"description": "The list of runtime rulesets associated with this policy.", "items": map[string]any{"type": "object"}, "type": "array"}})),
			mcp.WithObject("fileIntegrityMonitoringScan", mcp.Description("The Kubernetes File Integrity Monitoring Scanning configuration for creating a policy."), mcp.Properties(map[string]any{"rules": map[string]any{"description": "The list of Kubernetes FIM rules associated with this policy.", "items": map[string]any{"description": "References an existing FIM rule with a Kubernetes-specific condition that specifies where to apply it when creating a policy.", "type": "object"}, "type": "array"}, "schedule": map[string]any{"description": "The schedule for the FIM scan. If the schedule is not configured, the scheduled scan will be disabled, and the cron configuration will be...", "properties": map[string]any{"cron": map[string]any{"description": "The cron expression for the scheduled scan based on UTC. Only daily and weekly schedules are currently supported.", "type": "string"}}, "required": []string{"cron"}, "type": "object"}, "sha1Enabled": map[string]any{"description": "Whether SHA1 checksum is calculated for the container image. If true, the SHA1 checksum is calculated during the scan. Defaults to false.", "type": "boolean"}})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "xdrEnabled", Kind: kindBoolean}, {Name: "malwareScan", Kind: kindObject}, {Name: "secretScan", Kind: kindObject}, {Name: "platform", Kind: kindString}, {Name: "default", Kind: kindObject, Required: true}, {Name: "namespaced", Kind: kindArray}, {Name: "runtime", Kind: kindObject}, {Name: "fileIntegrityMonitoringScan", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityPoliciesCreate(params)
			return handleResponse(resp, err, "failed to create a Container Security policy")
		},
	}
}

func toolContainerSecurityPolicyDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_policy_delete",
			mcp.WithDescription("Delete a Container Security policy. Deletes the specified Container Security policy."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The The ID of the container Security policy.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityPolicyDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete a Container Security policy")
		},
	}
}

func toolContainerSecurityPolicyGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_policy_get",
			mcp.WithDescription("Get Container Security policy. Displays the details of the specified Container Security policy."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The The ID of the container Security policy.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityPolicyGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get Container Security policy")
		},
	}
}

func toolContainerSecurityPolicyUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_policy_update",
			mcp.WithDescription("Modify a Container Security policy. Modifies the settings of the specified Container Security policy."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The The ID of the container Security policy.")),
			mcp.WithString("description", mcp.Description("(optional) A description of the policy.")),
			mcp.WithBoolean("xdrEnabled", mcp.Description("If true, enables XDR telemetry.")),
			mcp.WithObject("malwareScan", mcp.Description("The Runtime Malware Scanning configuration for this policy."), mcp.Properties(map[string]any{"fileScan": map[string]any{"description": "File scan filter configuration. If omitted, the existing configuration is unchanged.", "properties": map[string]any{"filter": map[string]any{"description": "The file scan filter settings. Contains enabled flag and filter rules. Optional — if omitted, no file scan filter is applied.", "type": "object"}}, "type": "object"}, "imageScan": map[string]any{"description": "Image scan filter configuration. If omitted, the existing configuration is unchanged.", "properties": map[string]any{"filter": map[string]any{"description": "The image scan filter settings. Contains enabled flag, mode, and filter rules. Optional — if omitted, no image scan filter is applied.", "type": "object"}}, "type": "object"}, "mitigation": map[string]any{"description": "The mitigation action for malware.", "enum": []string{"log", "isolate", "terminate"}, "type": "string"}, "schedule": map[string]any{"properties": map[string]any{"cron": map[string]any{"description": "The cron expression for the scheduled scan based on UTC. Only daily and weekly schedules are currently supported.", "type": "string"}, "enabled": map[string]any{"description": "If true, enables scheduled scan. If the schedule has been configured and the new schedule is not provided, it will apply the configured...", "type": "boolean"}}, "type": "object"}})),
			mcp.WithObject("secretScan", mcp.Description("The Runtime Secret Scanning configuration for the policy."), mcp.Properties(map[string]any{"excludedPaths": map[string]any{"description": "The list of file paths excluded from Runtime Secret Scans. You can use wildcards to match multiple files or directories. For example:", "items": map[string]any{"type": "string"}, "type": "array"}, "mitigation": map[string]any{"description": "The mitigation or action to apply to the secret event type.", "enum": []string{"log"}, "type": "string"}, "schedule": map[string]any{"properties": map[string]any{"cron": map[string]any{"description": "The cron expression for the scheduled scan based on UTC. Only daily and weekly schedules are currently supported.", "type": "string"}, "enabled": map[string]any{"description": "Whether the Runtime Secret Scan is enabled. If true and a new schedule is not configured, the previously configured schedule is applied.", "type": "boolean"}, "skipIfRuleNotChanged": map[string]any{"description": "Whether scheduled Runtime Secret Scans will be skipped. If true, scheduled Runtime Secret Scans of the container image are skipped when the...", "type": "boolean"}}, "type": "object"}})),
			mcp.WithObject("default", mcp.Properties(map[string]any{"exceptions": map[string]any{"description": "The set of policy rules. The rules are OR together.", "items": map[string]any{"type": "object"}, "type": "array"}, "rules": map[string]any{"description": "The set of policy rules. The rules are OR together.", "items": map[string]any{"type": "object"}, "type": "array"}})),
			mcp.WithArray("namespaced", mcp.Description("The definition of all the namespaced policies. To clear all namespaced policies, set them to null. e.g. \"namespaced\": null Properties specified in a namespaced item override the corresponding parent configuration;."), mcp.Items(map[string]any{"properties": map[string]any{"exceptions": map[string]any{"description": "The set of policy rules. The rules are OR together.", "items": map[string]any{"type": "object"}, "type": "array"}, "malwareScan": map[string]any{"description": "Malware scan filter configuration. Only imageScan and fileScan filter fields are used.", "properties": map[string]any{"fileScan": map[string]any{"description": "File scan filter configuration.", "type": "object"}, "imageScan": map[string]any{"description": "Image scan filter configuration.", "type": "object"}}, "type": "object"}, "name": map[string]any{"description": "Descriptive name for the namespaced policy definition.", "type": "string"}, "namespaces": map[string]any{"description": "The namespaces that are associated with this policy definition.", "items": map[string]any{"type": "string"}, "type": "array"}, "rules": map[string]any{"description": "The set of policy rules. The rules are OR together.", "items": map[string]any{"type": "object"}, "type": "array"}}, "required": []string{"name", "namespaces", "rules"}, "type": "object"})),
			mcp.WithObject("runtime", mcp.Description("The runtime properties for this policy."), mcp.Properties(map[string]any{"customRuleset": map[string]any{"description": "The custom ruleset associated with this policy.", "properties": map[string]any{"id": map[string]any{"description": "The ID of the custom ruleset.", "type": "string"}}, "required": []string{"id"}, "type": "object"}, "rulesets": map[string]any{"description": "The list of runtime rulesets associated with this policy. To clear rulesets, set them to an empty array. e.g. \"rulesets\": []", "items": map[string]any{"type": "object"}, "type": "array"}})),
			mcp.WithObject("fileIntegrityMonitoringScan", mcp.Description("The Kubernetes FIM Scanning configuration for updating a policy."), mcp.Properties(map[string]any{"rules": map[string]any{"description": "The list of Kubernetes FIM rules associated with this policy.", "items": map[string]any{"description": "References an existing FIM rule with a Kubernetes-specific condition for updating a policy.", "type": "object"}, "type": "array"}, "schedule": map[string]any{"properties": map[string]any{"cron": map[string]any{"description": "The cron expression for the scheduled scan based on UTC. Only daily and weekly schedules are currently supported.", "type": "string"}, "enabled": map[string]any{"description": "If true, enables scheduled scan. If the schedule has been configured and the new schedule is not provided, it will apply the configured...", "type": "boolean"}}, "type": "object"}, "sha1Enabled": map[string]any{"description": "Whether SHA1 checksum is calculated for the container image. If true, the SHA1 checksum is calculated during the scan.", "type": "boolean"}})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "description", Kind: kindString}, {Name: "xdrEnabled", Kind: kindBoolean}, {Name: "malwareScan", Kind: kindObject}, {Name: "secretScan", Kind: kindObject}, {Name: "default", Kind: kindObject}, {Name: "namespaced", Kind: kindArray}, {Name: "runtime", Kind: kindObject}, {Name: "fileIntegrityMonitoringScan", Kind: kindObject}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityPolicyUpdate(id, params)
			return handleResponse(resp, err, "failed to modify a Container Security policy")
		},
	}
}

func toolContainerSecurityRulesetsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_rulesets_list",
			mcp.WithDescription("Get all rulesets. Displays all your Container Security rulesets in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page. One of: 25, 50, 100. Default: 25."), mcp.Enum("25", "50", "100")),
			mcp.WithString("orderBy", mcp.Description("The fields by which the results are sorted. You can indicate multiple fields separated by commas (,). Default: createdDateTime desc.")),
			mcp.WithString("filter", mcp.Description("The filter for retrieving a subset of the rulesets list. Include this parameter in every request that generates paginated output. Supported fields: name - The name of the ruleset; type - The type of the ruleset. Supported values: userManaged, clusterManaged; Supported operators; eq - Operator \"equal to\"; contains - Operator that allows you to search for a specified string in a field; or - Operator \"or\"; not - Operator \"not\". Example: contains(name,'rulesets_name') or name eq 'example_ruleset'")),
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

			resp, err := client.ContainerSecurityRulesetsList(params)
			return handleResponse(resp, err, "failed to get all rulesets")
		},
	}
}

func toolContainerSecurityRulesetsCreate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_rulesets_create",
			mcp.WithDescription("Create a Container Security ruleset. Creates a Container Security ruleset."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("name", mcp.Required(), mcp.Description("The name of the ruleset.")),
			mcp.WithString("description", mcp.Description("If you provided a description for the ruleset, it is returned here.")),
			mcp.WithString("labelMatchingCondition", mcp.Description("Specifies whether to include or exclude the pods that match the labels list. If \"any\" is selected, the ruleset is applied to pods that include any labels. If \"notAny\" is selected, the ruleset is applied to pods that do not include any labels. Default: any."), mcp.Enum("any", "notAny")),
			mcp.WithArray("labels", mcp.Description("The key and value of the container object. If no labels are specified, rulesets are applied to all pods despite the \"labelMatchingCondition\" value."), mcp.Items(map[string]any{"properties": map[string]any{"key": map[string]any{"description": "The key of the container object label.", "type": "string"}, "value": map[string]any{"description": "The value of the container object label.", "type": "string"}}, "required": []string{"key", "value"}, "type": "object"})),
			mcp.WithArray("rules", mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The unique ID assigned to the rule.", "type": "string"}, "mitigation": map[string]any{"enum": []string{"log", "isolate", "terminate"}, "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "name", Kind: kindString, Required: true}, {Name: "description", Kind: kindString}, {Name: "labelMatchingCondition", Kind: kindString}, {Name: "labels", Kind: kindArray}, {Name: "rules", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityRulesetsCreate(params)
			return handleResponse(resp, err, "failed to create a Container Security ruleset")
		},
	}
}

func toolContainerSecurityRulesetDelete(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_ruleset_delete",
			mcp.WithDescription("Delete a ruleset. Deletes the specified Container Security ruleset."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint:    toPtr(false),
				DestructiveHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the ruleset.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityRulesetDelete(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to delete a ruleset")
		},
	}
}

func toolContainerSecurityRulesetGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_ruleset_get",
			mcp.WithDescription("Get ruleset details. Displays the details of the specified Container Security ruleset."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the ruleset.")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityRulesetGet(id, v1client.RequestParams{})
			return handleResponse(resp, err, "failed to get ruleset details")
		},
	}
}

func toolContainerSecurityRulesetUpdate(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_ruleset_update",
			mcp.WithDescription("Modify a ruleset. Modifies specified Container Security ruleset."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
			mcp.WithString("id", mcp.Required(), mcp.Description("The ID of the ruleset.")),
			mcp.WithString("description", mcp.Description("If you provided a description for the ruleset, it is returned here.")),
			mcp.WithString("labelMatchingCondition", mcp.Description("Specifies whether to include or exclude the pods that match the labels list. If \"any\" is selected, the ruleset is applied to pods that include any labels. If \"notAny\" is selected, the ruleset is applied to pods that do not include any labels."), mcp.Enum("any", "notAny")),
			mcp.WithArray("labels", mcp.Description("The key and value of the container object. If no labels are specified, rulesets are applied to all pods despite the \"labelMatchingCondition\" value."), mcp.Items(map[string]any{"properties": map[string]any{"key": map[string]any{"description": "The key of the container object label.", "type": "string"}, "value": map[string]any{"description": "The value of the container object label.", "type": "string"}}, "required": []string{"key", "value"}, "type": "object"})),
			mcp.WithArray("rules", mcp.Items(map[string]any{"properties": map[string]any{"id": map[string]any{"description": "The unique ID assigned to the rule.", "type": "string"}, "mitigation": map[string]any{"enum": []string{"log", "isolate", "terminate"}, "type": "string"}}, "required": []string{"id"}, "type": "object"})),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			id, err := pathValue("id", args)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			params, err := buildRequestParams(args, requestSpec{
				Body:       bodyObject,
				BodyFields: []paramDef{{Name: "description", Kind: kindString}, {Name: "labelMatchingCondition", Kind: kindString}, {Name: "labels", Kind: kindArray}, {Name: "rules", Kind: kindArray}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.ContainerSecurityRulesetUpdate(id, params)
			return handleResponse(resp, err, "failed to modify a ruleset")
		},
	}
}

func toolContainerSecurityStartComplianceScan(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"container_security_start_compliance_scan",
			mcp.WithDescription("Start a compliance scan. Starts a compliance scan in your managed clusters."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(false),
			}),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			resp, err := client.ContainerSecurityStartComplianceScan(v1client.RequestParams{})
			return handleResponse(resp, err, "failed to start a compliance scan")
		},
	}
}
