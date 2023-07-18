// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ReadResponseData - Object containing the collection document.
type ReadResponseData struct {
}

type ReadResponse struct {
	// Object containing the collection document.
	Data *ReadResponseData `json:"data,omitempty"`
	// Has metadata related to the documents stored.
	Metadata *ResponseMetadata `json:"metadata,omitempty"`
	// An internal key, used for pagination.
	ResumeToken *string `json:"resume_token,omitempty"`
}

func (o *ReadResponse) GetData() *ReadResponseData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ReadResponse) GetMetadata() *ResponseMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ReadResponse) GetResumeToken() *string {
	if o == nil {
		return nil
	}
	return o.ResumeToken
}
