// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetAccessTokenResponse - The response of GetAccessToken which contains access_token and optionally refresh_token.
type GetAccessTokenResponse struct {
	// An Access Token.
	AccessToken *string `json:"access_token,omitempty"`
	// Access token expiration timeout in seconds.
	ExpiresIn *int `json:"expires_in,omitempty"`
	// The Refresh Token.
	RefreshToken *string `json:"refresh_token,omitempty"`
}
