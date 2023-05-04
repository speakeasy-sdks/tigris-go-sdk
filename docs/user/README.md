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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.User.GetMetadata(ctx, operations.ManagementGetUserMetadataRequest{
        GetUserMetadataRequest: shared.GetUserMetadataRequest{
            MetadataKey: tigris.String("consectetur"),
            Value: map[string]interface{}{
                "iste": "temporibus",
            },
        },
        MetadataKey: "accusantium",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUserMetadataResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.User.InsertMetadata(ctx, operations.ManagementInsertUserMetadataRequest{
        InsertUserMetadataRequest: shared.InsertUserMetadataRequest{
            MetadataKey: tigris.String("rem"),
            Value: map[string]interface{}{
                "laudantium": "eum",
            },
        },
        MetadataKey: "mollitia",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InsertUserMetadataResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.User.UpdateMetadata(ctx, operations.ManagementUpdateUserMetadataRequest{
        UpdateUserMetadataRequest: shared.UpdateUserMetadataRequest{
            MetadataKey: tigris.String("ab"),
            Value: map[string]interface{}{
                "non": "voluptatem",
                "dolor": "occaecati",
                "numquam": "impedit",
            },
        },
        MetadataKey: "explicabo",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateUserMetadataResponse != nil {
        // handle response
    }
}
```
