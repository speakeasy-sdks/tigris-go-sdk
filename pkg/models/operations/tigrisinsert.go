package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisInsertRequest struct {
	InsertRequest shared.InsertRequest `request:"mediaType=application/json"`
	Collection    string               `pathParam:"style=simple,explode=false,name=collection"`
	Project       string               `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisInsertResponse struct {
	ContentType    string
	InsertResponse *shared.InsertResponse
	Status         *shared.Status
	StatusCode     int
	RawResponse    *http.Response
}
