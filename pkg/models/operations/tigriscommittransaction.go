package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisCommitTransactionRequest struct {
	CommitTransactionRequest shared.CommitTransactionRequest `request:"mediaType=application/json"`
	Project                  string                          `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisCommitTransactionResponse struct {
	CommitTransactionResponse *shared.CommitTransactionResponse
	ContentType               string
	Status                    *shared.Status
	StatusCode                int
	RawResponse               *http.Response
}
