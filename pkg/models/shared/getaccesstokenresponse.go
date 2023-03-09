package shared

// GetAccessTokenResponse
// The response of GetAccessToken which contains access_token and optionally refresh_token.
type GetAccessTokenResponse struct {
	AccessToken  *string `json:"access_token,omitempty"`
	ExpiresIn    *int    `json:"expires_in,omitempty"`
	RefreshToken *string `json:"refresh_token,omitempty"`
}
