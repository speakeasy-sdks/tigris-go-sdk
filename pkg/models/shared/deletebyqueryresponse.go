// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeleteByQueryResponse struct {
	// The number of documents deleted.
	Count *int `json:"count,omitempty"`
}

func (o *DeleteByQueryResponse) GetCount() *int {
	if o == nil {
		return nil
	}
	return o.Count
}
