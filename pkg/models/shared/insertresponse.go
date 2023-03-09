package shared

type InsertResponse struct {
	Keys     []string          `json:"keys,omitempty"`
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
	Status   *string           `json:"status,omitempty"`
}
