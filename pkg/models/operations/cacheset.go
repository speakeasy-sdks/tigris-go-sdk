package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheSetRequest struct {
	SetRequest shared.SetRequest `request:"mediaType=application/json"`
	Key        string            `pathParam:"style=simple,explode=false,name=key"`
	Name       string            `pathParam:"style=simple,explode=false,name=name"`
	Project    string            `pathParam:"style=simple,explode=false,name=project"`
}

type CacheSetResponse struct {
	ContentType string
	SetResponse *shared.SetResponse
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
