package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisListProjectsResponse struct {
	ContentType          string
	ListProjectsResponse *shared.ListProjectsResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
