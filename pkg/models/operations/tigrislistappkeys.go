package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisListAppKeysPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisListAppKeysRequest struct {
	PathParams TigrisListAppKeysPathParams
}

type TigrisListAppKeysResponse struct {
	ContentType         string
	ListAppKeysResponse *shared.ListAppKeysResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
