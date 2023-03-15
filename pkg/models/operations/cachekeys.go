package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheKeysRequest struct {
	Count   *int64  `queryParam:"style=form,explode=true,name=count"`
	Cursor  *int64  `queryParam:"style=form,explode=true,name=cursor"`
	Name    string  `pathParam:"style=simple,explode=false,name=name"`
	Pattern *string `queryParam:"style=form,explode=true,name=pattern"`
	Project string  `pathParam:"style=simple,explode=false,name=project"`
}

type CacheKeysResponse struct {
	ContentType  string
	KeysResponse *shared.KeysResponse
	Status       *shared.Status
	StatusCode   int
	RawResponse  *http.Response
}
