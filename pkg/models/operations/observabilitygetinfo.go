package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type ObservabilityGetInfoResponse struct {
	ContentType     string
	GetInfoResponse *shared.GetInfoResponse
	Status          *shared.Status
	StatusCode      int
	RawResponse     *http.Response
}
