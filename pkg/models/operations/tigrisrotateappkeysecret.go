package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
