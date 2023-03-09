package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type AuthGetAccessTokenResponse struct {
	ContentType            string
	GetAccessTokenResponse *shared.GetAccessTokenResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
