// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ObservabilityQuotaLimitsResponse struct {
	ContentType string
	// OK
	QuotaLimitsResponse *shared.QuotaLimitsResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
