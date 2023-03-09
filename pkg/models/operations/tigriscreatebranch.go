package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisCreateBranchPathParams struct {
	Branch  string `pathParam:"style=simple,explode=false,name=branch"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisCreateBranchRequest struct {
	PathParams TigrisCreateBranchPathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
}

type TigrisCreateBranchResponse struct {
	ContentType          string
	CreateBranchResponse *shared.CreateBranchResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
