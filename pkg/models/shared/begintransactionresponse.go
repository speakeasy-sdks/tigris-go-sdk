package shared

// BeginTransactionResponse
// Start transaction returns transaction context  which uniquely identifies the transaction
type BeginTransactionResponse struct {
	TxCtx *TransactionCtx `json:"tx_ctx,omitempty"`
}
