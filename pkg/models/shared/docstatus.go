// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DocStatus struct {
	// The Error type defines a logical error model
	Error *Error `json:"error,omitempty"`
	// An id of the document.
	ID *string `json:"id,omitempty"`
}

func (o *DocStatus) GetError() *Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DocStatus) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
