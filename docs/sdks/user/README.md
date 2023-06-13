# User

## Overview

A User on the Tigris Platform.

### Available Operations

* [GetMetadata](#getmetadata) - Reads the User Metadata
* [InsertMetadata](#insertmetadata) - Inserts User Metadata
* [UpdateMetadata](#updatemetadata) - Updates User Metadata

## GetMetadata

GetUserMetadata inserts the user metadata object

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
    res, err := s.User.GetMetadata(ctx, operations.ManagementGetUserMetadataRequest{
        GetUserMetadataRequest: shared.GetUserMetadataRequest{
            MetadataKey: tigris.String("distinctio"),
            Value: &shared.GetUserMetadataRequestValue{},
        },
        MetadataKey: "quod",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUserMetadataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ManagementGetUserMetadataRequest](../../models/operations/managementgetusermetadatarequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.ManagementGetUserMetadataResponse](../../models/operations/managementgetusermetadataresponse.md), error**


## InsertMetadata

insertUserMetadata inserts the user metadata object

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
    res, err := s.User.InsertMetadata(ctx, operations.ManagementInsertUserMetadataRequest{
        InsertUserMetadataRequest: shared.InsertUserMetadataRequest{
            MetadataKey: tigris.String("odio"),
            Value: &shared.InsertUserMetadataRequestValue{},
        },
        MetadataKey: "similique",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InsertUserMetadataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ManagementInsertUserMetadataRequest](../../models/operations/managementinsertusermetadatarequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.ManagementInsertUserMetadataResponse](../../models/operations/managementinsertusermetadataresponse.md), error**


## UpdateMetadata

updateUserMetadata updates the user metadata object

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
    res, err := s.User.UpdateMetadata(ctx, operations.ManagementUpdateUserMetadataRequest{
        UpdateUserMetadataRequest: shared.UpdateUserMetadataRequest{
            MetadataKey: tigris.String("facilis"),
            Value: &shared.UpdateUserMetadataRequestValue{},
        },
        MetadataKey: "vero",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateUserMetadataResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ManagementUpdateUserMetadataRequest](../../models/operations/managementupdateusermetadatarequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.ManagementUpdateUserMetadataResponse](../../models/operations/managementupdateusermetadataresponse.md), error**
