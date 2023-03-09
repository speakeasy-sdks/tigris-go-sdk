package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type HealthAPIHealthResponse struct {
	ContentType         string
	HealthCheckResponse *shared.HealthCheckResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
