package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ObservabilityQuotaLimitsRequest struct {
	Request map[string]interface{} `request:"mediaType=application/json"`
}

type ObservabilityQuotaLimitsResponse struct {
	ContentType         string
	QuotaLimitsResponse *shared.QuotaLimitsResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
