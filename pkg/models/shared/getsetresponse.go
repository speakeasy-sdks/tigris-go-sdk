// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetSetResponse - OK
type GetSetResponse struct {
	// A detailed response message.
	Message *string `json:"message,omitempty"`
	// An old value if exists
	OldValue *string `json:"old_value,omitempty"`
	// An enum with value set as "set"
	Status *string `json:"status,omitempty"`
}

func (o *GetSetResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetSetResponse) GetOldValue() *string {
	if o == nil {
		return nil
	}
	return o.OldValue
}

func (o *GetSetResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
