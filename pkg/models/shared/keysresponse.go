package shared

type KeysResponse struct {
	Cursor *int64   `json:"cursor,omitempty"`
	Keys   []string `json:"keys,omitempty"`
}
