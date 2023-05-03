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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.CreateDocument(ctx, operations.SearchCreateByIDRequest{
        CreateByIDRequest: shared.CreateByIDRequest{
            Document: tigris.String("ex"),
            ID: tigris.String("d9f5fce6-c556-4146-83e2-50fb008c42e1"),
            Index: tigris.String("non"),
            Project: tigris.String("et"),
        },
        ID: "aac366c8-dd6b-4144-a907-474778a7bd46",
        Index: "suscipit",
        Project: "assumenda",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateByIDResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.CreateDocuments(ctx, operations.SearchCreateRequest{
        CreateDocumentRequest: shared.CreateDocumentRequest{
            Documents: []string{
                "praesentium",
            },
            Index: tigris.String("quisquam"),
            Project: tigris.String("veritatis"),
        },
        Index: "ipsa",
        Project: "id",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDocumentResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.DeleteDocuments(ctx, operations.SearchDeleteRequest{
        DeleteDocumentRequest: shared.DeleteDocumentRequest{
            Ids: []string{
                "neque",
                "quo",
                "illum",
            },
            Index: tigris.String("quo"),
            Project: tigris.String("fuga"),
        },
        Index: "eius",
        Project: "eos",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteDocumentResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.DeleteIndex(ctx, operations.SearchDeleteIndexRequest{
        DeleteIndexRequest: shared.DeleteIndexRequest{
            Name: tigris.String("Mrs. Virginia McGlynn"),
            Project: tigris.String("ipsam"),
        },
        Name: "Emily Satterfield",
        Project: "aperiam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteIndexResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.FindDocuments(ctx, operations.SearchSearchRequest{
        SearchIndexRequest: shared.SearchIndexRequest{
            Collation: &shared.Collation{
                Case: tigris.String("distinctio"),
            },
            ExcludeFields: []string{
                "dignissimos",
                "inventore",
                "nihil",
                "totam",
            },
            Facet: tigris.String("accusamus"),
            Filter: tigris.String("aliquam"),
            IncludeFields: []string{
                "occaecati",
                "commodi",
            },
            Index: tigris.String("sapiente"),
            Page: tigris.Int(174112),
            PageSize: tigris.Int(645570),
            Project: tigris.String("molestiae"),
            Q: tigris.String("accusantium"),
            SearchFields: []string{
                "eum",
                "quas",
                "praesentium",
                "consequuntur",
            },
            Sort: tigris.String("deleniti"),
        },
        Index: "fugit",
        Project: "fuga",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.SearchIndexResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.GetDocuments(ctx, operations.SearchGetRequest{
        Ids: []string{
            "incidunt",
            "atque",
            "explicabo",
        },
        Index: "minima",
        Project: "nisi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDocumentResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.GetIndex(ctx, operations.SearchGetIndexRequest{
        Name: "Jeannie Cronin",
        Project: "saepe",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetIndexResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.ListIndexes(ctx, operations.SearchListIndexesRequest{
        FilterBranch: tigris.String("occaecati"),
        FilterCollection: tigris.String("atque"),
        FilterType: tigris.String("et"),
        Project: "esse",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListIndexesResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.QueryDeleteDocuments(ctx, operations.SearchDeleteByQueryRequest{
        DeleteByQueryRequest: shared.DeleteByQueryRequest{
            Filter: tigris.String("eveniet"),
            Index: tigris.String("accusamus"),
            Project: tigris.String("veritatis"),
        },
        Index: "esse",
        Project: "quod",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteByQueryResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.ReplaceDocuments(ctx, operations.SearchCreateOrReplaceRequest{
        CreateOrReplaceDocumentRequest: shared.CreateOrReplaceDocumentRequest{
            Documents: []string{
                "vero",
                "aliquid",
                "quasi",
            },
            Index: tigris.String("saepe"),
            Project: tigris.String("vel"),
        },
        Index: "harum",
        Project: "molestiae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOrReplaceDocumentResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.UpdateDocuments(ctx, operations.SearchUpdateRequest{
        UpdateDocumentRequest: shared.UpdateDocumentRequest{
            Documents: []string{
                "occaecati",
                "minima",
                "distinctio",
            },
            Index: tigris.String("eligendi"),
            Project: tigris.String("sit"),
        },
        Index: "culpa",
        Project: "tempore",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDocumentResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Search.UpdateIndex(ctx, operations.SearchCreateOrUpdateIndexRequest{
        CreateOrUpdateIndexRequest: shared.CreateOrUpdateIndexRequest{
            Name: tigris.String("Miss Blanca Cronin"),
            OnlyCreate: tigris.Bool(false),
            Project: tigris.String("sapiente"),
            Schema: tigris.String("consectetur"),
        },
        Name: "Mattie McLaughlin",
        Project: "quas",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOrUpdateIndexResponse != nil {
        // handle response
    }
}
```
