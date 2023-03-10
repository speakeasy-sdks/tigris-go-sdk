package shared

// UpdateNamespaceMetadataResponse
// Update of namespace metadata response
type UpdateNamespaceMetadataResponse struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	NamespaceID *int64                 `json:"namespaceId,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
