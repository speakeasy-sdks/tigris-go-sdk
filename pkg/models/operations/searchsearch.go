package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchSearchRequest struct {
	SearchIndexRequest shared.SearchIndexRequest `request:"mediaType=application/json"`
	Index              string                    `pathParam:"style=simple,explode=false,name=index"`
	Project            string                    `pathParam:"style=simple,explode=false,name=project"`
}

type SearchSearchResponse struct {
	ContentType         string
	SearchIndexResponse *shared.SearchIndexResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
