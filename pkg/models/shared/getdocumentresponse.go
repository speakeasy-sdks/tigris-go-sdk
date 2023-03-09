package shared

type GetDocumentResponse struct {
	Documents []IndexDoc `json:"documents,omitempty"`
}
