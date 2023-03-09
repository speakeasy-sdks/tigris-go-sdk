package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ObservabilityGetInfoResponse struct {
	ContentType     string
	GetInfoResponse *shared.GetInfoResponse
	Status          *shared.Status
	StatusCode      int
	RawResponse     *http.Response
}
