package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisCommitTransactionPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisCommitTransactionRequest struct {
	PathParams TigrisCommitTransactionPathParams
	Request    shared.CommitTransactionRequest `request:"mediaType=application/json"`
}

type TigrisCommitTransactionResponse struct {
	CommitTransactionResponse *shared.CommitTransactionResponse
	ContentType               string
	Status                    *shared.Status
	StatusCode                int
	RawResponse               *http.Response
}
