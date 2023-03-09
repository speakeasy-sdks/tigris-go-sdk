package shared

// TransactionCtx
// Contains ID which uniquely identifies transaction This context is returned by BeginTransaction request and should be passed in the metadata (headers) of subsequent requests in order to run them in the context of the same transaction.
type TransactionCtx struct {
	ID     *string `json:"id,omitempty"`
	Origin *string `json:"origin,omitempty"`
}
