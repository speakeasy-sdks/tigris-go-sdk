// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GetIndexResponse struct {
	Index *IndexInfo `json:"index,omitempty"`
}

func (o *GetIndexResponse) GetIndex() *IndexInfo {
	if o == nil {
		return nil
	}
	return o.Index
}
