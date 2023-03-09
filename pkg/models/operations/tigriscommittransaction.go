package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
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
