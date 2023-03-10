package shared

// UpdateUserMetadataResponse
// Update of user metadata response
type UpdateUserMetadataResponse struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	NamespaceID *int64                 `json:"namespaceId,omitempty"`
	UserID      *string                `json:"userId,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
