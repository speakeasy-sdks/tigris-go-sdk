package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisRollbackTransactionPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisRollbackTransactionRequest struct {
	PathParams TigrisRollbackTransactionPathParams
	Request    shared.RollbackTransactionRequest `request:"mediaType=application/json"`
}

type TigrisRollbackTransactionResponse struct {
	ContentType                 string
	RollbackTransactionResponse *shared.RollbackTransactionResponse
	Status                      *shared.Status
	StatusCode                  int
	RawResponse                 *http.Response
}
