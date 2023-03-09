package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchDeleteByQueryPathParams struct {
	Index   string `pathParam:"style=simple,explode=false,name=index"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchDeleteByQueryRequest struct {
	PathParams SearchDeleteByQueryPathParams
	Request    shared.DeleteByQueryRequest `request:"mediaType=application/json"`
}

type SearchDeleteByQueryResponse struct {
	ContentType           string
	DeleteByQueryResponse *shared.DeleteByQueryResponse
	Status                *shared.Status
	StatusCode            int
	RawResponse           *http.Response
}
