package shared

// MetricSeries
// Represents series in timeseries based on input query.
type MetricSeries struct {
	DataPoints []DataPoint `json:"dataPoints,omitempty"`
	From       *int64      `json:"from,omitempty"`
	Metric     *string     `json:"metric,omitempty"`
	Scope      *string     `json:"scope,omitempty"`
	To         *int64      `json:"to,omitempty"`
}
