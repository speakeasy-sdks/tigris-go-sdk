package shared

// SearchResponse
// Response struct for search
type SearchResponse struct {
	Facets map[string]SearchFacet `json:"facets,omitempty"`
	Hits   []SearchHit            `json:"hits,omitempty"`
	Meta   *SearchMetadata        `json:"meta,omitempty"`
}
