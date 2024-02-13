# tigris-core

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/tigris-go-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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
		tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Cache.Create(ctx, operations.CacheCreateCacheRequest{
		CreateCacheRequest: shared.CreateCacheRequest{},
		Name:               "string",
		Project:            "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.CreateCacheResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Auth](docs/sdks/auth/README.md)

* [Get](docs/sdks/auth/README.md#get) - Access Token

### [System](docs/sdks/system/README.md)

* [GetHealth](docs/sdks/system/README.md#gethealth) - Health Check
* [GetServerInfo](docs/sdks/system/README.md#getserverinfo) - Information about the server
* [ObservabilityQuotaUsage](docs/sdks/system/README.md#observabilityquotausage) - Queries current namespace quota usage
* [QueryQuotaLimits](docs/sdks/system/README.md#queryquotalimits) - Queries current namespace quota limits
* [QueryTimeSeriesMetrics](docs/sdks/system/README.md#querytimeseriesmetrics) - Queries time series metrics

### [Namespace](docs/sdks/namespace/README.md)

* [Create](docs/sdks/namespace/README.md#create) - Creates a Namespace
* [Get](docs/sdks/namespace/README.md#get) - Describe the details of all namespaces
* [GetMetadata](docs/sdks/namespace/README.md#getmetadata) - Reads the Namespace Metadata
* [InsertMetadata](docs/sdks/namespace/README.md#insertmetadata) - Inserts Namespace Metadata
* [List](docs/sdks/namespace/README.md#list) - Lists all Namespaces
* [UpdateMetadata](docs/sdks/namespace/README.md#updatemetadata) - Updates Namespace Metadata

### [User](docs/sdks/user/README.md)

* [GetMetadata](docs/sdks/user/README.md#getmetadata) - Reads the User Metadata
* [InsertMetadata](docs/sdks/user/README.md#insertmetadata) - Inserts User Metadata
* [UpdateMetadata](docs/sdks/user/README.md#updatemetadata) - Updates User Metadata

### [Project](docs/sdks/project/README.md)

* [Create](docs/sdks/project/README.md#create) - Create Project
* [DeleteProject](docs/sdks/project/README.md#deleteproject) - Delete Project and all resources under project
* [List](docs/sdks/project/README.md#list) - List Projects

### [AppKey](docs/sdks/appkey/README.md)

* [Delete](docs/sdks/appkey/README.md#delete) - Deletes the app key
* [List](docs/sdks/appkey/README.md#list) - List all the app keys
* [Rotate](docs/sdks/appkey/README.md#rotate) - Rotates the app key secret
* [TigrisCreateAppKey](docs/sdks/appkey/README.md#tigriscreateappkey) - Creates the app key
* [Update](docs/sdks/appkey/README.md#update) - Updates the description of the app key

### [Cache](docs/sdks/cache/README.md)

* [Create](docs/sdks/cache/README.md#create) - Creates the cache
* [Delete](docs/sdks/cache/README.md#delete) - Deletes the cache
* [DeleteKeys](docs/sdks/cache/README.md#deletekeys) - Deletes an entry from cache
* [GetKey](docs/sdks/cache/README.md#getkey) - Reads an entry from cache
* [GetSetKey](docs/sdks/cache/README.md#getsetkey) - Sets an entry in the cache and returns the previous value if exists
* [List](docs/sdks/cache/README.md#list) - Lists all the caches for the given project
* [ListKeys](docs/sdks/cache/README.md#listkeys) - Lists all the key for this cache
* [SetKey](docs/sdks/cache/README.md#setkey) - Sets an entry in the cache

### [Database](docs/sdks/database/README.md)

* [BeginTransaction](docs/sdks/database/README.md#begintransaction) - Begin a transaction
* [CommitTransaction](docs/sdks/database/README.md#committransaction) - Commit a Transaction
* [CreateBranch](docs/sdks/database/README.md#createbranch) - Create a database branch
* [DeleteBranch](docs/sdks/database/README.md#deletebranch) - Delete a database branch
* [Describe](docs/sdks/database/README.md#describe) - Describe database
* [ListCollections](docs/sdks/database/README.md#listcollections) - List Collections
* [RollbackTransaction](docs/sdks/database/README.md#rollbacktransaction) - Rollback a transaction
* [TigrisListBranches](docs/sdks/database/README.md#tigrislistbranches) - List database branches

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

### [Channel](docs/sdks/channel/README.md)

* [Get](docs/sdks/channel/README.md#get) - Get the details about a channel
* [GetMessages](docs/sdks/channel/README.md#getmessages) - Get all messages for a channel
* [List](docs/sdks/channel/README.md#list) - Get all channels for your application project
* [ListSubscriptions](docs/sdks/channel/README.md#listsubscriptions) - Get the subscriptions details about a channel
* [PushMessages](docs/sdks/channel/README.md#pushmessages) - push messages to a single channel
* [RealtimePresence](docs/sdks/channel/README.md#realtimepresence) - Presence about the channel

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
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := tigrisgosdk.New(
		tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.Get(ctx)
	if err != nil {

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.preview.tigrisdata.cloud` | None |
| 1 | `http://localhost:8081` | None |

#### Example

```go
package main

import (
	"context"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := tigrisgosdk.New(
		tigrisgosdk.WithServerIndex(1),
		tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.GetAccessTokenResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := tigrisgosdk.New(
		tigrisgosdk.WithServerURL("https://api.preview.tigrisdata.cloud"),
		tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.GetAccessTokenResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BearerAuth` | http         | HTTP Bearer  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	tigrisgosdk "github.com/speakeasy-sdks/tigris-go-sdk"
	"log"
)

func main() {
	s := tigrisgosdk.New(
		tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.GetAccessTokenResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
