# QuotaLimitsResponse

Contains current quota limits


## Fields

| Field                                    | Type                                     | Required                                 | Description                              |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `ReadUnits`                              | **int64*                                 | :heavy_minus_sign:                       | Number of allowed read units per second  |
| `StorageSize`                            | **int64*                                 | :heavy_minus_sign:                       | Maximum number of bytes allowed to store |
| `WriteUnits`                             | **int64*                                 | :heavy_minus_sign:                       | Number of allowed write units per second |