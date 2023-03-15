package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheDelRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	Key         string                 `pathParam:"style=simple,explode=false,name=key"`
	Name        string                 `pathParam:"style=simple,explode=false,name=name"`
	Project     string                 `pathParam:"style=simple,explode=false,name=project"`
}

type CacheDelResponse struct {
	ContentType string
	DelResponse *shared.DelResponse
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
