package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type HealthAPIHealthResponse struct {
	ContentType         string
	HealthCheckResponse *shared.HealthCheckResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
