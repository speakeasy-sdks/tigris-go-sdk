package shared

type GetResponse struct {
	ExpiresInMs *int64  `json:"expires_in_ms,omitempty"`
	Value       *string `json:"value,omitempty"`
}
