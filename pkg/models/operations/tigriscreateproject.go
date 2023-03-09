package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisCreateProjectPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisCreateProjectRequest struct {
	PathParams TigrisCreateProjectPathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
}

type TigrisCreateProjectResponse struct {
	ContentType           string
	CreateProjectResponse *shared.CreateProjectResponse
	Status                *shared.Status
	StatusCode            int
	RawResponse           *http.Response
}
