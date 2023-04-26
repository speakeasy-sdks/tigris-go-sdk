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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.RealtimeGetRTChannelRequest{
        Channel: "eligendi",
        Project: "sint",
    }

    res, err := s.Channel.Get(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.RealtimeReadMessagesRequest{
        Channel: "aliquid",
        End: tigris.String("provident"),
        Event: tigris.String("necessitatibus"),
        Limit: tigris.Int64(572252),
        Project: "officia",
        SessionID: tigris.String("dolor"),
        SocketID: tigris.String("debitis"),
        Start: tigris.String("a"),
    }

    res, err := s.Channel.GetMessages(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.RealtimeGetRTChannelsRequest{
        Project: "dolorum",
    }

    res, err := s.Channel.List(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.RealtimeListSubscriptionsRequest{
        Channel: "in",
        Page: tigris.Int(449198),
        PageSize: tigris.Int(846409),
        Project: "maiores",
    }

    res, err := s.Channel.ListSubscriptions(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.RealtimeMessagesRequest{
        MessagesRequest: shared.MessagesRequest{
            Channel: tigris.String("rerum"),
            Messages: []shared.Message{
                shared.Message{
                    Data: tigris.String("magnam"),
                    ID: tigris.String("cd66ae39-5efb-49ba-88f3-a66997074ba4"),
                    Name: tigris.String("Laurie Mosciski"),
                    Sequence: tigris.String("vero"),
                },
            },
            Project: tigris.String("aspernatur"),
        },
        Channel: "architecto",
        Project: "magnam",
    }

    res, err := s.Channel.PushMessages(ctx, req)
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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.RealtimePresenceRequest{
        Channel: "et",
        Project: "excepturi",
    }

    res, err := s.Channel.RealtimePresence(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PresenceResponse != nil {
        // handle response
    }
}
```
