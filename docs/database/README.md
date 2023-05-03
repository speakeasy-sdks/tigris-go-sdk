# Database

## Overview

The Database section provide APIs that can be used to interact with the database. A single Database can have one or more collections. A database is created automatically for you when you create a project.

### Available Operations

* [BeginTransaction](#begintransaction) - Begin a transaction
* [CommitTransaction](#committransaction) - Commit a Transaction
* [CreateBranch](#createbranch) - Create a database branch
* [DeleteBranch](#deletebranch) - Delete a database branch
* [Describe](#describe) - Describe database
* [ListCollections](#listcollections) - List Collections
* [RollbackTransaction](#rollbacktransaction) - Rollback a transaction
* [TigrisListBranches](#tigrislistbranches) - List database branches

## BeginTransaction

Starts a new transaction and returns a transactional object. All reads/writes performed
 within a transaction will run with serializable isolation. Tigris offers global transactions,
 with ACID properties and strict serializability.

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
    res, err := s.Database.BeginTransaction(ctx, operations.TigrisBeginTransactionRequest{
        BeginTransactionRequest: shared.BeginTransactionRequest{
            Branch: tigris.String("maxime"),
            Options: map[string]interface{}{
                "soluta": "dicta",
                "laborum": "totam",
                "incidunt": "aspernatur",
                "dolores": "distinctio",
            },
        },
        Project: "facilis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BeginTransactionResponse != nil {
        // handle response
    }
}
```

## CommitTransaction

Atomically commit all the changes performed in the context of the transaction. Commit provides all
 or nothing semantics by ensuring no partial updates are in the project due to a transaction failure.

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
    res, err := s.Database.CommitTransaction(ctx, operations.TigrisCommitTransactionRequest{
        CommitTransactionRequest: shared.CommitTransactionRequest{
            Branch: tigris.String("aliquid"),
        },
        Project: "quam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CommitTransactionResponse != nil {
        // handle response
    }
}
```

## CreateBranch

Creates a new database branch, if not already existing.

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
    res, err := s.Database.CreateBranch(ctx, operations.TigrisCreateBranchRequest{
        RequestBody: map[string]interface{}{
            "temporibus": "qui",
            "neque": "fugit",
            "magni": "odio",
        },
        Branch: "sunt",
        Project: "ullam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateBranchResponse != nil {
        // handle response
    }
}
```

## DeleteBranch

Deletes a database branch, if exists.
 Throws 400 Bad Request if "main" branch is being deleted

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
    res, err := s.Database.DeleteBranch(ctx, operations.TigrisDeleteBranchRequest{
        RequestBody: map[string]interface{}{
            "hic": "voluptatem",
            "cumque": "soluta",
            "nobis": "et",
        },
        Branch: "saepe",
        Project: "ipsum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteBranchResponse != nil {
        // handle response
    }
}
```

## Describe

This API returns information related to the project along with all the collections inside the project.
 This can be used to retrieve the size of the project or to retrieve schemas, branches and the size of all the collections present in this project.

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
    res, err := s.Database.Describe(ctx, operations.TigrisDescribeDatabaseRequest{
        DescribeDatabaseRequest: shared.DescribeDatabaseRequest{
            Branch: tigris.String("veritatis"),
            Project: tigris.String("nobis"),
            SchemaFormat: tigris.String("quos"),
        },
        Project: "tempore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DescribeDatabaseResponse != nil {
        // handle response
    }
}
```

## ListCollections

List all the collections present in the project passed in the request.

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
    res, err := s.Database.ListCollections(ctx, operations.TigrisListCollectionsRequest{
        Branch: tigris.String("cupiditate"),
        Project: "aperiam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListCollectionsResponse != nil {
        // handle response
    }
}
```

## RollbackTransaction

Rollback transaction discards all the changes
 performed in the transaction

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
    res, err := s.Database.RollbackTransaction(ctx, operations.TigrisRollbackTransactionRequest{
        RollbackTransactionRequest: shared.RollbackTransactionRequest{
            Branch: tigris.String("delectus"),
        },
        Project: "dolorem",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RollbackTransactionResponse != nil {
        // handle response
    }
}
```

## TigrisListBranches

List database branches

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
    res, err := s.Database.TigrisListBranches(ctx, operations.TigrisListBranchesRequest{
        Project: "dolore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListBranchesResponse != nil {
        // handle response
    }
}
```
