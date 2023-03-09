package shared

// InsertNamespaceMetadataResponse
// Insertion of namespace metadata response
type InsertNamespaceMetadataResponse struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	NamespaceID *int64                 `json:"namespaceId,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
