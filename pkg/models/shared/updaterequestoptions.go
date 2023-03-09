package shared

// UpdateRequestOptions
// Additional options for update requests.
type UpdateRequestOptions struct {
	Collation    *Collation             `json:"collation,omitempty"`
	Limit        *int64                 `json:"limit,omitempty"`
	WriteOptions map[string]interface{} `json:"write_options,omitempty"`
}
