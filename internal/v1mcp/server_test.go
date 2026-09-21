package v1mcp

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var toolNamePattern = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

func TestNewMcpServerRegistersValidTools(t *testing.T) {
	for _, readOnly := range []bool{true, false} {
		s, err := NewMcpServer(ServerConfig{ApiKey: "key", Host: "api.example.com", ReadOnly: readOnly, Version: "test"})
		require.NoError(t, err)

		tools := s.ListTools()
		require.NotEmpty(t, tools)

		for name, st := range tools {
			require.Equal(t, name, st.Tool.Name)
			require.Regexp(t, toolNamePattern, name)
			require.LessOrEqual(t, len(name), 64, "tool name %q is too long for some MCP clients", name)
			require.NotEmpty(t, st.Tool.Description, name)
			require.NotNil(t, st.Tool.Annotations.ReadOnlyHint, name)
			require.NotNil(t, st.Handler, name)

			b, err := json.Marshal(st.Tool)
			require.NoError(t, err, name)
			require.True(t, json.Valid(b), name)

			if readOnly {
				require.True(t, *st.Tool.Annotations.ReadOnlyHint, "%s must be read only in read only mode", name)
			}
		}
	}
}

func TestNoBetaAPIs(t *testing.T) {
	s, err := NewMcpServer(ServerConfig{ApiKey: "key", Host: "api.example.com", Version: "test"})
	require.NoError(t, err)

	for name, st := range s.ListTools() {
		require.False(t, strings.HasPrefix(name, "cloud_posture_"), name)
		require.NotContains(t, strings.ToLower(st.Tool.Description), "beta", name)
	}
}

func TestToolsetSelection(t *testing.T) {
	cfg := ServerConfig{ApiKey: "key", Host: "api.example.com", ReadOnly: false, Version: "test"}

	all, err := NewMcpServer(cfg)
	if err != nil {
		t.Fatal(err)
	}

	cfg.Toolsets = []string{"sandbox", "iam"}
	s, err := NewMcpServer(cfg)
	if err != nil {
		t.Fatal(err)
	}

	got := s.ListTools()
	if len(got) == 0 || len(got) >= len(all.ListTools()) {
		t.Fatalf("expected a non-empty subset of %d tools, got %d", len(all.ListTools()), len(got))
	}
	for name := range got {
		if !strings.HasPrefix(name, "sandbox_") && !strings.HasPrefix(name, "iam_") {
			t.Errorf("tool %q does not belong to the selected toolsets", name)
		}
	}

	cfg.ReadOnly = true
	ro, err := NewMcpServer(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if len(ro.ListTools()) >= len(got) {
		t.Errorf("readonly should register fewer tools than write mode")
	}

	cfg.Toolsets = []string{"nope"}
	if _, err := NewMcpServer(cfg); err == nil {
		t.Error("expected an error for an unknown toolset")
	}
}

func TestParseToolsets(t *testing.T) {
	for _, in := range []string{"", "all", "iam,all", " ALL "} {
		got, err := ParseToolsets(in)
		if err != nil || got != nil {
			t.Errorf("ParseToolsets(%q) = %v, %v; want nil, nil", in, got, err)
		}
	}

	got, err := ParseToolsets(" IAM, sandbox,iam")
	if err != nil || len(got) != 2 || got[0] != "iam" || got[1] != "sandbox" {
		t.Errorf("unexpected result %v, %v", got, err)
	}

	if _, err := ParseToolsets("iam,bogus"); err == nil {
		t.Error("expected an error for an unknown toolset")
	}
}
