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
            BearerAuth: "",
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

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CacheCreateCacheRequest](../../models/operations/cachecreatecacherequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CacheCreateCacheResponse](../../models/operations/cachecreatecacheresponse.md), error**


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
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.Delete(ctx, operations.CacheDeleteCacheRequest{
        DeleteCacheRequest: shared.DeleteCacheRequest{},
        Name: "Angelica Dietrich",
        Project: "id",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteCacheResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CacheDeleteCacheRequest](../../models/operations/cachedeletecacherequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CacheDeleteCacheResponse](../../models/operations/cachedeletecacheresponse.md), error**


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
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.DeleteKeys(ctx, operations.CacheDelRequest{
        DelRequest: shared.DelRequest{},
        Key: "possimus",
        Name: "Joyce Mueller",
        Project: "quasi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DelResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.CacheDelRequest](../../models/operations/cachedelrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CacheDelResponse](../../models/operations/cachedelresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.GetKey(ctx, operations.CacheGetRequest{
        Key: "reiciendis",
        Name: "Caleb Koss",
        Project: "ipsa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.CacheGetRequest](../../models/operations/cachegetrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CacheGetResponse](../../models/operations/cachegetresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.GetSetKey(ctx, operations.CacheGetSetRequest{
        GetSetRequest: shared.GetSetRequest{
            Value: tigris.String("omnis"),
        },
        Key: "voluptate",
        Name: "Thomas Batz",
        Project: "maiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetSetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CacheGetSetRequest](../../models/operations/cachegetsetrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CacheGetSetResponse](../../models/operations/cachegetsetresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.List(ctx, operations.CacheListCachesRequest{
        Project: "dicta",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListCachesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CacheListCachesRequest](../../models/operations/cachelistcachesrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.CacheListCachesResponse](../../models/operations/cachelistcachesresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.ListKeys(ctx, operations.CacheKeysRequest{
        Count: tigris.Int64(359444),
        Cursor: tigris.Int64(296140),
        Name: "Mildred Pfeffer",
        Pattern: tigris.String("commodi"),
        Project: "repudiandae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.KeysResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.CacheKeysRequest](../../models/operations/cachekeysrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.CacheKeysResponse](../../models/operations/cachekeysresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Cache.SetKey(ctx, operations.CacheSetRequest{
        SetRequest: shared.SetRequest{
            Ex: tigris.Int64(64147),
            Nx: tigris.Bool(false),
            Px: tigris.Int64(216822),
            Value: tigris.String("quidem"),
            Xx: tigris.Bool(false),
        },
        Key: "molestias",
        Name: "Ervin Gleason",
        Project: "voluptates",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.CacheSetRequest](../../models/operations/cachesetrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CacheSetResponse](../../models/operations/cachesetresponse.md), error**

