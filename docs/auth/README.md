# Auth

## Overview

The auth section of API provides OAuth 2.0 APIs. Tigris supports pluggable OAuth provider. Pass the token in the headers for authentication, as an example `-H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6I"`(replace it with your token). All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.

### Available Operations

* [Get](#get) - Access Token

## Get

Endpoint for receiving access token from Tigris Server. The endpoint requires Grant Type(`grant_type`) which has
 two possible values <i>"REFRESH_TOKEN"</i> or <i>"CLIENT_CREDENTIALS"</i> based on which either Refresh token(`refresh_token`)
 needs to be set or client credentials(`client_id`, `client_secret`).

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Auth.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAccessTokenResponse != nil {
        // handle response
    }
}
```
