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
            Branch: tigris.String("eius"),
            OnlyCreate: tigris.Bool(false),
            Options: map[string]interface{}{
                "deleniti": "facilis",
                "in": "architecto",
                "architecto": "repudiandae",
                "ullam": "expedita",
            },
            Schema: map[string]interface{}{
                "repellat": "quibusdam",
                "sed": "saepe",
            },
        },
        Collection: "pariatur",
        Project: "accusantium",
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
            Branch: tigris.String("consequuntur"),
            Filter: map[string]interface{}{
                "natus": "magni",
                "sunt": "quo",
                "illum": "pariatur",
            },
            Options: &shared.DeleteRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("maxime"),
                },
                Limit: tigris.Int64(411397),
                WriteOptions: map[string]interface{}{
                    "odit": "ea",
                    "accusantium": "ab",
                    "maiores": "quidem",
                },
            },
        },
        Collection: "ipsam",
        Project: "voluptate",
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
            Branch: tigris.String("autem"),
            Collection: tigris.String("nam"),
            Options: map[string]interface{}{
                "pariatur": "nemo",
            },
            Project: tigris.String("voluptatibus"),
            SchemaFormat: tigris.String("perferendis"),
        },
        Collection: "fugiat",
        Project: "amet",
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
            Branch: tigris.String("aut"),
            Options: map[string]interface{}{
                "corporis": "hic",
                "libero": "nobis",
                "dolores": "quis",
                "totam": "dignissimos",
            },
        },
        Collection: "eaque",
        Project: "quis",
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
                "eos",
            },
            Branch: tigris.String("perferendis"),
            CreateCollection: tigris.Bool(false),
            Documents: []map[string]interface{}{
                map[string]interface{}{
                    "quam": "dolor",
                    "vero": "nostrum",
                    "hic": "recusandae",
                    "omnis": "facilis",
                },
            },
            Options: &shared.ImportRequestOptions{
                WriteOptions: map[string]interface{}{
                    "voluptatem": "porro",
                    "consequuntur": "blanditiis",
                    "error": "eaque",
                },
            },
            PrimaryKey: []string{
                "rerum",
                "adipisci",
                "asperiores",
            },
        },
        Collection: "earum",
        Project: "modi",
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
            Branch: tigris.String("iste"),
            Documents: []map[string]interface{}{
                map[string]interface{}{
                    "pariatur": "provident",
                    "nobis": "libero",
                    "delectus": "quaerat",
                },
                map[string]interface{}{
                    "aliquid": "dolorem",
                    "dolorem": "dolor",
                    "qui": "ipsum",
                },
                map[string]interface{}{
                    "excepturi": "cum",
                    "voluptate": "dignissimos",
                    "reiciendis": "amet",
                    "dolorum": "numquam",
                },
            },
            Options: &shared.InsertRequestOptions{
                WriteOptions: map[string]interface{}{
                    "ipsa": "ipsa",
                },
            },
        },
        Collection: "iure",
        Project: "odio",
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
            Branch: tigris.String("quaerat"),
            Fields: map[string]interface{}{
                "quidem": "voluptatibus",
                "voluptas": "natus",
                "eos": "atque",
                "sit": "fugiat",
            },
            Filter: map[string]interface{}{
                "soluta": "dolorum",
            },
            Options: &shared.ReadRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("iusto"),
                },
                Limit: tigris.Int64(453697),
                Offset: tigris.String("dolorum"),
                Skip: tigris.Int64(536579),
            },
            Sort: tigris.String("omnis"),
        },
        Collection: "necessitatibus",
        Project: "distinctio",
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
            Branch: tigris.String("asperiores"),
            Documents: []map[string]interface{}{
                map[string]interface{}{
                    "voluptate": "id",
                },
                map[string]interface{}{
                    "eius": "aspernatur",
                    "perferendis": "amet",
                    "optio": "accusamus",
                    "ad": "saepe",
                },
            },
            Options: &shared.ReplaceRequestOptions{
                WriteOptions: map[string]interface{}{
                    "deserunt": "provident",
                    "minima": "repellendus",
                },
            },
        },
        Collection: "totam",
        Project: "similique",
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
            Branch: tigris.String("alias"),
            Collation: &shared.Collation{
                Case: tigris.String("at"),
            },
            ExcludeFields: []string{
                "tempora",
                "vel",
            },
            Facet: map[string]interface{}{
                "officiis": "qui",
                "dolorum": "a",
                "esse": "harum",
                "iusto": "ipsum",
            },
            Fields: map[string]interface{}{
                "tenetur": "amet",
                "tempore": "accusamus",
                "numquam": "enim",
                "dolorem": "sapiente",
            },
            Filter: map[string]interface{}{
                "nihil": "sit",
                "expedita": "neque",
                "sed": "vel",
            },
            IncludeFields: []string{
                "voluptas",
                "deserunt",
                "quam",
            },
            Page: tigris.Int(214880),
            PageSize: tigris.Int(277628),
            Q: tigris.String("qui"),
            SearchFields: []string{
                "maxime",
                "pariatur",
                "soluta",
            },
            Sort: map[string]interface{}{
                "laborum": "totam",
            },
        },
        Collection: "incidunt",
        Project: "aspernatur",
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
            Branch: tigris.String("dolores"),
            Fields: map[string]interface{}{
                "facilis": "aliquid",
                "quam": "molestias",
                "temporibus": "qui",
            },
            Filter: map[string]interface{}{
                "fugit": "magni",
            },
            Options: &shared.UpdateRequestOptions{
                Collation: &shared.Collation{
                    Case: tigris.String("odio"),
                },
                Limit: tigris.Int64(124833),
                WriteOptions: map[string]interface{}{
                    "nam": "hic",
                    "voluptatem": "cumque",
                },
            },
        },
        Collection: "soluta",
        Project: "nobis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateResponse != nil {
        // handle response
    }
}
```
