package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type SearchDeletePathParams struct {
	Index   string `pathParam:"style=simple,explode=false,name=index"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchDeleteRequest struct {
	PathParams SearchDeletePathParams
	Request    shared.DeleteDocumentRequest `request:"mediaType=application/json"`
}

type SearchDeleteResponse struct {
	ContentType            string
	DeleteDocumentResponse *shared.DeleteDocumentResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
