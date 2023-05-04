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
    res, err := s.Cache.Create(ctx, operations.CacheCreateCacheRequest{
        CreateCacheRequest: shared.CreateCacheRequest{
            Options: &shared.CreateCacheOptions{
                TTLMs: tigris.Int64(548814),
            },
        },
        Name: "Kelvin Sporer",
        Project: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCacheResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->