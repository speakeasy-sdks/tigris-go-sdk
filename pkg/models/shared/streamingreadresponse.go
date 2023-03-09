package shared

type StreamingReadResponse struct {
	Error  *Error        `json:"error,omitempty"`
	Result *ReadResponse `json:"result,omitempty"`
}
