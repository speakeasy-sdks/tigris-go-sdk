package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDeleteProjectRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	Project     string                 `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDeleteProjectResponse struct {
	ContentType           string
	DeleteProjectResponse *shared.DeleteProjectResponse
	Status                *shared.Status
	StatusCode            int
	RawResponse           *http.Response
}
