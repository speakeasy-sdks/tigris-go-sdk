// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// StreamingReadResponse - OK
type StreamingReadResponse struct {
	// The Error type defines a logical error model
	Error  *Error        `json:"error,omitempty"`
	Result *ReadResponse `json:"result,omitempty"`
}

func (o *StreamingReadResponse) GetError() *Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *StreamingReadResponse) GetResult() *ReadResponse {
	if o == nil {
		return nil
	}
	return o.Result
}
