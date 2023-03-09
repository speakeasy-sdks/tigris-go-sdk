package shared

// UpdateAppKeyResponse
// Returns response for updating the app key description
type UpdateAppKeyResponse struct {
	UpdatedAppKey *AppKey `json:"updated_app_key,omitempty"`
}
