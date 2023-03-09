package shared

type CreateOrReplaceDocumentResponse struct {
	Status []DocStatus `json:"status,omitempty"`
}
