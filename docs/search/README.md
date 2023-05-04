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
            Document: tigris.String("deleniti"),
            ID: tigris.String("c42e141a-ac36-46c8-9d6b-144290747477"),
            Index: tigris.String("rem"),
            Project: tigris.String("fuga"),
        },
        ID: "7bd466d2-8c10-4ab3-8dca-4251904e523c",
        Index: "esse",
        Project: "recusandae",
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
                "distinctio",
            },
            Index: tigris.String("quod"),
            Project: tigris.String("dignissimos"),
        },
        Index: "inventore",
        Project: "nihil",
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
                "accusamus",
                "aliquam",
                "odio",
            },
            Index: tigris.String("occaecati"),
            Project: tigris.String("commodi"),
        },
        Index: "sapiente",
        Project: "dolores",
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
            Name: tigris.String("Fernando Barton"),
            Project: tigris.String("quas"),
        },
        Name: "Eugene Leuschke",
        Project: "mollitia",
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
                Case: tigris.String("incidunt"),
            },
            ExcludeFields: []string{
                "explicabo",
                "minima",
                "nisi",
            },
            Facet: tigris.String("fugit"),
            Filter: tigris.String("sapiente"),
            IncludeFields: []string{
                "ratione",
            },
            Index: tigris.String("explicabo"),
            Page: tigris.Int(903984),
            PageSize: tigris.Int(578922),
            Project: tigris.String("atque"),
            Q: tigris.String("et"),
            SearchFields: []string{
                "eveniet",
                "accusamus",
            },
            Sort: tigris.String("veritatis"),
        },
        Index: "esse",
        Project: "quod",
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
            "vero",
            "aliquid",
            "quasi",
        },
        Index: "saepe",
        Project: "vel",
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
        Name: "Javier Price",
        Project: "distinctio",
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
        FilterBranch: tigris.String("eligendi"),
        FilterCollection: tigris.String("sit"),
        FilterType: tigris.String("culpa"),
        Project: "tempore",
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
            Filter: tigris.String("adipisci"),
            Index: tigris.String("cumque"),
            Project: tigris.String("consequuntur"),
        },
        Index: "consequatur",
        Project: "minus",
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
                "sapiente",
                "consectetur",
            },
            Index: tigris.String("esse"),
            Project: tigris.String("blanditiis"),
        },
        Index: "provident",
        Project: "a",
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
                "quas",
                "esse",
                "quasi",
                "a",
            },
            Index: tigris.String("error"),
            Project: tigris.String("sint"),
        },
        Index: "pariatur",
        Project: "possimus",
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
            Name: tigris.String("Laverne Zemlak Sr."),
            OnlyCreate: tigris.Bool(false),
            Project: tigris.String("quasi"),
            Schema: tigris.String("similique"),
        },
        Name: "Dr. Gene Wiegand",
        Project: "in",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOrUpdateIndexResponse != nil {
        // handle response
    }
}
```
