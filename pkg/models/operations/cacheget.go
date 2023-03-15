package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CacheGetRequest struct {
	Key     string `pathParam:"style=simple,explode=false,name=key"`
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type CacheGetResponse struct {
	ContentType string
	GetResponse *shared.GetResponse
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
