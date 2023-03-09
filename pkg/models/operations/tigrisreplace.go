package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisReplacePathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisReplaceRequest struct {
	PathParams TigrisReplacePathParams
	Request    shared.ReplaceRequest `request:"mediaType=application/json"`
}

type TigrisReplaceResponse struct {
	ContentType     string
	ReplaceResponse *shared.ReplaceResponse
	Status          *shared.Status
	StatusCode      int
	RawResponse     *http.Response
}
