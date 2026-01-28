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

			resp, err := client.ContainerSecurityListK8Clusters(filter, qp)
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
			return handleStatusResponse(resp, err, http.StatusOK, "failed to list ecs clusters")
		},
	}
}
