// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisInsertRequest struct {
	InsertRequest shared.InsertRequest `request:"mediaType=application/json"`
	// Collection name where to insert documents.
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	// Project name whose db is under target to insert documents.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *TigrisInsertRequest) GetInsertRequest() shared.InsertRequest {
	if o == nil {
		return shared.InsertRequest{}
	}
	return o.InsertRequest
}

func (o *TigrisInsertRequest) GetCollection() string {
	if o == nil {
		return ""
	}
	return o.Collection
}

func (o *TigrisInsertRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type TigrisInsertResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	InsertResponse *shared.InsertResponse
	// Default error response
	Status *shared.Status
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *TigrisInsertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TigrisInsertResponse) GetInsertResponse() *shared.InsertResponse {
	if o == nil {
		return nil
	}
	return o.InsertResponse
}

func (o *TigrisInsertResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TigrisInsertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TigrisInsertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
