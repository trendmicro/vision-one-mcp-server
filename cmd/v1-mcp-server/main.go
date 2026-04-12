package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"slices"

	"github.com/trendmicro/vision-one-mcp-server/internal/v1mcp"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	readOnly := flag.Bool("readonly", true, "set readonly false to allow the MCP server to perform write operations.")
	v1Region := flag.String("region", "", "set the region of your vision one account.")
	showVersion := flag.Bool("version", false, "print version information")
	host := flag.String("host", "", "set the Trend Vision One endpoint you want to use. Only useful for interacting with internal environments.")
	transport := flag.String("transport", getEnvOrDefault("TRANSPORT", "stdio"), "transport type: stdio or http")
	addr := flag.String("addr", getEnvOrDefault("ADDR", ":8000"), "address to listen on when using http transport")

	flag.Parse()

	if *showVersion {
		printVersion()
		return nil
	}

	apiKey := os.Getenv("TREND_VISION_ONE_API_KEY")
	if apiKey == "" {
		return errors.New("TREND_VISION_ONE_API_KEY not set")
	}

	if *host != "" && *v1Region != "" {
		return errors.New("host and region cannot be used together")
	}

	if *v1Region != "" {
		if err := validateRegion(*v1Region); err != nil {
			return err
		}
	}

	if *transport != "stdio" && *transport != "http" {
		return fmt.Errorf("invalid transport %q, must be stdio or http", *transport)
	}

	version := getVersion()

	serverCfg := v1mcp.ServerConfig{
		ApiKey:   apiKey,
		ReadOnly: *readOnly,
		Region:   *v1Region,
		Version:  version,
		Host:     *host,
	}

	if *transport == "http" {
		return v1mcp.RunMcpHttpServer(serverCfg, *addr)
	}

	return v1mcp.RunMcpStdioServer(serverCfg)
}

func validateRegion(region string) error {
	validRegions := []string{
		"au",
		"eu",
		"in",
		"jp",
		"sg",
		"us",
		"mea",
	}

	b, _ := json.Marshal(validRegions)

	if region == "" {
		return fmt.Errorf("missing region, please provide one of %s", string(b))
	}

	if !slices.Contains(validRegions, region) {
		return fmt.Errorf("invalid region %q, provide on of %s", region, string(b))
	}

	return nil
}

func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	if info.Main.Version != "" {
		return info.Main.Version
	}

	return "unknown"
}

func printVersion() {
	fmt.Fprintf(os.Stderr, "%s\n", getVersion())
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
