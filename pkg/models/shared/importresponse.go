// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ImportResponse - OK
type ImportResponse struct {
	// an array returns the value of the primary keys.
	Keys []string `json:"keys,omitempty"`
	// Has metadata related to the documents stored.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
	// An enum with value set as "inserted"
	Status *string `json:"status,omitempty"`
}

func (o *ImportResponse) GetKeys() []string {
	if o == nil {
		return nil
	}
	return o.Keys
}

func (o *ImportResponse) GetMetadata() *ResponseMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ImportResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
