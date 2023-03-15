package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisCreateAppKeyRequest struct {
	CreateAppKeyRequest shared.CreateAppKeyRequest `request:"mediaType=application/json"`
	Project             string                     `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisCreateAppKeyResponse struct {
	ContentType          string
	CreateAppKeyResponse *shared.CreateAppKeyResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
