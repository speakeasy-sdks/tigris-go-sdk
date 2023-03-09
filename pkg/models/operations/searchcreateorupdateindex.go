package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchCreateOrUpdateIndexPathParams struct {
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchCreateOrUpdateIndexRequest struct {
	PathParams SearchCreateOrUpdateIndexPathParams
	Request    shared.CreateOrUpdateIndexRequest `request:"mediaType=application/json"`
}

type SearchCreateOrUpdateIndexResponse struct {
	ContentType                 string
	CreateOrUpdateIndexResponse *shared.CreateOrUpdateIndexResponse
	Status                      *shared.Status
	StatusCode                  int
	RawResponse                 *http.Response
}
