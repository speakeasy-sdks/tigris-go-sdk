package shared

// DeleteAppKeyResponse
// AppKeys returns the flag to convey if app key was deleted
type DeleteAppKeyResponse struct {
	Deleted *bool `json:"deleted,omitempty"`
}
