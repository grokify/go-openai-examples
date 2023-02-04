package main

import (
	"context"
	"fmt"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	gogpt "github.com/sashabaranov/go-gpt3"
)

const EnvOpenAIKey = "OPENAPI_API_KEY"

func main() {
	paths, err := config.LoadDotEnv([]string{}, 1)
	fmtutil.PrintJSON(paths)

	c := gogpt.NewClient("your token")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
