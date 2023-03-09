package shared

type QueryTimeSeriesMetricsRequestFunctionEnum string

const (
	QueryTimeSeriesMetricsRequestFunctionEnumRate  QueryTimeSeriesMetricsRequestFunctionEnum = "RATE"
	QueryTimeSeriesMetricsRequestFunctionEnumCount QueryTimeSeriesMetricsRequestFunctionEnum = "COUNT"
	QueryTimeSeriesMetricsRequestFunctionEnumNone  QueryTimeSeriesMetricsRequestFunctionEnum = "NONE"
)

type QueryTimeSeriesMetricsRequestSpaceAggregationEnum string

const (
	QueryTimeSeriesMetricsRequestSpaceAggregationEnumAvg QueryTimeSeriesMetricsRequestSpaceAggregationEnum = "AVG"
	QueryTimeSeriesMetricsRequestSpaceAggregationEnumMin QueryTimeSeriesMetricsRequestSpaceAggregationEnum = "MIN"
	QueryTimeSeriesMetricsRequestSpaceAggregationEnumMax QueryTimeSeriesMetricsRequestSpaceAggregationEnum = "MAX"
	QueryTimeSeriesMetricsRequestSpaceAggregationEnumSum QueryTimeSeriesMetricsRequestSpaceAggregationEnum = "SUM"
)

type QueryTimeSeriesMetricsRequestTigrisOperationEnum string

const (
	QueryTimeSeriesMetricsRequestTigrisOperationEnumAll      QueryTimeSeriesMetricsRequestTigrisOperationEnum = "ALL"
	QueryTimeSeriesMetricsRequestTigrisOperationEnumRead     QueryTimeSeriesMetricsRequestTigrisOperationEnum = "READ"
	QueryTimeSeriesMetricsRequestTigrisOperationEnumWrite    QueryTimeSeriesMetricsRequestTigrisOperationEnum = "WRITE"
	QueryTimeSeriesMetricsRequestTigrisOperationEnumMetadata QueryTimeSeriesMetricsRequestTigrisOperationEnum = "METADATA"
)

// QueryTimeSeriesMetricsRequest
// Requests the time series metrics
type QueryTimeSeriesMetricsRequest struct {
	AdditionalFunctions []AdditionalFunction                               `json:"additionalFunctions,omitempty"`
	Branch              *string                                            `json:"branch,omitempty"`
	Collection          *string                                            `json:"collection,omitempty"`
	Db                  *string                                            `json:"db,omitempty"`
	From                *int64                                             `json:"from,omitempty"`
	Function            *QueryTimeSeriesMetricsRequestFunctionEnum         `json:"function,omitempty"`
	MetricName          *string                                            `json:"metric_name,omitempty"`
	Quantile            *float32                                           `json:"quantile,omitempty"`
	SpaceAggregatedBy   []string                                           `json:"space_aggregated_by,omitempty"`
	SpaceAggregation    *QueryTimeSeriesMetricsRequestSpaceAggregationEnum `json:"space_aggregation,omitempty"`
	TigrisOperation     *QueryTimeSeriesMetricsRequestTigrisOperationEnum  `json:"tigris_operation,omitempty"`
	To                  *int64                                             `json:"to,omitempty"`
}
