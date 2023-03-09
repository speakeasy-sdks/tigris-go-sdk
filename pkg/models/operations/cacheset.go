package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheSetPathParams struct {
	Key     string `pathParam:"style=simple,explode=false,name=key"`
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheSetRequest struct {
	PathParams CacheSetPathParams
	Request    shared.SetRequest `request:"mediaType=application/json"`
}

type CacheSetResponse struct {
	ContentType string
	SetResponse *shared.SetResponse
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
