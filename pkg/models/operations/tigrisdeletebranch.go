package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDeleteBranchRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	Branch      string                 `pathParam:"style=simple,explode=false,name=branch"`
	Project     string                 `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDeleteBranchResponse struct {
	ContentType          string
	DeleteBranchResponse *shared.DeleteBranchResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
