package shared

type SearchMetadata struct {
	Found      *int64 `json:"found,omitempty"`
	Page       *Page  `json:"page,omitempty"`
	TotalPages *int   `json:"total_pages,omitempty"`
}
