package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListProjectsResponse struct {
	ContentType          string
	ListProjectsResponse *shared.ListProjectsResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
