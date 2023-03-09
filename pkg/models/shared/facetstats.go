package shared

// FacetStats
// Additional stats for faceted field
type FacetStats struct {
	Avg   *float64 `json:"avg,omitempty"`
	Count *int64   `json:"count,omitempty"`
	Max   *float64 `json:"max,omitempty"`
	Min   *float64 `json:"min,omitempty"`
	Sum   *float64 `json:"sum,omitempty"`
}
