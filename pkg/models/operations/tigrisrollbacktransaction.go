package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisRollbackTransactionRequest struct {
	RollbackTransactionRequest shared.RollbackTransactionRequest `request:"mediaType=application/json"`
	Project                    string                            `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisRollbackTransactionResponse struct {
	ContentType                 string
	RollbackTransactionResponse *shared.RollbackTransactionResponse
	Status                      *shared.Status
	StatusCode                  int
	RawResponse                 *http.Response
}
