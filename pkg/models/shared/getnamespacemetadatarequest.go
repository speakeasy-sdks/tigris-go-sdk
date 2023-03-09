package shared

// GetNamespaceMetadataRequest
// Request namespace metadata
type GetNamespaceMetadataRequest struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
