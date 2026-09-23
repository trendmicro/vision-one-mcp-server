package v1mcp

import (
	"fmt"
	"sort"
	"strings"

	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tools"
)

// AllToolsets is the keyword that selects every toolset.
const AllToolsets = "all"

type toolsetFactories = []func(*v1client.V1ApiClient) mcpserver.ServerTool

type toolsetDef struct {
	name  string
	read  toolsetFactories
	write toolsetFactories
}

// toolsetDefs lists every toolset in registration order.
var toolsetDefs = []toolsetDef{
	{"iam", tools.ToolsetsReadOnlyIAM, tools.ToolsetsWriteIAM},
	{"crem", tools.ToolsetsReadOnlyCREM, tools.ToolsetsWriteCREM},
	{"cloudrisk", tools.ToolsetsReadOnlyCloudRiskManagement, nil},
	{"workbench", tools.ToolsetsReadOnlyWorkench, tools.ToolsetsWriteWorkench},
	{"cam", tools.ToolsetsReadOnlyCAM, tools.ToolsetsWriteCAM},
	{"email", tools.ToolsetsReadOnlyEmail, nil},
	{"container", tools.ToolsetsReadOnlyContainer, tools.ToolsetsWriteContainer},
	{"endpoint", tools.ToolsetsReadOnlyEndpoint, tools.ToolsetsWriteEndpoint},
	{"ai", tools.ToolsetsReadOnlyAISecurity, nil},
	{"threatintel", tools.ToolsetsReadOnlyThreatIntel, tools.ToolsetsWriteThreatIntel},
	{"audit", tools.ToolsetsReadOnlyAudit, nil},
	{"business", tools.ToolsetsReadOnlyBusiness, tools.ToolsetsWriteBusiness},
	{"cases", tools.ToolsetsReadOnlyCaseManagement, tools.ToolsetsWriteCaseManagement},
	{"datalake", tools.ToolsetsReadOnlyDatalake, tools.ToolsetsWriteDatalake},
	{"dmm", tools.ToolsetsReadOnlyDMM, tools.ToolsetsWriteDMM},
	{"eiqs", tools.ToolsetsReadOnlyEIQS, nil},
	{"filesecurity", tools.ToolsetsReadOnlyFileSecurity, tools.ToolsetsWriteFileSecurity},
	{"healthcheck", tools.ToolsetsReadOnlyHealthcheck, nil},
	{"response", tools.ToolsetsReadOnlyResponse, tools.ToolsetsWriteResponse},
	{"sandbox", tools.ToolsetsReadOnlySandbox, tools.ToolsetsWriteSandbox},
	{"search", tools.ToolsetsReadOnlySearch, nil},
	{"awareness", tools.ToolsetsReadOnlySecurityAwareness, nil},
	{"playbooks", tools.ToolsetsReadOnlySecurityPlaybooks, tools.ToolsetsWriteSecurityPlaybooks},
	{"tags", tools.ToolsetsReadOnlyTagManagement, tools.ToolsetsWriteTagManagement},
}

// ToolsetNames returns the names accepted by ParseToolsets, sorted alphabetically.
func ToolsetNames() []string {
	names := make([]string, 0, len(toolsetDefs))
	for _, d := range toolsetDefs {
		names = append(names, d.name)
	}
	sort.Strings(names)
	return names
}

// ParseToolsets splits a comma separated list of toolset names.
// An empty value or the keyword "all" selects every toolset and yields nil.
func ParseToolsets(value string) ([]string, error) {
	var names []string
	seen := map[string]bool{}
	for _, part := range strings.Split(value, ",") {
		name := strings.ToLower(strings.TrimSpace(part))
		if name == "" {
			continue
		}
		if name == AllToolsets {
			return nil, nil
		}
		if !isToolsetName(name) {
			return nil, fmt.Errorf("unknown toolset %q, provide %q or a comma separated list of: %s", name, AllToolsets, strings.Join(ToolsetNames(), ", "))
		}
		if !seen[name] {
			seen[name] = true
			names = append(names, name)
		}
	}
	return names, nil
}

func isToolsetName(name string) bool {
	for _, d := range toolsetDefs {
		if d.name == name {
			return true
		}
	}
	return false
}

// selectedToolsets returns the definitions matching names, or all of them when names is empty.
func selectedToolsets(names []string) ([]toolsetDef, error) {
	if len(names) == 0 {
		return toolsetDefs, nil
	}
	want := map[string]bool{}
	for _, n := range names {
		n = strings.ToLower(strings.TrimSpace(n))
		if n == AllToolsets {
			return toolsetDefs, nil
		}
		if !isToolsetName(n) {
			return nil, fmt.Errorf("unknown toolset %q, valid toolsets: %s", n, strings.Join(ToolsetNames(), ", "))
		}
		want[n] = true
	}
	var out []toolsetDef
	for _, d := range toolsetDefs {
		if want[d.name] {
			out = append(out, d)
		}
	}
	return out, nil
}
