# Collection

## Overview

The Collections section provide APIs that can be used to manage collections. A collection can have one or more documents.

### Available Operations

* [Create](#create) - Create or update a collection
* [DeleteDocuments](#deletedocuments) - Delete Documents
* [Describe](#describe) - Describe Collection
* [Drop](#drop) - Drop Collection
* [ImportDocuments](#importdocuments) - Import Documents
* [InsertDocuments](#insertdocuments) - Insert Documents
* [ReadDocuments](#readdocuments) - Read Documents
* [ReplaceDocuments](#replacedocuments) - Insert or Replace Documents
* [SearchDocuments](#searchdocuments) - Search Documents.
* [UpdateDocuments](#updatedocuments) - Update Documents.

## Create

Creates a new collection or atomically upgrades the collection to the new schema provided in the request.
 Schema changes are applied atomically and immediately without any downtime.
 Tigris Offers two types of collections: <p></p>
    <li> `DOCUMENTS`: Offers rich CRUD APIs.
    <li> `MESSAGES`: Offers event streaming APIs.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.Create(ctx, operations.TigrisCreateOrUpdateCollectionRequest{
        CreateOrUpdateCollectionRequest: shared.CreateOrUpdateCollectionRequest{
            Branch: tigris.String("occaecati"),
            OnlyCreate: tigris.Bool(false),
            Options: &shared.CollectionOptions{},
            Schema: &shared.CreateOrUpdateCollectionRequestSchema{},
        },
        Collection: "enim",
        Project: "accusamus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOrUpdateCollectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.TigrisCreateOrUpdateCollectionRequest](../../models/operations/tigriscreateorupdatecollectionrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.TigrisCreateOrUpdateCollectionResponse](../../models/operations/tigriscreateorupdatecollectionresponse.md), error**


## DeleteDocuments

Delete a range of documents in the collection using the condition provided in the filter.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.DeleteDocuments(ctx, operations.TigrisDeleteRequest{
        DeleteRequest: shared.DeleteRequest{
            Branch: tigris.String("delectus"),
            Filter: &shared.DeleteRequestFilter{},
            Options: &shared.DeleteRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("quidem"),
                },
                Limit: tigris.Int64(588465),
                WriteOptions: &shared.WriteOptions{},
            },
        },
        Collection: "nam",
        Project: "id",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.TigrisDeleteRequest](../../models/operations/tigrisdeleterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.TigrisDeleteResponse](../../models/operations/tigrisdeleteresponse.md), error**


## Describe

Returns the information related to the collection. This can be used to retrieve the schema or size of the collection.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.Describe(ctx, operations.TigrisDescribeCollectionRequest{
        DescribeCollectionRequest: shared.DescribeCollectionRequest{
            Branch: tigris.String("blanditiis"),
            Collection: tigris.String("deleniti"),
            Options: &shared.CollectionOptions{},
            Project: tigris.String("sapiente"),
            SchemaFormat: tigris.String("amet"),
        },
        Collection: "deserunt",
        Project: "nisi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DescribeCollectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.TigrisDescribeCollectionRequest](../../models/operations/tigrisdescribecollectionrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.TigrisDescribeCollectionResponse](../../models/operations/tigrisdescribecollectionresponse.md), error**


## Drop

Drops the collection inside this project. This API deletes all of the
 documents inside this collection and any metadata associated with it.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.Drop(ctx, operations.TigrisDropCollectionRequest{
        DropCollectionRequest: shared.DropCollectionRequest{
            Branch: tigris.String("vel"),
            Options: &shared.CollectionOptions{},
        },
        Collection: "natus",
        Project: "omnis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DropCollectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.TigrisDropCollectionRequest](../../models/operations/tigrisdropcollectionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.TigrisDropCollectionResponse](../../models/operations/tigrisdropcollectionresponse.md), error**


## ImportDocuments

Imports documents into the collection.

 It automatically:
  * Detects the schema of the documents in the batch
  * Evolves the schema as soon as it's backward compatible
  * Creates collection with inferred schema (if requested)

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.ImportDocuments(ctx, operations.TigrisImportRequest{
        ImportRequest: shared.ImportRequest{
            Autogenerated: []string{
                "molestiae",
            },
            Branch: tigris.String("perferendis"),
            CreateCollection: tigris.Bool(false),
            Documents: []shared.ImportRequestDocuments{
                shared.ImportRequestDocuments{},
            },
            Options: &shared.ImportRequestOptions{
                WriteOptions: &shared.WriteOptions{},
            },
            PrimaryKey: []string{
                "nihil",
            },
        },
        Collection: "magnam",
        Project: "distinctio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.TigrisImportRequest](../../models/operations/tigrisimportrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.TigrisImportResponse](../../models/operations/tigrisimportresponse.md), error**


## InsertDocuments

Inserts new documents in the collection and returns an AlreadyExists error if any of the documents
 in the request already exists. Insert provides idempotency by returning an error if the document
 already exists. To replace documents, use REPLACE API instead of INSERT.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.InsertDocuments(ctx, operations.TigrisInsertRequest{
        InsertRequest: shared.InsertRequest{
            Branch: tigris.String("id"),
            Documents: []shared.InsertRequestDocuments{
                shared.InsertRequestDocuments{},
            },
            Options: &shared.InsertRequestOptions{
                WriteOptions: &shared.WriteOptions{},
            },
        },
        Collection: "labore",
        Project: "labore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InsertResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.TigrisInsertRequest](../../models/operations/tigrisinsertrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.TigrisInsertResponse](../../models/operations/tigrisinsertresponse.md), error**


## ReadDocuments

Reads a range of documents from the collection, or messages from a collection in case of event streaming. Tigris does not require you to
 index any fields and automatically index all the fields which means you can filter by any field in the document.
 An empty filter will trigger reading all the documents inside this collection. The API supports pagination that
 can be used by passing `Limit/Skip` parameters. The `skip` parameter skips the number of documents from the start and
 the `limit` parameter is used to specify the number of documents to read. You can find more detailed documentation
 of the Read API <a href="https://docs.tigrisdata.com/overview/query" title="here">here</a>.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.ReadDocuments(ctx, operations.TigrisReadRequest{
        ReadRequest: shared.ReadRequest{
            Branch: tigris.String("suscipit"),
            Fields: &shared.ReadRequestFields{},
            Filter: &shared.ReadRequestFilter{},
            Options: &shared.ReadRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("natus"),
                },
                Limit: tigris.Int64(749170),
                Offset: tigris.String("eum"),
                Skip: tigris.Int64(878453),
            },
            Sort: tigris.String("aspernatur"),
        },
        Collection: "architecto",
        Project: "magnam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StreamingReadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.TigrisReadRequest](../../models/operations/tigrisreadrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.TigrisReadResponse](../../models/operations/tigrisreadresponse.md), error**


## ReplaceDocuments

Inserts the documents or replaces the existing documents in the collections.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.ReplaceDocuments(ctx, operations.TigrisReplaceRequest{
        ReplaceRequest: shared.ReplaceRequest{
            Branch: tigris.String("et"),
            Documents: []shared.ReplaceRequestDocuments{
                shared.ReplaceRequestDocuments{},
            },
            Options: &shared.ReplaceRequestOptions{
                WriteOptions: &shared.WriteOptions{},
            },
        },
        Collection: "excepturi",
        Project: "ullam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReplaceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.TigrisReplaceRequest](../../models/operations/tigrisreplacerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.TigrisReplaceResponse](../../models/operations/tigrisreplaceresponse.md), error**


## SearchDocuments

Searches a collection for the documents matching the query, or messages in case of event streaming. A search can be
 a term search or a phrase search. Search API allows filtering the result set using filters as documented <a href="https://docs.tigrisdata.com/overview/query#specification-1" title="here">here</a>.
 You can also perform a faceted search by passing the fields in the facet parameter.
 You can find more detailed documentation of the Search API with multiple examples <a href="https://docs.tigrisdata.com/overview/search" title="here">here</a>.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.SearchDocuments(ctx, operations.TigrisSearchRequest{
        SearchRequest: shared.SearchRequest{
            Branch: tigris.String("provident"),
            Collation: &shared.Collation{
                Case: tigris.String("quos"),
            },
            ExcludeFields: []string{
                "sint",
            },
            Facet: &shared.SearchRequestFacet{},
            Fields: &shared.SearchRequestFields{},
            Filter: &shared.SearchRequestFilter{},
            IncludeFields: []string{
                "accusantium",
            },
            Page: tigris.Int(653201),
            PageSize: tigris.Int(968962),
            Q: tigris.String("mollitia"),
            SearchFields: []string{
                "ad",
            },
            Sort: &shared.SearchRequestSort{},
        },
        Collection: "eum",
        Project: "dolor",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StreamingSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.TigrisSearchRequest](../../models/operations/tigrissearchrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.TigrisSearchResponse](../../models/operations/tigrissearchresponse.md), error**


## UpdateDocuments

Update a range of documents in the collection using the condition provided in the filter.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.UpdateDocuments(ctx, operations.TigrisUpdateRequest{
        UpdateRequest: shared.UpdateRequest{
            Branch: tigris.String("necessitatibus"),
            Fields: &shared.UpdateRequestFields{},
            Filter: &shared.UpdateRequestFilter{},
            Options: &shared.UpdateRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("odit"),
                },
                Limit: tigris.Int64(367562),
                WriteOptions: &shared.WriteOptions{},
            },
        },
        Collection: "quasi",
        Project: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.TigrisUpdateRequest](../../models/operations/tigrisupdaterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.TigrisUpdateResponse](../../models/operations/tigrisupdateresponse.md), error**

