package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type SearchListIndexesPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchListIndexesQueryParams struct {
	FilterBranch     *string `queryParam:"style=form,explode=true,name=filter.branch"`
	FilterCollection *string `queryParam:"style=form,explode=true,name=filter.collection"`
	FilterType       *string `queryParam:"style=form,explode=true,name=filter.type"`
}

type SearchListIndexesRequest struct {
	PathParams  SearchListIndexesPathParams
	QueryParams SearchListIndexesQueryParams
}

type SearchListIndexesResponse struct {
	ContentType         string
	ListIndexesResponse *shared.ListIndexesResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
