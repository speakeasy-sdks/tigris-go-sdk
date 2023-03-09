package shared

// ListAppKeysResponse
// ListAppKeysResponse returns one or more visible app keys to user
type ListAppKeysResponse struct {
	AppKeys []AppKey `json:"app_keys,omitempty"`
}
