package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
