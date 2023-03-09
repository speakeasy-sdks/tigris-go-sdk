package shared

type UpdateResponse struct {
	Metadata      *ResponseMetadata `json:"metadata,omitempty"`
	ModifiedCount *int              `json:"modified_count,omitempty"`
	Status        *string           `json:"status,omitempty"`
}
