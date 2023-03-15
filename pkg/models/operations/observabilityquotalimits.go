package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ObservabilityQuotaLimitsResponse struct {
	ContentType         string
	QuotaLimitsResponse *shared.QuotaLimitsResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
