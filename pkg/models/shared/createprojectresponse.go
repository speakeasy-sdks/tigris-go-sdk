// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateProjectResponse - OK
type CreateProjectResponse struct {
	// A detailed response message.
	Message *string `json:"message,omitempty"`
	// An enum with value set as "created".
	Status *string `json:"status,omitempty"`
}

func (o *CreateProjectResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateProjectResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
