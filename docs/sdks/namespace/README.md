# Namespace

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
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Namespace.Create(ctx, shared.CreateNamespaceRequest{
        Code: tigris.Int64(982575),
        ID: tigris.String("b576b0d5-f0d3-40c5-bbb2-587053202c73"),
        Name: tigris.String("Dean Welch"),
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
	"github.com/speakeasy-sdks/tigris-go-sdk"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
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
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Namespace.GetMetadata(ctx, operations.ManagementGetNamespaceMetadataRequest{
        GetNamespaceMetadataRequest: shared.GetNamespaceMetadataRequest{
            MetadataKey: tigris.String("facilis"),
            Value: &shared.GetNamespaceMetadataRequestValue{},
        },
        MetadataKey: "perspiciatis",
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
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Namespace.InsertMetadata(ctx, operations.ManagementInsertNamespaceMetadataRequest{
        InsertNamespaceMetadataRequest: shared.InsertNamespaceMetadataRequest{
            MetadataKey: tigris.String("voluptatem"),
            Value: &shared.InsertNamespaceMetadataRequestValue{},
        },
        MetadataKey: "porro",
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
	"github.com/speakeasy-sdks/tigris-go-sdk"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
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
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Namespace.UpdateMetadata(ctx, operations.ManagementUpdateNamespaceMetadataRequest{
        UpdateNamespaceMetadataRequest: shared.UpdateNamespaceMetadataRequest{
            MetadataKey: tigris.String("consequuntur"),
            Value: &shared.UpdateNamespaceMetadataRequestValue{},
        },
        MetadataKey: "blanditiis",
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

