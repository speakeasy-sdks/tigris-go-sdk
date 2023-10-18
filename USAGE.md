<!-- Start SDK Example Usage -->


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
		tigrisgosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Cache.Create(ctx, operations.CacheCreateCacheRequest{
		CreateCacheRequest: shared.CreateCacheRequest{
			Options: &shared.CreateCacheOptions{},
		},
		Name:    "neural",
		Project: "Tasty",
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