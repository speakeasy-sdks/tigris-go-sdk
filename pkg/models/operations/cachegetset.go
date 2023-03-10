package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheGetSetPathParams struct {
	Key     string `pathParam:"style=simple,explode=false,name=key"`
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheGetSetRequest struct {
	PathParams CacheGetSetPathParams
	Request    shared.GetSetRequest `request:"mediaType=application/json"`
}

type CacheGetSetResponse struct {
	ContentType    string
	GetSetResponse *shared.GetSetResponse
	Status         *shared.Status
	StatusCode     int
	RawResponse    *http.Response
}
