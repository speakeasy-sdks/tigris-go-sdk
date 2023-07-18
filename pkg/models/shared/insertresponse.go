// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// InsertResponse - OK
type InsertResponse struct {
	// an array returns the value of the primary keys.
	Keys []string `json:"keys,omitempty"`
	// Has metadata related to the documents stored.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
	// An enum with value set as "inserted"
	Status *string `json:"status,omitempty"`
}

func (o *InsertResponse) GetKeys() []string {
	if o == nil {
		return nil
	}
	return o.Keys
}

func (o *InsertResponse) GetMetadata() *ResponseMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *InsertResponse) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}
