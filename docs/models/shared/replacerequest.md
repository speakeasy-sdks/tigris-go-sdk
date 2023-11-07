# ReplaceRequest


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Branch`                                                                           | **string*                                                                          | :heavy_minus_sign:                                                                 | Optionally specify a database branch name to perform operation on                  |
| `Documents`                                                                        | [][shared.ReplaceRequestDocuments](../../models/shared/replacerequestdocuments.md) | :heavy_minus_sign:                                                                 | Array of documents to be replaced. Each document is a JSON object.                 |
| `Options`                                                                          | [*shared.ReplaceRequestOptions](../../models/shared/replacerequestoptions.md)      | :heavy_minus_sign:                                                                 | Additional options for replace requests.                                           |