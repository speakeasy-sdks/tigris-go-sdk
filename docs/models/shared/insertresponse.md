# InsertResponse

OK


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Keys`                                                       | []*string*                                                   | :heavy_minus_sign:                                           | an array returns the value of the primary keys.              |
| `Metadata`                                                   | [*ResponseMetadata](../../models/shared/responsemetadata.md) | :heavy_minus_sign:                                           | Has metadata related to the documents stored.                |
| `Status`                                                     | **string*                                                    | :heavy_minus_sign:                                           | An enum with value set as "inserted"                         |