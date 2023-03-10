package shared

type RollupFunctionAggregatorEnum string

const (
	RollupFunctionAggregatorEnumRollupAggregatorSum   RollupFunctionAggregatorEnum = "ROLLUP_AGGREGATOR_SUM"
	RollupFunctionAggregatorEnumRollupAggregatorCount RollupFunctionAggregatorEnum = "ROLLUP_AGGREGATOR_COUNT"
	RollupFunctionAggregatorEnumRollupAggregatorMin   RollupFunctionAggregatorEnum = "ROLLUP_AGGREGATOR_MIN"
	RollupFunctionAggregatorEnumRollupAggregatorMax   RollupFunctionAggregatorEnum = "ROLLUP_AGGREGATOR_MAX"
	RollupFunctionAggregatorEnumRollupAggregatorAvg   RollupFunctionAggregatorEnum = "ROLLUP_AGGREGATOR_AVG"
)

// RollupFunction
// Rollup function aggregates the slices of metrics returned by original query and lets you operate on the slices using aggregator and constructs the bigger slice of your choice of interval (specified in seconds).
type RollupFunction struct {
	Aggregator *RollupFunctionAggregatorEnum `json:"aggregator,omitempty"`
	Interval   *int64                        `json:"interval,omitempty"`
}
