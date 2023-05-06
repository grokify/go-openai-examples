package openapiexamples

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/grokify/go-openai-examples/examples"
	"github.com/grokify/goauth"
	"github.com/grokify/mogo/config"
	openai "github.com/sashabaranov/go-openai"
)

const EnvOpenAIKey = "OPENAI_API_KEY"

type Options struct {
	EnvFile      string `short:"e" long:"env" description:".env file"`
	CredsFile    string `short:"c" long:"credsfile" description:"Credentials file"`
	CredsAccount string `short:"a" long:"account" credsaccount:"Credentials account"`
	Prompt       string `short:"p" long:"prompt" description:"Prompt"`
	Key          string `short:"k" long:"key" description:"API Key"`
}

func (opts Options) PromptOrExample(idx int, def string) string {
	if strings.TrimSpace(opts.Prompt) != "" {
		return opts.Prompt
	}
	exs := examples.ExamplesData()
	if idx < 0 || idx >= len(exs) {
		return exs.PromptFirstOrDefault(def)
	}
	return exs.PromptOrDefault(idx, def)
}

func (opts Options) NewClient() (*openai.Client, error) {
	apiKey, err := opts.APIKey()
	if err != nil {
		return nil, err
	}
	return openai.NewClient(apiKey), nil
}

func (opts Options) APIKey() (string, error) {
	if opts.Key != "" {
		return opts.Key, nil
	} else if opts.EnvFile != "" {
		_, err := config.LoadDotEnv([]string{opts.EnvFile}, 1)
		if err != nil {
			return "", err
		}
		return os.Getenv(EnvOpenAIKey), nil
	} else if opts.CredsFile != "" {
		credsSet, err := goauth.ReadFileCredentialsSet(opts.CredsFile, false)
		if err != nil {
			return "", err
		}
		if opts.CredsAccount == "" {
			accts := credsSet.Accounts()
			return "", fmt.Errorf("CredsAccount is required. Accounts: (%s)", strings.Join(accts, ", "))
		}
		creds, err := credsSet.Get(opts.CredsAccount)
		if err != nil {
			return "", err
		}
		if creds.OAuth2.Token == nil {
			return "", errors.New("OAuth2 token not set")
		}
		return creds.OAuth2.Token.AccessToken, nil
	} else {
		return os.Getenv(EnvOpenAIKey), nil
	}
	return "", fmt.Errorf("required API Key is missing, use '-e' or '-c'")
}
