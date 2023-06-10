# CreateDocumentRequest


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Documents`                                                                      | []*string*                                                                       | :heavy_minus_sign:                                                               | An array of documents to be created or replaced. Each document is a JSON object. |
| `Index`                                                                          | **string*                                                                        | :heavy_minus_sign:                                                               | index name where to create documents.                                            |
| `Project`                                                                        | **string*                                                                        | :heavy_minus_sign:                                                               | Tigris project name.                                                             |