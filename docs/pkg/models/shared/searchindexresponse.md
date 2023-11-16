# SearchIndexResponse

Response struct for search


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Facets`                                                                   | map[string][shared.SearchFacet](../../../pkg/models/shared/searchfacet.md) | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Hits`                                                                     | [][shared.IndexDoc](../../../pkg/models/shared/indexdoc.md)                | :heavy_minus_sign:                                                         | N/A                                                                        |
| `Meta`                                                                     | [*shared.SearchMetadata](../../../pkg/models/shared/searchmetadata.md)     | :heavy_minus_sign:                                                         | N/A                                                                        |