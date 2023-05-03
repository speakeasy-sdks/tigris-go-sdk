# Project

## Overview

Every Tigris projects comes with a transactional document database built on FoundationDB, one of the most resilient and battle-tested open source distributed key-value store. A database is created automatically for you when you create a project.

### Available Operations

* [Create](#create) - Create Project
* [DeleteProject](#deleteproject) - Delete Project and all resources under project
* [List](#list) - List Projects

## Create

Creates a new project. Returns an AlreadyExists error with a status code 409 if the project already exists.

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
    res, err := s.Project.Create(ctx, operations.TigrisCreateProjectRequest{
        RequestBody: map[string]interface{}{
            "aspernatur": "vel",
            "possimus": "magnam",
        },
        Project: "ratione",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateProjectResponse != nil {
        // handle response
    }
}
```

## DeleteProject

Delete Project deletes all the collections in this project along with all of the documents, and associated metadata for these collections.

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
    res, err := s.Project.DeleteProject(ctx, operations.TigrisDeleteProjectRequest{
        RequestBody: map[string]interface{}{
            "laudantium": "dicta",
            "dolor": "maiores",
        },
        Project: "quasi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteProjectResponse != nil {
        // handle response
    }
}
```

## List

List returns all the projects.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Project.List(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListProjectsResponse != nil {
        // handle response
    }
}
```
