package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListBranchesRequest struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisListBranchesResponse struct {
	ContentType          string
	ListBranchesResponse *shared.ListBranchesResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
