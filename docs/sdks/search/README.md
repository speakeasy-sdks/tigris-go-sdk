# Search

## Overview

The search section provides you APIs that can be used to implement powerful apps with search experiences. You can manage storing documents on your own or you can simply integrate it with your database.

### Available Operations

* [CreateDocument](#createdocument) - Create a single document
* [CreateDocuments](#createdocuments) - Create multiple documents
* [DeleteDocuments](#deletedocuments) - Delete documents by ids
* [DeleteIndex](#deleteindex) - Deletes search index
* [FindDocuments](#finddocuments) - Search Documents.
* [GetDocuments](#getdocuments) - Get a single or multiple documents
* [GetIndex](#getindex) - Get information about a search index
* [ListIndexes](#listindexes) - List search indexes
* [QueryDeleteDocuments](#querydeletedocuments) - Delete documents by query
* [ReplaceDocuments](#replacedocuments) - Create or replace documents in an index
* [UpdateDocuments](#updatedocuments) - Update documents in an index
* [UpdateIndex](#updateindex) - Creates or updates search index

## CreateDocument

CreateById is used for indexing a single document. The API expects a single document. An "id" is optional
 and the server can automatically generate it for you in case it is missing. In cases an id is provided in
 the document and the document already exists then that document will not be indexed and an error is returned
 with HTTP status code 409.

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
    res, err := s.Search.CreateDocument(ctx, operations.SearchCreateByIDRequest{
        CreateByIDRequest: shared.CreateByIDRequest{
            Document: tigris.String("occaecati"),
            ID: tigris.String("b3fe49a8-d9cb-4f48-a333-23f9b77f3a41"),
            Index: tigris.String("ipsa"),
            Project: tigris.String("ipsa"),
        },
        ID: "674ebf69-280d-41ba-b7a8-9ebf737ae420",
        Index: "amet",
        Project: "optio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateByIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.SearchCreateByIDRequest](../../models/operations/searchcreatebyidrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.SearchCreateByIDResponse](../../models/operations/searchcreatebyidresponse.md), error**


## CreateDocuments

Create is used for indexing a single or multiple documents. The API expects an array of documents.
 Each document is a JSON object. An "id" is optional and the server can automatically generate it for you in
 case it is missing. In cases when an id is provided in the document and the document already exists then that
 document will not be indexed and in the response there will be an error corresponding to that document id other
 documents will succeed. Returns an array of status indicating the status of each document.

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
    res, err := s.Search.CreateDocuments(ctx, operations.SearchCreateRequest{
        CreateDocumentRequest: shared.CreateDocumentRequest{
            Documents: []string{
                "ad",
                "saepe",
                "suscipit",
                "deserunt",
            },
            Index: tigris.String("provident"),
            Project: tigris.String("minima"),
        },
        Index: "repellendus",
        Project: "totam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.SearchCreateRequest](../../models/operations/searchcreaterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.SearchCreateResponse](../../models/operations/searchcreateresponse.md), error**


## DeleteDocuments

Delete one or more documents by id. Returns an array of status indicating the status of each document. Each status
 has an error field that is set to null in case document is deleted successfully otherwise it will non null with
 an error code and message.

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
    res, err := s.Search.DeleteDocuments(ctx, operations.SearchDeleteRequest{
        DeleteDocumentRequest: shared.DeleteDocumentRequest{
            Ids: []string{
                "alias",
                "at",
                "quaerat",
            },
            Index: tigris.String("tempora"),
            Project: tigris.String("vel"),
        },
        Index: "quod",
        Project: "officiis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.SearchDeleteRequest](../../models/operations/searchdeleterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.SearchDeleteResponse](../../models/operations/searchdeleteresponse.md), error**


## DeleteIndex

Deletes search index

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
    res, err := s.Search.DeleteIndex(ctx, operations.SearchDeleteIndexRequest{
        DeleteIndexRequest: shared.DeleteIndexRequest{
            Name: tigris.String("Jan Wilderman"),
            Project: tigris.String("iusto"),
        },
        Name: "Rosalie White",
        Project: "accusamus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteIndexResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.SearchDeleteIndexRequest](../../models/operations/searchdeleteindexrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.SearchDeleteIndexResponse](../../models/operations/searchdeleteindexresponse.md), error**


## FindDocuments

Searches an index for the documents matching the query. A search can be a term search or a phrase search.
 Search API allows filtering the result set using filters as documented
 <a href="https://docs.tigrisdata.com/overview/query#specification-1" title="here">here</a>. You can also perform
 a faceted search by passing the fields in the facet parameter. You can find more detailed documentation of the
 Search API with multiple examples <a href="https://docs.tigrisdata.com/overview/search" title="here">here</a>.

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
    res, err := s.Search.FindDocuments(ctx, operations.SearchSearchRequest{
        SearchIndexRequest: shared.SearchIndexRequest{
            Collation: &shared.Collation{
                Case: tigris.String("numquam"),
            },
            ExcludeFields: []string{
                "dolorem",
                "sapiente",
            },
            Facet: tigris.String("totam"),
            Filter: tigris.String("nihil"),
            IncludeFields: []string{
                "expedita",
            },
            Index: tigris.String("neque"),
            Page: tigris.Int(153694),
            PageSize: tigris.Int(424685),
            Project: tigris.String("libero"),
            Q: tigris.String("voluptas"),
            SearchFields: []string{
                "quam",
                "ipsum",
                "incidunt",
            },
            Sort: tigris.String("qui"),
        },
        Index: "cupiditate",
        Project: "maxime",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchIndexResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.SearchSearchRequest](../../models/operations/searchsearchrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.SearchSearchResponse](../../models/operations/searchsearchresponse.md), error**


## GetDocuments

Retrieves one or more documents by id. The response is an array of documents in the same order it is requests.
 A null is returned for the documents that are not found.

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
    res, err := s.Search.GetDocuments(ctx, operations.SearchGetRequest{
        Ids: []string{
            "soluta",
            "dicta",
            "laborum",
            "totam",
        },
        Index: "incidunt",
        Project: "aspernatur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.SearchGetRequest](../../models/operations/searchgetrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.SearchGetResponse](../../models/operations/searchgetresponse.md), error**


## GetIndex

Get information about a search index

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
    res, err := s.Search.GetIndex(ctx, operations.SearchGetIndexRequest{
        Name: "Verna Purdy",
        Project: "molestias",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetIndexResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SearchGetIndexRequest](../../models/operations/searchgetindexrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SearchGetIndexResponse](../../models/operations/searchgetindexresponse.md), error**


## ListIndexes

List search indexes

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
    res, err := s.Search.ListIndexes(ctx, operations.SearchListIndexesRequest{
        FilterBranch: tigris.String("temporibus"),
        FilterCollection: tigris.String("qui"),
        FilterType: tigris.String("neque"),
        Project: "fugit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListIndexesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.SearchListIndexesRequest](../../models/operations/searchlistindexesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.SearchListIndexesResponse](../../models/operations/searchlistindexesresponse.md), error**


## QueryDeleteDocuments

DeleteByQuery is used to delete documents that match the filter. A filter is required. To delete document by id,
 you can pass the filter as follows ```{"id": "test"}```. Returns a count of number of documents deleted.

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
    res, err := s.Search.QueryDeleteDocuments(ctx, operations.SearchDeleteByQueryRequest{
        DeleteByQueryRequest: shared.DeleteByQueryRequest{
            Filter: tigris.String("magni"),
            Index: tigris.String("odio"),
            Project: tigris.String("sunt"),
        },
        Index: "ullam",
        Project: "nam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteByQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.SearchDeleteByQueryRequest](../../models/operations/searchdeletebyqueryrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.SearchDeleteByQueryResponse](../../models/operations/searchdeletebyqueryresponse.md), error**


## ReplaceDocuments

Creates or replaces one or more documents. Each document is a JSON object. A document is replaced
 if it already exists. An "id" is generated automatically in case it is missing in the document. The
 document is created if "id" doesn't exists otherwise it is replaced. Returns an array of status indicating
 the status of each document.

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
    res, err := s.Search.ReplaceDocuments(ctx, operations.SearchCreateOrReplaceRequest{
        CreateOrReplaceDocumentRequest: shared.CreateOrReplaceDocumentRequest{
            Documents: []string{
                "voluptatem",
                "cumque",
                "soluta",
                "nobis",
            },
            Index: tigris.String("et"),
            Project: tigris.String("saepe"),
        },
        Index: "ipsum",
        Project: "veritatis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOrReplaceDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.SearchCreateOrReplaceRequest](../../models/operations/searchcreateorreplacerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.SearchCreateOrReplaceResponse](../../models/operations/searchcreateorreplaceresponse.md), error**


## UpdateDocuments

Updates one or more documents by "id". Each document is required to have the
 "id" field in it. Returns an array of status indicating the status of each document. Each status
 has an error field that is set to null in case document is updated successfully otherwise the error
 field is set with a code and message.

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
    res, err := s.Search.UpdateDocuments(ctx, operations.SearchUpdateRequest{
        UpdateDocumentRequest: shared.UpdateDocumentRequest{
            Documents: []string{
                "quos",
                "tempore",
                "cupiditate",
            },
            Index: tigris.String("aperiam"),
            Project: tigris.String("delectus"),
        },
        Index: "dolorem",
        Project: "dolore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDocumentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.SearchUpdateRequest](../../models/operations/searchupdaterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.SearchUpdateResponse](../../models/operations/searchupdateresponse.md), error**


## UpdateIndex

Creates or updates search index

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
    res, err := s.Search.UpdateIndex(ctx, operations.SearchCreateOrUpdateIndexRequest{
        CreateOrUpdateIndexRequest: shared.CreateOrUpdateIndexRequest{
            Name: tigris.String("Mr. Josephine Pagac V"),
            OnlyCreate: tigris.Bool(false),
            Project: tigris.String("itaque"),
            Schema: tigris.String("consequatur"),
        },
        Name: "Marcos Schaden",
        Project: "facilis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOrUpdateIndexResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.SearchCreateOrUpdateIndexRequest](../../models/operations/searchcreateorupdateindexrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.SearchCreateOrUpdateIndexResponse](../../models/operations/searchcreateorupdateindexresponse.md), error**
