# Channel
(*Channel*)

## Overview

The realtime section provide APIs that can be used realtime operations.

### Available Operations

* [Get](#get) - Get the details about a channel
* [GetMessages](#getmessages) - Get all messages for a channel
* [List](#list) - Get all channels for your application project
* [ListSubscriptions](#listsubscriptions) - Get the subscriptions details about a channel
* [PushMessages](#pushmessages) - push messages to a single channel
* [RealtimePresence](#realtimepresence) - Presence about the channel

## Get

Get the details about a channel

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
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Channel.Get(ctx, operations.RealtimeGetRTChannelRequest{
        Channel: "string",
        Project: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRTChannelResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RealtimeGetRTChannelRequest](../../pkg/models/operations/realtimegetrtchannelrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.RealtimeGetRTChannelResponse](../../pkg/models/operations/realtimegetrtchannelresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetMessages

Get all messages for a channel

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
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Channel.GetMessages(ctx, operations.RealtimeReadMessagesRequest{
        Channel: "string",
        Project: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReadMessagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.RealtimeReadMessagesRequest](../../pkg/models/operations/realtimereadmessagesrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.RealtimeReadMessagesResponse](../../pkg/models/operations/realtimereadmessagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## List

Get all channels for your application project

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
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Channel.List(ctx, operations.RealtimeGetRTChannelsRequest{
        Project: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRTChannelsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.RealtimeGetRTChannelsRequest](../../pkg/models/operations/realtimegetrtchannelsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.RealtimeGetRTChannelsResponse](../../pkg/models/operations/realtimegetrtchannelsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListSubscriptions

Get the subscriptions details about a channel

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
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Channel.ListSubscriptions(ctx, operations.RealtimeListSubscriptionsRequest{
        Channel: "string",
        Project: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.RealtimeListSubscriptionsRequest](../../pkg/models/operations/realtimelistsubscriptionsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.RealtimeListSubscriptionsResponse](../../pkg/models/operations/realtimelistsubscriptionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PushMessages

push messages to a single channel

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
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Channel.PushMessages(ctx, operations.RealtimeMessagesRequest{
        MessagesRequest: shared.MessagesRequest{
            Messages: []shared.Message{
                shared.Message{},
            },
        },
        Channel: "string",
        Project: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessagesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RealtimeMessagesRequest](../../pkg/models/operations/realtimemessagesrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.RealtimeMessagesResponse](../../pkg/models/operations/realtimemessagesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## RealtimePresence

Presence about the channel

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
        tigrisgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Channel.RealtimePresence(ctx, operations.RealtimePresenceRequest{
        Channel: "string",
        Project: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PresenceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.RealtimePresenceRequest](../../pkg/models/operations/realtimepresencerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.RealtimePresenceResponse](../../pkg/models/operations/realtimepresenceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
