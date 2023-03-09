package shared

// AdditionalFunction
// Additional function to apply on metrics query
type AdditionalFunction struct {
	Rollup *RollupFunction `json:"rollup,omitempty"`
}
