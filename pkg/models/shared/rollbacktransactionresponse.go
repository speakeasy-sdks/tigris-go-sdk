// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RollbackTransactionResponse - OK
type RollbackTransactionResponse struct {
	// Status of rollback transaction operation.
	Status *string `json:"status,omitempty"`
}

func (o *RollbackTransactionResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
