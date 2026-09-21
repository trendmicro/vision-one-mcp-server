# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Trend Vision One MCP Server - A Go-based Model Context Protocol (MCP) server that bridges AI tooling (Claude, VSCode + GitHub Copilot) with Trend Vision One security platform APIs. Enables natural language interaction with security services like Workbench alerts, Cloud Risk Management, endpoint management, attack surface discovery, and AI security guardrails.

## Build & Test Commands

```bash
# Build
make mcpserver              # Build to ./bin/v1-mcp-server
go build -o ./bin/v1-mcp-server ./cmd/v1-mcp-server/main.go

# Test
go test -v ./...            # Run all tests

# Lint & Format
./script/check-gofmt        # Check formatting
./script/lint               # Run golangci-lint
gofmt -s -w ./              # Auto-format code

# Run locally
./bin/v1-mcp-server -region us  # Requires TREND_VISION_ONE_API_KEY env var
```

## Architecture

```
cmd/v1-mcp-server/main.go     # Entry point, CLI flags, region validation
internal/v1mcp/server.go      # MCP server setup, tool registration
internal/v1mcp/tools/*.go     # Tool handlers (one file per domain)
internal/v1client/*.go        # HTTP client, API endpoint methods
```

**Request Flow:** MCP Request → Tool Handler → Parameter Extraction → v1client API call → HTTP → Response → MCP Result

## Key Conventions

### Tool Implementation Pattern
Each tool is a factory function returning `mcpserver.ServerTool`:
```go
func toolDomainResourceAction(client *v1client.V1ApiClient) mcpserver.ServerTool {
    return mcpserver.ServerTool{
        Tool: mcp.NewTool("domain_resource_action",
            mcp.WithDescription("..."),
            mcp.WithString("param", mcp.Description("...")),
            mcp.WithToolAnnotation(mcp.ToolAnnotation{ReadOnlyHint: toPtr(true)}),
        ),
        Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
            // Extract parameters, call API, return result
        },
    }
}
```

### Read-Only vs Write Tools
- Tools must be annotated with `ReadOnlyHint: toPtr(true/false)`
- Server validates annotations match toolset registration (panics on mismatch)
- Add read tools to `ToolsetsReadOnly{Domain}`, write tools to `ToolsetsWrite{Domain}` in respective `tools/*.go` files
- Register toolsets in `toolsets.go` (`toolsetDefs`); each has a name selectable with the `-toolsets` flag (comma list, `all` by default).
- **REQUIRED:** Update `README.md` Tools section with new tools (include tool name, description, and mode)
- Only expose generally available APIs. Do not add tools for beta or preview endpoints (paths starting with `beta/`)

### API Paths
Paths must NOT start with `/`. The client's `Parse` method handles URL joining:
```go
// Correct
c.searchAndFilter("v3.0/iam/apiKeys", filter, params)
// Wrong
c.searchAndFilter("/v3.0/iam/apiKeys", filter, params)
```

### Parameter Helpers (tools/tools.go)
```go
requiredValue[T](property string, vals map[string]any) (T, error)
optionalValue[T](property string, vals map[string]any) (T, error)
optionalIntValue(property string, vals map[string]any) (int, error)
optionalTimeValue(property string, vals map[string]any) (time.Time, error)
handleStatusResponse(r *http.Response, err error, expectedStatusCode int, msg string) (*mcp.CallToolResult, error)
```

### Generic Client Functions (v1client/v1client.go)
Always prefer using these generic functions over manual request construction:
```go
// For GET requests with filter header and query parameters (most common)
c.searchAndFilter(path string, filter string, queryParams any) (*http.Response, error)

// For simple GET requests without filters or query parameters
c.genericGet(path string) (*http.Response, error)
```

Example usage:
```go
// Correct - use searchAndFilter for list/search endpoints
func (c *V1ApiClient) SomeListEndpoint(filter string, queryParams QueryParameters) (*http.Response, error) {
    return c.searchAndFilter("v3.0/domain/resources", filter, queryParams)
}

// Correct - use genericGet for simple GET by ID
func (c *V1ApiClient) SomeGetEndpoint(id string) (*http.Response, error) {
    return c.genericGet(fmt.Sprintf("v3.0/domain/resources/%s", id))
}

// Wrong - don't manually construct requests when generic functions work
func (c *V1ApiClient) SomeListEndpoint(filter string, queryParams QueryParameters) (*http.Response, error) {
    p, err := query.Values(queryParams)  // Unnecessary boilerplate
    // ... more manual construction
}
```

### Request Options (v1client/v1client.go)
For cases where generic functions don't apply (custom headers, POST/PUT/DELETE):
```go
withHeader(name, value string) requestOptionFunc  // Add custom headers to requests
withFilter(filter string) requestOptionFunc       // Add TMV1-Filter header
withUrlParameters(params url.Values) requestOptionFunc  // Add query parameters
withContentTypeJSON() requestOptionFunc           // Add JSON content-type header
```
Used by domains requiring custom headers (e.g., AI Security uses `TMV1-Application-Name`, `TMV1-Request-Type`, `Prefer`).

### Request Params Pattern (v1client/request.go, tools/request.go)
Most endpoints, including every write endpoint, use `RequestParams` instead of typed query structs:
```go
// v1client: path arguments are typed, everything else travels in RequestParams
func (c *V1ApiClient) SandboxTaskGet(id string, p RequestParams) (*http.Response, error) {
    return c.do(http.MethodGet, pathf("v3.0/sandbox/tasks/%s", id), p)
}

// tools: requestSpec maps tool arguments to query parameters, headers and the request body
params, err := buildRequestParams(args, requestSpec{
    Query:   []paramDef{{Name: "top", Kind: kindString}},
    Headers: []headerDef{{Arg: "filter", Header: "TMV1-Filter"}},
})
resp, err := client.SandboxTaskGet(id, params)
return handleResponse(resp, err, "failed to get submission status")
```
- `pathf` URL-escapes every path argument. Read path arguments with `pathValue`.
- `Body` is `bodyObject` (the `BodyFields` are sent as JSON), `bodyArray` (the `items` argument is sent as the JSON array) or `bodyMultipart` (`BodyFields` are form fields and `FileFields` hold the path of a local file to upload).
- `handleResponse` treats every 2xx status as success (bulk endpoints return 207), returns `ETag`, `Operation-Location` and `Location` headers alongside the body, and returns binary responses as embedded resources.
- Header arguments are named `filter` (`TMV1-Filter`), `query` (`TMV1-Query`), `ifMatch` (`If-Match`) and `accept` (`Accept`).

### Tool Naming Convention
`{domain}_{resource}_{action}` - e.g., `iam_api_keys_list`, `workbench_alert_detail_get`

### Comments
Only add comments when something truly needs explanation. Comments should explain "why", not "what". If the code is clear from function names, variable names, and structure, no comment is needed.

## Domain Organization

| Domain | Client File | Tools File | API Prefix |
|--------|-------------|------------|------------|
| AI Security | `v1client/aisecurity.go` | `tools/aisecurity.go` | `v3.0/aiSecurity/` |
| IAM | `v1client/iam.go` | `tools/iam.go` | `v3.0/iam/` |
| Workbench | `v1client/workbench.go` | `tools/workbench.go` | `v3.0/workbench/` |
| OAT | `v1client/oat.go` | `tools/workbench.go` | `v3.0/oat/` |
| CREM | `v1client/crem.go` | `tools/crem.go` | `v3.0/asrm/` |
| CAM | `v1client/cam.go` | `tools/cam.go` | `v3.0/cam/` |
| Email | `v1client/email.go` | `tools/email.go` | `v3.0/email/` |
| Container | `v1client/container.go` | `tools/container.go` | `v3.0/containerSecurity/` |
| Endpoint | `v1client/endpoint.go` | `tools/endpoint.go` | `v3.0/endpointSecurity/` |
| Threat Intel | `v1client/threatintel.go` | `tools/threatintel.go` | `v3.0/threatintel/` |
| Audit Logs | `v1client/audit.go` | `tools/audit.go` | `v3.0/audit/` |
| Business Information | `v1client/business.go` | `tools/business.go` | `v3.0/business/` |
| Case Management | `v1client/casemanagement.go` | `tools/casemanagement.go` | `v3.0/caseManagement/` |
| Data Pipelines (Datalake) | `v1client/datalake.go` | `tools/datalake.go` | `v3.0/datalake/` |
| Detection Model Management | `v1client/dmm.go` | `tools/dmm.go` | `v3.0/dmm/` |
| Endpoint Inventory (EIQS) | `v1client/eiqs.go` | `tools/eiqs.go` | `v3.0/eiqs/` |
| File Security | `v1client/filesecurity.go` | `tools/filesecurity.go` | `v3.0/fileSecurity/` |
| Health Check | `v1client/healthcheck.go` | `tools/healthcheck.go` | `v3.0/healthcheck/` |
| Response Management | `v1client/response.go` | `tools/response.go` | `v3.0/response/` |
| Sandbox Analysis | `v1client/sandbox.go` | `tools/sandbox.go` | `v3.0/sandbox/` |
| Search | `v1client/search.go` | `tools/search.go` | `v3.0/search/` |
| Security Awareness | `v1client/securityawareness.go` | `tools/securityawareness.go` | `v3.0/securityAwareness/` |
| Security Playbooks | `v1client/securityplaybooks.go` | `tools/securityplaybooks.go` | `v3.0/securityPlaybooks/` |
| Tag Management | `v1client/tagmanagement.go` | `tools/tagmanagement.go` | `v3.0/tagManagement/` |

**Note:** OAT (Observed Attack Techniques) has its own client file but tools are registered under the Workbench toolset. The CREM domain covers everything under `v3.0/asrm/`, including vulnerability and risk event endpoints.

## Contributing

Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages.
