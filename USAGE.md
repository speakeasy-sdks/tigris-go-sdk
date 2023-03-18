<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/tigris-go-sdk"
    "github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    req := operations.TigrisDeleteAppKeyRequest{
        DeleteAppKeyRequest: shared.DeleteAppKeyRequest{
            ID: "unde",
        },
        Project: "deserunt",
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