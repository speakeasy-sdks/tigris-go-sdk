package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisImportPathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisImportRequest struct {
	PathParams TigrisImportPathParams
	Request    shared.ImportRequest `request:"mediaType=application/json"`
}

type TigrisImportResponse struct {
	ContentType    string
	ImportResponse *shared.ImportResponse
	Status         *shared.Status
	StatusCode     int
	RawResponse    *http.Response
}
