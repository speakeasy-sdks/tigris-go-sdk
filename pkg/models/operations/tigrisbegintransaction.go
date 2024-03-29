// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisBeginTransactionRequest struct {
	BeginTransactionRequest shared.BeginTransactionRequest `request:"mediaType=application/json"`
	// Project name whose DB this transaction belongs to.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *TigrisBeginTransactionRequest) GetBeginTransactionRequest() shared.BeginTransactionRequest {
	if o == nil {
		return shared.BeginTransactionRequest{}
	}
	return o.BeginTransactionRequest
}

func (o *TigrisBeginTransactionRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type TigrisBeginTransactionResponse struct {
	// OK
	BeginTransactionResponse *shared.BeginTransactionResponse
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TigrisBeginTransactionResponse) GetBeginTransactionResponse() *shared.BeginTransactionResponse {
	if o == nil {
		return nil
	}
	return o.BeginTransactionResponse
}

func (o *TigrisBeginTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TigrisBeginTransactionResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TigrisBeginTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TigrisBeginTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
