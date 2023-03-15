package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisRotateAppKeySecretRequest struct {
	RotateAppKeyRequest shared.RotateAppKeyRequest `request:"mediaType=application/json"`
	Project             string                     `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisRotateAppKeySecretResponse struct {
	ContentType          string
	RotateAppKeyResponse *shared.RotateAppKeyResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
