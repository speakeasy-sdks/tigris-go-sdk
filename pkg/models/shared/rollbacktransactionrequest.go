package shared

// RollbackTransactionRequest
// Rollback transaction with the given ID
type RollbackTransactionRequest struct {
	Branch *string `json:"branch,omitempty"`
}
