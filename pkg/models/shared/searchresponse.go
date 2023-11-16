// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SearchResponse - Response struct for search
type SearchResponse struct {
	Facets map[string]SearchFacet `json:"facets,omitempty"`
	Hits   []SearchHit            `json:"hits,omitempty"`
	Meta   *SearchMetadata        `json:"meta,omitempty"`
}

func (o *SearchResponse) GetFacets() map[string]SearchFacet {
	if o == nil {
		return nil
	}
	return o.Facets
}

func (o *SearchResponse) GetHits() []SearchHit {
	if o == nil {
		return nil
	}
	return o.Hits
}

func (o *SearchResponse) GetMeta() *SearchMetadata {
	if o == nil {
		return nil
	}
	return o.Meta
}
