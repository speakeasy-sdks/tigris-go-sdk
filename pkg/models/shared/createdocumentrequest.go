// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateDocumentRequest struct {
	// An array of documents to be created or replaced. Each document is a JSON object.
	Documents []string `json:"documents,omitempty"`
	// index name where to create documents.
	Index *string `json:"index,omitempty"`
	// Tigris project name.
	Project *string `json:"project,omitempty"`
}

func (o *CreateDocumentRequest) GetDocuments() []string {
	if o == nil {
		return nil
	}
	return o.Documents
}

func (o *CreateDocumentRequest) GetIndex() *string {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *CreateDocumentRequest) GetProject() *string {
	if o == nil {
		return nil
	}
	return o.Project
}
