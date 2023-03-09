package shared

type UpdateDocumentRequest struct {
	Documents []string `json:"documents,omitempty"`
	Index     *string  `json:"index,omitempty"`
	Project   *string  `json:"project,omitempty"`
}
