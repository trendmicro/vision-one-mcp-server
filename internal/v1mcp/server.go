package v1mcp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp/tools"
)

type ServerConfig struct {
	ApiKey   string
	ReadOnly bool
	Version  string
	Region   string
}

func NewMcpServer(cfg ServerConfig) *mcpserver.MCPServer {
	s := mcpserver.NewMCPServer(
		"v1mcp",
		cfg.Version,
		mcpserver.WithLogging(),
	)

	client := v1client.NewV1ApiClient(cfg.ApiKey, cfg.Region)
	client.UserAgent = fmt.Sprintf("trend-vision-one-mcp-server/%s", cfg.Version)

	addReadOnlyToolset(s, client, tools.ToolsetsReadOnlyIAM)
	addReadOnlyToolset(s, client, tools.ToolsetsReadOnlyCREM)
	addReadOnlyToolset(s, client, tools.ToolsetsReadOnlyCloudPosture)
	addReadOnlyToolset(s, client, tools.ToolsetsReadOnlyWorkench)
	addReadOnlyToolset(s, client, tools.ToolsetsReadOnlyCAM)
	addReadOnlyToolset(s, client, tools.ToolsetsReadOnlyEmail)

	if !cfg.ReadOnly {
		addWriteToolset(s, client, tools.ToolsetsWriteCloudPosture)
		addWriteToolset(s, client, tools.ToolsetsWriteIAM)
	}

	return s
}

func RunMcpStdioServer(cfg ServerConfig) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s := NewMcpServer(cfg)

	stdioServer := mcpserver.NewStdioServer(s)

	serverError := make(chan error)
	go func() {
		serverError <- stdioServer.Listen(ctx, os.Stdin, os.Stdout)
	}()

	fmt.Fprintf(os.Stderr, "server listening...\n")

	select {
	case <-ctx.Done():
		fmt.Fprintf(os.Stderr, "shutting down server...\n")
	case e := <-serverError:
		return fmt.Errorf("server encountered error: %w\n", e)
	}

	return nil
}

func addReadOnlyToolset(
	s *mcpserver.MCPServer,
	client *v1client.V1ApiClient,
	servertools []func(*v1client.V1ApiClient) mcpserver.ServerTool,
) {
	for _, getTool := range servertools {
		addReadTools(s, getTool(client))
	}
}

func addWriteToolset(
	s *mcpserver.MCPServer,
	client *v1client.V1ApiClient,
	servertools []func(*v1client.V1ApiClient) mcpserver.ServerTool,
) {
	for _, getTool := range servertools {
		addWriteTools(s, getTool(client))
	}
}

func addWriteTools(s *mcpserver.MCPServer, serverTools ...mcpserver.ServerTool) {
	for _, tool := range serverTools {
		if *tool.Tool.Annotations.ReadOnlyHint {
			panic(fmt.Sprintf("tool %q shouldn't be marked as being readonly"))
		}
	}
	s.AddTools(serverTools...)
}

func addReadTools(s *mcpserver.MCPServer, serverTools ...mcpserver.ServerTool) {
	for _, tool := range serverTools {
		if !*tool.Tool.Annotations.ReadOnlyHint {
			panic(fmt.Sprintf("tool %q should be marked as readonly", tool.Tool.Name))
		}
	}
	s.AddTools(serverTools...)
}
