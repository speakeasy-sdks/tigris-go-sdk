package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisUpdateAppKeyRequest struct {
	UpdateAppKeyRequest shared.UpdateAppKeyRequest `request:"mediaType=application/json"`
	Project             string                     `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisUpdateAppKeyResponse struct {
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	Status               *shared.Status
	UpdateAppKeyResponse *shared.UpdateAppKeyResponse
}
