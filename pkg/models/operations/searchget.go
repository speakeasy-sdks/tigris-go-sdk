package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
