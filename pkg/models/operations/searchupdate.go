package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchUpdatePathParams struct {
	Index   string `pathParam:"style=simple,explode=false,name=index"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchUpdateRequest struct {
	PathParams SearchUpdatePathParams
	Request    shared.UpdateDocumentRequest `request:"mediaType=application/json"`
}

type SearchUpdateResponse struct {
	ContentType            string
	StatusCode             int
	RawResponse            *http.Response
	Status                 *shared.Status
	UpdateDocumentResponse *shared.UpdateDocumentResponse
}
