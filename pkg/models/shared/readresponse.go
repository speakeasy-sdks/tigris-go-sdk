package shared

type ReadResponse struct {
	Data        map[string]interface{} `json:"data,omitempty"`
	Metadata    *ResponseMetadata      `json:"metadata,omitempty"`
	ResumeToken *string                `json:"resume_token,omitempty"`
}
