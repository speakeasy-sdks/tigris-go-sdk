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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.TigrisDeleteAppKeyRequest{
        DeleteAppKeyRequest: shared.DeleteAppKeyRequest{
            ID: tigris.String("05dfc2dd-f7cc-478c-a1ba-928fc816742c"),
        },
        Project: "cum",
    }

    res, err := s.AppKey.Delete(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.TigrisListAppKeysRequest{
        Project: "esse",
    }

    res, err := s.AppKey.List(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.TigrisRotateAppKeySecretRequest{
        RotateAppKeyRequest: shared.RotateAppKeyRequest{
            ID: tigris.String("39205929-396f-4ea7-996e-b10faaa2352c"),
            Project: tigris.String("enim"),
        },
        Project: "omnis",
    }

    res, err := s.AppKey.Rotate(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.TigrisCreateAppKeyRequest{
        CreateAppKeyRequest: shared.CreateAppKeyRequest{
            Description: tigris.String("nemo"),
            Name: tigris.String("Velma Batz"),
        },
        Project: "doloribus",
    }

    res, err := s.AppKey.TigrisCreateAppKey(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.TigrisUpdateAppKeyRequest{
        UpdateAppKeyRequest: shared.UpdateAppKeyRequest{
            Description: tigris.String("sapiente"),
            ID: tigris.String("1a3a2fa9-4677-4392-91aa-52c3f5ad019d"),
            Name: tigris.String("Ryan Witting"),
        },
        Project: "nihil",
    }

    res, err := s.AppKey.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAppKeyResponse != nil {
        // handle response
    }
}
```
