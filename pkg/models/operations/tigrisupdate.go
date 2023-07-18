// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisUpdateRequest struct {
	UpdateRequest shared.UpdateRequest `request:"mediaType=application/json"`
	// Collection name where to update documents
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	// Project name whose db is under target  to update documents
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *TigrisUpdateRequest) GetUpdateRequest() shared.UpdateRequest {
	if o == nil {
		return shared.UpdateRequest{}
	}
	return o.UpdateRequest
}

func (o *TigrisUpdateRequest) GetCollection() string {
	if o == nil {
		return ""
	}
	return o.Collection
}

func (o *TigrisUpdateRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type TigrisUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default error response
	Status *shared.Status
	// OK
	UpdateResponse *shared.UpdateResponse
}

func (o *TigrisUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TigrisUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TigrisUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TigrisUpdateResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TigrisUpdateResponse) GetUpdateResponse() *shared.UpdateResponse {
	if o == nil {
		return nil
	}
	return o.UpdateResponse
}
