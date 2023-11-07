# SearchIndexResponse

Response struct for search


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Facets`                                                            | map[string][shared.SearchFacet](../../models/shared/searchfacet.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Hits`                                                              | [][shared.IndexDoc](../../models/shared/indexdoc.md)                | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Meta`                                                              | [*shared.SearchMetadata](../../models/shared/searchmetadata.md)     | :heavy_minus_sign:                                                  | N/A                                                                 |