package shared

// InsertRequestOptions
// additional options for insert requests.
type InsertRequestOptions struct {
	WriteOptions map[string]interface{} `json:"write_options,omitempty"`
}
