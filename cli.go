package openapiexamples

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/grokify/go-openai-examples/examples"
	"github.com/grokify/goauth/credentials"
	"github.com/grokify/mogo/config"
	openai "github.com/sashabaranov/go-openai"
)

const EnvOpenAIKey = "OPENAI_API_KEY"

type Options struct {
	UseEnv       []bool `short:"e" long:"env" description:"List subscriptions"`
	CredsFile    string `short:"c" long:"credsfile" description:"Credentials file"`
	CredsAccount string `short:"a" long:"account" credsaccount:"Credentials account"`
	Prompt       string `short:"p" long:"prompt" description:"Prompt"`
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
	if len(opts.UseEnv) > 0 {
		_, err := config.LoadDotEnv([]string{}, 1)
		if err != nil {
			return "", err
		}
		return os.Getenv(EnvOpenAIKey), nil
	} else if opts.CredsFile != "" {
		credsSet, err := credentials.ReadFileCredentialsSet(opts.CredsFile, false)
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
	}
	return "", fmt.Errorf("required API Key is missing, use '-e' or '-c'")
}
