// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeleteBranchResponse struct {
	// A detailed response message.
	Message *string `json:"message,omitempty"`
	// An enum with value set as "deleted".
	Status *string `json:"status,omitempty"`
}

func (o *DeleteBranchResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteBranchResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
