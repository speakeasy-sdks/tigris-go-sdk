package shared

type SearchFacet struct {
	Counts []FacetCount `json:"counts,omitempty"`
	Stats  *FacetStats  `json:"stats,omitempty"`
}
