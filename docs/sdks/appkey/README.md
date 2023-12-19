# AppKey
(*AppKey*)

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
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.AppKey.Delete(ctx, operations.TigrisDeleteAppKeyRequest{
        DeleteAppKeyRequest: shared.DeleteAppKeyRequest{},
        Project: "string",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.TigrisDeleteAppKeyRequest](../../pkg/models/operations/tigrisdeleteappkeyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.TigrisDeleteAppKeyResponse](../../pkg/models/operations/tigrisdeleteappkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Lists all app keys visible to requesting actor.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.AppKey.List(ctx, operations.TigrisListAppKeysRequest{
        Project: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.TigrisListAppKeysRequest](../../pkg/models/operations/tigrislistappkeysrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.TigrisListAppKeysResponse](../../pkg/models/operations/tigrislistappkeysresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Rotate

Endpoint is used to rotate the secret for the app key.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.AppKey.Rotate(ctx, operations.TigrisRotateAppKeySecretRequest{
        RotateAppKeyRequest: shared.RotateAppKeyRequest{},
        Project: "string",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.TigrisRotateAppKeySecretRequest](../../pkg/models/operations/tigrisrotateappkeysecretrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.TigrisRotateAppKeySecretResponse](../../pkg/models/operations/tigrisrotateappkeysecretresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## TigrisCreateAppKey

Create an app key.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.AppKey.TigrisCreateAppKey(ctx, operations.TigrisCreateAppKeyRequest{
        CreateAppKeyRequest: shared.CreateAppKeyRequest{},
        Project: "string",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.TigrisCreateAppKeyRequest](../../pkg/models/operations/tigriscreateappkeyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.TigrisCreateAppKeyResponse](../../pkg/models/operations/tigriscreateappkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Update the description of an app key.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.AppKey.Update(ctx, operations.TigrisUpdateAppKeyRequest{
        UpdateAppKeyRequest: shared.UpdateAppKeyRequest{},
        Project: "string",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.TigrisUpdateAppKeyRequest](../../pkg/models/operations/tigrisupdateappkeyrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.TigrisUpdateAppKeyResponse](../../pkg/models/operations/tigrisupdateappkeyresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
