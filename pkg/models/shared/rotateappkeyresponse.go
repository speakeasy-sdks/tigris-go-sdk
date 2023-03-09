package shared

// RotateAppKeyResponse
// RotateAppKeyResponse returns the new app key with rotated secret
type RotateAppKeyResponse struct {
	AppKey *AppKey `json:"app_key,omitempty"`
}
