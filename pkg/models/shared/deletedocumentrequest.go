package shared

type DeleteDocumentRequest struct {
	Ids     []string `json:"ids,omitempty"`
	Index   *string  `json:"index,omitempty"`
	Project *string  `json:"project,omitempty"`
}
