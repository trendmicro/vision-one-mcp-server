package tools

import (
	"context"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/trendmicro/vision-one-mcp-server/internal/v1client"

	mcpserver "github.com/mark3labs/mcp-go/server"
)

var ToolsetsReadOnlyAISecurity = []func(*v1client.V1ApiClient) mcpserver.ServerTool{
	toolAISecurityApplyGuardrails,
}

func toolAISecurityApplyGuardrails(client *v1client.V1ApiClient) mcpserver.ServerTool {
	return mcpserver.ServerTool{
		Tool: mcp.NewTool(
			"aisecurity_guardrails_apply",
			mcp.WithDescription("Evaluates prompts against AI guard policies and returns the recommended action (Allow/Block) with reasons for any policy violations detected"),
			mcp.WithToolAnnotation(mcp.ToolAnnotation{
				ReadOnlyHint: toPtr(true),
			}),
			mcp.WithString("applicationName",
				mcp.Required(),
				mcp.Description("The name of the AI application whose prompts are being evaluated (max 64 characters)"),
			),
			mcp.WithString("prompt",
				mcp.Description("A simple text prompt to evaluate (max 1024 characters). Use this for SimpleRequestGuard request type."),
			),
			mcp.WithArray("messages",
				mcp.Description("A list of chat messages for OpenAI chat completion format. Each message should have 'role' (system/user/assistant) and 'content' fields. Use this for OpenAIChatCompletionRequestV1 request type."),
				mcp.Items(map[string]any{
					"type": "object",
					"properties": map[string]any{
						"role": map[string]any{
							"type":        "string",
							"enum":        []string{"system", "user", "assistant"},
							"description": "The role of the entity that creates the message",
						},
						"content": map[string]any{
							"type":        "string",
							"description": "The text content of the message",
						},
					},
					"required": []string{"role", "content"},
				}),
			),
			mcp.WithString("model",
				mcp.Description("The AI model identifier when using OpenAI chat completion format (e.g., 'us.meta.llama3-1-70b-instruct-v1:0')"),
			),
			mcp.WithString("requestType",
				mcp.Description("The type of request being evaluated"),
				mcp.Enum("SimpleRequestGuard", "OpenAIChatCompletionRequestV1", "OpenAIChatCompletionResponseV1"),
			),
			mcp.WithString("prefer",
				mcp.Description("Controls response detail level. 'return=representation' for detailed evaluation including harmful content, sensitive information, and prompt attacks. 'return=minimal' for shorter response with just action and reasons."),
				mcp.Enum("return=representation", "return=minimal"),
			),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			applicationName, err := requiredValue[string]("applicationName", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			prompt, err := optionalValue[string]("prompt", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			model, err := optionalValue[string]("model", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			requestType, err := optionalValue[string]("requestType", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			prefer, err := optionalValue[string]("prefer", request.GetArguments())
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			// Parse messages if provided
			var messages []v1client.AISecurityChatMessage
			if rawMessages, ok := request.GetArguments()["messages"].([]any); ok && len(rawMessages) > 0 {
				for _, rawMsg := range rawMessages {
					msgMap, ok := rawMsg.(map[string]any)
					if !ok {
						return mcp.NewToolResultError("each message must be an object with 'role' and 'content' fields"), nil
					}

					role, ok := msgMap["role"].(string)
					if !ok {
						return mcp.NewToolResultError("message 'role' must be a string"), nil
					}

					content, ok := msgMap["content"].(string)
					if !ok {
						return mcp.NewToolResultError("message 'content' must be a string"), nil
					}

					messages = append(messages, v1client.AISecurityChatMessage{
						Role:    role,
						Content: content,
					})
				}
			}

			// Validate that either prompt or messages is provided
			if prompt == "" && len(messages) == 0 {
				return mcp.NewToolResultError("either 'prompt' or 'messages' must be provided"), nil
			}

			input := v1client.AISecurityApplyGuardrailsInput{
				Prompt:   prompt,
				Model:    model,
				Messages: messages,
			}

			opts := v1client.AISecurityApplyGuardrailsOptions{
				ApplicationName: applicationName,
				RequestType:     requestType,
				Prefer:          prefer,
			}

			resp, err := client.AISecurityApplyGuardrails(input, opts)
			return handleStatusResponse(resp, err, http.StatusOK, "failed to apply guardrails")
		},
	}
}
