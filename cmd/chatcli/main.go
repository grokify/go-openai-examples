package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/type/stringsutil"
	"github.com/grokify/ringgpt/examples"
	gogpt "github.com/sashabaranov/go-gpt3"
)

const EnvOpenAIKey = "OPENAI_API_KEY"

func main() {
	paths, err := config.LoadDotEnv([]string{}, 1)
	logutil.FatalErr(err)
	fmtutil.MustPrintJSON(paths)

	apiKey := os.Getenv(EnvOpenAIKey)
	fmt.Printf("APIKEY (%s)\n", apiKey)

	c := gogpt.NewClient(apiKey)
	ctx := context.Background()

	exs := examples.Examples()

	prompt := exs[0].Prompt

	fmt.Printf("PROMPT: [%s]\n", prompt)

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3Ada,
		MaxTokens: 500,
		Prompt:    prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	logutil.FatalErr(err)

	fmt.Println(stringsutil.CondenseLines(resp.Choices[0].Text))

	/*
		searchReq := gogpt.SearchRequest{
			Documents: []string{"White House", "hospital", "school"},
			Query:     "the president",
		}
		searchResp, err := c.Search(ctx, "ada", searchReq)
		if err != nil {
			return
		}
		fmt.Println(searchResp.SearchResults)
	*/
}
