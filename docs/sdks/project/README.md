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
    res, err := s.Project.Create(ctx, operations.TigrisCreateProjectRequest{
        CreateProjectRequest: shared.CreateProjectRequest{},
        Project: "nostrum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.TigrisCreateProjectRequest](../../models/operations/tigriscreateprojectrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.TigrisCreateProjectResponse](../../models/operations/tigriscreateprojectresponse.md), error**


## DeleteProject

Delete Project deletes all the collections in this project along with all of the documents, and associated metadata for these collections.

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
    res, err := s.Project.DeleteProject(ctx, operations.TigrisDeleteProjectRequest{
        DeleteProjectRequest: shared.DeleteProjectRequest{},
        Project: "hic",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.TigrisDeleteProjectRequest](../../models/operations/tigrisdeleteprojectrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.TigrisDeleteProjectResponse](../../models/operations/tigrisdeleteprojectresponse.md), error**


## List

List returns all the projects.

### Example Usage

```go
package main

import(
	"context"
	"log"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.TigrisListProjectsResponse](../../models/operations/tigrislistprojectsresponse.md), error**

