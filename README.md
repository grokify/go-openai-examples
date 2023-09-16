# Go OpenAI Examples

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

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

The `github.com/grokify/goauth` credentials set file is a single format to handle multiple types of authentication. When using a credentials set file, use the `-c` and `-a` parameters.

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
          "access_token": "<yourOpenAIAPIKey>",
        }
      }
    }
  }
}
```

 [used-by-svg]: https://sourcegraph.com/github.com/grokify/go-openai-examples/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/go-openai-examples?badge
 [build-status-svg]: https://github.com/grokify/goauth/workflows/test/badge.svg
 [build-status-url]: https://github.com/grokify/goauth/actions/workflows/test.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-openai-examples
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-openai-examples
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-openai-examples
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-openai-examples
 [loc-svg]: https://tokei.rs/b1/github/grokify/go-openai-examples
 [repo-url]: https://github.com/grokify/go-openai-examples
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-openai-examples/blob/master/LICENSE