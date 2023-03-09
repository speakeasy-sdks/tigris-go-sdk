package shared

// ReplaceRequestOptions
// Additional options for replace requests.
type ReplaceRequestOptions struct {
	WriteOptions map[string]interface{} `json:"write_options,omitempty"`
}
