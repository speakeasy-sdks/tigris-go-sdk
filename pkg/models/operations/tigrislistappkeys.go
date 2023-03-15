package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListAppKeysRequest struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisListAppKeysResponse struct {
	ContentType         string
	ListAppKeysResponse *shared.ListAppKeysResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
