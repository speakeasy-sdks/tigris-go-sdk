package shared

// InsertNamespaceMetadataRequest
// Request insertion of namespace metadata
type InsertNamespaceMetadataRequest struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
