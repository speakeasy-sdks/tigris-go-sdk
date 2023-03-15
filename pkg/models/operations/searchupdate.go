package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchUpdateRequest struct {
	UpdateDocumentRequest shared.UpdateDocumentRequest `request:"mediaType=application/json"`
	Index                 string                       `pathParam:"style=simple,explode=false,name=index"`
	Project               string                       `pathParam:"style=simple,explode=false,name=project"`
}

type SearchUpdateResponse struct {
	ContentType            string
	StatusCode             int
	RawResponse            *http.Response
	Status                 *shared.Status
	UpdateDocumentResponse *shared.UpdateDocumentResponse
}
