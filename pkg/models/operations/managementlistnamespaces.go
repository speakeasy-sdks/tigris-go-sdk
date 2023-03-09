package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type ManagementListNamespacesResponse struct {
	ContentType            string
	ListNamespacesResponse *shared.ListNamespacesResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
