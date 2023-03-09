package shared

// UpdateUserMetadataRequest
// Request update of user metadata
type UpdateUserMetadataRequest struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
