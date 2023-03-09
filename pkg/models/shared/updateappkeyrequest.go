package shared

// UpdateAppKeyRequest
// To update the description of the app key
type UpdateAppKeyRequest struct {
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}
