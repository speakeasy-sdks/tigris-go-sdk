# AppKey

## Overview

The application keys section provide APIs that can be used to manage application keys for your project. A single project can have one or more application keys.

### Available Operations

* [Delete](#delete) - Deletes the app key
* [List](#list) - List all the app keys
* [Rotate](#rotate) - Rotates the app key secret
* [TigrisCreateAppKey](#tigriscreateappkey) - Creates the app key
* [Update](#update) - Updates the description of the app key

## Delete

Delete an app key.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.Delete(ctx, operations.TigrisDeleteAppKeyRequest{
        DeleteAppKeyRequest: shared.DeleteAppKeyRequest{
            ID: tigris.String("d69a674e-0f46-47cc-8796-ed151a05dfc2"),
        },
        Project: "at",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAppKeyResponse != nil {
        // handle response
    }
}
```

## List

Lists all app keys visible to requesting actor.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.List(ctx, operations.TigrisListAppKeysRequest{
        Project: "at",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppKeysResponse != nil {
        // handle response
    }
}
```

## Rotate

Endpoint is used to rotate the secret for the app key.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.Rotate(ctx, operations.TigrisRotateAppKeySecretRequest{
        RotateAppKeyRequest: shared.RotateAppKeyRequest{
            ID: tigris.String("f7cc78ca-1ba9-428f-8816-742cb7392059"),
            Project: tigris.String("sed"),
        },
        Project: "iste",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RotateAppKeyResponse != nil {
        // handle response
    }
}
```

## TigrisCreateAppKey

Create an app key.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.TigrisCreateAppKey(ctx, operations.TigrisCreateAppKeyRequest{
        CreateAppKeyRequest: shared.CreateAppKeyRequest{
            Description: tigris.String("dolor"),
            Name: tigris.String("Lester Welch"),
        },
        Project: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAppKeyResponse != nil {
        // handle response
    }
}
```

## Update

Update the description of an app key.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.Update(ctx, operations.TigrisUpdateAppKeyRequest{
        UpdateAppKeyRequest: shared.UpdateAppKeyRequest{
            Description: tigris.String("corporis"),
            ID: tigris.String("96eb10fa-aa23-452c-9955-907aff1a3a2f"),
            Name: tigris.String("Tracy Fritsch"),
        },
        Project: "molestiae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAppKeyResponse != nil {
        // handle response
    }
}
```
