package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
