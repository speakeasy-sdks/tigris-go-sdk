package shared

// CreateAppKeyResponse
// CreateAppKeyResponse returns created app keys
type CreateAppKeyResponse struct {
	CreatedAppKey *AppKey `json:"created_app_key,omitempty"`
}
