# Search
(*Search*)

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
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.CreateDocument(ctx, operations.SearchCreateByIDRequest{
        CreateByIDRequest: shared.CreateByIDRequest{},
        ID: "<ID>",
        Index: "string",
        Project: "string",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.SearchCreateByIDRequest](../../pkg/models/operations/searchcreatebyidrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.SearchCreateByIDResponse](../../pkg/models/operations/searchcreatebyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.CreateDocuments(ctx, operations.SearchCreateRequest{
        CreateDocumentRequest: shared.CreateDocumentRequest{
            Documents: []string{
                "string",
            },
        },
        Index: "string",
        Project: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SearchCreateRequest](../../pkg/models/operations/searchcreaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SearchCreateResponse](../../pkg/models/operations/searchcreateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteDocuments

Delete one or more documents by id. Returns an array of status indicating the status of each document. Each status
 has an error field that is set to null in case document is deleted successfully otherwise it will non null with
 an error code and message.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.DeleteDocuments(ctx, operations.SearchDeleteRequest{
        DeleteDocumentRequest: shared.DeleteDocumentRequest{
            Ids: []string{
                "string",
            },
        },
        Index: "string",
        Project: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SearchDeleteRequest](../../pkg/models/operations/searchdeleterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SearchDeleteResponse](../../pkg/models/operations/searchdeleteresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteIndex

Deletes search index

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.DeleteIndex(ctx, operations.SearchDeleteIndexRequest{
        DeleteIndexRequest: shared.DeleteIndexRequest{},
        Name: "string",
        Project: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.SearchDeleteIndexRequest](../../pkg/models/operations/searchdeleteindexrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.SearchDeleteIndexResponse](../../pkg/models/operations/searchdeleteindexresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.FindDocuments(ctx, operations.SearchSearchRequest{
        SearchIndexRequest: shared.SearchIndexRequest{
            Collation: &shared.Collation{},
            ExcludeFields: []string{
                "string",
            },
            IncludeFields: []string{
                "string",
            },
            SearchFields: []string{
                "string",
            },
        },
        Index: "string",
        Project: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SearchSearchRequest](../../pkg/models/operations/searchsearchrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SearchSearchResponse](../../pkg/models/operations/searchsearchresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetDocuments

Retrieves one or more documents by id. The response is an array of documents in the same order it is requests.
 A null is returned for the documents that are not found.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.GetDocuments(ctx, operations.SearchGetRequest{
        Ids: []string{
            "string",
        },
        Index: "string",
        Project: "string",
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.SearchGetRequest](../../pkg/models/operations/searchgetrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.SearchGetResponse](../../pkg/models/operations/searchgetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetIndex

Get information about a search index

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.GetIndex(ctx, operations.SearchGetIndexRequest{
        Name: "string",
        Project: "string",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.SearchGetIndexRequest](../../pkg/models/operations/searchgetindexrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.SearchGetIndexResponse](../../pkg/models/operations/searchgetindexresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListIndexes

List search indexes

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.ListIndexes(ctx, operations.SearchListIndexesRequest{
        Project: "string",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.SearchListIndexesRequest](../../pkg/models/operations/searchlistindexesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.SearchListIndexesResponse](../../pkg/models/operations/searchlistindexesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## QueryDeleteDocuments

DeleteByQuery is used to delete documents that match the filter. A filter is required. To delete document by id,
 you can pass the filter as follows ```{"id": "test"}```. Returns a count of number of documents deleted.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.QueryDeleteDocuments(ctx, operations.SearchDeleteByQueryRequest{
        DeleteByQueryRequest: shared.DeleteByQueryRequest{},
        Index: "string",
        Project: "string",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.SearchDeleteByQueryRequest](../../pkg/models/operations/searchdeletebyqueryrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.SearchDeleteByQueryResponse](../../pkg/models/operations/searchdeletebyqueryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ReplaceDocuments

Creates or replaces one or more documents. Each document is a JSON object. A document is replaced
 if it already exists. An "id" is generated automatically in case it is missing in the document. The
 document is created if "id" doesn't exists otherwise it is replaced. Returns an array of status indicating
 the status of each document.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.ReplaceDocuments(ctx, operations.SearchCreateOrReplaceRequest{
        CreateOrReplaceDocumentRequest: shared.CreateOrReplaceDocumentRequest{
            Documents: []string{
                "string",
            },
        },
        Index: "string",
        Project: "string",
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.SearchCreateOrReplaceRequest](../../pkg/models/operations/searchcreateorreplacerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.SearchCreateOrReplaceResponse](../../pkg/models/operations/searchcreateorreplaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateDocuments

Updates one or more documents by "id". Each document is required to have the
 "id" field in it. Returns an array of status indicating the status of each document. Each status
 has an error field that is set to null in case document is updated successfully otherwise the error
 field is set with a code and message.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.UpdateDocuments(ctx, operations.SearchUpdateRequest{
        UpdateDocumentRequest: shared.UpdateDocumentRequest{
            Documents: []string{
                "string",
            },
        },
        Index: "string",
        Project: "string",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SearchUpdateRequest](../../pkg/models/operations/searchupdaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SearchUpdateResponse](../../pkg/models/operations/searchupdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateIndex

Creates or updates search index

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Search.UpdateIndex(ctx, operations.SearchCreateOrUpdateIndexRequest{
        CreateOrUpdateIndexRequest: shared.CreateOrUpdateIndexRequest{},
        Name: "string",
        Project: "string",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.SearchCreateOrUpdateIndexRequest](../../pkg/models/operations/searchcreateorupdateindexrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.SearchCreateOrUpdateIndexResponse](../../pkg/models/operations/searchcreateorupdateindexresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
