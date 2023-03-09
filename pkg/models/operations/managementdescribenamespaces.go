package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementDescribeNamespacesResponse struct {
	ContentType                string
	DescribeNamespacesResponse *shared.DescribeNamespacesResponse
	Status                     *shared.Status
	StatusCode                 int
	RawResponse                *http.Response
}
