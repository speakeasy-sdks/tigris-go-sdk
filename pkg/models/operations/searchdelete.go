package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchDeleteRequest struct {
	DeleteDocumentRequest shared.DeleteDocumentRequest `request:"mediaType=application/json"`
	Index                 string                       `pathParam:"style=simple,explode=false,name=index"`
	Project               string                       `pathParam:"style=simple,explode=false,name=project"`
}

type SearchDeleteResponse struct {
	ContentType            string
	DeleteDocumentResponse *shared.DeleteDocumentResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
