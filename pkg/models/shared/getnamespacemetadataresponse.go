// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GetNamespaceMetadataResponseValue struct {
}

// GetNamespaceMetadataResponse - Namespace metadata response
type GetNamespaceMetadataResponse struct {
	MetadataKey *string                            `json:"metadataKey,omitempty"`
	NamespaceID *int64                             `json:"namespaceId,omitempty"`
	Value       *GetNamespaceMetadataResponseValue `json:"value,omitempty"`
}
