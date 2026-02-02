# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Trend Vision One MCP Server - A Go-based Model Context Protocol (MCP) server that bridges AI tooling (Claude, VSCode + GitHub Copilot) with Trend Vision One security platform APIs. Enables natural language interaction with security services like Workbench alerts, Cloud Posture, endpoint management, attack surface discovery, and AI security guardrails.

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
- Register toolsets in `server.go`

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

### Request Options (v1client/v1client.go)
```go
withHeader(name, value string) requestOptionFunc  // Add custom headers to requests
```
Used by domains requiring custom headers (e.g., AI Security uses `TMV1-Application-Name`, `TMV1-Request-Type`, `Prefer`).

### Tool Naming Convention
`{domain}_{resource}_{action}` - e.g., `iam_api_keys_list`, `workbench_alert_detail_get`

## Domain Organization

| Domain | Client File | Tools File | API Prefix |
|--------|-------------|------------|------------|
| AI Security | `v1client/aisecurity.go` | `tools/aisecurity.go` | `v3.0/aiSecurity/` |
| IAM | `v1client/iam.go` | `tools/iam.go` | `v3.0/iam/` |
| Workbench | `v1client/workbench.go` | `tools/workbench.go` | `v3.0/workbench/` |
| OAT | `v1client/oat.go` | `tools/workbench.go` | `v3.0/oat/` |
| Cloud Posture | `v1client/cloudposture.go` | `tools/cloudposture.go` | `v3.0/asrm/` |
| CREM | `v1client/crem.go` | `tools/crem.go` | `v3.0/asrm/` |
| CAM | `v1client/cam.go` | `tools/cam.go` | `v3.0/cam/` |
| Email | `v1client/email.go` | `tools/email.go` | `v3.0/email/` |
| Container | `v1client/container.go` | `tools/container.go` | `v3.0/containerSecurity/` |
| Endpoint | `v1client/endpoint.go` | `tools/endpoint.go` | `v3.0/endpointSecurity/` |

**Note:** OAT (Observed Attack Techniques) has its own client file but tools are registered under the Workbench toolset.

## Contributing

Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for commit messages.
