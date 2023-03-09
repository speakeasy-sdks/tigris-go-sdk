package shared

// QueryTimeSeriesMetricsResponse
// QueryTimeSeriesMetric responds with this type.
type QueryTimeSeriesMetricsResponse struct {
	From   *int64         `json:"from,omitempty"`
	Query  *string        `json:"query,omitempty"`
	Series []MetricSeries `json:"series,omitempty"`
	To     *int64         `json:"to,omitempty"`
}
