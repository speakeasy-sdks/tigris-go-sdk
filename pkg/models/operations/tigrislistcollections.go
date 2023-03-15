package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListCollectionsRequest struct {
	Branch  *string `queryParam:"style=form,explode=true,name=branch"`
	Project string  `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisListCollectionsResponse struct {
	ContentType             string
	ListCollectionsResponse *shared.ListCollectionsResponse
	Status                  *shared.Status
	StatusCode              int
	RawResponse             *http.Response
}
