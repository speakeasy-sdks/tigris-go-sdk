package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheDeleteCacheRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	Name        string                 `pathParam:"style=simple,explode=false,name=name"`
	Project     string                 `pathParam:"style=simple,explode=false,name=project"`
}

type CacheDeleteCacheResponse struct {
	ContentType         string
	DeleteCacheResponse *shared.DeleteCacheResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
