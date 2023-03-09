package shared

// GetUserMetadataResponse
// User metadata response
type GetUserMetadataResponse struct {
	MetadataKey *string `json:"metadataKey,omitempty"`
	NamespaceID *int64  `json:"namespaceId,omitempty"`
	UserID      *string `json:"userId,omitempty"`
	Value       *string `json:"value,omitempty"`
}
