package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisRotateAppKeySecretPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisRotateAppKeySecretRequest struct {
	PathParams TigrisRotateAppKeySecretPathParams
	Request    shared.RotateAppKeyRequest `request:"mediaType=application/json"`
}

type TigrisRotateAppKeySecretResponse struct {
	ContentType          string
	RotateAppKeyResponse *shared.RotateAppKeyResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
