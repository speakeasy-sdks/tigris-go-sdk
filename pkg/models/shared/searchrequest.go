package shared

type SearchRequest struct {
	Branch        *string                `json:"branch,omitempty"`
	Collation     *Collation             `json:"collation,omitempty"`
	ExcludeFields []string               `json:"exclude_fields,omitempty"`
	Facet         map[string]interface{} `json:"facet,omitempty"`
	Fields        map[string]interface{} `json:"fields,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
	IncludeFields []string               `json:"include_fields,omitempty"`
	Page          *int                   `json:"page,omitempty"`
	PageSize      *int                   `json:"page_size,omitempty"`
	Q             *string                `json:"q,omitempty"`
	SearchFields  []string               `json:"search_fields,omitempty"`
	Sort          map[string]interface{} `json:"sort,omitempty"`
}
