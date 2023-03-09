package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisUpdateAppKeyPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisUpdateAppKeyRequest struct {
	PathParams TigrisUpdateAppKeyPathParams
	Request    shared.UpdateAppKeyRequest `request:"mediaType=application/json"`
}

type TigrisUpdateAppKeyResponse struct {
	ContentType          string
	StatusCode           int
	RawResponse          *http.Response
	Status               *shared.Status
	UpdateAppKeyResponse *shared.UpdateAppKeyResponse
}
