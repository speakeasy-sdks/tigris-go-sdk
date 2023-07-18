// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateCacheResponse - OK
type CreateCacheResponse struct {
	// A detailed response message.
	Message *string `json:"message,omitempty"`
	// An enum with value set as "created"
	Status *string `json:"status,omitempty"`
}

func (o *CreateCacheResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateCacheResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
