# RollupFunction

Rollup function aggregates the slices of metrics returned by original query and lets you operate on the slices using aggregator and constructs the bigger slice of your choice of interval (specified in seconds).


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Aggregator`                                                                 | [*RollupFunctionAggregator](../../models/shared/rollupfunctionaggregator.md) | :heavy_minus_sign:                                                           | N/A                                                                          |
| `Interval`                                                                   | **int64*                                                                     | :heavy_minus_sign:                                                           | N/A                                                                          |