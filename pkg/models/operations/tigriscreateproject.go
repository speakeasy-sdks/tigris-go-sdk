package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisCreateProjectRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	Project     string                 `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisCreateProjectResponse struct {
	ContentType           string
	CreateProjectResponse *shared.CreateProjectResponse
	Status                *shared.Status
	StatusCode            int
	RawResponse           *http.Response
}
