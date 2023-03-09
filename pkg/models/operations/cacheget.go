package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type CacheGetPathParams struct {
	Key     string `pathParam:"style=simple,explode=false,name=key"`
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheGetRequest struct {
	PathParams CacheGetPathParams
}

type CacheGetResponse struct {
	ContentType string
	GetResponse *shared.GetResponse
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
