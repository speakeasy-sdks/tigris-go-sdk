package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDeleteAppKeyRequest struct {
	DeleteAppKeyRequest shared.DeleteAppKeyRequest `request:"mediaType=application/json"`
	Project             string                     `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDeleteAppKeyResponse struct {
	ContentType          string
	DeleteAppKeyResponse *shared.DeleteAppKeyResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
