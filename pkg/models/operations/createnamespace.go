package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type CreateNamespaceRequest struct {
	Request shared.CreateNamespaceRequest `request:"mediaType=application/json"`
}

type CreateNamespaceResponse struct {
	ContentType             string
	CreateNamespaceResponse *shared.CreateNamespaceResponse
	Status                  *shared.Status
	StatusCode              int
	RawResponse             *http.Response
}
