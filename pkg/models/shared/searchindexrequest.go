package shared

type SearchIndexRequest struct {
	Collation     *Collation `json:"collation,omitempty"`
	ExcludeFields []string   `json:"exclude_fields,omitempty"`
	Facet         *string    `json:"facet,omitempty"`
	Filter        *string    `json:"filter,omitempty"`
	IncludeFields []string   `json:"include_fields,omitempty"`
	Index         *string    `json:"index,omitempty"`
	Page          *int       `json:"page,omitempty"`
	PageSize      *int       `json:"page_size,omitempty"`
	Project       *string    `json:"project,omitempty"`
	Q             *string    `json:"q,omitempty"`
	SearchFields  []string   `json:"search_fields,omitempty"`
	Sort          *string    `json:"sort,omitempty"`
}
