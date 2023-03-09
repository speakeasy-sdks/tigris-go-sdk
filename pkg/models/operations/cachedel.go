package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type CacheDelPathParams struct {
	Key     string `pathParam:"style=simple,explode=false,name=key"`
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheDelRequest struct {
	PathParams CacheDelPathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
}

type CacheDelResponse struct {
	ContentType string
	DelResponse *shared.DelResponse
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
