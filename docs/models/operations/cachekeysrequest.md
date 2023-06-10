# CacheKeysRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `Count`                                                           | **int64*                                                          | :heavy_minus_sign:                                                | optional - count of keys to return a stream response line.        |
| `Cursor`                                                          | **int64*                                                          | :heavy_minus_sign:                                                | optional - cursor - skip this argument if no cursor is associated |
| `Name`                                                            | *string*                                                          | :heavy_check_mark:                                                | cache name                                                        |
| `Pattern`                                                         | **string*                                                         | :heavy_minus_sign:                                                | optional key pattern                                              |
| `Project`                                                         | *string*                                                          | :heavy_check_mark:                                                | Tigris project name                                               |