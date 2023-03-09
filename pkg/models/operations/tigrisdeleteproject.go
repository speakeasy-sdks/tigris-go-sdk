package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDeleteProjectPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDeleteProjectRequest struct {
	PathParams TigrisDeleteProjectPathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
}

type TigrisDeleteProjectResponse struct {
	ContentType           string
	DeleteProjectResponse *shared.DeleteProjectResponse
	Status                *shared.Status
	StatusCode            int
	RawResponse           *http.Response
}
