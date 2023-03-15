package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisImportRequest struct {
	ImportRequest shared.ImportRequest `request:"mediaType=application/json"`
	Collection    string               `pathParam:"style=simple,explode=false,name=collection"`
	Project       string               `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisImportResponse struct {
	ContentType    string
	ImportResponse *shared.ImportResponse
	Status         *shared.Status
	StatusCode     int
	RawResponse    *http.Response
}
