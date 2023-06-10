# Channel

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
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/operations"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.Get(ctx, operations.RealtimeGetRTChannelRequest{
        Channel: "quasi",
        Project: "repudiandae",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RealtimeGetRTChannelRequest](../../models/operations/realtimegetrtchannelrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.RealtimeGetRTChannelResponse](../../models/operations/realtimegetrtchannelresponse.md), error**


## GetMessages

Get all messages for a channel

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.GetMessages(ctx, operations.RealtimeReadMessagesRequest{
        Channel: "sint",
        End: tigris.String("veritatis"),
        Event: tigris.String("itaque"),
        Limit: tigris.Int64(277718),
        Project: "enim",
        SessionID: tigris.String("consequatur"),
        SocketID: tigris.String("est"),
        Start: tigris.String("quibusdam"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.RealtimeReadMessagesRequest](../../models/operations/realtimereadmessagesrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.RealtimeReadMessagesResponse](../../models/operations/realtimereadmessagesresponse.md), error**


## List

Get all channels for your application project

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.List(ctx, operations.RealtimeGetRTChannelsRequest{
        Project: "explicabo",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.RealtimeGetRTChannelsRequest](../../models/operations/realtimegetrtchannelsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.RealtimeGetRTChannelsResponse](../../models/operations/realtimegetrtchannelsresponse.md), error**


## ListSubscriptions

Get the subscriptions details about a channel

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.ListSubscriptions(ctx, operations.RealtimeListSubscriptionsRequest{
        Channel: "deserunt",
        Page: tigris.Int(716327),
        PageSize: tigris.Int(841386),
        Project: "labore",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.RealtimeListSubscriptionsRequest](../../models/operations/realtimelistsubscriptionsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.RealtimeListSubscriptionsResponse](../../models/operations/realtimelistsubscriptionsresponse.md), error**


## PushMessages

push messages to a single channel

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.PushMessages(ctx, operations.RealtimeMessagesRequest{
        MessagesRequest: shared.MessagesRequest{
            Channel: tigris.String("modi"),
            Messages: []shared.Message{
                shared.Message{
                    Data: tigris.String("aliquid"),
                    ID: tigris.String("9802d502-a94b-4b4f-a3c9-69e9a3efa77d"),
                    Name: tigris.String("Jean Buckridge"),
                    Sequence: tigris.String("facere"),
                },
            },
            Project: tigris.String("ea"),
        },
        Channel: "aliquid",
        Project: "laborum",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RealtimeMessagesRequest](../../models/operations/realtimemessagesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RealtimeMessagesResponse](../../models/operations/realtimemessagesresponse.md), error**


## RealtimePresence

Presence about the channel

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.RealtimePresence(ctx, operations.RealtimePresenceRequest{
        Channel: "accusamus",
        Project: "non",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.RealtimePresenceRequest](../../models/operations/realtimepresencerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.RealtimePresenceResponse](../../models/operations/realtimepresenceresponse.md), error**

