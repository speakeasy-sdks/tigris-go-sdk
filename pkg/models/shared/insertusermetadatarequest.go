package shared

// InsertUserMetadataRequest
// Request insertion of user metadata
type InsertUserMetadataRequest struct {
	MetadataKey *string                `json:"metadataKey,omitempty"`
	Value       map[string]interface{} `json:"value,omitempty"`
}
