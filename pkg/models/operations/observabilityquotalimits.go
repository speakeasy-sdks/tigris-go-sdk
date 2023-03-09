package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
