package shared

// InsertUserMetadataResponse
// Insertion of user metadata response
type InsertUserMetadataResponse struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	NamespaceID *int64                 `json:"namespaceId,omitempty"`
	UserID      *string                `json:"userId,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
