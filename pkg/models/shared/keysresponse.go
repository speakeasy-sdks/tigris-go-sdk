// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type KeysResponse struct {
	// cursor - 0 is the keys scan is finished, non-zero cursor can be passed in next keys request to continue the scan this is useful if streaming breaks and user wants to resume stream
	Cursor *int64 `json:"cursor,omitempty"`
	// keys
	Keys []string `json:"keys,omitempty"`
}

func (o *KeysResponse) GetCursor() *int64 {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *KeysResponse) GetKeys() []string {
	if o == nil {
		return nil
	}
	return o.Keys
}
