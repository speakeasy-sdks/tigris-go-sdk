// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateDocumentRequest struct {
	// An array of documents. Each document should have "id" present which will be used by Tigris for updating the document.
	Documents []string `json:"documents,omitempty"`
	// Index name where to create documents.
	Index *string `json:"index,omitempty"`
	// Project name whose db is under target to insert documents.
	Project *string `json:"project,omitempty"`
}
