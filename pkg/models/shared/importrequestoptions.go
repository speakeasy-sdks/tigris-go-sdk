package shared

// ImportRequestOptions
// additional options for import requests.
type ImportRequestOptions struct {
	WriteOptions map[string]interface{} `json:"write_options,omitempty"`
}
