package shared

// CommitTransactionRequest
// Commit transaction with the given ID
type CommitTransactionRequest struct {
	Branch *string `json:"branch,omitempty"`
}
