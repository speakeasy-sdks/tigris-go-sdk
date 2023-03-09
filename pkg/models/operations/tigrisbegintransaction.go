package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisBeginTransactionPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisBeginTransactionRequest struct {
	PathParams TigrisBeginTransactionPathParams
	Request    shared.BeginTransactionRequest `request:"mediaType=application/json"`
}

type TigrisBeginTransactionResponse struct {
	BeginTransactionResponse *shared.BeginTransactionResponse
	ContentType              string
	Status                   *shared.Status
	StatusCode               int
	RawResponse              *http.Response
}
