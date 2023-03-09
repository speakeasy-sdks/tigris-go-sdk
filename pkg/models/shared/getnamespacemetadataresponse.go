package shared

// GetNamespaceMetadataResponse
// Namespace metadata response
type GetNamespaceMetadataResponse struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	NamespaceID *int64                 `json:"namespaceId,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
