package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchGetPathParams struct {
	Index   string `pathParam:"style=simple,explode=false,name=index"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchGetQueryParams struct {
	Ids []string `queryParam:"style=form,explode=true,name=ids"`
}

type SearchGetRequest struct {
	PathParams  SearchGetPathParams
	QueryParams SearchGetQueryParams
}

type SearchGetResponse struct {
	ContentType         string
	GetDocumentResponse *shared.GetDocumentResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
