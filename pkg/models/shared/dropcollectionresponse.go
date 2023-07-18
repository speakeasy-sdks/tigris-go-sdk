// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DropCollectionResponse - OK
type DropCollectionResponse struct {
	// A detailed response message.
	Message *string `json:"message,omitempty"`
	// An enum with value set as "dropped".
	Status *string `json:"status,omitempty"`
}

func (o *DropCollectionResponse) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DropCollectionResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
