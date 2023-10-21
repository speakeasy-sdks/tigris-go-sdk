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

import (
	"context"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := tigrisgosdk.New(
		tigrisgosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Cache.Create(ctx, operations.CacheCreateCacheRequest{
		CreateCacheRequest: shared.CreateCacheRequest{
			Options: &shared.CreateCacheOptions{},
		},
		Name:    "string",
		Project: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateCacheResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AppKey](docs/sdks/appkey/README.md)

* [Delete](docs/sdks/appkey/README.md#delete) - Deletes the app key
* [List](docs/sdks/appkey/README.md#list) - List all the app keys
* [Rotate](docs/sdks/appkey/README.md#rotate) - Rotates the app key secret
* [TigrisCreateAppKey](docs/sdks/appkey/README.md#tigriscreateappkey) - Creates the app key
* [Update](docs/sdks/appkey/README.md#update) - Updates the description of the app key

### [Auth](docs/sdks/auth/README.md)

* [Get](docs/sdks/auth/README.md#get) - Access Token

### [Cache](docs/sdks/cache/README.md)

* [Create](docs/sdks/cache/README.md#create) - Creates the cache
* [Delete](docs/sdks/cache/README.md#delete) - Deletes the cache
* [DeleteKeys](docs/sdks/cache/README.md#deletekeys) - Deletes an entry from cache
* [GetKey](docs/sdks/cache/README.md#getkey) - Reads an entry from cache
* [GetSetKey](docs/sdks/cache/README.md#getsetkey) - Sets an entry in the cache and returns the previous value if exists
* [List](docs/sdks/cache/README.md#list) - Lists all the caches for the given project
* [ListKeys](docs/sdks/cache/README.md#listkeys) - Lists all the key for this cache
* [SetKey](docs/sdks/cache/README.md#setkey) - Sets an entry in the cache

### [Channel](docs/sdks/channel/README.md)

* [Get](docs/sdks/channel/README.md#get) - Get the details about a channel
* [GetMessages](docs/sdks/channel/README.md#getmessages) - Get all messages for a channel
* [List](docs/sdks/channel/README.md#list) - Get all channels for your application project
* [ListSubscriptions](docs/sdks/channel/README.md#listsubscriptions) - Get the subscriptions details about a channel
* [PushMessages](docs/sdks/channel/README.md#pushmessages) - push messages to a single channel
* [RealtimePresence](docs/sdks/channel/README.md#realtimepresence) - Presence about the channel

### [Collection](docs/sdks/collection/README.md)

* [Create](docs/sdks/collection/README.md#create) - Create or update a collection
* [DeleteDocuments](docs/sdks/collection/README.md#deletedocuments) - Delete Documents
* [Describe](docs/sdks/collection/README.md#describe) - Describe Collection
* [Drop](docs/sdks/collection/README.md#drop) - Drop Collection
* [ImportDocuments](docs/sdks/collection/README.md#importdocuments) - Import Documents
* [InsertDocuments](docs/sdks/collection/README.md#insertdocuments) - Insert Documents
* [ReadDocuments](docs/sdks/collection/README.md#readdocuments) - Read Documents
* [ReplaceDocuments](docs/sdks/collection/README.md#replacedocuments) - Insert or Replace Documents
* [SearchDocuments](docs/sdks/collection/README.md#searchdocuments) - Search Documents.
* [UpdateDocuments](docs/sdks/collection/README.md#updatedocuments) - Update Documents.

### [Database](docs/sdks/database/README.md)

* [BeginTransaction](docs/sdks/database/README.md#begintransaction) - Begin a transaction
* [CommitTransaction](docs/sdks/database/README.md#committransaction) - Commit a Transaction
* [CreateBranch](docs/sdks/database/README.md#createbranch) - Create a database branch
* [DeleteBranch](docs/sdks/database/README.md#deletebranch) - Delete a database branch
* [Describe](docs/sdks/database/README.md#describe) - Describe database
* [ListCollections](docs/sdks/database/README.md#listcollections) - List Collections
* [RollbackTransaction](docs/sdks/database/README.md#rollbacktransaction) - Rollback a transaction
* [TigrisListBranches](docs/sdks/database/README.md#tigrislistbranches) - List database branches

### [Namespace](docs/sdks/namespace/README.md)

* [Create](docs/sdks/namespace/README.md#create) - Creates a Namespace
* [Get](docs/sdks/namespace/README.md#get) - Describe the details of all namespaces
* [GetMetadata](docs/sdks/namespace/README.md#getmetadata) - Reads the Namespace Metadata
* [InsertMetadata](docs/sdks/namespace/README.md#insertmetadata) - Inserts Namespace Metadata
* [List](docs/sdks/namespace/README.md#list) - Lists all Namespaces
* [UpdateMetadata](docs/sdks/namespace/README.md#updatemetadata) - Updates Namespace Metadata

### [Project](docs/sdks/project/README.md)

* [Create](docs/sdks/project/README.md#create) - Create Project
* [DeleteProject](docs/sdks/project/README.md#deleteproject) - Delete Project and all resources under project
* [List](docs/sdks/project/README.md#list) - List Projects

### [Search](docs/sdks/search/README.md)

* [CreateDocument](docs/sdks/search/README.md#createdocument) - Create a single document
* [CreateDocuments](docs/sdks/search/README.md#createdocuments) - Create multiple documents
* [DeleteDocuments](docs/sdks/search/README.md#deletedocuments) - Delete documents by ids
* [DeleteIndex](docs/sdks/search/README.md#deleteindex) - Deletes search index
* [FindDocuments](docs/sdks/search/README.md#finddocuments) - Search Documents.
* [GetDocuments](docs/sdks/search/README.md#getdocuments) - Get a single or multiple documents
* [GetIndex](docs/sdks/search/README.md#getindex) - Get information about a search index
* [ListIndexes](docs/sdks/search/README.md#listindexes) - List search indexes
* [QueryDeleteDocuments](docs/sdks/search/README.md#querydeletedocuments) - Delete documents by query
* [ReplaceDocuments](docs/sdks/search/README.md#replacedocuments) - Create or replace documents in an index
* [UpdateDocuments](docs/sdks/search/README.md#updatedocuments) - Update documents in an index
* [UpdateIndex](docs/sdks/search/README.md#updateindex) - Creates or updates search index

### [System](docs/sdks/system/README.md)

* [GetHealth](docs/sdks/system/README.md#gethealth) - Health Check
* [GetServerInfo](docs/sdks/system/README.md#getserverinfo) - Information about the server
* [ObservabilityQuotaUsage](docs/sdks/system/README.md#observabilityquotausage) - Queries current namespace quota usage
* [QueryQuotaLimits](docs/sdks/system/README.md#queryquotalimits) - Queries current namespace quota limits
* [QueryTimeSeriesMetrics](docs/sdks/system/README.md#querytimeseriesmetrics) - Queries time series metrics

### [User](docs/sdks/user/README.md)

* [GetMetadata](docs/sdks/user/README.md#getmetadata) - Reads the User Metadata
* [InsertMetadata](docs/sdks/user/README.md#insertmetadata) - Inserts User Metadata
* [UpdateMetadata](docs/sdks/user/README.md#updatemetadata) - Updates User Metadata
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
