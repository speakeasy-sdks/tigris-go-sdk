package shared

// DataPoint
// Represents the data point in timeseries.
type DataPoint struct {
	Timestamp *int64   `json:"timestamp,omitempty"`
	Value     *float64 `json:"value,omitempty"`
}
