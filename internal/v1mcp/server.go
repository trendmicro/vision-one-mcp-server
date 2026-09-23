package v1mcp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mcpserver "github.com/mark3labs/mcp-go/server"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"
)

type ServerConfig struct {
	ApiKey   string
	ReadOnly bool
	Version  string
	Region   string
	Host     string
	// Toolsets limits the registered tools to the named toolsets. Empty registers all of them.
	Toolsets []string
}

func NewMcpServer(cfg ServerConfig) (*mcpserver.MCPServer, error) {
	s := mcpserver.NewMCPServer(
		"v1mcp",
		cfg.Version,
		mcpserver.WithLogging(),
	)

	client, err := v1client.NewV1ApiClient(v1client.ClientOptions{
		Host:   cfg.Host,
		Region: cfg.Region,
		ApiKey: cfg.ApiKey,
	})
	if err != nil {
		return nil, err
	}
	client.UserAgent = fmt.Sprintf("trend-vision-one-mcp-server/%s", cfg.Version)

	selected, err := selectedToolsets(cfg.Toolsets)
	if err != nil {
		return nil, err
	}

	for _, d := range selected {
		addReadOnlyToolset(s, client, d.read)
		if !cfg.ReadOnly {
			addWriteToolset(s, client, d.write)
		}
	}

	return s, nil
}

func RunMcpStdioServer(cfg ServerConfig) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	s, err := NewMcpServer(cfg)
	if err != nil {
		return fmt.Errorf("error creating mcp server: %w", err)
	}

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
		return fmt.Errorf("server encountered error: %w", e)
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
			panic(fmt.Sprintf("tool %q shouldn't be marked as being readonly", tool.Tool.Name))
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
