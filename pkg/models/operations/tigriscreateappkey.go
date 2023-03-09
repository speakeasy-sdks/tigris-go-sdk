package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisCreateAppKeyPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisCreateAppKeyRequest struct {
	PathParams TigrisCreateAppKeyPathParams
	Request    shared.CreateAppKeyRequest `request:"mediaType=application/json"`
}

type TigrisCreateAppKeyResponse struct {
	ContentType          string
	CreateAppKeyResponse *shared.CreateAppKeyResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
