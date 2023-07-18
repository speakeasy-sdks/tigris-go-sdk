// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisReplaceRequest struct {
	ReplaceRequest shared.ReplaceRequest `request:"mediaType=application/json"`
	// Collection name where to replace documents.
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	// Project name whose db is under target to replace documents.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

func (o *TigrisReplaceRequest) GetReplaceRequest() shared.ReplaceRequest {
	if o == nil {
		return shared.ReplaceRequest{}
	}
	return o.ReplaceRequest
}

func (o *TigrisReplaceRequest) GetCollection() string {
	if o == nil {
		return ""
	}
	return o.Collection
}

func (o *TigrisReplaceRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type TigrisReplaceResponse struct {
	ContentType string
	// OK
	ReplaceResponse *shared.ReplaceResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}

func (o *TigrisReplaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TigrisReplaceResponse) GetReplaceResponse() *shared.ReplaceResponse {
	if o == nil {
		return nil
	}
	return o.ReplaceResponse
}

func (o *TigrisReplaceResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TigrisReplaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TigrisReplaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
