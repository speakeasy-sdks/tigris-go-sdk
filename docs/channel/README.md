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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.Get(ctx, operations.RealtimeGetRTChannelRequest{
        Channel: "modi",
        Project: "qui",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRTChannelResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.GetMessages(ctx, operations.RealtimeReadMessagesRequest{
        Channel: "aliquid",
        End: tigris.String("cupiditate"),
        Event: tigris.String("quos"),
        Limit: tigris.Int64(20107),
        Project: "magni",
        SessionID: tigris.String("assumenda"),
        SocketID: tigris.String("ipsam"),
        Start: tigris.String("alias"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ReadMessagesResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.List(ctx, operations.RealtimeGetRTChannelsRequest{
        Project: "fugit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRTChannelsResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.ListSubscriptions(ctx, operations.RealtimeListSubscriptionsRequest{
        Channel: "dolorum",
        Page: tigris.Int(569618),
        PageSize: tigris.Int(270008),
        Project: "facilis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSubscriptionResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.PushMessages(ctx, operations.RealtimeMessagesRequest{
        MessagesRequest: shared.MessagesRequest{
            Channel: tigris.String("tempore"),
            Messages: []shared.Message{
                shared.Message{
                    Data: tigris.String("delectus"),
                    ID: tigris.String("63c969e9-a3ef-4a77-9fb1-4cd66ae395ef"),
                    Name: tigris.String("Rene Reinger"),
                    Sequence: tigris.String("deleniti"),
                },
                shared.Message{
                    Data: tigris.String("sapiente"),
                    ID: tigris.String("3a669970-74ba-4446-9b6e-2141959890af"),
                    Name: tigris.String("Tommy Kemmer"),
                    Sequence: tigris.String("odit"),
                },
            },
            Project: tigris.String("nemo"),
        },
        Channel: "quasi",
        Project: "iure",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.MessagesResponse != nil {
        // handle response
    }
}
```

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Channel.RealtimePresence(ctx, operations.RealtimePresenceRequest{
        Channel: "doloribus",
        Project: "debitis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PresenceResponse != nil {
        // handle response
    }
}
```
