package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
