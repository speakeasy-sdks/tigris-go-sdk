package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchCreatePathParams struct {
	Index   string `pathParam:"style=simple,explode=false,name=index"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchCreateRequest struct {
	PathParams SearchCreatePathParams
	Request    shared.CreateDocumentRequest `request:"mediaType=application/json"`
}

type SearchCreateResponse struct {
	ContentType            string
	CreateDocumentResponse *shared.CreateDocumentResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
