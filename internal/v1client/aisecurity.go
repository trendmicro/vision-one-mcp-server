package v1client

import (
	"net/http"
)

// AISecurityApplyGuardrailsInput represents the input for the applyGuardrails API.
// It supports three request types:
// - SimpleRequestGuard: just a prompt string
// - OpenAIChatCompletionRequestV1: OpenAI chat completion request format
// - OpenAIChatCompletionResponseV1: OpenAI chat completion response format
type AISecurityApplyGuardrailsInput struct {
	// For SimpleRequestGuard
	Prompt string `json:"prompt,omitempty"`

	// For OpenAIChatCompletionRequestV1
	Model    string                  `json:"model,omitempty"`
	Messages []AISecurityChatMessage `json:"messages,omitempty"`

	// For OpenAIChatCompletionResponseV1
	ID      string                 `json:"id,omitempty"`
	Object  string                 `json:"object,omitempty"`
	Created int64                  `json:"created,omitempty"`
	Choices []AISecurityChatChoice `json:"choices,omitempty"`
	Usage   *AISecurityUsage       `json:"usage,omitempty"`
}

// AISecurityChatMessage represents a message in the chat conversation.
type AISecurityChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// AISecurityChatChoice represents a choice in the OpenAI chat completion response.
type AISecurityChatChoice struct {
	Index        int                     `json:"index"`
	Message      AISecurityChoiceMessage `json:"message"`
	FinishReason string                  `json:"finish_reason"`
}

// AISecurityChoiceMessage represents a message in a chat choice.
type AISecurityChoiceMessage struct {
	Role    string  `json:"role"`
	Content string  `json:"content"`
	Refusal *string `json:"refusal,omitempty"`
}

// AISecurityUsage represents token usage statistics.
type AISecurityUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// AISecurityApplyGuardrailsOptions contains the optional headers for the applyGuardrails API.
type AISecurityApplyGuardrailsOptions struct {
	// ApplicationName is required - the name of the AI application whose prompts are being evaluated
	ApplicationName string
	// RequestType is the type of request being evaluated (SimpleRequestGuard, OpenAIChatCompletionRequestV1, OpenAIChatCompletionResponseV1)
	RequestType string
	// Prefer controls response detail level (return=representation for detailed, return=minimal for short)
	Prefer string
}

// AISecurityApplyGuardrails evaluates prompts against AI guard policies.
func (c *V1ApiClient) AISecurityApplyGuardrails(input AISecurityApplyGuardrailsInput, opts AISecurityApplyGuardrailsOptions) (*http.Response, error) {
	return c.genericJSONPost(
		"v3.0/aiSecurity/applyGuardrails",
		input,
		withHeader("TMV1-Application-Name", opts.ApplicationName),
		withHeader("TMV1-Request-Type", opts.RequestType),
		withHeader("Prefer", opts.Prefer),
	)
}
