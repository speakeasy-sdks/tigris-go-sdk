package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ObservabilityQueryTimeSeriesMetricsResponse struct {
	ContentType                    string
	QueryTimeSeriesMetricsResponse *shared.QueryTimeSeriesMetricsResponse
	Status                         *shared.Status
	StatusCode                     int
	RawResponse                    *http.Response
}
