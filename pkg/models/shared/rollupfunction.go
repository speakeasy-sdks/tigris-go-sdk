// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Aggregator string

const (
	AggregatorRollupAggregatorSum   Aggregator = "ROLLUP_AGGREGATOR_SUM"
	AggregatorRollupAggregatorCount Aggregator = "ROLLUP_AGGREGATOR_COUNT"
	AggregatorRollupAggregatorMin   Aggregator = "ROLLUP_AGGREGATOR_MIN"
	AggregatorRollupAggregatorMax   Aggregator = "ROLLUP_AGGREGATOR_MAX"
	AggregatorRollupAggregatorAvg   Aggregator = "ROLLUP_AGGREGATOR_AVG"
)

func (e Aggregator) ToPointer() *Aggregator {
	return &e
}

func (e *Aggregator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ROLLUP_AGGREGATOR_SUM":
		fallthrough
	case "ROLLUP_AGGREGATOR_COUNT":
		fallthrough
	case "ROLLUP_AGGREGATOR_MIN":
		fallthrough
	case "ROLLUP_AGGREGATOR_MAX":
		fallthrough
	case "ROLLUP_AGGREGATOR_AVG":
		*e = Aggregator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Aggregator: %v", v)
	}
}

// RollupFunction - Rollup function aggregates the slices of metrics returned by original query and lets you operate on the slices using aggregator and constructs the bigger slice of your choice of interval (specified in seconds).
type RollupFunction struct {
	Aggregator *Aggregator `json:"aggregator,omitempty"`
	Interval   *int64      `json:"interval,omitempty"`
}

func (o *RollupFunction) GetAggregator() *Aggregator {
	if o == nil {
		return nil
	}
	return o.Aggregator
}

func (o *RollupFunction) GetInterval() *int64 {
	if o == nil {
		return nil
	}
	return o.Interval
}
