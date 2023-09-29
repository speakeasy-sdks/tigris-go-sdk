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
	"context"
	"log"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Namespace.Create(ctx, shared.CreateNamespaceRequest{
        Code: tigrisgosdk.Int64(481196),
        ID: tigrisgosdk.String("<ID>"),
        Name: tigrisgosdk.String("Tasty island Southwest"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateNamespaceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateNamespaceRequest](../../models/shared/createnamespacerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateNamespaceResponse](../../models/operations/createnamespaceresponse.md), error**


## Get

Get details for all namespaces

### Example Usage

```go
package main

import(
	"context"
	"log"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
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

**[*operations.ManagementDescribeNamespacesResponse](../../models/operations/managementdescribenamespacesresponse.md), error**


## GetMetadata

GetNamespaceMetadata inserts the user metadata object

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
    res, err := s.Namespace.GetMetadata(ctx, operations.ManagementGetNamespaceMetadataRequest{
        GetNamespaceMetadataRequest: shared.GetNamespaceMetadataRequest{
            MetadataKey: tigrisgosdk.String("or Plastic"),
            Value: &shared.GetNamespaceMetadataRequestValue{},
        },
        MetadataKey: "offensively Electric",
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

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ManagementGetNamespaceMetadataRequest](../../models/operations/managementgetnamespacemetadatarequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.ManagementGetNamespaceMetadataResponse](../../models/operations/managementgetnamespacemetadataresponse.md), error**


## InsertMetadata

InsertNamespaceMetadata inserts the namespace metadata object

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
    res, err := s.Namespace.InsertMetadata(ctx, operations.ManagementInsertNamespaceMetadataRequest{
        InsertNamespaceMetadataRequest: shared.InsertNamespaceMetadataRequest{
            MetadataKey: tigrisgosdk.String("Curve Liaison calculate"),
            Value: &shared.InsertNamespaceMetadataRequestValue{},
        },
        MetadataKey: "female Tantalum",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ManagementInsertNamespaceMetadataRequest](../../models/operations/managementinsertnamespacemetadatarequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.ManagementInsertNamespaceMetadataResponse](../../models/operations/managementinsertnamespacemetadataresponse.md), error**


## List

List all namespace

### Example Usage

```go
package main

import(
	"context"
	"log"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
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

**[*operations.ManagementListNamespacesResponse](../../models/operations/managementlistnamespacesresponse.md), error**


## UpdateMetadata

UpdateNamespaceMetadata updates the user metadata object

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
    res, err := s.Namespace.UpdateMetadata(ctx, operations.ManagementUpdateNamespaceMetadataRequest{
        UpdateNamespaceMetadataRequest: shared.UpdateNamespaceMetadataRequest{
            MetadataKey: tigrisgosdk.String("platforms Concrete Tempe"),
            Value: &shared.UpdateNamespaceMetadataRequestValue{},
        },
        MetadataKey: "transmitting Silicon North",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ManagementUpdateNamespaceMetadataRequest](../../models/operations/managementupdatenamespacemetadatarequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.ManagementUpdateNamespaceMetadataResponse](../../models/operations/managementupdatenamespacemetadataresponse.md), error**

