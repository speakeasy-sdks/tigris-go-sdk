// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisCommitTransactionRequest struct {
	CommitTransactionRequest shared.CommitTransactionRequest `request:"mediaType=application/json"`
	// Project name whose DB this transaction belongs to.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *TigrisCommitTransactionRequest) GetCommitTransactionRequest() shared.CommitTransactionRequest {
	if o == nil {
		return shared.CommitTransactionRequest{}
	}
	return o.CommitTransactionRequest
}

func (o *TigrisCommitTransactionRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type TigrisCommitTransactionResponse struct {
	// OK
	CommitTransactionResponse *shared.CommitTransactionResponse
	ContentType               string
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}

func (o *TigrisCommitTransactionResponse) GetCommitTransactionResponse() *shared.CommitTransactionResponse {
	if o == nil {
		return nil
	}
	return o.CommitTransactionResponse
}

func (o *TigrisCommitTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TigrisCommitTransactionResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TigrisCommitTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TigrisCommitTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
