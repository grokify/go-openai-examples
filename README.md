# Go OpenAI Examples

This is a repo to contain various experiments with OpenAI, such as ChatGPT.

## `openaichat` CLI app

`openaichat` provides a CLI method for interacting with the OpenAI Chat API.

### Installation

```
% go install github.com/grokify/go-openai-examples/cmd/openaichat
```

### Usage

Use the `-p` prompt parameter to send your request to ChatGPT.

There are several ways to provide the OpenAI token to the `openaichat` CLI app as shown below.

#### Usage with Environment Variable

If no CLI parameters are used, the token is assumed to be in the `OPENAI_API_KEY` environment variable.

```
% OPENAI_API_KEY = my_key openaichat -p 'hello ChatGPT!'
```

#### Usage with `.env` file

When using an `.env` file, use the `-e` parameter, and the the CLI app will look for the `OPENAI_API_KEY` environment variable.

```
% openaichat -e /path/to/envfile -p 'hello ChatGPT!'
```

#### Usage with `goauth` credentials set file

The `github.com/grokify/goauth` credentials set file is a single format to handle multiple types of authentication.n When using a credentials set file, use the `-c` and `-a` parameters.

```
% openaichat -c /path/to/credentialsfile.json` -a yourAccountName -p 'hello ChatGPT!'
```

This code will look for the static access token as an OAuth 2.0 credential due to use of the bearer token format.

If your credentials file looks like the following, you can use `-a OPENAI`.

```json
{
  "credentials": {
    "OPENAI": {
      "oauth2": {
        "token": {
          "access_token": "<yourOpenAIToken",
        }
      }
    }
  }
}

```