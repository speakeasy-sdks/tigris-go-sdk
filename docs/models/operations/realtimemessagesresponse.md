# RealtimeMessagesResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ContentType`                                                       | *string*                                                            | :heavy_check_mark:                                                  | N/A                                                                 |
| `MessagesResponse`                                                  | [*shared.MessagesResponse](../../models/shared/messagesresponse.md) | :heavy_minus_sign:                                                  | OK                                                                  |
| `Status`                                                            | [*shared.Status](../../models/shared/status.md)                     | :heavy_minus_sign:                                                  | Default error response                                              |
| `StatusCode`                                                        | *int*                                                               | :heavy_check_mark:                                                  | N/A                                                                 |
| `RawResponse`                                                       | [*http.Response](https://pkg.go.dev/net/http#Response)              | :heavy_minus_sign:                                                  | N/A                                                                 |