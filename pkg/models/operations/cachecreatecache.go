package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheCreateCachePathParams struct {
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheCreateCacheRequest struct {
	PathParams CacheCreateCachePathParams
	Request    shared.CreateCacheRequest `request:"mediaType=application/json"`
}

type CacheCreateCacheResponse struct {
	ContentType         string
	CreateCacheResponse *shared.CreateCacheResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
