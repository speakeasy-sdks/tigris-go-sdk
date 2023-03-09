package shared

// UpdateNamespaceMetadataRequest
// Request update of namespace metadata
type UpdateNamespaceMetadataRequest struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
