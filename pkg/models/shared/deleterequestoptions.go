package shared

// DeleteRequestOptions
// Additional options for deleted requests.
type DeleteRequestOptions struct {
	Collation    *Collation             `json:"collation,omitempty"`
	Limit        *int64                 `json:"limit,omitempty"`
	WriteOptions map[string]interface{} `json:"write_options,omitempty"`
}
