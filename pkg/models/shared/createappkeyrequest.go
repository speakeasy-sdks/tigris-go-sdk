package shared

// CreateAppKeyRequest
// Request creation of user app key
type CreateAppKeyRequest struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}
