package openapiexamples

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/grokify/mogo/type/stringsutil"
	openai "github.com/sashabaranov/go-openai"
)

func ChatRequest(client *openai.Client, prompt, model string) (string, error) {
	if client == nil {
		return "", errors.New("client is nil")
	}
	prompt = strings.TrimSpace(prompt)
	if prompt == "" {
		return "", errors.New("prompt is empty")
	}
	model = strings.TrimSpace(model)
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}
	if model == openai.GPT3Dot5Turbo {
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: prompt,
					},
				},
			},
		)
		if err != nil {
			return "", err
		}
		return stringsutil.CondenseLines(resp.Choices[0].Message.Content, "\n"), nil
	} else if model == openai.GPT3Ada {
		req := openai.CompletionRequest{
			Model:     model,
			MaxTokens: 500,
			Prompt:    prompt,
		}
		resp, err := client.CreateCompletion(context.Background(), req)
		if err != nil {
			return "", err
		}
		return stringsutil.CondenseLines(resp.Choices[0].Text, "\n"), nil

	}
	return "", fmt.Errorf("model unknown (%s)", model)
}
