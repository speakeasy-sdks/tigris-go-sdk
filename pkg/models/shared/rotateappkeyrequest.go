package shared

// RotateAppKeyRequest
// Request rotation of an app key secret
type RotateAppKeyRequest struct {
	ID      *string `json:"id,omitempty"`
	Project *string `json:"project,omitempty"`
}
