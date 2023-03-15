package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisBeginTransactionRequest struct {
	BeginTransactionRequest shared.BeginTransactionRequest `request:"mediaType=application/json"`
	Project                 string                         `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisBeginTransactionResponse struct {
	BeginTransactionResponse *shared.BeginTransactionResponse
	ContentType              string
	Status                   *shared.Status
	StatusCode               int
	RawResponse              *http.Response
}
