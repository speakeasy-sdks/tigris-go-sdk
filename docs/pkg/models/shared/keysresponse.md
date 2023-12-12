# KeysResponse


## Fields

| Field                                                                                                                                                                               | Type                                                                                                                                                                                | Required                                                                                                                                                                            | Description                                                                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Cursor`                                                                                                                                                                            | **int64*                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                  | cursor - 0 is the keys scan is finished, non-zero cursor can be passed in next keys request to continue the scan this is useful if streaming breaks and user wants to resume stream |
| `Keys`                                                                                                                                                                              | []*string*                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                  | keys                                                                                                                                                                                |