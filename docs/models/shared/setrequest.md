# SetRequest


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Ex`                                               | **int64*                                           | :heavy_minus_sign:                                 | optional - ttl specific to this key in second      |
| `Nx`                                               | **bool*                                            | :heavy_minus_sign:                                 | set only if the key doesn't exist                  |
| `Px`                                               | **int64*                                           | :heavy_minus_sign:                                 | optional - ttl specific to this key in millisecond |
| `Value`                                            | **string*                                          | :heavy_minus_sign:                                 | free form byte[] value                             |
| `Xx`                                               | **bool*                                            | :heavy_minus_sign:                                 | set only if the key exist                          |