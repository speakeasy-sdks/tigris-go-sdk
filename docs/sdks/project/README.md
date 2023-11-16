# Project
(*Project*)

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
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Project.Create(ctx, operations.TigrisCreateProjectRequest{
        CreateProjectRequest: shared.CreateProjectRequest{},
        Project: "string",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.TigrisCreateProjectRequest](../../pkg/models/operations/tigriscreateprojectrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.TigrisCreateProjectResponse](../../pkg/models/operations/tigriscreateprojectresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## DeleteProject

Delete Project deletes all the collections in this project along with all of the documents, and associated metadata for these collections.

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
        tigrisgosdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Project.DeleteProject(ctx, operations.TigrisDeleteProjectRequest{
        DeleteProjectRequest: shared.DeleteProjectRequest{},
        Project: "string",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.TigrisDeleteProjectRequest](../../pkg/models/operations/tigrisdeleteprojectrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.TigrisDeleteProjectResponse](../../pkg/models/operations/tigrisdeleteprojectresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## List

List returns all the projects.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"context"
	"log"
)

func main() {
    s := tigrisgosdk.New(
        tigrisgosdk.WithSecurity(""),
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

**[*operations.TigrisListProjectsResponse](../../pkg/models/operations/tigrislistprojectsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
