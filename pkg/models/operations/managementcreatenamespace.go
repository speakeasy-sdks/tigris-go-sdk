package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type ManagementCreateNamespaceRequest struct {
	Request shared.CreateNamespaceRequest `request:"mediaType=application/json"`
}

type ManagementCreateNamespaceResponse struct {
	ContentType             string
	CreateNamespaceResponse *shared.CreateNamespaceResponse
	Status                  *shared.Status
	StatusCode              int
	RawResponse             *http.Response
}
