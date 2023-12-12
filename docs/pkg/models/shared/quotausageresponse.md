# QuotaUsageResponse

Contains current quota usage


## Fields

| Field                                                                                                          | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ReadUnits`                                                                                                    | **int64*                                                                                                       | :heavy_minus_sign:                                                                                             | Number of read units used per second                                                                           |
| `ReadUnitsThrottled`                                                                                           | **int64*                                                                                                       | :heavy_minus_sign:                                                                                             | Number of read units throttled per second. Units which was rejected with "resource exhausted error".           |
| `StorageSize`                                                                                                  | **int64*                                                                                                       | :heavy_minus_sign:                                                                                             | Number of bytes stored                                                                                         |
| `StorageSizeThrottled`                                                                                         | **int64*                                                                                                       | :heavy_minus_sign:                                                                                             | Number of bytes throttled. Number of bytes which were attempted to write in excess of quota and were rejected. |
| `WriteUnits`                                                                                                   | **int64*                                                                                                       | :heavy_minus_sign:                                                                                             | Number of write units used per second                                                                          |
| `WriteUnitsThrottled`                                                                                          | **int64*                                                                                                       | :heavy_minus_sign:                                                                                             | Number of write units throttled per second. Units which was rejected with "resource exhausted error".          |