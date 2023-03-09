package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type SearchSearchPathParams struct {
	Index   string `pathParam:"style=simple,explode=false,name=index"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchSearchRequest struct {
	PathParams SearchSearchPathParams
	Request    shared.SearchIndexRequest `request:"mediaType=application/json"`
}

type SearchSearchResponse struct {
	ContentType         string
	SearchIndexResponse *shared.SearchIndexResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
