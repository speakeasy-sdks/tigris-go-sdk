// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ObservabilityQuotaUsageResponse struct {
	ContentType string
	// OK
	QuotaUsageResponse *shared.QuotaUsageResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
