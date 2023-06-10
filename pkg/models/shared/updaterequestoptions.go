// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateRequestOptions - Additional options for update requests.
type UpdateRequestOptions struct {
	// A collation allows you to specify string comparison rules. Default is case-sensitive, to override it you can set this option to 'ci' that will apply to all the text fields in the filters.
	Collation *Collation `json:"collation,omitempty"`
	// Limit the number of documents to be updated
	Limit *int64 `json:"limit,omitempty"`
	// Additional options to modify write requests.
	WriteOptions *WriteOptions `json:"write_options,omitempty"`
}
