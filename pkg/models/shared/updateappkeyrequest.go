// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateAppKeyRequest - To update the description of the app key
type UpdateAppKeyRequest struct {
	// A new human readable app description
	Description *string `json:"description,omitempty"`
	// app key id - this is not allowed to update
	ID *string `json:"id,omitempty"`
	// A new human readable app name
	Name *string `json:"name,omitempty"`
}

func (o *UpdateAppKeyRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateAppKeyRequest) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateAppKeyRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
