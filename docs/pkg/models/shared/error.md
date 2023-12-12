# Error

The Error type defines a logical error model


## Fields

| Field                                                                                                                                                   | Type                                                                                                                                                    | Required                                                                                                                                                | Description                                                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Code`                                                                                                                                                  | [*shared.Code](../../../pkg/models/shared/code.md)                                                                                                      | :heavy_minus_sign:                                                                                                                                      | The status code is a short, machine parsable string, which uniquely identifies the error type. Tigris to HTTP code mapping [here](/reference/http-code) |
| `Message`                                                                                                                                               | **string*                                                                                                                                               | :heavy_minus_sign:                                                                                                                                      | A developer-facing descriptive error message                                                                                                            |