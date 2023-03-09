package shared

// ReadRequestOptions
// Options that can be used to modify the results, for example "limit" to control the number of documents returned by the server.
type ReadRequestOptions struct {
	Collation *Collation `json:"collation,omitempty"`
	Limit     *int64     `json:"limit,omitempty"`
	Offset    *string    `json:"offset,omitempty"`
	Skip      *int64     `json:"skip,omitempty"`
}
