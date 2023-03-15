package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateNamespaceResponse struct {
	ContentType             string
	CreateNamespaceResponse *shared.CreateNamespaceResponse
	Status                  *shared.Status
	StatusCode              int
	RawResponse             *http.Response
}
