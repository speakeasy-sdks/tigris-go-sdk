# Namespace
(*Namespace*)

## Overview

The Management section provide APIs that can be used to manage users, and app_keys.

### Available Operations

* [Create](#create) - Creates a Namespace
* [Get](#get) - Describe the details of all namespaces
* [GetMetadata](#getmetadata) - Reads the Namespace Metadata
* [InsertMetadata](#insertmetadata) - Inserts Namespace Metadata
* [List](#list) - Lists all Namespaces
* [UpdateMetadata](#updatemetadata) - Updates Namespace Metadata

## Create

Creates a new namespace, if it does not exist

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Namespace.Create(ctx, shared.CreateNamespaceRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateNamespaceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.CreateNamespaceRequest](../../pkg/models/shared/createnamespacerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateNamespaceResponse](../../pkg/models/operations/createnamespaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Get details for all namespaces

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Namespace.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.DescribeNamespacesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ManagementDescribeNamespacesResponse](../../pkg/models/operations/managementdescribenamespacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetMetadata

GetNamespaceMetadata inserts the user metadata object

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
    res, err := s.Namespace.GetMetadata(ctx, operations.ManagementGetNamespaceMetadataRequest{
        GetNamespaceMetadataRequest: shared.GetNamespaceMetadataRequest{
            Value: &shared.Value{},
        },
        MetadataKey: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetNamespaceMetadataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.ManagementGetNamespaceMetadataRequest](../../pkg/models/operations/managementgetnamespacemetadatarequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |


### Response

**[*operations.ManagementGetNamespaceMetadataResponse](../../pkg/models/operations/managementgetnamespacemetadataresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## InsertMetadata

InsertNamespaceMetadata inserts the namespace metadata object

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
    res, err := s.Namespace.InsertMetadata(ctx, operations.ManagementInsertNamespaceMetadataRequest{
        InsertNamespaceMetadataRequest: shared.InsertNamespaceMetadataRequest{
            Value: &shared.InsertNamespaceMetadataRequestValue{},
        },
        MetadataKey: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InsertNamespaceMetadataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ManagementInsertNamespaceMetadataRequest](../../pkg/models/operations/managementinsertnamespacemetadatarequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.ManagementInsertNamespaceMetadataResponse](../../pkg/models/operations/managementinsertnamespacemetadataresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List all namespace

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Namespace.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListNamespacesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ManagementListNamespacesResponse](../../pkg/models/operations/managementlistnamespacesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateMetadata

UpdateNamespaceMetadata updates the user metadata object

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
    res, err := s.Namespace.UpdateMetadata(ctx, operations.ManagementUpdateNamespaceMetadataRequest{
        UpdateNamespaceMetadataRequest: shared.UpdateNamespaceMetadataRequest{
            Value: &shared.UpdateNamespaceMetadataRequestValue{},
        },
        MetadataKey: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateNamespaceMetadataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ManagementUpdateNamespaceMetadataRequest](../../pkg/models/operations/managementupdatenamespacemetadatarequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.ManagementUpdateNamespaceMetadataResponse](../../pkg/models/operations/managementupdatenamespacemetadataresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
