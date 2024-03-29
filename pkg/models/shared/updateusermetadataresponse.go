// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateUserMetadataResponseValue struct {
}

// UpdateUserMetadataResponse - Update of user metadata response
type UpdateUserMetadataResponse struct {
	MetadataKey *string                          `json:"metadataKey,omitempty"`
	NamespaceID *int64                           `json:"namespaceId,omitempty"`
	UserID      *string                          `json:"userId,omitempty"`
	Value       *UpdateUserMetadataResponseValue `json:"value,omitempty"`
}

func (o *UpdateUserMetadataResponse) GetMetadataKey() *string {
	if o == nil {
		return nil
	}
	return o.MetadataKey
}

func (o *UpdateUserMetadataResponse) GetNamespaceID() *int64 {
	if o == nil {
		return nil
	}
	return o.NamespaceID
}

func (o *UpdateUserMetadataResponse) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *UpdateUserMetadataResponse) GetValue() *UpdateUserMetadataResponseValue {
	if o == nil {
		return nil
	}
	return o.Value
}
