package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ObservabilityQuotaUsageResponse struct {
	ContentType        string
	QuotaUsageResponse *shared.QuotaUsageResponse
	Status             *shared.Status
	StatusCode         int
	RawResponse        *http.Response
}
