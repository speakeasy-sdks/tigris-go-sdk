package shared

type UpdateDocumentResponse struct {
	Status []DocStatus `json:"status,omitempty"`
}
