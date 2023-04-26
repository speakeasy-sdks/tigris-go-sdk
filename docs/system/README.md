# System

## Overview

The Observability section has APIs that provides full visibility into the health, metrics, and monitoring of the Server.

### Available Operations

* [GetHealth](#gethealth) - Health Check
* [GetServerInfo](#getserverinfo) - Information about the server
* [ObservabilityQuotaUsage](#observabilityquotausage) - Queries current namespace quota usage
* [QueryQuotaLimits](#queryquotalimits) - Queries current namespace quota limits
* [QueryTimeSeriesMetrics](#querytimeseriesmetrics) - Queries time series metrics

## GetHealth

This endpoint can be used to check the liveness of the server.

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
    res, err := s.System.GetHealth(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.HealthCheckResponse != nil {
        // handle response
    }
}
```

## GetServerInfo

Provides the information about the server. This information includes returning the server version, etc.

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
    res, err := s.System.GetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetInfoResponse != nil {
        // handle response
    }
}
```

## ObservabilityQuotaUsage

Returns current namespace quota limits

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
    req := map[string]interface{}{
        "quasi": "a",
        "error": "sint",
    }

    res, err := s.System.ObservabilityQuotaUsage(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.QuotaUsageResponse != nil {
        // handle response
    }
}
```

## QueryQuotaLimits

Returns current namespace quota limits

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
    req := map[string]interface{}{
        "possimus": "quia",
        "eveniet": "asperiores",
        "facere": "veritatis",
        "consequuntur": "quasi",
    }

    res, err := s.System.QueryQuotaLimits(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.QuotaLimitsResponse != nil {
        // handle response
    }
}
```

## QueryTimeSeriesMetrics

Queries time series metrics

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/tigris-go-sdk"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
)

func main() {
    s := tigris.New(
        tigris.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.QueryTimeSeriesMetricsRequest{
        AdditionalFunctions: []shared.AdditionalFunction{
            shared.AdditionalFunction{
                Rollup: &shared.RollupFunction{
                    Aggregator: shared.RollupFunctionAggregatorEnumRollupAggregatorMax.ToPointer(),
                    Interval: tigris.Int64(398434),
                },
            },
            shared.AdditionalFunction{
                Rollup: &shared.RollupFunction{
                    Aggregator: shared.RollupFunctionAggregatorEnumRollupAggregatorAvg.ToPointer(),
                    Interval: tigris.Int64(62713),
                },
            },
            shared.AdditionalFunction{
                Rollup: &shared.RollupFunction{
                    Aggregator: shared.RollupFunctionAggregatorEnumRollupAggregatorAvg.ToPointer(),
                    Interval: tigris.Int64(424032),
                },
            },
        },
        Branch: tigris.String("in"),
        Collection: tigris.String("eius"),
        Db: tigris.String("libero"),
        From: tigris.Int64(849039),
        Function: shared.QueryTimeSeriesMetricsRequestFunctionEnumNone.ToPointer(),
        MetricName: tigris.String("accusantium"),
        Quantile: tigris.Float32(3069.86),
        SpaceAggregatedBy: []string{
            "dicta",
            "ullam",
            "reprehenderit",
            "ullam",
        },
        SpaceAggregation: shared.QueryTimeSeriesMetricsRequestSpaceAggregationEnumMin.ToPointer(),
        TigrisOperation: shared.QueryTimeSeriesMetricsRequestTigrisOperationEnumAll.ToPointer(),
        To: tigris.Int64(531849),
    }

    res, err := s.System.QueryTimeSeriesMetrics(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.QueryTimeSeriesMetricsResponse != nil {
        // handle response
    }
}
```
