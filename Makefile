mcpserver:
	go build -o ./bin/v1-mcp-server ./cmd/v1-mcp-server/main.go
install:
	go install -ldflags="-X main.version=dev"
