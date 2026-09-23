package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

var ToolsetsReadOnlySearch = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolSearchActivityStatisticsGet,
	toolSearchCloudActivitiesList,
	toolSearchContainerActivitiesList,
	toolSearchDetectionsList,
	toolSearchEmailActivitiesList,
	toolSearchEndpointActivitiesList,
	toolSearchIdentityActivitiesList,
	toolSearchMobileActivitiesList,
	toolSearchNetworkActivitiesList,
	toolSearchSensorStatisticsGet,
}

func toolSearchActivityStatisticsGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_activity_statistics_get",
			mcp.WithDescription("Query activity data statistics. Query the connected product status for the specified time range."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("period", mcp.Description("The time period for the TrendAI Vision One™ connected product status. Default: 7 days. Default: 7d."), mcp.Enum("24h", "7d", "30d")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "period", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchActivityStatisticsGet(params)
			return handleResponse(resp, err, "failed to query activity data statistics")
		},
	}
}

func toolSearchCloudActivitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_cloud_activities_list",
			mcp.WithDescription("Get cloud activity data. Displays search results from the Cloud Activity Data source in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format. Default:startDateTime defaults to 24 hours before the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page in default mode. One of: 50, 100, 500, 1000, 5000. Default: 500."), mcp.Enum("50", "100", "500", "1000", "5000")),
			mcp.WithString("mode", mcp.Description("The type of data returned by the query results (\"default\" displays all data returned by the query, \"countOnly\" displays the number of records returned, and \"performance\" limits the returned records to 500.): * 'default' - Displays the data returned by the... Default: default."), mcp.Enum("default", "countOnly", "performance")),
			mcp.WithString("select", mcp.Description("The list of Trend Cloud One - AWS CloudTrail fields to include in the search results. Note: If no fields are specified, the query returns all supported fields. Supported fields: eventID; eventName; eventSource; readOnly; requestParameters; resources; responseElements; sourceIPAddress; userAgent; userIdentity; vpcEndpointId; uuid; productCode; tags; The list of Trend Cloud One - AWS CloudTrail fields to include in the search results. Note: If no fields are specified, the query returns all supported fields. Supported fields: dst; src; pname; spt; dpt; start; end; vpcFlowLogsVersion; cloudAccountId; networkInterfaceId; ipProto; packets; bytes; action; logStatus; vpcId; subnetId; instanceId; tcpFlags; flowType; pktSrcAddr; pktDstAddr; regionCode; azId; subLocationType; subLocationId; pktSrcCloudServiceName; pktDstCloudServiceName; flowDirection; trafficPath; productCode; filterRiskLevel. Default: empty. Example: eventName,tags,uuid")),
			mcp.WithString("query", mcp.Required(), mcp.Description("The statement that retrieves a subset of the collected cloud activity data from Trend Cloud One - AWS CloudTrail. Supported fields: 'eventID' - Match type support: Partial match; 'eventName' - Match type support: Partial match; 'eventSource' - Match type support: Partial match; 'requestParameters' - Match type support: Partial match; 'resources' - Match type support: Partial match; 'responseElements' - Match type support: Partial match; 'sourceIPAddress' - Match type support: Partial match; 'userAgent' - Match type support: Partial match; 'userIdentity' - Match type support: Partial match; 'vpcEndpointId' - Match type support: Partial match; 'uuid' - Match type support: Partial match; 'productCode' - Match type support: Partial match; 'tags' - Match type support: Partial match; Supported operators: ':' - Abbreviation of the operator \"equal to\"; 'and' - Operator \"and\"; 'or' - Operator \"or\"; 'not' - Operator \"not\"; '( )' - Symbols for grouping operands with their correct operator. For more information about query syntax and operators, see the online help. The statement that retrieves a subset of the collected cloud activity data from XDR for Cloud - AWS VPC Flow Logs. Supported fields: ' dst ' - Match type support: Partial match; ' src ' - Match type support: Partial match; ' pname ' - Match type support: Partial match; ' dpt ' - Match type support: Full match. Example: eventName:xxxxxxxxxxxxx")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "mode", Kind: kindString}, {Name: "select", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchCloudActivitiesList(params)
			return handleResponse(resp, err, "failed to get cloud activity data")
		},
	}
}

func toolSearchContainerActivitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_container_activities_list",
			mcp.WithDescription("Get Container Activity Data. Displays search results from the Container Activity Data source in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format. Default:startDateTime defaults to 24 hours before the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page in default mode. One of: 50, 100, 500, 1000, 5000. Default: 500."), mcp.Enum("50", "100", "500", "1000", "5000")),
			mcp.WithString("mode", mcp.Description("The type of data returned by the query results (\"default\" displays all data returned by the query, \"countOnly\" displays the number of records returned, and \"performance\" limits the returned records to 500.): * 'default' - Displays the data returned by the... Default: default."), mcp.Enum("default", "countOnly", "performance")),
			mcp.WithString("select", mcp.Description("The list of fields included in search results (If no fields are specified, the query returns all supported fields.). Default: empty. Example: dpt,dst,objectFilePath")),
			mcp.WithString("query", mcp.Required(), mcp.Description("Statement that allows you to retrieve a subset of the collected secure access activity data. Supported fields: 'eventId' - Full match; 'eventSubId' - Full match; 'eventTime' - Full match; 'objectFilePath' - Partial match; 'srcFilePath' - Partial match; 'tags' - Partial match; 'uuid' - Partial match; 'productCode' - Partial match; 'filterRiskLevel' - Partial match; 'clusterId' - Partial match; 'clusterName' - Partial match; 'k8sNamespace' - Partial match; 'containerName' - Partial match; 'containerId' - Partial match; 'containerImage' - Partial match; 'processCmd' - Partial match; 'parentCmd' - Partial match; 'processFilePath' - Partial match; 'parentFilePath' - Partial match; 'processName' - Partial match; 'processPid' - Full match; 'parentPid' - Full match; 'dpt' - Full match; 'dst' - Partial match; 'spt' - Full match; 'src' - Partial match; 'objectCmd' - Partial match; 'objectUser' - Partial match; 'pname' - Partial match; Supported operators: ':' - Abbreviation of the operator \"equal to\"; 'and' - Operator \"and\"; 'or' - Operator \"or\"; 'not' - Operator \"not\"; '( )' - Symbols for grouping operands with their correct operator. For more information about query syntax and operators, see the online help. Note: For more information about supported fields, see the online help. Example: objectFilePath:\"/opt/nimsoft/probes/system/processes/processes\" and eventId:3")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "mode", Kind: kindString}, {Name: "select", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchContainerActivitiesList(params)
			return handleResponse(resp, err, "failed to get Container Activity Data")
		},
	}
}

func toolSearchDetectionsList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_detections_list",
			mcp.WithDescription("Get detection data. Displays search results from the Detection Data source in a paginated list. The primary purpose of the \"Get Detection Data\" API is to provide customers with a querying capability, rather than exporting all logs."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format. Default:startDateTime defaults to 24 hours before the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page in default mode. One of: 50, 100, 500, 1000, 5000. Default: 500."), mcp.Enum("50", "100", "500", "1000", "5000")),
			mcp.WithString("mode", mcp.Description("The type of data returned by the query results (\"default\" displays all data returned by the query, \"countOnly\" displays the number of records returned, and \"performance\" limits the returned records to 500.): * 'default' - Displays the data returned by the... Default: default."), mcp.Enum("default", "countOnly", "performance")),
			mcp.WithString("select", mcp.Description("The list of fields included in search results (If no fields are specified, the query returns all supported fields.). Default: empty. Example: hostName,userDomain,endpointGUID")),
			mcp.WithString("query", mcp.Required(), mcp.Description("Statement that allows you to retrieve a subset of the collected detection data. Supported fields: 'hostName' - Match type support: Partial match; 'interestedHost' - Match type support: Partial match; 'shost' - Match type support: Partial match; 'dhost' - Match type support: Partial match; 'endpointHostName' - Match type support: Partial match; 'userDomain' - Match type support: Partial match; 'endpointGUID' - Match type support: Partial match; 'request' - Match type support: Partial match; 'src' - Match type support: Partial match; 'dst' - Match type support: Partial match; 'interestedIp' - Match type support: Partial match; 'endpointIp' - Match type support: Partial match; 'peerIp' - Match type support: Partial match; 'denyListIp' - Match type support: Partial match; 'dpt' - Match type support: Full match; 'spt' - Match type support: Full match; 'fileName' - Match type support: Partial match; 'objectFileName' - Match type support: Partial match; 'compressedFileName' - Match type support: Partial match; 'attachmentFileName' - Match type support: Partial match; 'filePath' - Match type support: Partial match; 'filePathName' - Match type support: Partial match; 'objectFilePath' - Match type support: Partial match; 'fileHash' - Match type support: Full match; 'attachmentFileHash' - Match type support: Full match; 'attachmentFileHashSha1' - Match type support: Full match. Example: spt:443 or attachmentFileName:test.txt")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "mode", Kind: kindString}, {Name: "select", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchDetectionsList(params)
			return handleResponse(resp, err, "failed to get detection data")
		},
	}
}

func toolSearchEmailActivitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_email_activities_list",
			mcp.WithDescription("Get email activity data. Displays search results from the Email Activity Data source in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page in default mode. One of: 50, 100, 500, 1000, 5000. Default: 500."), mcp.Enum("50", "100", "500", "1000", "5000")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format. Default:startDateTime defaults to 24 hours before the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("mode", mcp.Description("The type of data returned by the query results (\"default\" displays all data returned by the query, \"countOnly\" displays the number of records returned, and \"performance\" limits the returned records to 500.): * 'default' - Displays the data returned by the... Default: default."), mcp.Enum("default", "countOnly", "performance")),
			mcp.WithString("select", mcp.Description("List of fields to include in the search results. If no fields are specified, the query returns all supported fields. Default: empty. Example: mailMsgSubject,mailFromAddresses,mailToAddresses")),
			mcp.WithString("query", mcp.Required(), mcp.Description("Statement that allows you to retrieve a subset of the collected email activity data. Supported fields: 'uuid' - Match type support: Partial match; 'tags' - Match type support: Partial match; 'pname' - Match type support: Partial match; 'msgUuid' - Match type support: Partial match; 'mailDirection' - Match type support: Partial match; 'mailFromAddresses' - Match type support: Partial match; 'mailToAddresses' - Match type support: Partial match; 'mailMsgSubject' - Match type support: Partial match; 'mailMsgId' - Match type support: Partial match; 'mailCcAddresses' - Match type support: Partial match; 'mailBccAddresses' - Match type support: Partial match; 'mailSenderIp' - Match type support: Partial match; 'mailAttachmentHash' - Match type support: Partial match; 'attachmentFileName' - Match type support: Partial match; 'attachmentSha1' - Match type support: Full match; 'attachmentMd5' - Match type support: Full match; 'attachmentSha256' - Match type support: Full match; 'attachmentUrls' - Match type support: Partial match; 'orgId' - Match type support: Partial match; 'mailUrlsVisibleLink' - Match type support: Partial match; 'mailUrlsRealLink' - Match type support: Partial match; Supported operators: ':' - Abbreviation of the operator \"equal to\"; 'and' - Operator \"and\"; 'or' - Operator \"or\"; 'not' - Operator \"not\". Example: mailMsgSubject:spam or mailSenderIp:\"192.169.1.1\"")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "mode", Kind: kindString}, {Name: "select", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchEmailActivitiesList(params)
			return handleResponse(resp, err, "failed to get email activity data")
		},
	}
}

func toolSearchEndpointActivitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_endpoint_activities_list",
			mcp.WithDescription("Get endpoint activity data. Displays search results from the Endpoint Activity Data source in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format. Default:startDateTime defaults to 24 hours before the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page in default mode. One of: 50, 100, 500, 1000, 5000. Default: 500."), mcp.Enum("50", "100", "500", "1000", "5000")),
			mcp.WithString("mode", mcp.Description("The type of data returned by the query results (\"default\" displays all data returned by the query, \"countOnly\" displays the number of records returned, and \"performance\" limits the returned records to 500.): * 'default' - Displays the data returned by the... Default: default."), mcp.Enum("default", "countOnly", "performance")),
			mcp.WithString("select", mcp.Description("The list of fields included in search results (If no fields are specified, the query returns all supported fields.). Default: empty. Example: dpt,dst,endpointHostName")),
			mcp.WithString("query", mcp.Required(), mcp.Description("The filter for retrieving a subset of the collected endpoint activity data. Supported fields: dpt - Full match; dst - Partial match; endpointGuid - Partial match; endpointHostName - Partial match; endpointIp - Partial match; eventId - Full match; eventSubId - Full match; hostName - Partial match; logonUser - Partial match; objectAppName - Partial match; objectCmd - Partial match; objectFileHashMd5 - Full match; objectFileHashSha1 - Full match; objectFileHashSha256 - Full match; objectFilePath - Partial match; objectIp - Partial match; objectIps - Partial match; objectPid - Full match; objectPort - Full match; objectProcessHashId - Full match; objectRawDataStr - Partial match; objectRegistryData - Partial match; objectRegistryKeyHandle - Partial match; objectRegistryValue - Partial match; objectSigner - Partial match; objectSignerValid - Partial match; objectUser - Partial match; parentCmd - Partial match; parentFileHashMd5 - Full match; parentFileHashSha1 - Full match; parentFileHashSha256 - Full match; parentFilePath - Partial match; parentPid - Full match; pname - Partial match; processCmd - Partial match; processFileHashMd5 - Full match; processFileHashSha1 - Full match; processFileHashSha256 - Full match; processFilePath - Partial match; processHashId - Full match; processPid - Full match; request - Partial match; spt - Full match; src - Partial match. Example: dpt:443 or src:\"192.169.1.1\"")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "mode", Kind: kindString}, {Name: "select", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchEndpointActivitiesList(params)
			return handleResponse(resp, err, "failed to get endpoint activity data")
		},
	}
}

func toolSearchIdentityActivitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_identity_activities_list",
			mcp.WithDescription("Get Identity and Access Activity Data. Displays search results from the Identity and Access Activity Data source in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page in default mode. One of: 50, 100, 500, 1000, 5000. Default: 500."), mcp.Enum("50", "100", "500", "1000", "5000")),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format. Default:startDateTime defaults to 24 hours before the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("mode", mcp.Description("The type of data returned by the query results (\"default\" displays all data returned by the query, \"countOnly\" displays the number of records returned, and \"performance\" limits the returned records to 500.): * 'default' - Displays the data returned by the... Default: default."), mcp.Enum("default", "countOnly", "performance")),
			mcp.WithString("select", mcp.Description("The list of fields to include in the search results. Default: The query returns all supported fields if no fields are specified. Default: empty. Example: uuid,result,tags")),
			mcp.WithString("query", mcp.Required(), mcp.Description("The statement that allows you to retrieve a subset of the collected Identity and Access Activity Data. Supported fields: 'uuid' - Partial match; 'tags' - Partial match; 'productCode' - Partial match; 'filterRiskLevel' - Partial match; 'eventId' - Partial match; 'eventName' - Partial match; 'idpName' - Partial match; 'idpId' - Partial match; 'locationCountry' - Partial match; 'locationCity' - Partial match; 'locationState' - Partial match; 'locationLongitude' - Partial match; 'locationLatitude' - Partial match; 'clientId' - Partial match; 'clientDisplayName' - Partial match; 'clientOS' - Partial match; 'clientBrowser' - Partial match; 'clientApp' - Partial match; 'ipAddress' - Partial match; 'userId' - Partial match; 'userDisplayName' - Partial match; 'statusDetail' - Partial match; 'status' - Partial match; 'statusReason' - Partial match; 'targetResourceId' - Partial match; 'targetResourceDisplayName' - Partial match; 'requestMethod' - Partial match; 'eventAdditionalDetails' - Partial match; 'eventCategory' - Partial match; 'initiatedByAppId' - Partial match; 'initiatedByAppDisplayName' - Partial match; 'initiatedByServicePrincipalId' - Partial match; 'initiatedByServicePrincipalName' - Partial match; 'initiatedByUserId' - Partial match; 'initiatedByUserDisplayName' - Partial match; 'initiatedByUserHomeTenantId' - Partial match; 'initiatedByUserHomeTenantName' - Partial match. Example: clientDisplayName:* or ipAddress:\"192.169.1.1\"")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "top", Kind: kindString}, {Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "mode", Kind: kindString}, {Name: "select", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchIdentityActivitiesList(params)
			return handleResponse(resp, err, "failed to get Identity and Access Activity Data")
		},
	}
}

func toolSearchMobileActivitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_mobile_activities_list",
			mcp.WithDescription("Get mobile activity data. Displays search results from the mobile Activity Data source in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format. Default:startDateTime defaults to 24 hours before the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page in default mode. One of: 50, 100, 500, 1000, 5000. Default: 500."), mcp.Enum("50", "100", "500", "1000", "5000")),
			mcp.WithString("mode", mcp.Description("The type of data returned by the query results (\"default\" displays all data returned by the query, \"countOnly\" displays the number of records returned, and \"performance\" limits the returned records to 500.): * 'default' - Displays the data returned by the... Default: default."), mcp.Enum("default", "countOnly", "performance")),
			mcp.WithString("select", mcp.Description("The list of fields included in search results (If no fields are specified, the query returns all supported fields.). Default: empty. Example: endpointGuid,endpointHostName")),
			mcp.WithString("query", mcp.Required(), mcp.Description("Statement that allows you to retrieve a subset of the collected secure access activity data. Supported fields: 'tags' - Match type support: Partial match; 'productCode' - Match type support: Partial match; 'uuid' - Match type support: Partial match; 'filterRiskLevel' - Match type support: Partial match; 'endpointGuid' - Match type support: Partial match; 'endpointHostName' - Match type support: Partial match; 'endpointIp' - Match type support: Partial match; 'eventId' - Match type support: Full match; 'eventSubId' - Match type support: Full match; 'logonUser' - Match type support: Partial match; 'objectFileHashSha256' - Match type support: Partial match; 'objectFilePath' - Match type support: Partial match; 'pname' - Match type support: Partial match; 'request' - Match type support: Partial match; 'srcFileHashSha256' - Match type support: Partial match; 'srcFilePath' - Match type support: Partial match; 'endpointModel' - Match type support: Partial match; 'osName' - Match type support: Partial match; 'appLabel' - Match type support: Partial match; 'appPkgName' - Match type support: Partial match; 'appPublicKeySha1' - Match type support: Partial match; 'objectAppDexSha256' - Match type support: Partial match; 'objectAppLabel' - Match type support: Partial match; 'objectAppPackageName' - Match type support: Partial match; 'objectAppPublicKeySha1' - Match type support: Partial mat. Example: appPkgName:org.mozilla.firefox AND osName:Android")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "mode", Kind: kindString}, {Name: "select", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchMobileActivitiesList(params)
			return handleResponse(resp, err, "failed to get mobile activity data")
		},
	}
}

func toolSearchNetworkActivitiesList(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_network_activities_list",
			mcp.WithDescription("Get network activity data. Displays search results from the Network Activity Data source in a paginated list."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("startDateTime", mcp.Description("The start of the data retrieval range in ISO 8601 format. Default:startDateTime defaults to 24 hours before the request is made.")),
			mcp.WithString("endDateTime", mcp.Description("Timestamp in ISO 8601 format that indicates the end of the data retrieval time range. If no value is specified, 'endDateTime' defaults to the time the request is made.")),
			mcp.WithString("top", mcp.Description("The number of records displayed on a page in default mode. One of: 50, 100, 500, 1000, 5000. Default: 500."), mcp.Enum("50", "100", "500", "1000", "5000")),
			mcp.WithString("mode", mcp.Description("The type of data returned by the query results (\"default\" displays all data returned by the query, \"countOnly\" displays the number of records returned, and \"performance\" limits the returned records to 500.): * 'default' - Displays the data returned by the... Default: default."), mcp.Enum("default", "countOnly", "performance")),
			mcp.WithString("select", mcp.Description("The list of fields included in search results (If no fields are specified, the query returns all supported fields.). Default: empty. Example: dpt,dst,endpointHostName")),
			mcp.WithString("query", mcp.Required(), mcp.Description("Statement that allows you to retrieve a subset of the collected secure access activity data. Supported fields: 'tags' - Partial match; 'productCode' - Partial match; 'uuid' - Partial match; 'osName' - Partial match; 'filterRiskLevel' - Partial match; 'endpointHostName' - Partial match; 'dst' - Partial match; 'src' - Partial match; 'endpointGuid' - Partial match; 'principalName' - Partial match; 'request' - Partial match; 'application' - Partial match; 'ruleName' - Partial match; 'clientIp' - Partial match; 'requestBase' - Partial match; 'eventSubName' - Partial match; 'fileHash' - Partial match; 'fileHashSha256' - Partial match; 'fileName' - Partial match; 'dpt' - Full match; 'fileName' - Partial match; 'act' - Partial match; 'eventName' - Partial match; 'malName' - Partial match; 'detectionType' - Partial match; 'profile' - Partial match; 'deviceGUID' - Partial match; 'flowId' - Partial match; 'serverIp' - Partial match; 'clientPort' - Full match; 'serverPort' - Full match; 'app' - Partial match; 'requestMethod' - Partial match; 'respMethod' - Partial match; 'requestClientApplication' - Partial match; 'httpReferer' - Partial match; 'httpXForwardedForIp' - Partial match; 'resolvedUrlIp' - Partial match; 'resolvedUrlPort' - Full match; 'respCode' - Partial match; 'httpLocation' - Partial match; 'userDomain' - Partial match; 'suid' - Partial match; 'hostName' - Partial match. Example: dpt:443 or src:\"192.169.1.1\"")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query:   []paramDef{{Name: "startDateTime", Kind: kindString}, {Name: "endDateTime", Kind: kindString}, {Name: "top", Kind: kindString}, {Name: "mode", Kind: kindString}, {Name: "select", Kind: kindString}},
				Headers: []headerDef{{Arg: "query", Header: "TMV1-Query", Required: true}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchNetworkActivitiesList(params)
			return handleResponse(resp, err, "failed to get network activity data")
		},
	}
}

func toolSearchSensorStatisticsGet(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"search_sensor_statistics_get",
			mcp.WithDescription("Query endpoint sensor statistics. Query the endpoint sensor status for the specified time range."),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("period", mcp.Description("The time period for the TrendAI Vision One™ connected product status. Default: 7 days. Default: 7d."), mcp.Enum("24h", "7d", "30d")),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			params, err := buildRequestParams(args, requestSpec{
				Query: []paramDef{{Name: "period", Kind: kindString}},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			resp, err := client.SearchSensorStatisticsGet(params)
			return handleResponse(resp, err, "failed to query endpoint sensor statistics")
		},
	}
}
