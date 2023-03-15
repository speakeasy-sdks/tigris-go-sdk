package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisUpdateRequest struct {
	UpdateRequest shared.UpdateRequest `request:"mediaType=application/json"`
	Collection    string               `pathParam:"style=simple,explode=false,name=collection"`
	Project       string               `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisUpdateResponse struct {
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	Status         *shared.Status
	UpdateResponse *shared.UpdateResponse
}
