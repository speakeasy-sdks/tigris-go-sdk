package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchGetRequest struct {
	Ids     []string `queryParam:"style=form,explode=true,name=ids"`
	Index   string   `pathParam:"style=simple,explode=false,name=index"`
	Project string   `pathParam:"style=simple,explode=false,name=project"`
}

type SearchGetResponse struct {
	ContentType         string
	GetDocumentResponse *shared.GetDocumentResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
