package main

import (
	"fmt"

	"github.com/grokify/mogo/log/logutil"
	"github.com/jessevdk/go-flags"

	openaiexamples "github.com/grokify/go-openai-examples"
)

func main() {
	opts := openaiexamples.Options{}
	_, err := flags.Parse(&opts)
	logutil.FatalErr(err)

	client, err := opts.NewClient()
	logutil.FatalErr(err)

	prompt := opts.PromptOrExample(-1, "")

	fmt.Printf("REQUEST: [%s]\n", prompt)

	answer, err := openaiexamples.ChatRequest(client, prompt, "")
	logutil.FatalErr(err)

	fmt.Printf("===== BEGIN RESPONSE =====\n%s\n===== END RESPONSE =====\n", answer)

	/*
		searchReq := openai.SearchRequest{
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
