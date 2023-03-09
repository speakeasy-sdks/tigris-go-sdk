package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDeletePathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDeleteRequest struct {
	PathParams TigrisDeletePathParams
	Request    shared.DeleteRequest `request:"mediaType=application/json"`
}

type TigrisDeleteResponse struct {
	ContentType    string
	DeleteResponse *shared.DeleteResponse
	Status         *shared.Status
	StatusCode     int
	RawResponse    *http.Response
}
