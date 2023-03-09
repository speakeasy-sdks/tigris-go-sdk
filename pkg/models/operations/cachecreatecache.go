package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
