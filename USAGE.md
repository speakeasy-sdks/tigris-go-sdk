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
    s := tigris.New(tigris.WithSecurity(
        shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        },
    ))
    
    req := operations.TigrisDeleteAppKeyRequest{
        PathParams: operations.TigrisDeleteAppKeyPathParams{
            Project: "unde",
        },
        Request: shared.DeleteAppKeyRequest{
            ID: "deserunt",
        },
    }

    ctx := context.Background()
    res, err := s.AppKey.Delete(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAppKeyResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->