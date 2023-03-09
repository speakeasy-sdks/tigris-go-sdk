package shared

// BeginTransactionRequest
// Start new transaction in project specified by "project".
type BeginTransactionRequest struct {
	Branch  *string                `json:"branch,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
}
