package shared

// Page
// Pagination metadata for SearchResponse
type Page struct {
	Current *int `json:"current,omitempty"`
	Size    *int `json:"size,omitempty"`
}
