package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type CacheListCachesPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheListCachesRequest struct {
	PathParams CacheListCachesPathParams
}

type CacheListCachesResponse struct {
	ContentType        string
	ListCachesResponse *shared.ListCachesResponse
	Status             *shared.Status
	StatusCode         int
	RawResponse        *http.Response
}
