package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchListIndexesRequest struct {
	FilterBranch     *string `queryParam:"style=form,explode=true,name=filter.branch"`
	FilterCollection *string `queryParam:"style=form,explode=true,name=filter.collection"`
	FilterType       *string `queryParam:"style=form,explode=true,name=filter.type"`
	Project          string  `pathParam:"style=simple,explode=false,name=project"`
}

type SearchListIndexesResponse struct {
	ContentType         string
	ListIndexesResponse *shared.ListIndexesResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
