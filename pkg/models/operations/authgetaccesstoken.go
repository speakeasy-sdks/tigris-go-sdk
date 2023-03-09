package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type AuthGetAccessTokenResponse struct {
	ContentType            string
	GetAccessTokenResponse *shared.GetAccessTokenResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
