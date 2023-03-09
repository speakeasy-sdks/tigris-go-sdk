package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchCreateByIDPathParams struct {
	ID      string `pathParam:"style=simple,explode=false,name=id"`
	Index   string `pathParam:"style=simple,explode=false,name=index"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchCreateByIDRequest struct {
	PathParams SearchCreateByIDPathParams
	Request    shared.CreateByIDRequest `request:"mediaType=application/json"`
}

type SearchCreateByIDResponse struct {
	ContentType        string
	CreateByIDResponse *shared.CreateByIDResponse
	Status             *shared.Status
	StatusCode         int
	RawResponse        *http.Response
}
