# UpdateResponse


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Metadata`                                                                 | [*shared.ResponseMetadata](../../../pkg/models/shared/responsemetadata.md) | :heavy_minus_sign:                                                         | Has metadata related to the documents stored.                              |
| `ModifiedCount`                                                            | **int*                                                                     | :heavy_minus_sign:                                                         | Returns the number of documents modified.                                  |
| `Status`                                                                   | **string*                                                                  | :heavy_minus_sign:                                                         | an enum with value set as "updated".                                       |