package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheListCachesRequest struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheListCachesResponse struct {
	ContentType        string
	ListCachesResponse *shared.ListCachesResponse
	Status             *shared.Status
	StatusCode         int
	RawResponse        *http.Response
}
