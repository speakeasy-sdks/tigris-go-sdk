# tigris-core

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/tigris-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
    res, err := s.AppKey.Delete(ctx, operations.TigrisDeleteAppKeyRequest{
        DeleteAppKeyRequest: shared.DeleteAppKeyRequest{
            ID: tigris.String("89bd9d8d-69a6-474e-8f46-7cc8796ed151"),
        },
        Project: "deserunt",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteAppKeyResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AppKey](docs/appkey/README.md)

* [Delete](docs/appkey/README.md#delete) - Deletes the app key
* [List](docs/appkey/README.md#list) - List all the app keys
* [Rotate](docs/appkey/README.md#rotate) - Rotates the app key secret
* [TigrisCreateAppKey](docs/appkey/README.md#tigriscreateappkey) - Creates the app key
* [Update](docs/appkey/README.md#update) - Updates the description of the app key

### [Auth](docs/auth/README.md)

* [Get](docs/auth/README.md#get) - Access Token

### [Cache](docs/cache/README.md)

* [Create](docs/cache/README.md#create) - Creates the cache
* [Delete](docs/cache/README.md#delete) - Deletes the cache
* [DeleteKeys](docs/cache/README.md#deletekeys) - Deletes an entry from cache
* [GetKey](docs/cache/README.md#getkey) - Reads an entry from cache
* [GetSetKey](docs/cache/README.md#getsetkey) - Sets an entry in the cache and returns the previous value if exists
* [List](docs/cache/README.md#list) - Lists all the caches for the given project
* [ListKeys](docs/cache/README.md#listkeys) - Lists all the key for this cache
* [SetKey](docs/cache/README.md#setkey) - Sets an entry in the cache

### [Channel](docs/channel/README.md)

* [Get](docs/channel/README.md#get) - Get the details about a channel
* [GetMessages](docs/channel/README.md#getmessages) - Get all messages for a channel
* [List](docs/channel/README.md#list) - Get all channels for your application project
* [ListSubscriptions](docs/channel/README.md#listsubscriptions) - Get the subscriptions details about a channel
* [PushMessages](docs/channel/README.md#pushmessages) - push messages to a single channel
* [RealtimePresence](docs/channel/README.md#realtimepresence) - Presence about the channel

### [Collection](docs/collection/README.md)

* [Create](docs/collection/README.md#create) - Create or update a collection
* [DeleteDocuments](docs/collection/README.md#deletedocuments) - Delete Documents
* [Describe](docs/collection/README.md#describe) - Describe Collection
* [Drop](docs/collection/README.md#drop) - Drop Collection
* [ImportDocuments](docs/collection/README.md#importdocuments) - Import Documents
* [InsertDocuments](docs/collection/README.md#insertdocuments) - Insert Documents
* [ReadDocuments](docs/collection/README.md#readdocuments) - Read Documents
* [ReplaceDocuments](docs/collection/README.md#replacedocuments) - Insert or Replace Documents
* [SearchDocuments](docs/collection/README.md#searchdocuments) - Search Documents.
* [UpdateDocuments](docs/collection/README.md#updatedocuments) - Update Documents.

### [Database](docs/database/README.md)

* [BeginTransaction](docs/database/README.md#begintransaction) - Begin a transaction
* [CommitTransaction](docs/database/README.md#committransaction) - Commit a Transaction
* [CreateBranch](docs/database/README.md#createbranch) - Create a database branch
* [DeleteBranch](docs/database/README.md#deletebranch) - Delete a database branch
* [Describe](docs/database/README.md#describe) - Describe database
* [ListCollections](docs/database/README.md#listcollections) - List Collections
* [RollbackTransaction](docs/database/README.md#rollbacktransaction) - Rollback a transaction
* [TigrisListBranches](docs/database/README.md#tigrislistbranches) - List database branches

### [Namespace](docs/namespace/README.md)

* [Create](docs/namespace/README.md#create) - Creates a Namespace
* [Get](docs/namespace/README.md#get) - Describe the details of all namespaces
* [GetMetadata](docs/namespace/README.md#getmetadata) - Reads the Namespace Metadata
* [InsertMetadata](docs/namespace/README.md#insertmetadata) - Inserts Namespace Metadata
* [List](docs/namespace/README.md#list) - Lists all Namespaces
* [UpdateMetadata](docs/namespace/README.md#updatemetadata) - Updates Namespace Metadata

### [Project](docs/project/README.md)

* [Create](docs/project/README.md#create) - Create Project
* [DeleteProject](docs/project/README.md#deleteproject) - Delete Project and all resources under project
* [List](docs/project/README.md#list) - List Projects

### [Search](docs/search/README.md)

* [CreateDocument](docs/search/README.md#createdocument) - Create a single document
* [CreateDocuments](docs/search/README.md#createdocuments) - Create multiple documents
* [DeleteDocuments](docs/search/README.md#deletedocuments) - Delete documents by ids
* [DeleteIndex](docs/search/README.md#deleteindex) - Deletes search index
* [FindDocuments](docs/search/README.md#finddocuments) - Search Documents.
* [GetDocuments](docs/search/README.md#getdocuments) - Get a single or multiple documents
* [GetIndex](docs/search/README.md#getindex) - Get information about a search index
* [ListIndexes](docs/search/README.md#listindexes) - List search indexes
* [QueryDeleteDocuments](docs/search/README.md#querydeletedocuments) - Delete documents by query
* [ReplaceDocuments](docs/search/README.md#replacedocuments) - Create or replace documents in an index
* [UpdateDocuments](docs/search/README.md#updatedocuments) - Update documents in an index
* [UpdateIndex](docs/search/README.md#updateindex) - Creates or updates search index

### [System](docs/system/README.md)

* [GetHealth](docs/system/README.md#gethealth) - Health Check
* [GetServerInfo](docs/system/README.md#getserverinfo) - Information about the server
* [ObservabilityQuotaUsage](docs/system/README.md#observabilityquotausage) - Queries current namespace quota usage
* [QueryQuotaLimits](docs/system/README.md#queryquotalimits) - Queries current namespace quota limits
* [QueryTimeSeriesMetrics](docs/system/README.md#querytimeseriesmetrics) - Queries time series metrics

### [User](docs/user/README.md)

* [GetMetadata](docs/user/README.md#getmetadata) - Reads the User Metadata
* [InsertMetadata](docs/user/README.md#insertmetadata) - Inserts User Metadata
* [UpdateMetadata](docs/user/README.md#updatemetadata) - Updates User Metadata
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
