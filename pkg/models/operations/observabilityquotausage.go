package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type ObservabilityQuotaUsageRequest struct {
	Request map[string]interface{} `request:"mediaType=application/json"`
}

type ObservabilityQuotaUsageResponse struct {
	ContentType        string
	QuotaUsageResponse *shared.QuotaUsageResponse
	Status             *shared.Status
	StatusCode         int
	RawResponse        *http.Response
}
