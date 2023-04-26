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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CacheCreateCacheRequest{
        CreateCacheRequest: shared.CreateCacheRequest{
            Options: &shared.CreateCacheOptions{
                TTLMs: tigris.Int64(509624),
            },
        },
        Name: "Jose Moen",
        Project: "perferendis",
    }

    res, err := s.Cache.Create(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CacheDeleteCacheRequest{
        RequestBody: map[string]interface{}{
            "reprehenderit": "ut",
        },
        Name: "Willie Hessel",
        Project: "dicta",
    }

    res, err := s.Cache.Delete(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CacheDelRequest{
        RequestBody: map[string]interface{}{
            "enim": "accusamus",
            "commodi": "repudiandae",
            "quae": "ipsum",
        },
        Key: "quidem",
        Name: "Andy Streich",
        Project: "rem",
    }

    res, err := s.Cache.DeleteKeys(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CacheGetRequest{
        Key: "voluptates",
        Name: "Dr. Casey Mayer",
        Project: "enim",
    }

    res, err := s.Cache.GetKey(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CacheGetSetRequest{
        GetSetRequest: shared.GetSetRequest{
            Value: tigris.String("consequatur"),
        },
        Key: "est",
        Name: "Benjamin O'Connell",
        Project: "labore",
    }

    res, err := s.Cache.GetSetKey(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CacheListCachesRequest{
        Project: "modi",
    }

    res, err := s.Cache.List(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CacheKeysRequest{
        Count: tigris.Int64(183191),
        Cursor: tigris.Int64(397821),
        Name: "Isaac Aufderhar",
        Pattern: tigris.String("ipsam"),
        Project: "alias",
    }

    res, err := s.Cache.ListKeys(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.CacheSetRequest{
        SetRequest: shared.SetRequest{
            Ex: tigris.Int64(146441),
            Nx: tigris.Bool(false),
            Px: tigris.Int64(677817),
            Value: tigris.String("excepturi"),
            Xx: tigris.Bool(false),
        },
        Key: "tempora",
        Name: "Geoffrey Green",
        Project: "non",
    }

    res, err := s.Cache.SetKey(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SetResponse != nil {
        // handle response
    }
}
```
