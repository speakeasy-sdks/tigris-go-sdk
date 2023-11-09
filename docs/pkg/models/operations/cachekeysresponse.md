# CacheKeysResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ContentType`                                                      | *string*                                                           | :heavy_check_mark:                                                 | HTTP response content type for this operation                      |
| `KeysResponse`                                                     | [*shared.KeysResponse](../../../pkg/models/shared/keysresponse.md) | :heavy_minus_sign:                                                 | OK                                                                 |
| `Status`                                                           | [*shared.Status](../../../pkg/models/shared/status.md)             | :heavy_minus_sign:                                                 | Default error response                                             |
| `StatusCode`                                                       | *int*                                                              | :heavy_check_mark:                                                 | HTTP response status code for this operation                       |
| `RawResponse`                                                      | [*http.Response](https://pkg.go.dev/net/http#Response)             | :heavy_minus_sign:                                                 | Raw HTTP response; suitable for custom response parsing            |