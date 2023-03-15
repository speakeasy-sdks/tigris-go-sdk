package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchGetIndexRequest struct {
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchGetIndexResponse struct {
	ContentType      string
	GetIndexResponse *shared.GetIndexResponse
	Status           *shared.Status
	StatusCode       int
	RawResponse      *http.Response
}
