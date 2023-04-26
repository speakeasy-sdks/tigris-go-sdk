<!-- Start SDK Example Usage -->
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
    req := operations.TigrisDeleteAppKeyRequest{
        DeleteAppKeyRequest: shared.DeleteAppKeyRequest{
            ID: tigris.String("89bd9d8d-69a6-474e-8f46-7cc8796ed151"),
        },
        Project: "deserunt",
    }

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