package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchCreateOrUpdateIndexRequest struct {
	CreateOrUpdateIndexRequest shared.CreateOrUpdateIndexRequest `request:"mediaType=application/json"`
	Name                       string                            `pathParam:"style=simple,explode=false,name=name"`
	Project                    string                            `pathParam:"style=simple,explode=false,name=project"`
}

type SearchCreateOrUpdateIndexResponse struct {
	ContentType                 string
	CreateOrUpdateIndexResponse *shared.CreateOrUpdateIndexResponse
	Status                      *shared.Status
	StatusCode                  int
	RawResponse                 *http.Response
}
