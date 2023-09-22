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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.Delete(ctx, operations.TigrisDeleteAppKeyRequest{
        DeleteAppKeyRequest: shared.DeleteAppKeyRequest{
            ID: tigrisgosdk.String("e0f467cc-8796-4ed1-91a0-5dfc2ddf7cc7"),
        },
        Project: "totam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAppKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.TigrisDeleteAppKeyRequest](../../models/operations/tigrisdeleteappkeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.TigrisDeleteAppKeyResponse](../../models/operations/tigrisdeleteappkeyresponse.md), error**


## List

Lists all app keys visible to requesting actor.

### Example Usage

```go
package main

import(
	"context"
	"log"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.List(ctx, operations.TigrisListAppKeysRequest{
        Project: "porro",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListAppKeysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.TigrisListAppKeysRequest](../../models/operations/tigrislistappkeysrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.TigrisListAppKeysResponse](../../models/operations/tigrislistappkeysresponse.md), error**


## Rotate

Endpoint is used to rotate the secret for the app key.

### Example Usage

```go
package main

import(
	"context"
	"log"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.Rotate(ctx, operations.TigrisRotateAppKeySecretRequest{
        RotateAppKeyRequest: shared.RotateAppKeyRequest{
            ID: tigrisgosdk.String("a1ba928f-c816-4742-8b73-9205929396fe"),
            Project: tigrisgosdk.String("fuga"),
        },
        Project: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RotateAppKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.TigrisRotateAppKeySecretRequest](../../models/operations/tigrisrotateappkeysecretrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.TigrisRotateAppKeySecretResponse](../../models/operations/tigrisrotateappkeysecretresponse.md), error**


## TigrisCreateAppKey

Create an app key.

### Example Usage

```go
package main

import(
	"context"
	"log"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.TigrisCreateAppKey(ctx, operations.TigrisCreateAppKeyRequest{
        CreateAppKeyRequest: shared.CreateAppKeyRequest{
            Description: tigrisgosdk.String("corporis"),
            Name: tigrisgosdk.String("Brad Turcotte Jr."),
        },
        Project: "reiciendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAppKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.TigrisCreateAppKeyRequest](../../models/operations/tigriscreateappkeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.TigrisCreateAppKeyResponse](../../models/operations/tigriscreateappkeyresponse.md), error**


## Update

Update the description of an app key.

### Example Usage

```go
package main

import(
	"context"
	"log"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.AppKey.Update(ctx, operations.TigrisUpdateAppKeyRequest{
        UpdateAppKeyRequest: shared.UpdateAppKeyRequest{
            Description: tigrisgosdk.String("est"),
            ID: tigrisgosdk.String("aa2352c5-9559-407a-bf1a-3a2fa9467739"),
            Name: tigrisgosdk.String("Beatrice Brown"),
        },
        Project: "enim",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateAppKeyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.TigrisUpdateAppKeyRequest](../../models/operations/tigrisupdateappkeyrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.TigrisUpdateAppKeyResponse](../../models/operations/tigrisupdateappkeyresponse.md), error**

