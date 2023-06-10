# InsertRequest


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Branch`                                                                  | **string*                                                                 | :heavy_minus_sign:                                                        | Optionally specify a database branch name to perform operation on         |
| `Documents`                                                               | [][InsertRequestDocuments](../../models/shared/insertrequestdocuments.md) | :heavy_minus_sign:                                                        | Array of documents to insert. Each document is a JSON object.             |
| `Options`                                                                 | [*InsertRequestOptions](../../models/shared/insertrequestoptions.md)      | :heavy_minus_sign:                                                        | additional options for insert requests.                                   |