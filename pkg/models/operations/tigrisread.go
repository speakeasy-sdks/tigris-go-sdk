package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisReadPathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisReadRequest struct {
	PathParams TigrisReadPathParams
	Request    shared.ReadRequest `request:"mediaType=application/json"`
}

type TigrisReadResponse struct {
	ContentType           string
	StatusCode            int
	RawResponse           *http.Response
	Status                *shared.Status
	StreamingReadResponse *shared.StreamingReadResponse
}
