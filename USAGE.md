<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "tigris-core"
    "tigris-core/pkg/models/shared"
    "tigris-core/pkg/models/operations"
)

func main() {
    s := sdk.New(sdk.WithSecurity(
        shared.Security{
            BearerAuth: shared.SchemeBearerAuth{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        },
    ))
    
    req := operations.TigrisCreateAppKeyRequest{
        PathParams: operations.TigrisCreateAppKeyPathParams{
            Project: "unde",
        },
        Request: shared.CreateAppKeyRequest{
            Description: "deserunt",
            Name: "porro",
        },
    }

    ctx := context.Background()
    res, err := s.ApplicationKeys.TigrisCreateAppKey(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAppKeyResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->