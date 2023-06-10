# SearchIndexResponse

Response struct for search


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Facets`                                                     | map[string][SearchFacet](../../models/shared/searchfacet.md) | :heavy_minus_sign:                                           | N/A                                                          |
| `Hits`                                                       | [][IndexDoc](../../models/shared/indexdoc.md)                | :heavy_minus_sign:                                           | N/A                                                          |
| `Meta`                                                       | [*SearchMetadata](../../models/shared/searchmetadata.md)     | :heavy_minus_sign:                                           | N/A                                                          |