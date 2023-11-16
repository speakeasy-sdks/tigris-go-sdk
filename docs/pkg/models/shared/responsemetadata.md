# ResponseMetadata

Has metadata related to the documents stored.


## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `CreatedAt`                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                       | :heavy_minus_sign:                                                                               | Time at which the document was inserted/replaced. Measured in nano-seconds since the Unix epoch. |
| `DeletedAt`                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                       | :heavy_minus_sign:                                                                               | Time at which the document was deleted. Measured in nano-seconds since the Unix epoch.           |
| `UpdatedAt`                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                       | :heavy_minus_sign:                                                                               | Time at which the document was updated. Measured in nano-seconds since the Unix epoch.           |