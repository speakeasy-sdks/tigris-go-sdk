// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GetDocumentResponse struct {
	// An array of documents.
	Documents []IndexDoc `json:"documents,omitempty"`
}

func (o *GetDocumentResponse) GetDocuments() []IndexDoc {
	if o == nil {
		return nil
	}
	return o.Documents
}
