package shared

type DocStatus struct {
	Error *Error  `json:"error,omitempty"`
	ID    *string `json:"id,omitempty"`
}
