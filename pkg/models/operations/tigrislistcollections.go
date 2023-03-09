package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListCollectionsPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisListCollectionsQueryParams struct {
	Branch *string `queryParam:"style=form,explode=true,name=branch"`
}

type TigrisListCollectionsRequest struct {
	PathParams  TigrisListCollectionsPathParams
	QueryParams TigrisListCollectionsQueryParams
}

type TigrisListCollectionsResponse struct {
	ContentType             string
	ListCollectionsResponse *shared.ListCollectionsResponse
	Status                  *shared.Status
	StatusCode              int
	RawResponse             *http.Response
}
