# DocMeta

Contains metadata related to the search hit, has information about document created_at/updated_at as well.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `CreatedAt`                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                       | :heavy_minus_sign:                                                                               | Time at which the document was inserted/replaced. Measured in nano-seconds since the Unix epoch. |
| `UpdatedAt`                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                       | :heavy_minus_sign:                                                                               | Time at which the document was updated. Measured in nano-seconds since the Unix epoch.           |