package shared

type DeleteResponse struct {
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
	Status   *string           `json:"status,omitempty"`
}
