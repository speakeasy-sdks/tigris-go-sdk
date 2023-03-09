package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
