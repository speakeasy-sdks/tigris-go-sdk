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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.CreateNamespaceRequest{
        Code: tigris.Int64(286915),
        ID: tigris.String("3a1108e0-adcf-44b9-a187-9fce953f73ef"),
        Name: tigris.String("Darla Rau"),
    }

    res, err := s.Namespace.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateNamespaceResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ManagementGetNamespaceMetadataRequest{
        GetNamespaceMetadataRequest: shared.GetNamespaceMetadataRequest{
            MetadataKey: tigris.String("similique"),
            Value: map[string]interface{}{
                "vero": "ducimus",
                "dolore": "quibusdam",
                "illum": "sequi",
            },
        },
        MetadataKey: "natus",
    }

    res, err := s.Namespace.GetMetadata(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetNamespaceMetadataResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ManagementInsertNamespaceMetadataRequest{
        InsertNamespaceMetadataRequest: shared.InsertNamespaceMetadataRequest{
            MetadataKey: tigris.String("impedit"),
            Value: map[string]interface{}{
                "voluptatibus": "exercitationem",
            },
        },
        MetadataKey: "nulla",
    }

    res, err := s.Namespace.InsertMetadata(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.InsertNamespaceMetadataResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.ManagementUpdateNamespaceMetadataRequest{
        UpdateNamespaceMetadataRequest: shared.UpdateNamespaceMetadataRequest{
            MetadataKey: tigris.String("fugit"),
            Value: map[string]interface{}{
                "maiores": "doloribus",
                "iusto": "eligendi",
                "ducimus": "alias",
                "officia": "tempora",
            },
        },
        MetadataKey: "ipsam",
    }

    res, err := s.Namespace.UpdateMetadata(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateNamespaceMetadataResponse != nil {
        // handle response
    }
}
```
