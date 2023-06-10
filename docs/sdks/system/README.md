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
            BearerAuth: "",
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.HealthAPIHealthResponse](../../models/operations/healthapihealthresponse.md), error**


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
            BearerAuth: "",
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ObservabilityGetInfoResponse](../../models/operations/observabilitygetinforesponse.md), error**


## ObservabilityQuotaUsage

Returns current namespace quota limits

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.System.ObservabilityQuotaUsage(ctx, shared.QuotaUsageRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.QuotaUsageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.QuotaUsageRequest](../../models/shared/quotausagerequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.ObservabilityQuotaUsageResponse](../../models/operations/observabilityquotausageresponse.md), error**


## QueryQuotaLimits

Returns current namespace quota limits

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.System.QueryQuotaLimits(ctx, shared.QuotaLimitsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.QuotaLimitsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.QuotaLimitsRequest](../../models/shared/quotalimitsrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.ObservabilityQuotaLimitsResponse](../../models/operations/observabilityquotalimitsresponse.md), error**


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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.System.QueryTimeSeriesMetrics(ctx, shared.QueryTimeSeriesMetricsRequest{
        AdditionalFunctions: []shared.AdditionalFunction{
            shared.AdditionalFunction{
                Rollup: &shared.RollupFunction{
                    Aggregator: shared.RollupFunctionAggregatorRollupAggregatorSum.ToPointer(),
                    Interval: tigris.Int64(63955),
                },
            },
            shared.AdditionalFunction{
                Rollup: &shared.RollupFunction{
                    Aggregator: shared.RollupFunctionAggregatorRollupAggregatorMin.ToPointer(),
                    Interval: tigris.Int64(485628),
                },
            },
            shared.AdditionalFunction{
                Rollup: &shared.RollupFunction{
                    Aggregator: shared.RollupFunctionAggregatorRollupAggregatorMin.ToPointer(),
                    Interval: tigris.Int64(977496),
                },
            },
        },
        Branch: tigris.String("quisquam"),
        Collection: tigris.String("vero"),
        Db: tigris.String("omnis"),
        From: tigris.Int64(338159),
        Function: shared.QueryTimeSeriesMetricsRequestFunctionRate.ToPointer(),
        MetricName: tigris.String("delectus"),
        Quantile: tigris.Float32(4551.69),
        SpaceAggregatedBy: []string{
            "vero",
        },
        SpaceAggregation: shared.QueryTimeSeriesMetricsRequestSpaceAggregationSum.ToPointer(),
        TigrisOperation: shared.QueryTimeSeriesMetricsRequestTigrisOperationRead.ToPointer(),
        To: tigris.Int64(941378),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.QueryTimeSeriesMetricsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.QueryTimeSeriesMetricsRequest](../../models/shared/querytimeseriesmetricsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ObservabilityQueryTimeSeriesMetricsResponse](../../models/operations/observabilityquerytimeseriesmetricsresponse.md), error**

