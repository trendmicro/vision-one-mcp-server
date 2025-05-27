package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

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

	flag.Parse()

	if *showVersion {
		printVersion()
		return nil
	}

	if err := validateRegion(*v1Region); err != nil {
		return err
	}

	apiKey := os.Getenv("TREND_VISION_ONE_API_KEY")
	if apiKey == "" {
		return errors.New("TREND_VISION_ONE_API_KEY not set")
	}

	version := getVersion()

	serverCfg := v1mcp.ServerConfig{
		ApiKey:   apiKey,
		ReadOnly: *readOnly,
		Region:   *v1Region,
		Version:  version,
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
		return errors.New(fmt.Sprintf("missing region, please provide one of %s", string(b)))
	}

	for _, r := range validRegions {
		if r == region {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("invalid region %q, provide on of %s", region, string(b)))
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
