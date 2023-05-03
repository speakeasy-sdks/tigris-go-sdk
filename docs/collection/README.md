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
    res, err := s.Collection.Create(ctx, operations.TigrisCreateOrUpdateCollectionRequest{
        CreateOrUpdateCollectionRequest: shared.CreateOrUpdateCollectionRequest{
            Branch: tigris.String("ullam"),
            OnlyCreate: tigris.Bool(false),
            Options: map[string]interface{}{
                "quos": "sint",
                "accusantium": "mollitia",
                "reiciendis": "mollitia",
            },
            Schema: map[string]interface{}{
                "eum": "dolor",
                "necessitatibus": "odit",
            },
        },
        Collection: "nemo",
        Project: "quasi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOrUpdateCollectionResponse != nil {
        // handle response
    }
}
```

## DeleteDocuments

Delete a range of documents in the collection using the condition provided in the filter.

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
    res, err := s.Collection.DeleteDocuments(ctx, operations.TigrisDeleteRequest{
        DeleteRequest: shared.DeleteRequest{
            Branch: tigris.String("iure"),
            Filter: map[string]interface{}{
                "debitis": "eius",
                "maxime": "deleniti",
                "facilis": "in",
                "architecto": "architecto",
            },
            Options: &shared.DeleteRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("repudiandae"),
                },
                Limit: tigris.Int64(352312),
                WriteOptions: map[string]interface{}{
                    "nihil": "repellat",
                    "quibusdam": "sed",
                    "saepe": "pariatur",
                },
            },
        },
        Collection: "accusantium",
        Project: "consequuntur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteResponse != nil {
        // handle response
    }
}
```

## Describe

Returns the information related to the collection. This can be used to retrieve the schema or size of the collection.

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
    res, err := s.Collection.Describe(ctx, operations.TigrisDescribeCollectionRequest{
        DescribeCollectionRequest: shared.DescribeCollectionRequest{
            Branch: tigris.String("praesentium"),
            Collection: tigris.String("natus"),
            Options: map[string]interface{}{
                "sunt": "quo",
            },
            Project: tigris.String("illum"),
            SchemaFormat: tigris.String("pariatur"),
        },
        Collection: "maxime",
        Project: "ea",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DescribeCollectionResponse != nil {
        // handle response
    }
}
```

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
    res, err := s.Collection.Drop(ctx, operations.TigrisDropCollectionRequest{
        DropCollectionRequest: shared.DropCollectionRequest{
            Branch: tigris.String("excepturi"),
            Options: map[string]interface{}{
                "ea": "accusantium",
            },
        },
        Collection: "ab",
        Project: "maiores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DropCollectionResponse != nil {
        // handle response
    }
}
```

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
    res, err := s.Collection.ImportDocuments(ctx, operations.TigrisImportRequest{
        ImportRequest: shared.ImportRequest{
            Autogenerated: []string{
                "ipsam",
                "voluptate",
                "autem",
            },
            Branch: tigris.String("nam"),
            CreateCollection: tigris.Bool(false),
            Documents: []map[string]interface{}{
                map[string]interface{}{
                    "nemo": "voluptatibus",
                    "perferendis": "fugiat",
                    "amet": "aut",
                    "cumque": "corporis",
                },
            },
            Options: &shared.ImportRequestOptions{
                WriteOptions: map[string]interface{}{
                    "libero": "nobis",
                    "dolores": "quis",
                    "totam": "dignissimos",
                    "eaque": "quis",
                },
            },
            PrimaryKey: []string{
                "eos",
            },
        },
        Collection: "perferendis",
        Project: "dolores",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ImportResponse != nil {
        // handle response
    }
}
```

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
    res, err := s.Collection.InsertDocuments(ctx, operations.TigrisInsertRequest{
        InsertRequest: shared.InsertRequest{
            Branch: tigris.String("minus"),
            Documents: []map[string]interface{}{
                map[string]interface{}{
                    "vero": "nostrum",
                },
                map[string]interface{}{
                    "recusandae": "omnis",
                    "facilis": "perspiciatis",
                    "voluptatem": "porro",
                    "consequuntur": "blanditiis",
                },
            },
            Options: &shared.InsertRequestOptions{
                WriteOptions: map[string]interface{}{
                    "eaque": "occaecati",
                    "rerum": "adipisci",
                    "asperiores": "earum",
                },
            },
        },
        Collection: "modi",
        Project: "iste",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InsertResponse != nil {
        // handle response
    }
}
```

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
    res, err := s.Collection.ReadDocuments(ctx, operations.TigrisReadRequest{
        ReadRequest: shared.ReadRequest{
            Branch: tigris.String("dolorum"),
            Fields: map[string]interface{}{
                "pariatur": "provident",
                "nobis": "libero",
                "delectus": "quaerat",
            },
            Filter: map[string]interface{}{
                "aliquid": "dolorem",
                "dolorem": "dolor",
                "qui": "ipsum",
            },
            Options: &shared.ReadRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("hic"),
                },
                Limit: tigris.Int64(569574),
                Offset: tigris.String("cum"),
                Skip: tigris.Int64(452109),
            },
            Sort: tigris.String("dignissimos"),
        },
        Collection: "reiciendis",
        Project: "amet",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StreamingReadResponse != nil {
        // handle response
    }
}
```

## ReplaceDocuments

Inserts the documents or replaces the existing documents in the collections.

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
    res, err := s.Collection.ReplaceDocuments(ctx, operations.TigrisReplaceRequest{
        ReplaceRequest: shared.ReplaceRequest{
            Branch: tigris.String("dolorum"),
            Documents: []map[string]interface{}{
                map[string]interface{}{
                    "ipsa": "ipsa",
                },
                map[string]interface{}{
                    "odio": "quaerat",
                    "accusamus": "quidem",
                },
            },
            Options: &shared.ReplaceRequestOptions{
                WriteOptions: map[string]interface{}{
                    "voluptas": "natus",
                    "eos": "atque",
                    "sit": "fugiat",
                    "ab": "soluta",
                },
            },
        },
        Collection: "dolorum",
        Project: "iusto",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReplaceResponse != nil {
        // handle response
    }
}
```

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
    res, err := s.Collection.SearchDocuments(ctx, operations.TigrisSearchRequest{
        SearchRequest: shared.SearchRequest{
            Branch: tigris.String("voluptate"),
            Collation: &shared.Collation{
                Case: tigris.String("dolorum"),
            },
            ExcludeFields: []string{
                "omnis",
                "necessitatibus",
                "distinctio",
            },
            Facet: map[string]interface{}{
                "nihil": "ipsum",
                "voluptate": "id",
                "saepe": "eius",
                "aspernatur": "perferendis",
            },
            Fields: map[string]interface{}{
                "optio": "accusamus",
            },
            Filter: map[string]interface{}{
                "saepe": "suscipit",
                "deserunt": "provident",
            },
            IncludeFields: []string{
                "repellendus",
                "totam",
            },
            Page: tigris.Int(628982),
            PageSize: tigris.Int(55),
            Q: tigris.String("at"),
            SearchFields: []string{
                "tempora",
                "vel",
            },
            Sort: map[string]interface{}{
                "officiis": "qui",
                "dolorum": "a",
                "esse": "harum",
                "iusto": "ipsum",
            },
        },
        Collection: "quisquam",
        Project: "tenetur",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StreamingSearchResponse != nil {
        // handle response
    }
}
```

## UpdateDocuments

Update a range of documents in the collection using the condition provided in the filter.

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
    res, err := s.Collection.UpdateDocuments(ctx, operations.TigrisUpdateRequest{
        UpdateRequest: shared.UpdateRequest{
            Branch: tigris.String("amet"),
            Fields: map[string]interface{}{
                "accusamus": "numquam",
                "enim": "dolorem",
                "sapiente": "totam",
            },
            Filter: map[string]interface{}{
                "sit": "expedita",
                "neque": "sed",
            },
            Options: &shared.UpdateRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("vel"),
                },
                Limit: tigris.Int64(730442),
                WriteOptions: map[string]interface{}{
                    "deserunt": "quam",
                    "ipsum": "incidunt",
                },
            },
        },
        Collection: "qui",
        Project: "cupiditate",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateResponse != nil {
        // handle response
    }
}
```
