// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SearchIndexResponse - Response struct for search
type SearchIndexResponse struct {
	Facets map[string]SearchFacet `json:"facets,omitempty"`
	Hits   []IndexDoc             `json:"hits,omitempty"`
	Meta   *SearchMetadata        `json:"meta,omitempty"`
}
