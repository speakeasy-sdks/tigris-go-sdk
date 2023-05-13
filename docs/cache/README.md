# Cache

## Overview

The cache section provide APIs that can be used to perform cache operations.

### Available Operations

* [Create](#create) - Creates the cache
* [Delete](#delete) - Deletes the cache
* [DeleteKeys](#deletekeys) - Deletes an entry from cache
* [GetKey](#getkey) - Reads an entry from cache
* [GetSetKey](#getsetkey) - Sets an entry in the cache and returns the previous value if exists
* [List](#list) - Lists all the caches for the given project
* [ListKeys](#listkeys) - Lists all the key for this cache
* [SetKey](#setkey) - Sets an entry in the cache

## Create

Creates the cache

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.Create(ctx, operations.CacheCreateCacheRequest{
        CreateCacheRequest: shared.CreateCacheRequest{
            Options: &shared.CreateCacheOptions{
                TTLMs: tigris.Int64(244425),
            },
        },
        Name: "Miss Eugene Hauck",
        Project: "enim",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateCacheResponse != nil {
        // handle response
    }
}
```

## Delete

Deletes the cache

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.Delete(ctx, operations.CacheDeleteCacheRequest{
        RequestBody: map[string]interface{}{
            "quo": "sequi",
        },
        Name: "Vernon Ondricka Sr.",
        Project: "error",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteCacheResponse != nil {
        // handle response
    }
}
```

## DeleteKeys

Deletes an entry from cache

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.DeleteKeys(ctx, operations.CacheDelRequest{
        RequestBody: map[string]interface{}{
            "laborum": "quasi",
            "reiciendis": "voluptatibus",
            "vero": "nihil",
            "praesentium": "voluptatibus",
        },
        Key: "ipsa",
        Name: "Mr. Jared Ritchie",
        Project: "ut",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DelResponse != nil {
        // handle response
    }
}
```

## GetKey

Reads an entry from cache

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.GetKey(ctx, operations.CacheGetRequest{
        Key: "maiores",
        Name: "Stacy Gulgowski MD",
        Project: "enim",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetResponse != nil {
        // handle response
    }
}
```

## GetSetKey

Sets an entry in the cache and returns the previous value if exists

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.GetSetKey(ctx, operations.CacheGetSetRequest{
        GetSetRequest: shared.GetSetRequest{
            Value: tigris.String("accusamus"),
        },
        Key: "commodi",
        Name: "Eric Emmerich",
        Project: "excepturi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSetResponse != nil {
        // handle response
    }
}
```

## List

Lists all the caches for the given project

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.List(ctx, operations.CacheListCachesRequest{
        Project: "pariatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListCachesResponse != nil {
        // handle response
    }
}
```

## ListKeys

Lists all the key for this cache

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.ListKeys(ctx, operations.CacheKeysRequest{
        Count: tigris.Int64(265389),
        Cursor: tigris.Int64(508969),
        Name: "Grady Botsford",
        Pattern: tigris.String("veritatis"),
        Project: "itaque",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KeysResponse != nil {
        // handle response
    }
}
```

## SetKey

Sets an entry in the cache

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.SetKey(ctx, operations.CacheSetRequest{
        SetRequest: shared.SetRequest{
            Ex: tigris.Int64(277718),
            Nx: tigris.Bool(false),
            Px: tigris.Int64(318569),
            Value: tigris.String("consequatur"),
            Xx: tigris.Bool(false),
        },
        Key: "est",
        Name: "Benjamin O'Connell",
        Project: "labore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SetResponse != nil {
        // handle response
    }
}
```
