package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisSearchPathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisSearchRequest struct {
	PathParams TigrisSearchPathParams
	Request    shared.SearchRequest `request:"mediaType=application/json"`
}

type TigrisSearchResponse struct {
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
	Status                  *shared.Status
	StreamingSearchResponse *shared.StreamingSearchResponse
}
