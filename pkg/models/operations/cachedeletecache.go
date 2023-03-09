package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type CacheDeleteCachePathParams struct {
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheDeleteCacheRequest struct {
	PathParams CacheDeleteCachePathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
}

type CacheDeleteCacheResponse struct {
	ContentType         string
	DeleteCacheResponse *shared.DeleteCacheResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
