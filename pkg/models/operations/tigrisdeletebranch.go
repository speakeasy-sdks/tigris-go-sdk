package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisDeleteBranchPathParams struct {
	Branch  string `pathParam:"style=simple,explode=false,name=branch"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDeleteBranchRequest struct {
	PathParams TigrisDeleteBranchPathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
}

type TigrisDeleteBranchResponse struct {
	ContentType          string
	DeleteBranchResponse *shared.DeleteBranchResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
