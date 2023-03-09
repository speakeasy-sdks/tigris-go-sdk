package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementListNamespacesResponse struct {
	ContentType            string
	ListNamespacesResponse *shared.ListNamespacesResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
