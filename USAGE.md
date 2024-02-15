<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := tigrisgosdk.New(
		tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Cache.Create(ctx, operations.CacheCreateCacheRequest{
		CreateCacheRequest: shared.CreateCacheRequest{},
		Name:               "<value>",
		Project:            "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateCacheResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->