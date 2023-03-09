package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
