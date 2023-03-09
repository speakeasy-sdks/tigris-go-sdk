package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchGetIndexPathParams struct {
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchGetIndexRequest struct {
	PathParams SearchGetIndexPathParams
}

type SearchGetIndexResponse struct {
	ContentType      string
	GetIndexResponse *shared.GetIndexResponse
	Status           *shared.Status
	StatusCode       int
	RawResponse      *http.Response
}
