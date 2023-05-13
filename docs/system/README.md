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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.System.ObservabilityQuotaUsage(ctx, map[string]interface{}{
        "libero": "illum",
        "soluta": "accusantium",
    })
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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.System.QueryQuotaLimits(ctx, map[string]interface{}{
        "sapiente": "dicta",
        "ullam": "reprehenderit",
    })
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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.System.QueryTimeSeriesMetrics(ctx, shared.QueryTimeSeriesMetricsRequest{
        AdditionalFunctions: []shared.AdditionalFunction{
            shared.AdditionalFunction{
                Rollup: &shared.RollupFunction{
                    Aggregator: shared.RollupFunctionAggregatorEnumRollupAggregatorCount.ToPointer(),
                    Interval: tigris.Int64(16328),
                },
            },
            shared.AdditionalFunction{
                Rollup: &shared.RollupFunction{
                    Aggregator: shared.RollupFunctionAggregatorEnumRollupAggregatorMin.ToPointer(),
                    Interval: tigris.Int64(185232),
                },
            },
        },
        Branch: tigris.String("quibusdam"),
        Collection: tigris.String("ex"),
        Db: tigris.String("deleniti"),
        From: tigris.Int64(929292),
        Function: shared.QueryTimeSeriesMetricsRequestFunctionEnumNone.ToPointer(),
        MetricName: tigris.String("architecto"),
        Quantile: tigris.Float32(6091.78),
        SpaceAggregatedBy: []string{
            "quasi",
            "at",
            "et",
            "voluptate",
        },
        SpaceAggregation: shared.QueryTimeSeriesMetricsRequestSpaceAggregationEnumAvg.ToPointer(),
        TigrisOperation: shared.QueryTimeSeriesMetricsRequestTigrisOperationEnumRead.ToPointer(),
        To: tigris.Int64(86532),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.QueryTimeSeriesMetricsResponse != nil {
        // handle response
    }
}
```
