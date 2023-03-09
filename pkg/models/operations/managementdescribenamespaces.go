package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type ManagementDescribeNamespacesResponse struct {
	ContentType                string
	DescribeNamespacesResponse *shared.DescribeNamespacesResponse
	Status                     *shared.Status
	StatusCode                 int
	RawResponse                *http.Response
}
