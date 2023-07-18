// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type AuthGetAccessTokenResponse struct {
	ContentType string
	// OK
	GetAccessTokenResponse *shared.GetAccessTokenResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}

func (o *AuthGetAccessTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AuthGetAccessTokenResponse) GetGetAccessTokenResponse() *shared.GetAccessTokenResponse {
	if o == nil {
		return nil
	}
	return o.GetAccessTokenResponse
}

func (o *AuthGetAccessTokenResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AuthGetAccessTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AuthGetAccessTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
