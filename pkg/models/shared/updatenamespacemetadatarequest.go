// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateNamespaceMetadataRequestValue struct {
}

// UpdateNamespaceMetadataRequest - Request update of namespace metadata
type UpdateNamespaceMetadataRequest struct {
	MetadataKey *string                              `json:"metadataKey,omitempty"`
	Value       *UpdateNamespaceMetadataRequestValue `json:"value,omitempty"`
}

func (o *UpdateNamespaceMetadataRequest) GetMetadataKey() *string {
	if o == nil {
		return nil
	}
	return o.MetadataKey
}

func (o *UpdateNamespaceMetadataRequest) GetValue() *UpdateNamespaceMetadataRequestValue {
	if o == nil {
		return nil
	}
	return o.Value
}
