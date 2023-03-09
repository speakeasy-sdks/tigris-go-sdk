package shared

type DeleteDocumentResponse struct {
	Status []DocStatus `json:"status,omitempty"`
}
