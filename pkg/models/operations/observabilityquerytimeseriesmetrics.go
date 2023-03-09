package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ObservabilityQueryTimeSeriesMetricsRequest struct {
	Request shared.QueryTimeSeriesMetricsRequest `request:"mediaType=application/json"`
}

type ObservabilityQueryTimeSeriesMetricsResponse struct {
	ContentType                    string
	QueryTimeSeriesMetricsResponse *shared.QueryTimeSeriesMetricsResponse
	Status                         *shared.Status
	StatusCode                     int
	RawResponse                    *http.Response
}
