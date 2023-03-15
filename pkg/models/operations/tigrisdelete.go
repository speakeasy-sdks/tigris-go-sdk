package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDeleteRequest struct {
	DeleteRequest shared.DeleteRequest `request:"mediaType=application/json"`
	Collection    string               `pathParam:"style=simple,explode=false,name=collection"`
	Project       string               `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDeleteResponse struct {
	ContentType    string
	DeleteResponse *shared.DeleteResponse
	Status         *shared.Status
	StatusCode     int
	RawResponse    *http.Response
}
