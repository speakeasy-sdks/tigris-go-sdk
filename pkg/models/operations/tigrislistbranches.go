package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListBranchesPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisListBranchesRequest struct {
	PathParams TigrisListBranchesPathParams
}

type TigrisListBranchesResponse struct {
	ContentType          string
	ListBranchesResponse *shared.ListBranchesResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
