package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchCreateRequest struct {
	CreateDocumentRequest shared.CreateDocumentRequest `request:"mediaType=application/json"`
	Index                 string                       `pathParam:"style=simple,explode=false,name=index"`
	Project               string                       `pathParam:"style=simple,explode=false,name=project"`
}

type SearchCreateResponse struct {
	ContentType            string
	CreateDocumentResponse *shared.CreateDocumentResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
