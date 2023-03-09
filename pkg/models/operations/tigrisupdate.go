package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisUpdatePathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisUpdateRequest struct {
	PathParams TigrisUpdatePathParams
	Request    shared.UpdateRequest `request:"mediaType=application/json"`
}

type TigrisUpdateResponse struct {
	ContentType    string
	StatusCode     int
	RawResponse    *http.Response
	Status         *shared.Status
	UpdateResponse *shared.UpdateResponse
}
