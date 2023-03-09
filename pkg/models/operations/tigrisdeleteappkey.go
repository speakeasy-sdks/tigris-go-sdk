package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisDeleteAppKeyPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDeleteAppKeyRequest struct {
	PathParams TigrisDeleteAppKeyPathParams
	Request    shared.DeleteAppKeyRequest `request:"mediaType=application/json"`
}

type TigrisDeleteAppKeyResponse struct {
	ContentType          string
	DeleteAppKeyResponse *shared.DeleteAppKeyResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
