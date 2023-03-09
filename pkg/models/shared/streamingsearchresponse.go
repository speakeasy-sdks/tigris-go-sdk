package shared

type StreamingSearchResponse struct {
	Error  *Error          `json:"error,omitempty"`
	Result *SearchResponse `json:"result,omitempty"`
}
