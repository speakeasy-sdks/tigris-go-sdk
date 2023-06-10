# SearchResponse

Response struct for search


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Facets`                                                     | map[string][SearchFacet](../../models/shared/searchfacet.md) | :heavy_minus_sign:                                           | N/A                                                          |
| `Hits`                                                       | [][SearchHit](../../models/shared/searchhit.md)              | :heavy_minus_sign:                                           | N/A                                                          |
| `Meta`                                                       | [*SearchMetadata](../../models/shared/searchmetadata.md)     | :heavy_minus_sign:                                           | N/A                                                          |