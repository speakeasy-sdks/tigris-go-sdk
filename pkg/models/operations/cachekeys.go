package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type CacheKeysPathParams struct {
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheKeysQueryParams struct {
	Count   *int64  `queryParam:"style=form,explode=true,name=count"`
	Cursor  *int64  `queryParam:"style=form,explode=true,name=cursor"`
	Pattern *string `queryParam:"style=form,explode=true,name=pattern"`
}

type CacheKeysRequest struct {
	PathParams  CacheKeysPathParams
	QueryParams CacheKeysQueryParams
}

type CacheKeysResponse struct {
	ContentType  string
	KeysResponse *shared.KeysResponse
	Status       *shared.Status
	StatusCode   int
	RawResponse  *http.Response
}
