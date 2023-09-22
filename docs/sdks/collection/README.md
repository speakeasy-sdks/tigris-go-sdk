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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.Create(ctx, operations.TigrisCreateOrUpdateCollectionRequest{
        CreateOrUpdateCollectionRequest: shared.CreateOrUpdateCollectionRequest{
            Branch: tigrisgosdk.String("id"),
            OnlyCreate: tigrisgosdk.Bool(false),
            Options: &shared.CollectionOptions{},
            Schema: &shared.CreateOrUpdateCollectionRequestSchema{},
        },
        Collection: "blanditiis",
        Project: "deleniti",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.DeleteDocuments(ctx, operations.TigrisDeleteRequest{
        DeleteRequest: shared.DeleteRequest{
            Branch: tigrisgosdk.String("sapiente"),
            Filter: &shared.DeleteRequestFilter{},
            Options: &shared.DeleteRequestOptions{
                Collation: &shared.Collation{
                    Case: tigrisgosdk.String("amet"),
                },
                Limit: tigrisgosdk.Int64(643990),
                WriteOptions: &shared.WriteOptions{},
            },
        },
        Collection: "nisi",
        Project: "vel",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.Describe(ctx, operations.TigrisDescribeCollectionRequest{
        DescribeCollectionRequest: shared.DescribeCollectionRequest{
            Branch: tigrisgosdk.String("natus"),
            Collection: tigrisgosdk.String("omnis"),
            Options: &shared.CollectionOptions{},
            Project: tigrisgosdk.String("molestiae"),
            SchemaFormat: tigrisgosdk.String("perferendis"),
        },
        Collection: "nihil",
        Project: "magnam",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.Drop(ctx, operations.TigrisDropCollectionRequest{
        DropCollectionRequest: shared.DropCollectionRequest{
            Branch: tigrisgosdk.String("distinctio"),
            Options: &shared.CollectionOptions{},
        },
        Collection: "id",
        Project: "labore",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.ImportDocuments(ctx, operations.TigrisImportRequest{
        ImportRequest: shared.ImportRequest{
            Autogenerated: []string{
                "labore",
            },
            Branch: tigrisgosdk.String("suscipit"),
            CreateCollection: tigrisgosdk.Bool(false),
            Documents: []shared.ImportRequestDocuments{
                shared.ImportRequestDocuments{},
            },
            Options: &shared.ImportRequestOptions{
                WriteOptions: &shared.WriteOptions{},
            },
            PrimaryKey: []string{
                "natus",
            },
        },
        Collection: "nobis",
        Project: "eum",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.InsertDocuments(ctx, operations.TigrisInsertRequest{
        InsertRequest: shared.InsertRequest{
            Branch: tigrisgosdk.String("vero"),
            Documents: []shared.InsertRequestDocuments{
                shared.InsertRequestDocuments{},
            },
            Options: &shared.InsertRequestOptions{
                WriteOptions: &shared.WriteOptions{},
            },
        },
        Collection: "aspernatur",
        Project: "architecto",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.ReadDocuments(ctx, operations.TigrisReadRequest{
        ReadRequest: shared.ReadRequest{
            Branch: tigrisgosdk.String("magnam"),
            Fields: &shared.ReadRequestFields{},
            Filter: &shared.ReadRequestFilter{},
            Options: &shared.ReadRequestOptions{
                Collation: &shared.Collation{
                    Case: tigrisgosdk.String("et"),
                },
                Limit: tigrisgosdk.Int64(569965),
                Offset: tigrisgosdk.String("ullam"),
                Skip: tigrisgosdk.Int64(590873),
            },
            Sort: tigrisgosdk.String("quos"),
        },
        Collection: "sint",
        Project: "accusantium",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.ReplaceDocuments(ctx, operations.TigrisReplaceRequest{
        ReplaceRequest: shared.ReplaceRequest{
            Branch: tigrisgosdk.String("mollitia"),
            Documents: []shared.ReplaceRequestDocuments{
                shared.ReplaceRequestDocuments{},
            },
            Options: &shared.ReplaceRequestOptions{
                WriteOptions: &shared.WriteOptions{},
            },
        },
        Collection: "reiciendis",
        Project: "mollitia",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.SearchDocuments(ctx, operations.TigrisSearchRequest{
        SearchRequest: shared.SearchRequest{
            Branch: tigrisgosdk.String("ad"),
            Collation: &shared.Collation{
                Case: tigrisgosdk.String("eum"),
            },
            ExcludeFields: []string{
                "dolor",
            },
            Facet: &shared.SearchRequestFacet{},
            Fields: &shared.SearchRequestFields{},
            Filter: &shared.SearchRequestFilter{},
            IncludeFields: []string{
                "necessitatibus",
            },
            Page: tigrisgosdk.Int(141264),
            PageSize: tigrisgosdk.Int(367562),
            Q: tigrisgosdk.String("quasi"),
            SearchFields: []string{
                "iure",
            },
            Sort: &shared.SearchRequestSort{},
        },
        Collection: "doloribus",
        Project: "debitis",
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
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Collection.UpdateDocuments(ctx, operations.TigrisUpdateRequest{
        UpdateRequest: shared.UpdateRequest{
            Branch: tigrisgosdk.String("eius"),
            Fields: &shared.UpdateRequestFields{},
            Filter: &shared.UpdateRequestFilter{},
            Options: &shared.UpdateRequestOptions{
                Collation: &shared.Collation{
                    Case: tigrisgosdk.String("maxime"),
                },
                Limit: tigrisgosdk.Int64(537023),
                WriteOptions: &shared.WriteOptions{},
            },
        },
        Collection: "facilis",
        Project: "in",
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

